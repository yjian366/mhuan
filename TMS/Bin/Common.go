package Bin

import (
	"TMS/Accessory"
	"errors"
	"github.com/gin-gonic/gin"
)

func VerifyIdentity(c *gin.Context)(token Utils.Token,err error){

	if token.TokenExists(c) {
		_, err = UserAssisant.GetUserByEmail(token.Data.Email)

		if !token.IsValid() {
			err=errors.New("身份失效")
			return
		}

		if err != nil {
			err=errors.New("不存在该用户")
			return
		}

	}else{
		err=errors.New("未登录用户")
		return
	}
	return
}
