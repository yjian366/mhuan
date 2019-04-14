package Utils

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"
)

const(
	Key="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

type Token struct{

	Data struct{
		Email 		string
		Time  		time.Time
		Priority 	int64
	}

	token string
	expire int
}



func(this *Token)GetToken()string{
	return this.token
}

func(this *Token)SetToken(args string){
	this.token=args
}

func(this *Token)SetExpire(args int){
	this.expire=args
}

func(this *Token)GetExpire()int{
	return this.expire
}

func(this *Token)Encoding()(args string,err error){

	producer:=base64.NewEncoding(Key)
	v,err:=this.marshalData()
	if err!=nil{
		log.Println(err.Error())
		return "",err
	}
	this.token=producer.EncodeToString(v)

	return this.token,nil
}

func(this *Token)Decoding(value ...string)(err error){

	if value==nil{
		producer:=base64.NewEncoding(Key)
		v,err:=producer.DecodeString(this.token)
		if err!=nil{
			log.Println(err.Error())
			return err
		}
		bson.Unmarshal(v,&this.Data)
		return nil
	}else{
		producer:=base64.NewEncoding(Key)
		v,err:=producer.DecodeString(value[0])
		if err!=nil{
			log.Println(err.Error())
			return err
		}
		bson.Unmarshal(v,&this.Data)
		return nil
	}


}
func(this *Token)marshalData()([]byte,error){
	data,err:=bson.Marshal(this.Data)
	if err!=nil{
		log.Println(err.Error())
		return nil,err
	}
	return data,nil
}

func(this *Token)IsValid()bool{

	if this.Data.Time==(time.Time{}){
		err:=errors.New("the time is empty")
		log.Println(err.Error())
		return false
	}
	if time.Now().Sub(this.Data.Time)<=time.Duration(time.Minute*60){
		return false
	}else{
		return true
	}
}

func(this *Token)GenKey()string{
	b := make([]byte, 48)
	if _, err := rand.Read(b); err != nil {
		log.Println(err.Error())
	}
	value:=base64.URLEncoding.EncodeToString(b)
	return value
}

func(this *Token)WriteToCookie(arg *gin.Context)(err error){

	if this.token==""{
		this.Encoding()
	}

	cookie:=http.Cookie{}
	cookie.Name="token"

	cookie.Value=url.QueryEscape(this.token)
	cookie.HttpOnly=true
	if this.expire==0{
		cookie.MaxAge=3600
	}else{
		cookie.MaxAge=this.expire
	}


	http.SetCookie(arg.Writer,&cookie)
	return
}
func(this *Token)ClearCookie(arg *gin.Context)(err error){
	value,err:=arg.Cookie("token")
	if err!=nil{
		return err

	}
	if value==""{
		return
	}else{
		cookie:=http.Cookie{}
		cookie.Name="token"

		cookie.Value=url.QueryEscape(this.token)
		cookie.HttpOnly=true
		cookie.MaxAge=0

		http.SetCookie(arg.Writer,&cookie)

		return
	}

}

func(this *Token)UpdateToken(arg *gin.Context)(err error){

	err=this.ClearCookie(arg)
	if err!=nil{
		return
	}
	err=this.WriteToCookie(arg)
	this.Data.Time=time.Now()
	if err!=nil{
		return
	}
	return
}

func(this *Token)ReadFromCookie(arg *gin.Context)(err error){

	value,err:=arg.Cookie("token")
	if err!=nil{
		return err

	}
	value,err=url.QueryUnescape(value)
	if err!=nil{
		return err
	}

	this.Decoding(value)
	return nil
}

func(this *Token)TokenExists(c *gin.Context)(is bool){
	err:=this.ReadFromCookie(c)
	if err!=nil{
		return false
	}else{
		return true
	}
}


type TokenProvide struct{
	Token
	lock sync.Mutex
	tokens map[string]time.Time
}

func(this *TokenProvide)Encrypt(value string)(string){
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.tokens==nil{
		this.tokens=make(map[string]time.Time)
	}

	args,_:=this.Encoding()
	this.tokens[args]=time.Now()

	return args
}

func(this *TokenProvide)Decrypt(value string)([]byte,error){

	return this.Decrypt(value)
}

func(this *TokenProvide)QueryNames()(args []string){
	this.lock.Lock()
	defer this.lock.Unlock()

	for key,_:=range this.tokens{
		args=append(args,key)
	}
	return
}

func(this *TokenProvide)Del(args string){
	this.lock.Lock()
	defer this.lock.Unlock()

	for key,_:=range this.tokens{
		if key==args{
			delete(this.tokens,key)
		}
	}
}

func(this *TokenProvide)GC(args int64){

	this.lock.Lock()
	defer this.lock.Unlock()


	for key,value:=range this.tokens{

		if value.Unix()+args<time.Now().Unix(){
			delete(this.tokens,key)
		}
	}

}

type TokenManager struct {
	TokenProvide
	cookieName string
	lock        sync.Mutex
	expireTime int64		//以秒为单位
}

func(this *TokenManager)SetExpireTime(args int64){
	this.expireTime=args
}

func(this *TokenManager)GetExpireTime()(args int64){
	return this.expireTime
}


func(this *TokenManager)AutoGC(){

	futureTimeUinx:=time.Now().Unix()+this.expireTime
	futureTime:=time.Unix(futureTimeUinx,0)

	if this.expireTime==0 {
		this.expireTime = 3600
		futureTimeUinx = time.Now().Unix() + this.expireTime
		futureTime = time.Unix(futureTimeUinx, 0)

	}

	go func(){
		for{
			abc:=time.NewTimer(futureTime.Sub(time.Now()))
			select{
			case <-abc.C:{
				this.GC(0)
			}
			}
		}
	}()

}
