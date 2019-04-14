package Utils
import (
	"container/list"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"
	"math/rand"
)


var pder =new(SessionProvider)

var provides = make(map[string]*SessionProvider)

var GlobalSessionManager *SessionManager

func init() {

	pder.list=list.New()
	pder.sessions = make(map[string]*list.Element, 0)
	Register("memory",pder)
	GlobalSessionManager, _ = NewSessionManager("memory", "SessionID", 3600)
}


func Register(name string, provider *SessionProvider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	provides[name] = provider
}


type Session struct {
	sid          string                      //session id唯一标示
	timeAccessed time.Time                   //最后访问时间
	value        map[interface{}]interface{} //session里面存储的值
}

func (st *Session) Set(key, value interface{}) error {
	st.value[key] = value
	pder.Update(st.sid)
	return nil
}

func (st *Session) Get(key interface{}) interface{} {
	pder.Update(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}
func (st *Session) Delete(key interface{}) error {
	delete(st.value, key)
	pder.Update(st.sid)
	return nil
}

func (st *Session) SessionID() string {
	return st.sid
}

type SessionProvider struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做gc
}

func (this *SessionProvider) Create(sid string) (*Session, error) {
	this.lock.Lock()
	defer this.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &Session{sid: sid, timeAccessed: time.Now(), value: v}
	element := this.list.PushBack(newsess)
	this.sessions[sid] = element
	return newsess, nil
}

func (this *SessionProvider) Get(sid string) (*Session, error) {
	if element, ok := this.sessions[sid]; ok {
		return element.Value.(*Session), nil
	} else {
		sess, err := this.Create(sid)
		return sess, err
	}
	return nil, nil
}

func (this *SessionProvider) Del(sid string) error {
	if element, ok := this.sessions[sid]; ok {
		delete(this.sessions, sid)
		this.list.Remove(element)
		return nil
	}
	return nil
}

func (this *SessionProvider) GC(maxlifetime int64) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for {
		element := this.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*Session).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			this.list.Remove(element)
			delete(this.sessions, element.Value.(*Session).sid)

		} else {
			break
		}
	}
}

func (this *SessionProvider) Update(sid string) error {
	this.lock.Lock()
	defer this.lock.Unlock()
	if element, ok := this.sessions[sid]; ok {
		element.Value.(*Session).timeAccessed = time.Now()
		this.list.MoveToFront(element)
		return nil
	}
	return nil
}

func (this *SessionProvider) QueryNames() (SessionNames []string) {
	this.lock.Lock()
	defer this.lock.Unlock()
	for key,_:=range this.sessions{
		SessionNames=append(SessionNames,key)
	}
	return SessionNames

}



type SessionManager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    *SessionProvider
	maxLifeTime int64
}

func (this *SessionManager) sessionId() string {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

func (this *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session *Session) {
	this.lock.Lock()
	defer this.lock.Unlock()
	cookie, err := r.Cookie(this.cookieName)
	if err != nil || cookie.Value == "" {
		sid := this.sessionId()
		session, _ = this.provider.Create(sid)
		cookie := http.Cookie{Name: this.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(this.maxLifeTime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = this.provider.Get(sid)

	}
	return
}

func (this *SessionManager) SessionDestroy(w http.ResponseWriter, r *http.Request){
	cookie, err := r.Cookie(this.cookieName)
	if err != nil || cookie.Value == "" {
		return
	} else {
		this.lock.Lock()
		defer this.lock.Unlock()
		this.provider.Del(cookie.Value)
		expiration := time.Now()
		cookie := http.Cookie{Name: this.cookieName, Path: "/", HttpOnly: true, Expires: expiration, MaxAge: -1}
		http.SetCookie(w, &cookie)
	}
}

func (this *SessionManager) GC() {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.provider.GC(this.maxLifeTime)
	time.AfterFunc(time.Duration(this.maxLifeTime), func() { this.GC() })
}

//func (this *SessionManager) AutoGC() {
//
//	for{
//
//		futureTime:=time.Now().Unix()+this.maxLifeTime
//
//		ft:=time.Unix(futureTime,0)
//
//		myTimer:=time.NewTimer(ft.Sub(time.Now()))
//		select{
//			case <-myTimer.C:
//				this.provider.GC(0)
//		}
//	}
//
//}

func NewSessionManager(provideName, cookieName string, maxLifeTime int64) (*SessionManager, error) {
	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &SessionManager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}

//func login(w http.ResponseWriter, r *http.Request) {
//	sess := globalSessions.SessionStart(w, r)
//	r.ParseForm()
//	if r.Method == "GET" {
//		t, _ := template.ParseFiles("login.gtpl")
//		w.Header().Set("Content-Type", "text/html")
//		t.Execute(w, sess.Get("username"))
//	} else {
//		sess.Set("username", r.Form["username"])
//		http.Redirect(w, r, "/", 302)
//	}
//}