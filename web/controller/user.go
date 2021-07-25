package controller

import (
	"bj38web/web/model"
	getCaptcha "bj38web/web/proto/getCaptcha"
	user "bj38web/web/proto/user"
	"bj38web/web/utils"
	"context"
	"encoding/json"
	"fmt"
	"github.com/afocus/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"image/png"
	"net/http"
)

func GetSeesion(ctx *gin.Context) {
	// 获取Session信息
	// 初始化错误返回Map
	resp := make(map[string]interface{})
	// 获取 Session 数据
	s := sessions.Default(ctx) // 初始化 Session 对象
	userName := s.Get("name")
	if userName == nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		var nameData struct {
			Name string `json:"name"`
		}
		nameData.Name = userName.(string)
		resp["data"] = nameData
	}
	ctx.JSON(http.StatusOK, resp)
}

//func GetServer() micro.Service{
//	consulReg := consul.NewRegistry()
//	consulSer := micro.NewService(
//		micro.Registry(consulReg),
//	)
//	return consulSer
//}

func GetImageCd(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("go.micro.srv.getCaptcha", utils.GloabMicro.Client())
	// 调用函数
	resp, err := microClient.Call(context.TODO(), &getCaptcha.Request{
		Uuid: uuid,
	})
	if err != nil {
		fmt.Println("microClient.Call err", err)
		return
	}
	var img captcha.Image
	json.Unmarshal(resp.Img, &img)

	// 图片写出到浏览器
	png.Encode(ctx.Writer, img)
	fmt.Println(uuid)
}

func GetSmscd(ctx *gin.Context) {
	// 获取短信验证码
	phone := ctx.Param("phone")
	// 拆分 GET 请求中 的 URL === 格式: 资源路径?k=v&k=v&k=v
	imgCode := ctx.Query("text")
	uuid := ctx.Query("id")
	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("go.micro.srv.getCaptcha", utils.GloabMicro.Client())
	resp, err := microClient.GetSmscd(context.TODO(), &getCaptcha.SmscdRequest{
		Uuid:   uuid,
		Mobile: phone,
		Text:   imgCode,
	})
	if err != nil {
		fmt.Println("microClient.GetSmscd err", err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func PostRet(ctx *gin.Context) {
	var regData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
		SmsCode  string `json:"sms_code"`
	}
	ctx.Bind(&regData)

	// 初始化客户端
	microClient := user.NewUserService("go.micro.srv.user", utils.GloabMicro.Client())

	// 调用远程函数
	resp, err := microClient.Register(context.TODO(), &user.RegReq{
		Mobile:   regData.Mobile,
		SmsCode:  regData.SmsCode,
		Password: regData.PassWord,
	})
	if err != nil {
		fmt.Println("注册用户, 找不到远程服务!", err)
		return
	}
	// 写给浏览器
	ctx.JSON(http.StatusOK, resp)

}

func GetArea(ctx *gin.Context) {
	// 从缓存中获取数据
	var areas []model.Area
	conn := model.RedisPool.Get()
	data, _ := redis.Bytes(conn.Do("get", "area"))
	if len(data) == 0 {
		fmt.Println("")
		// 先从Mysql中获取数据
		model.GlobalConn.Find(&areas)
		areaBuf, _ := json.Marshal(areas)
		conn.Do("set", "area", areaBuf)
	} else {
		json.Unmarshal(data, &areas)
	}

	resp := make(map[string]interface{})
	resp["error"] = 200
	resp["errmes"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = areas
	ctx.JSON(http.StatusOK, resp)
}

func PostLogin(ctx *gin.Context) {
	var LoginData struct {
		Mobile   string `json:"mobile"`
		PassWord string `json:"password"`
	}
	ctx.Bind(&LoginData)
	resp := make(map[string]interface{})
	name, err := model.Login(LoginData.Mobile, LoginData.PassWord)
	if err == nil {
		// 登录成功
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
		// 初始化session
		s := sessions.Default(ctx)
		s.Set("name", name)
		s.Save()
	} else {
		// 登录失败
		resp["errno"] = utils.RECODE_LOGINERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_LOGINERR)
	}
	ctx.JSON(http.StatusOK, resp)
}

func DeleteSession(ctx *gin.Context) {
	resp := make(map[string]interface{})
	// 初始化session
	s := sessions.Default(ctx)
	s.Delete("name")
	err := s.Save()
	if err != nil {
		resp["errno"] = utils.RECODE_IOERR // 没有合适错误,使用 IO 错误!
		resp["errmsg"] = utils.RecodeText(utils.RECODE_IOERR)
	} else {
		resp["errno"] = utils.RECODE_OK
		resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	}
	ctx.JSON(http.StatusOK, resp)
}

func GetUserInfo(ctx *gin.Context) {
	// 获取用户基本信息
	resp := make(map[string]interface{})
	s := sessions.Default(ctx) // Session 初始化
	userName := s.Get("name")  // 根据key 获取Session
	if userName == nil {       // 用户没登录, 但进入该页面, 恶意进入.
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return // 如果出错, 报错, 退出
	}
	// 根据用户名, 获取 用户信息  ---- 查 MySQL 数据库  user 表.
	user, err := model.GetUserInfo(userName.(string)) // 类型断言，传参
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return // 如果出错, 报错, 退出
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)

	temp := make(map[string]interface{})
	temp["user_id"] = user.ID
	temp["name"] = user.Name
	temp["mobile"] = user.Mobile
	temp["real_name"] = user.Real_name
	temp["id_card"] = user.Id_card
	temp["avatar_url"] = user.Avatar_url

	resp["data"] = temp
	ctx.JSON(http.StatusOK, resp)
}

func PutUserInfo(ctx *gin.Context) {
	// 获取当前用户名
	// 获取新用户名
	// 更新用户名
	s := sessions.Default(ctx) // 初始化Session 对象
	userName := s.Get("name")
	var nameData struct {
		Name string `json:"name"`
	}
	ctx.Bind(&nameData)
	// 更新用户名
	resp := make(map[string]interface{})
	defer ctx.JSON(http.StatusOK, resp)
	// 更新数据库中的 name
	err := model.UpdateUserName(nameData.Name, userName.(string))
	if err != nil {
		resp["errno"] = utils.RECODE_DBERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_DBERR)
		return
	}
	// 更新 Session 数据
	s.Set("name", nameData.Name)
	err = s.Save() // 必须保存
	if err != nil {
		resp["errno"] = utils.RECODE_SESSIONERR
		resp["errmsg"] = utils.RecodeText(utils.RECODE_SESSIONERR)
		return
	}
	resp["errno"] = utils.RECODE_OK
	resp["errmsg"] = utils.RecodeText(utils.RECODE_OK)
	resp["data"] = nameData
}

// 上传头像
func PostAvatar(ctx *gin.Context) {
	// 获取图片文件, 静态文件对象
	file, _ := ctx.FormFile("avatar")
	// 上传文件到项目中
	err := ctx.SaveUploadedFile(file, "test/"+file.Filename)
	fmt.Println(err)
}
