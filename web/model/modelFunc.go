package model

import (
	"crypto/md5"
	"encoding/hex"
)

func Login(mobile string, pwd string) (string, error) {
	var user User
	m5 := md5.New()
	m5.Write([]byte(pwd))
	pwd_hash := hex.EncodeToString(m5.Sum(nil))
	err := GlobalConn.Where("mobile=?", mobile).Where("password_hash=?", pwd_hash).Find(&user).Error
	return user.Name, err
}

// 获取用户信息
func GetUserInfo(userName string) (User, error) {
	// 实现SQL: select * from user where name = userName;
	var user User
	err := GlobalConn.Where("name = ?", userName).First(&user).Error
	return user, err
}

// 更新用户名
func UpdateUserName(newName, oldName string) error {
	// update user set name = 'itcast' where name = 旧用户名
	return GlobalConn.Model(new(User)).Where("name = ?", oldName).Update("name", newName).Error
}
