package Bin

import (
	"TMS/Accessory"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"strings"
)

func QueryUser(c *gin.Context){
	var token Utils.Token

	if token.TokenExists(c) {
		_, err := UserAssisant.GetUserByEmail(token.Data.Email)

		if token.IsValid() {
			Utils.Answer{Code: 1, Descript: "未登录用户", More: "null"}.OutputJson(c)
			return
		}

		if err != nil {
			Utils.Answer{Code: 1, Descript: "不存在该用户", More: "null"}.OutputJson(c)
			Utils.Error("%s", err)
			return
		}

		value := c.Param("userInfo")

		if value == "" {
			Utils.Answer{Code: 1, Descript: "请转入有效数据", More: "null"}.OutputJson(c)
			return
		}
		value = strings.TrimSpace(value)
		if strings.Contains(value, "@") {

			user, err := UserAssisant.GetUserByEmail(value)

			if err != nil {
				Utils.Answer{Code: 1, Descript: fmt.Sprintf("不存在用户：%s", value), More: "null"}.OutputJson(c)
				Utils.Error("%s", err)
				return
			}else{
				Utils.Answer{Code: 0, Descript: "success", More: user}.OutputJson(c)
				token.UpdateToken(c)
				return
			}
		} else {
			user, err := UserAssisant.GetUserByName(value)
			if err != nil {
				Utils.Answer{Code: 1, Descript: fmt.Sprintf("不存在用户：%s", value), More: "null"}.OutputJson(c)
				Utils.Error("%s", err)
				return
			} else {
				Utils.Answer{Code: 0, Descript: "success", More: user}.OutputJson(c)
				token.UpdateToken(c)
				return
			}
		}

	}else{
		Utils.Answer{Code:1,Descript:"未登录用户",More:""}.OutputJson(c)
		return
	}
}

func QueryDepartmentMember(c *gin.Context) {

	var token Utils.Token

	if token.TokenExists(c){
		_,err:=UserAssisant.GetUserByEmail(token.Data.Email)

		if token.IsValid(){
			Utils.Answer{Code:1,Descript:"cookie到期",More:"null"}.OutputJson(c)
			Utils.Error("%s","cookie到期")
			return
		}

		if err!=nil{
			Utils.Answer{Code:1,Descript:"不存在该用户",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}

		//if user["user"].(map[string]string)["email"]!=token.Data.Email{
		//
		//	Utils.Answer{Code:1,Descript:"数据匹配失败",More:"null"}.OutputJson(c)
		//	return
		//}

		value:=c.Param("department")

		if value==""{
			Utils.Answer{Code:1,Descript:"请转入有效数据",More:"null"}.OutputJson(c)
			return
		}
		value=strings.TrimSpace(value)
		value=strings.ToUpper(value)

		alls,err:=UserAssisant.GetDepartment(value)

		if err!=nil{
			Utils.Answer{Code:1,Descript:"获取数据失败",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}
		if len(alls)==0{
			Utils.Answer{Code:1,Descript:fmt.Sprintf("不存%s部门人员",value),More:"null"}.OutputJson(c)
			return
		}

		Utils.Answer{Code:0,Descript:"sucess",More:alls}.OutputJson(c)
		token.UpdateToken(c)
	}else{
		Utils.Answer{Code:1,Descript:"未登录用户",More:""}.OutputJson(c)

		return
	}

}

func QueryTeamMember(c *gin.Context){

	var token Utils.Token

	if token.TokenExists(c){
		_,err:=UserAssisant.GetUserByEmail(token.Data.Email)

		if token.IsValid(){
			Utils.Answer{Code:1,Descript:"cookie到期",More:"null"}.OutputJson(c)
			return
		}

		if err!=nil{
			Utils.Answer{Code:1,Descript:"不存在该用户",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}

		value:=c.Param("team")


		if value==""{
			Utils.Answer{Code:1,Descript:"请转入有效数据",More:"null"}.OutputJson(c)
			return
		}
		value=strings.TrimSpace(value)

		if v,err:=strconv.Atoi(value);err==nil {
			switch v {
			case 1:
				value = "项目一组"
			case 2:
				value = "项目二组"
			case 3:
				value = "项目三组"
			case 4:
				value = "认证组"
			case 5:
				value = "配置集成组"
			case 6:
				value = "惠州一组"
			case 7:
				value = "惠州二组"
			case 8:
				value = "自动化组"
			default:
				Utils.Answer{Code:1,Descript:"请转入有效数据1至8",More:"null"}.OutputJson(c)
				return
			}
		}
		alls,err:=UserAssisant.GetTeam(value)

		if err!=nil{
			Utils.Answer{Code:1,Descript:"获取数据失败",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}
		if len(alls)==0{
			Utils.Answer{Code:1,Descript:fmt.Sprintf("不存%s人员",value),More:"null"}.OutputJson(c)
			return
		}

		Utils.Answer{Code:0,Descript:"sucess",More:alls}.OutputJson(c)
		token.UpdateToken(c)
	}else{
		Utils.Answer{Code:1,Descript:"示登录用户",More:""}.OutputJson(c)

		return
	}

}



func QueryTeamRoleMember(c *gin.Context){

	var token Utils.Token

	if token.TokenExists(c){
		_,err:=UserAssisant.GetUserByEmail(token.Data.Email)

		if token.IsValid(){
			Utils.Answer{Code:1,Descript:"cookie到期",More:"null"}.OutputJson(c)
			return
		}

		if err!=nil{
			Utils.Answer{Code:1,Descript:"不存在该用户",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}

		team:=c.Param("team")
		role:=c.Param("role")

		switch strings.ToLower(role){
			case "tester":
				role="Tester"
			case "testleader":
				role="TestLeader"
			case "testteamer":
				role="TestTeamer"
			case "testmanager":
				role="TestManager"
			case "softleader":
				role="SoftLeader"
			default:
				Utils.Answer{Code:1,Descript:fmt.Sprintf("测试角色%s无效",role),More:"null"}.OutputJson(c)
				return

		}

		if team==""{
			Utils.Answer{Code:1,Descript:"请转入有效数据",More:"null"}.OutputJson(c)
			return
		}

		team=strings.TrimSpace(team)
		role=strings.TrimSpace(role)

		if v,err:=strconv.Atoi(team);err==nil {
			switch v {
			case 1:
				team = "项目一组"
			case 2:
				team = "项目二组"
			case 3:
				team = "项目三组"
			case 4:
				team = "认证组"
			case 5:
				team = "配置集成组"
			case 6:
				team = "惠州一组"
			case 7:
				team = "惠州二组"
			case 8:
				team = "自动化组"
			default:
				Utils.Answer{Code:1,Descript:"请转入有效数据1至8",More:"null"}.OutputJson(c)
				return
			}
		}

		alls,err:=UserAssisant.GetTeamRole(team,role)

		if err!=nil{
			Utils.Answer{Code:1,Descript:"获取数据失败",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}
		if len(alls)==0{
			Utils.Answer{Code:1,Descript:fmt.Sprintf("不存%s人员",team),More:"null"}.OutputJson(c)
			return
		}

		Utils.Answer{Code:0,Descript:"sucess",More:alls}.OutputJson(c)
		token.UpdateToken(c)
	}else{
		Utils.Answer{Code:1,Descript:"示登录用户",More:""}.OutputJson(c)

		return
	}

}

func QueryDepartmentRoleMember(c *gin.Context){

	var token Utils.Token

	if token.TokenExists(c){
		_,err:=UserAssisant.GetUserByEmail(token.Data.Email)

		if token.IsValid(){
			Utils.Answer{Code:1,Descript:"cookie到期",More:"null"}.OutputJson(c)
			return
		}

		if err!=nil{
			Utils.Answer{Code:1,Descript:"不存在该用户",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}

		department:=c.Param("department")
		role:=c.Param("role")

		switch strings.ToLower(role){
			case "tester":
				role="Tester"
			case "testleader":
				role="TestLeader"
			case "testmanager":
				role="TestManager"
			case "testteamer":
				role="TestTeamer"
			case "softleader":
				role="SoftLeader"
			default:
				Utils.Answer{Code:1,Descript:fmt.Sprintf("测试角色%s无效",role),More:"null"}.OutputJson(c)
				return
		}

		if department==""{
			Utils.Answer{Code:1,Descript:"请转入有效数据",More:"null"}.OutputJson(c)
			return
		}

		department=strings.ToUpper(department)
		department=strings.TrimSpace(department)
		role=strings.TrimSpace(role)

		alls,err:=UserAssisant.GetDepartmentRole(department,role)

		if err!=nil{
			Utils.Answer{Code:1,Descript:"获取数据失败",More:"null"}.OutputJson(c)
			Utils.Error("%s",err)
			return
		}
		if len(alls)==0{
			Utils.Answer{Code:1,Descript:fmt.Sprintf("不存%s，%s人员",department,role),More:"null"}.OutputJson(c)
			return
		}

		Utils.Answer{Code:0,Descript:"sucess",More:alls}.OutputJson(c)
		token.UpdateToken(c)
	}else{
		Utils.Answer{Code:1,Descript:"示登录用户",More:""}.OutputJson(c)

		return
	}

}