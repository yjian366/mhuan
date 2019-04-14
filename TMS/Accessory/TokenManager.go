package Utils
/*
import (
	"encoding/base64"
	"sync"
	"time"
)

package Utils

import (
"encoding/base64"
"sync"
"time"
)

const(
	Key="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)


var GlobalTokenManager TokenManager
func init(){

	GlobalTokenManager.cookieName="TokenID"

}

type Token struct{}

func(this *Token)Encoding(value string)(args string){

	producer:=base64.NewEncoding(Key)
	return producer.EncodeToString([]byte(value))
}

func(this *Token)Decoding(value string)([]byte,error){
	producer:=base64.NewEncoding(Key)
	return producer.DecodeString(value)
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

	args:=this.Encoding(value)
	this.tokens[args]=time.Now()

	return args
}

func(this *TokenProvide)Decrypt(value string)([]byte,error){

	return this.Decoding(value)
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

*/