package service

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

/*
 * This is based on:
 * https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.2.html
 */

var globalSessions *SessionManager

//  initialize the session manager (init is run automatically)
func init() {
	var err error
	globalSessions, err = NewManager("memory", "gosessionid", 3600)
	if err != nil {
		fmt.Printf("Error creating session manager: %s", err.Error())
	}
}

//Session -- keep track of web session
type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}

//SessionProvider -- overly abstract for what we want -- remove this
type SessionProvider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var providers = make(map[string]SessionProvider)

//SessionManager -- overly abstract for what we want -- remove this
type SessionManager struct {
	cookieName  string
	lock        sync.Mutex // protects session
	provider    SessionProvider
	maxlifetime int64
}

//NewManager -- return a session manager of the given name (overly abstract -- remove)
func NewManager(providerName, cookieName string, maxlifetime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provider %q (forgotten import?)", providerName)
	}
	return &SessionManager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}

// Register makes a session provider available by the provided name.
// If a Register is called twice with the same name or if the driver is nil,
// it panics. We only plan to have a single provider, so all this can go.
func Register(name string, provider SessionProvider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("session: Register called twice for provider " + name)
	}
	providers[name] = provider
}

//sessionID -- make an ID as a 32 byte random number
func (manager *SessionManager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//SessionStart -- get the session cookie (if it exists) or make a new sessionID,
//then return the session.
func (manager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}

//SessionEnd -- delete the session from the server, then delete the cookie.
func (manager *SessionManager) SessionEnd(w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {

	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		_ = manager.provider.SessionDestroy(sid)
	}
	cookie = &http.Cookie{Name: manager.cookieName, Value: "deleted", Path: "/", HttpOnly: true, Expires: time.Unix(0, 0)}
	http.SetCookie(w, cookie)
	return
}
