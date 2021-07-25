package handler

import (
	"bj38web/service/user/model"
	"bj38web/service/user/utils"
	"context"
	"fmt"

	user "bj38web/service/user/proto/user"
)

type User struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *User) Register(ctx context.Context, req *user.RegReq, rsp *user.Response) error {
	err := model.CheckSmsCode(req.Mobile, req.SmsCode)
	if err == nil {
		// 如果校验正确. 注册用户. 将数据写入到 MySQL数据库.
		err = model.RegisterUser(req.Mobile, req.Password)
		fmt.Println(err)
		if err != nil {
			rsp.Errno = utils.RECODE_DBERR
			rsp.Errmsg = utils.RecodeText(utils.RECODE_DBERR)
		} else {
			rsp.Errno = utils.RECODE_OK
			rsp.Errmsg = utils.RecodeText(utils.RECODE_OK)
		}
	}

	return nil
}
