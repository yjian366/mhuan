package Bin

import (
	"TMS/Accessory"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"net/url"
	"strconv"
	"unicode/utf8"
)



func CreateTaskForGBK(c *gin.Context){

	token,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	value,err:=c.GetRawData()
	_,err=url.ParseQuery(string(value))

	if err!=nil{
		Utils.Answer{Code:1,Descript:"请求方式错误，需要传递json数据",More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}
	var task bson.M
	v,err:=GBKDecode(value)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	err=json.Unmarshal(v,&task)

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	TaskAssisant.Init()
	err=TaskAssisant.Add(task)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	Utils.Answer{Code:0,Descript:"添加任务成功",More:token.Data.Email}.OutputJson(c)
	return


}
func CreateTaskNoAuth(c *gin.Context){

	value,err:=c.GetRawData()
	_,err=url.ParseQuery(string(value))

	if err!=nil{
		Utils.Answer{Code:1,Descript:"请求方式错误，需要传递json数据",More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}
	var task bson.M
	//v,err:=GBKDecode(value)
	if !utf8.Valid(value){
		err=errors.New("请使用UTF-8编码上传数据")
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	err=json.Unmarshal(value,&task)

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	TaskAssisant.Init()
	err=TaskAssisant.Add(task)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	Utils.Answer{Code:0,Descript:"添加任务成功",More:""}.OutputJson(c)
	return


}

func CreateTask(c *gin.Context){

	token,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	value,err:=c.GetRawData()
	_,err=url.ParseQuery(string(value))

	if err!=nil{
		Utils.Answer{Code:1,Descript:"请求方式错误，需要传递json数据",More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}
	var task bson.M
	//v,err:=GBKDecode(value)
	if !utf8.Valid(value){
		err=errors.New("请使用UTF-8编码上传数据")
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	err=json.Unmarshal(value,&task)

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	TaskAssisant.Init()
	err=TaskAssisant.Add(task)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err)
		return
	}

	Utils.Answer{Code:0,Descript:"添加任务成功",More:token.Data.Email}.OutputJson(c)
	return


}

func GetAllTasks(c *gin.Context){
	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}

	value,err:=TaskAssisant.GetAllTasks()

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByCreater(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	name:=c.Param("name")

	value,err:=TaskAssisant.GetTasksByCondition(map[string]interface{}{"creater":name})

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByExecutor(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	name:=c.Param("name")

	value,err:=TaskAssisant.GetTasksByCondition(map[string]interface{}{"execute_owner":name})

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByModule(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	name:=c.Param("name")

	value,err:=TaskAssisant.GetTasksByCondition(map[string]interface{}{"module_name":name})

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}


func GetTasksBySystemID(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	name:=c.Param("id")

	value,err:=TaskAssisant.GetTasksByCondition(map[string]interface{}{"system_id":name})

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}
func GetTasksByTaskName(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	name:=c.Param("name")

	value,err:=TaskAssisant.GetTasksByCondition(map[string]interface{}{"name":name})

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByTaskType(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	name:=c.Param("name")

	value,err:=TaskAssisant.GetTasksByCondition(map[string]interface{}{"type":name})

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByTaskStatus(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	name:=c.Param("status")

	value,err:=TaskAssisant.GetTasksByCondition(map[string]interface{}{"complete_statement_notes":name})

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByTaskGreatTime(c *gin.Context){


	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	args:=c.Param("time")

	value,err:=TaskAssisant.GetTasksByCondition(nil)

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}

	if len(value)!=0{
		var nValue []map[string]interface{}
		for _,y:=range value{
			n,err:=strconv.Atoi(args)
			if err!=nil{
				Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
				Utils.Error("%s",err.Error())
				return
			}

			//m,err:=strconv.Atoi(y["create_time"].(string))
			//if err!=nil{
			//	Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
			//	Utils.Error("%s",err.Error())
			//	return
			//}


			if int(y["create_time"].(float64))>n{
				nValue=append(nValue,y)
			}
		}
		Utils.Answer{Code:0,Descript:"",More:nValue}.OutputJson(c)
		return
	}

	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByTaskGreatETime(c *gin.Context){


	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	args:=c.Param("time")

	value,err:=TaskAssisant.GetTasksByCondition(nil)

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}

	if len(value)!=0{
		var nValue []map[string]interface{}
		for _,y:=range value{
			n,err:=strconv.Atoi(args)
			if err!=nil{
				Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
				Utils.Error("%s",err.Error())
				return
			}

			//m,err:=strconv.Atoi(y["create_time"].(string))
			//if err!=nil{
			//	Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
			//	Utils.Error("%s",err.Error())
			//	return
			//}


			if int(y["create_time"].(float64))>=n{
				nValue=append(nValue,y)
			}
		}
		Utils.Answer{Code:0,Descript:"",More:nValue}.OutputJson(c)
		return
	}

	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}

func GetTasksByTaskLessTime(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	args:=c.Param("time")

	value,err:=TaskAssisant.GetTasksByCondition(nil)

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}

	if len(value)!=0{
		var nValue []map[string]interface{}
		for _,y:=range value{
			n,err:=strconv.Atoi(args)
			if err!=nil{
				Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
				Utils.Error("%s",err.Error())
				return
			}

			//m,err:=strconv.Atoi(y["create_time"].(string))
			//if err!=nil{
			//	Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
			//	Utils.Error("%s",err.Error())
			//	return
			//}


			if int(y["create_time"].(float64))<n{
				nValue=append(nValue,y)
			}
		}
		Utils.Answer{Code:0,Descript:"",More:nValue}.OutputJson(c)
		return
	}

	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}
func GetTasksByTaskLessETime(c *gin.Context){

	_,err:=VerifyIdentity(c)
	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}
	args:=c.Param("time")

	value,err:=TaskAssisant.GetTasksByCondition(nil)

	if err!=nil{
		Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
		Utils.Error("%s",err.Error())
		return
	}

	if len(value)!=0{
		var nValue []map[string]interface{}
		for _,y:=range value{
			n,err:=strconv.Atoi(args)
			if err!=nil{
				Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
				Utils.Error("%s",err.Error())
				return
			}

			//m,err:=strconv.Atoi(y["create_time"].(string))
			//if err!=nil{
			//	Utils.Answer{Code:1,Descript:err.Error(),More:""}.OutputJson(c)
			//	Utils.Error("%s",err.Error())
			//	return
			//}


			if int(y["create_time"].(float64))<=n{
				nValue=append(nValue,y)
			}
		}
		Utils.Answer{Code:0,Descript:"",More:nValue}.OutputJson(c)
		return
	}

	Utils.Answer{Code:0,Descript:"",More:value}.OutputJson(c)
	return
}