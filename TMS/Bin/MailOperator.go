package Bin

import (
	"TMS/Accessory"
	"TMS/Accessory/Mail"
	"fmt"
	"github.com/gin-gonic/gin"
)

func SendMail(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	var ms=mail.MailMessage{}
	err=c.BindJSON(&ms)
	fmt.Println(ms)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	abc:=mail.New()
	abc.SetToMail(ms.To.Mail)
	abc.SetToName(ms.To.Name)
	abc.SetSubject(ms.Subject)
	abc.SetBody(fmt.Sprintf( "<p>Hi%s:<br>  <p style='margin: 1cm'> %s 于 %s 给你创建任务“%s”</p></p>",ms.To.Name,ms.Body.TaskCreater,ms.Body.Date,ms.Body.TaskName))
	err=abc.Send()

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}
	Utils.Answer{Code:0,Descript:"success",More:""}.OutputJson(c)

}
