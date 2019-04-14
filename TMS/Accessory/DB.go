package Utils

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2"
	"reflect"
)
var GlobalDB mdb
type mdb struct{

	db string
	collections string
	data map[string]interface{}
	datas []map[string]interface{}
	lock  chan bool
}

func(this *mdb)GetDB()string{
	return this.db
}

func(this *mdb)GetCollection()string{
	return this.collections
}

func (this *mdb)SetDB(db string){
	this.db=db
}

func(this *mdb)SetCollection(c string){
	this.collections=c
}

func(this *mdb)Lock(){

	if this.lock==nil{
		this.lock=make(chan bool,1)
	}
	this.lock<-true
}
func(this *mdb)UnLock(){

	if this.lock==nil{
		this.lock=make(chan bool,1)
	}
	<-this.lock
}

func(this *mdb)GetResult()map[string]interface{}{
	this.Lock()
	defer this.UnLock()
	return this.data
}

func(this *mdb)GetResults()[]map[string]interface{}{
	this.Lock()
	defer this.UnLock()
	return this.datas
}

func(this *mdb) QueryAll(args interface{})(err error,result []map[string]interface{}){
	this.Lock()
	defer this.UnLock()
	switch reflect.ValueOf(args).Kind() {
		case reflect.Map,reflect.Struct,reflect.Invalid:
		default:
			err=errors.New("the params need map or struct")
			return
	}

	sess,err:=mgo.Dial("0.0.0.0")
	result=this.datas

	defer sess.Close()

	if err!=nil{
		return
	}

	c:=sess.DB(this.db).C(this.collections)
	if c==nil{
		err=errors.New(fmt.Sprint("the %s collection not exists",this.collections))
		return
	}
	this.datas=nil
	err=c.Find(args).All(&this.datas)
	if this.datas!=nil{
		return nil,this.datas
	}
	return err,this.datas
}

func(this *mdb)QueryOne(args interface{})(err error,result map[string]interface{}) {
	this.Lock()
	defer this.UnLock()

	switch reflect.ValueOf(args).Kind() {
	case reflect.Map,reflect.Struct:
	default:
		err=errors.New("the params need map or struct")
		return
	}

	sess,err:=mgo.Dial("0.0.0.0")
	result=this.data
	defer sess.Close()

	if err!=nil{
		return
	}

	c:=sess.DB(this.db).C(this.collections)
	if c==nil{
		err=errors.New(fmt.Sprint("the %s collection not exists",this.collections))
		return
	}
	this.data=nil
	Error("%s",args)
	err=c.Find(args).One(&this.data)
	Error("%s",err)
	Error("%s",this.data)
	if this.data!=nil{
		return nil,this.data
	}

	return err,this.data
}

func(this *mdb)Insert(args...interface{})(err error){
	this.Lock()
	defer this.UnLock()
	for _,v:=range args{
		switch reflect.ValueOf(v).Kind() {
		case reflect.Map,reflect.Struct:
		default:
			err=errors.New("the params need map or struct")
			return
		}

	}

	sess,err:=mgo.Dial("0.0.0.0")

	defer sess.Close()

	if err!=nil{
		return
	}

	c:=sess.DB(this.db).C(this.collections)
	err=c.Insert(args...)

	return
}

func(this *mdb)Update(selector interface{},update interface{})(err error){
	this.Lock()
	defer this.UnLock()
	switch reflect.ValueOf(selector).Kind() {
	case reflect.Map,reflect.Struct:
	default:
		err=errors.New("the params need map or struct")
		return
	}

	switch reflect.ValueOf(update).Kind() {
	case reflect.Map,reflect.Struct:
	default:
		err=errors.New("the params need map or struct")
		return
	}

	sess,err:=mgo.Dial("0.0.0.0")

	defer sess.Close()

	if err!=nil{
		return
	}

	c:=sess.DB(this.db).C(this.collections)
	err=c.Update(selector,update)
	return
}

func(this *mdb)Delete(selector interface{})(err error){
	this.Lock()
	defer this.UnLock()
	switch reflect.ValueOf(selector).Kind() {
	case reflect.Map,reflect.Struct:
	default:
		err=errors.New("the params need map or struct")
		return
	}

	sess,err:=mgo.Dial("0.0.0.0")

	defer sess.Close()

	if err!=nil{
		return
	}

	c:=sess.DB(this.db).C(this.collections)
	err=c.Remove(selector)
	return
}

func(this *mdb)CreateCollection(collectionName string)(err error){
	this.Lock()
	defer this.UnLock()
	sess,err:=mgo.Dial("0.0.0.0")

	defer sess.Close()

	if err!=nil{
		return
	}

	sess.DB(this.db).C(collectionName)

	return
}
