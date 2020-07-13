package service

import (
	"crypto/md5"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	_const "kbase/common/const"
	"kbase/common/response"
	"kbase/dao"
	. "kbase/model/pojo"
	"strconv"
)

/**
注册
*/
func Registered(author Author) response.Response {
	if author.Name == "" || author.Password == "" || author.Email == "" {
		return response.IsSuccess().SetMsg("参数不合法")

	}

	if result := dao.FindAuthorByEmail(author.Email); result.ID != 0 {
		return response.IsSuccess().SetMsg("邮件地址已经存在")

	}
	// MD5加密
	newPassword := md5.Sum([]byte(author.Password))
	author.Password = fmt.Sprintf("%x", newPassword)
	result := dao.Registered(author)

	if result.Error != nil {
		return response.IsSuccess().SetMsg(result.Error.Error())

	}
	return response.IsSuccess().SetMsg("创建成功")
}

/**
登陆
*/
func Login(email string, password string) response.Response {
	if email == "" || password == "" {
		return response.IsSuccess().SetMsg("参数不合法")
	}

	result := dao.FindAuthorByEmail(email)

	if result.ID == 0 {
		return response.IsSuccess().SetMsg("邮件地址不存在")
	}

	if result.Password != fmt.Sprintf("%x", md5.Sum([]byte(password))) {
		return response.IsSuccess().SetMsg("密码错误")
	}

	//生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorId": result.ID,
	})
	res := make(map[string]string)
	tokenString, err := token.SignedString([]byte(_const.SecretKey))
	if err != nil {
		return response.IsFalse()
	}
	res["token"] = tokenString
	res["authorId"] = strconv.Itoa(int(result.ID))
	return response.IsSuccess().SetMsg("登陆成功").SetData(res)
}
