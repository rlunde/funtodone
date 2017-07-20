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

var globalSessionManager *SessionManager

//  initialize the session manager (init is run automatically)
func init() {
	var err error
	globalSessionManager = &SessionManager{
		cookieName:  "gosessionid",
		maxlifetime: 3600,
	}
	if err != nil {
		fmt.Printf("Error creating session manager: %s", err.Error())
	}
}

//Session -- keep track of web session
type Session struct {
	sessionID  string
	lastAccess int64 // unix time of last access
}

//Set -- store a value of any type in a session
func (session *Session) Set(key, value interface{}) error {
	return nil
}

//Get -- get a value of any type from a session
func (session *Session) Get(key interface{}) interface{} {
	return nil
}

//Delete -- delete a key/value pair from a session
func (session *Session) Delete(key, value interface{}) error {
	return nil
}

//NewSession -- create a session and return a reference to it
func NewSession() (session *Session) {
	return nil
}

//SessionManager -- overly abstract for what we want -- remove this
type SessionManager struct {
	cookieName    string
	lock          sync.Mutex // protects session
	maxlifetime   int64
	sessionConfig SessionConfig
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
func (manager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session *Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session, _ = SessionInit(manager, sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = SessionRead(manager, sid)
	}
	return
}

//SessionEnd -- delete the session from the server, then delete the cookie.
func (manager *SessionManager) SessionEnd(mgr *SessionManager, w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {

	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		_ = SessionDestroy(mgr, sid)
	}
	cookie = &http.Cookie{Name: manager.cookieName, Value: "deleted", Path: "/", HttpOnly: true, Expires: time.Unix(0, 0)}
	http.SetCookie(w, cookie)
	return
}
