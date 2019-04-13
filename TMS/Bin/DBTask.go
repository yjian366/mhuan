package Bin

import (
	. "TMS/Accessory"
	. "TMS/Comment"
	"strings"
	"sync"
)

var TaskAssisant DBTask

type DBTask struct{
	lock sync.Mutex

}

func init(){
	GlobalDB.SetDB(DB)
	GlobalDB.SetCollection(TASK_COLLECTION)

}

func(this *DBTask)Init(){
	GlobalDB.SetDB(DB)
	GlobalDB.SetCollection(TASK_COLLECTION)
}

func(this *DBTask)Add(result...interface{})(err error){
	this.lock.Lock()
	defer this.lock.Unlock()

	this.Init()
	err=GlobalDB.Insert(result...)

	return
}


func(this *DBTask)InsertTask(mail string)(result map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()

	if !strings.Contains(mail,"@"){
		mail+="@tcl.com"
	}

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

func(this *DBTask)GetAllTasks()(result []map[string]interface{},err error){

	this.lock.Lock()
	defer  this.lock.Unlock()

	this.Init()
	err,result=GlobalDB.QueryAll(nil)

	return
}

func(this *DBTask)GetTasksByCondition(args interface{})(result []map[string]interface{},err error){

	this.lock.Lock()
	defer  this.lock.Unlock()

	this.Init()
	err,result=GlobalDB.QueryAll(args)

	return
}

func(this *DBTask)GetTaskByID(phone string)(result map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()

	err,result=GlobalDB.QueryOne(map[string]interface{}{"phone":phone})
	return

}

func(this *DBTask)GetTaskByName(name string)(result map[string]interface{},err error){

	this.lock.Lock()
	defer this.lock.Unlock()
	err,result=GlobalDB.QueryOne(map[string]interface{}{"name":name})
	return

}
