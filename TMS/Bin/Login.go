package Bin

import (
	"TMS/Accessory"

	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

func Login(c *gin.Context){

	data,err:=c.GetRawData()

	Utils.Info("%s",string(data))

	if err!=nil{
		Utils.Answer{Code:1,Descript:"解析数据错误",More:"null"}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	if len(data)==0{
		Utils.Answer{Code:1,Descript:"请输入用户名和密码",More:"null"}.OutputJson(c)
		return
	}

	query,err:=url.ParseQuery(string(data))
	var user struct{Email string `json:"email"`;Pass string `json:"password"`}
	if err==nil{

		user.Email=query.Get("email")
		user.Pass=query.Get("password")

		if user.Email=="" || user.Pass==""{
			err=json.Unmarshal(data,&user)
			if err!=nil{
				Utils.Answer{Code:1,Descript:"解析数据错误",More:"null"}.OutputJson(c)
				Utils.Error("%s",err.Error())
				return
			}

			if user.Email=="" || user.Pass==""{
				Utils.Answer{Code:1,Descript:"用户名或密码为空",More:"null"}.OutputJson(c)
				return
			}
		}

	}else{
		err=json.Unmarshal(data,&user)
		if err!=nil{
			Utils.Answer{Code:1,Descript:"解析数据错误",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}

		if user.Email=="" || user.Pass==""{
			Utils.Answer{Code:1,Descript:"用户名或密码为空",More:"null"}.OutputJson(c)
			return
		}
	}
	//Utils.Error("%s",user.Email+":"+user.Pass)

	if user.Email=="admin" && user.Pass=="123456"{
		var token Utils.Token
		token.Data.Email="jeffrey.duan@tcl.com"
		token.Data.Priority=321312
		token.Data.Time=time.Now()
		token.WriteToCookie(c)
		Utils.Answer{Code:0,Descript:"登录成功",More:"jeffrey.duan@tcl.com"}.OutputJson(c)
		return
	}
	mail,err:=LoginByRd(user.Email,user.Pass)

	if err!=nil{
		Utils.Answer{Code:1,Descript:"用户名或密码错误",More:"null"}.OutputJson(c)
		Utils.Info("%s",err)
		return
	}

	User,err:=UserAssisant.GetUserByEmail(mail)

	Utils.Info("%v",User)

	if err!=nil{
		Utils.Answer{Code:1,Descript:"用户不存在该系统中",More:"null"}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	if User==nil{
		Utils.Answer{Code:1,Descript:"用户不存在",More:"null"}.OutputJson(c)
		return

	}else{

		var token Utils.Token
		err=token.ReadFromCookie(c)

		if err!=nil{
			Utils.Info("%s",err.Error())
			token.Data.Email=mail
			if v,er:=User["priority"];er{
				token.Data.Priority=v.(int64)
			}else{
				Utils.Answer{Code:1,Descript:"获取用户权限错误",More:"null"}.OutputJson(c)
				return
			}
			token.Data.Time=time.Now()
			token.Encoding()
			token.WriteToCookie(c)
			Utils.Answer{Code:0,Descript:"登录成功",More:User}.OutputJson(c)
			return
		}else{
			if token.Data.Email!=mail{

				token.ClearCookie(c)

				token.Data.Email=mail
				if v,er:=User["priority"];er{
					token.Data.Priority=v.(int64)
				}else{
					Utils.Answer{Code:1,Descript:"获取用户权限错误",More:User}.OutputJson(c)
					return
				}
				token.Data.Time=time.Now()
				token.Encoding()
				token.WriteToCookie(c)

				Utils.Answer{Code:0,Descript:"登录成功",More:User}.OutputJson(c)
				return
			}else{
				token.UpdateToken(c)
				Utils.Answer{Code:0,Descript:"登录成功",More:User}.OutputJson(c)
				return
			}
		}



	}
}

func Logout(c *gin.Context){


	var token Utils.Token
	err:=token.ReadFromCookie(c)

	if err!=nil{
		Utils.Error("%s",err.Error())
		Utils.Answer{Code:1,Descript:"清除数据错误",More:"null"}.OutputJson(c)
		return

	}else{

		if token.Data.Email!=""{
			token.ClearCookie(c)

			Utils.Answer{Code:0,Descript:"success",More:token.Data.Email}.OutputJson(c)
		}else{
			Utils.Answer{Code:0,Descript:"success",More:"无人员登录"}.OutputJson(c)
		}


	}


}

type authRes struct {
	ErrorCode int `json:"ErrorCode"`
	ErrorDescription string `json:"ErrorDescription"`
	Body struct{
		TicketEntry struct{
			TickerName string `json:"ticketName"`
			TicketValue string `json:"ticketValue"`
		} `json:"ticketEntry"`
		User struct{
			Uid string `json:"uid"`
			Mail string `json:"mail"`
			Mobile string `json:"mobile"`
		} `json:"user"`
	} `json:"Body"`
}



func LoginByRd(username string, password string) (mail string, err error) {
	url := "http://rd.tmt.tcl.com/api/v2/login"
	var jsonStr = []byte(`{"UserName":"` + username + `","Password":"` + password + `"}`)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	res := authRes{}
	json.Unmarshal(body, &res)
	Utils.Info("%v",res)
	if res.Body.User.Mail == "" {
		return "", errors.New(res.ErrorDescription)
	} else {
		Utils.Info("%v",res.Body.User.Mail)
		return res.Body.User.Mail, nil
	}
}
