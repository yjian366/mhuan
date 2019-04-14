package Bin

import (
	"github.com/gin-gonic/gin"
)

const (
	ERROR_OK                = 0    //成功
	ERROR_PARAMERR          = iota+ 1001 //参数错误
	ERROR_CONVERT_DATA				//解析数据错误
	ERROR_NOTEXIST       			//资源不存在
	ERROR_PASSWORD          		//密码错误
	ERROR_USERNOTEXIST      		//用户不存在
	ERROR_TOKENEXPIRE       		//token过期
	ERROR_NOTSUPPORTMAILBOX 		//不支持的邮箱类型
	ERROR_SENDMAILFAIL     			//邮件发送失败
	ERROR_PERMISSIONDIE 			//权限不足
	ERROR_VERIFYCODE       			//验证码错误
	ERROR_FILEEXTERR       			//文件类型错误
	ERROR_CREATEFILEFAIL			//创建文件失败
	ERROR_ALREADYEXIST 				//已经存在
	ERROR_TIMEOUT 					//超时
	ERROR_OPERATE_FAIL 				//操作失败
	ERROR_INVALID_DATA			//校验不合法
	ERROR_UNKNOWN           = 1112000 //未知错误

)

var descOfErrorCode = map[int]string{
	ERROR_OK:                "success",
	ERROR_PARAMERR:          "param error",
	ERROR_NOTEXIST:          "not exist",
	ERROR_PASSWORD:          "username or password error",
	ERROR_USERNOTEXIST:      "user not exist",
	ERROR_TOKENEXPIRE:       "token expire",
	ERROR_NOTSUPPORTMAILBOX: "not support mailbox type",
	ERROR_SENDMAILFAIL:      "send mail fail",
	ERROR_PERMISSIONDIE:     "permission die",
	ERROR_VERIFYCODE:        "verify code error",
	ERROR_FILEEXTERR:        "file ext type error",
	ERROR_CREATEFILEFAIL:    "create file fail",
	ERROR_ALREADYEXIST:      "already exist",
	ERROR_TIMEOUT:           "timeout",
	ERROR_OPERATE_FAIL:      "operate fail",
	ERROR_INVALID_DATA:      "invalid data",
	ERROR_UNKNOWN:           "unknown error",
	ERROR_CONVERT_DATA:		 "解析数据错误",
}

type ResponseBuilder interface {
	Build() gin.H
}

type StringBuilder struct {
	Name  string
	Value string
}

func (s StringBuilder) Build() gin.H {
	return gin.H{
		s.Name: s.Value,
	}
}


type InterfaceBuilder struct {
	Name  string
	Value interface{}
}

func (s InterfaceBuilder) Build() gin.H {

	return gin.H{
		s.Name:s.Value,
	}
}

func ObjectResponse(errorCode int, data ResponseBuilder) gin.H {

	return gin.H{
			"error_code": errorCode,
			"error_desc": descOfErrorCode[errorCode],
			"data":       data.Build(),

	}

}


func ObjectArrayResponse(errorCode int, data ...ResponseBuilder) gin.H {


	if len(data) >0 {

		abc:=[]gin.H{}
		for _,y:=range data{
			abc=append(abc,y.Build())
		}

		return gin.H{
			"error_code": errorCode,
			"error_desc": descOfErrorCode[errorCode],
			"data":       abc,
		}
	} else{

		return gin.H{
			"error_code": errorCode,
			"error_desc": descOfErrorCode[errorCode],
			"data":       nil,
		}

	}
}

func StringResponse(errorCode int, str string) gin.H {
	return gin.H{
		"error_code": errorCode,
		"error_desc": descOfErrorCode[errorCode],
		"data":       str,
	}
}

func StringArrayRespone(errorCode int, strSlice...string) gin.H {

	return  gin.H{
		"error_code": errorCode,
		"error_desc": descOfErrorCode[errorCode],
		"data":       strSlice,
	}
}

func CustomResponse(errorCode int, data interface{}) gin.H {
	return gin.H{
		"error_code": errorCode,
		"error_desc": descOfErrorCode[errorCode],
		"data":       data,
	}
}

func CustomArraysResponse(errorCode int, data...interface{}) gin.H {
	return gin.H{
		"error_code": errorCode,
		"error_desc": descOfErrorCode[errorCode],
		"data":       data,
	}
}