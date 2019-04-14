package Bin

import (
	. "TMS/Accessory"
	. "TMS/Comment"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"sync"
)

var UserAssisant DBUser

type DBUser struct{
	lock sync.Mutex

}

func init(){
	GlobalDB.SetDB(DB)
	GlobalDB.SetCollection(USER_COLLECTION)

}
func(this *DBUser)Init(){
	GlobalDB.SetDB(DB)
	GlobalDB.SetCollection(USER_COLLECTION)
}

func(this *DBUser)GetUserByEmail(mail string)(result map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()
	this.Init()
	if !strings.Contains(mail,"@"){
		mail+="@tcl.com"
	}

	//err,result=GlobalDB.QueryOne(bson.M{"user":bson.M{"email":mail,"password":""}})
	//$where:function(){return this.age>7}


	//if user.Email=="" || user.Pass==""{
	//	err=json.Unmarshal(data,&user)
	//	if err!=nil{
	//		Utils.Answer{Code:1,Descript:"解析数据错误",More:"null"}.OutputJson(c)
	//		Utils.Error("%s",err.Error())
	//		return
	//	}
	//
	//	if user.Email=="" || user.Pass==""{
	//		Utils.Answer{Code:1,Descript:"用户名或密码为空",More:"null"}.OutputJson(c)
	//		return
	//	}
	//}
	err,result=GlobalDB.QueryOne(struct{
		User struct{
			Email string
			Password string
		}}{User:struct{
			Email string
			Password string
		}{Email:mail,Password:""}})
	if err!=nil{
		Error("%s",err.Error())
	}

	return

}

func(this *DBUser)GetUserByPhone(phone string)(result map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()

	err,result=GlobalDB.QueryOne(map[string]interface{}{"phone":phone})
	return

}

func(this *DBUser)GetUserByName(name string)(result map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()
	err,result=GlobalDB.QueryOne(map[string]interface{}{"name":name})
	return

}

func(this *DBUser)GetUserCount()(result []map[string]interface{},err error){
	this.lock.Lock()
	defer this.lock.Unlock()

	err,result=GlobalDB.QueryAll(nil)

	return
}

func(this *DBUser)GetRoleCount(role int)(result []map[string]interface{},err error){
	this.lock.Lock()
	defer this.lock.Unlock()

	err,result=GlobalDB.QueryAll(bson.M{"role":bson.M{"$eq":role}})

	return
}


func(this *DBUser)GetTeamCount(team int)(result []map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()

	err,result=GlobalDB.QueryAll(bson.M{"team":bson.M{"$eq":team}})

	return

}

func(this *DBUser)GetDepartment(department string)(result []map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()

	err,result=GlobalDB.QueryAll(bson.M{"department":bson.M{"$eq":department}})

	return

}
func(this *DBUser)GetTeam(team string)(result []map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()
	Info("%s",team)
	err,result=GlobalDB.QueryAll(bson.M{"team":bson.M{"$eq":team}})

	return

}
func(this *DBUser)GetTeamRole(team string,role string)(result []map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()
	Info("%s",team)
	err,result=GlobalDB.QueryAll(bson.M{"team":bson.M{"$eq":team},"title":bson.M{"$eq":role}})

	return

}
func(this *DBUser)GetDepartmentRole(team string,role string)(result []map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()
	Info("%s",team)
	err,result=GlobalDB.QueryAll(bson.M{"department":bson.M{"$eq":team},"title":bson.M{"$eq":role}})

	return

}

func(this *DBUser)Add(result...interface{})(err error){
	this.lock.Lock()
	defer this.lock.Unlock()

	err=GlobalDB.Insert(result...)
	return
}

func(this *DBUser)Delete(condition bson.M)(err error){
	this.lock.Lock()
	defer this.lock.Unlock()

	err=GlobalDB.Delete(condition)

	return
}

func(this *DBUser)Update(old interface{},new interface{})(err error){
	this.lock.Lock()
	defer this.lock.Unlock()

	err=GlobalDB.Update(old,new)

	return
}