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

	"gopkg.in/mgo.v2/bson"
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
	_id        bson.ObjectId // this ties it tightly to mongo -- better to have a wrapper struct
	sessionID  string
	lastAccess int64                       // unix time of last access
	m          map[interface{}]interface{} // holds a map of any key to any value
}

//Set -- store a value of any type in a session
func (session *Session) Set(key, value interface{}) error {
	session.m[key] = value
	return nil
}

//Get -- get a value of any type from a session
func (session *Session) Get(key interface{}) interface{} {
	return session.m[key]
}

//Delete -- delete a key/value pair from a session
func (session *Session) Delete(key, value interface{}) error {
	delete(session.m, key)
	return nil
}

//NewSession return a new session with the map and lastAccess initialized
func NewSession(sid string) (session Session) {
	session = Session{sessionID: sid,
		m:          make(map[interface{}]interface{}),
		lastAccess: time.Now().Unix()}
	return
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
func (manager *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)

	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session = NewSession(sid)
		err = SessionInit(manager, &session)
		if err != nil {
			return session, err
		}
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session = NewSession(sid)
		err = SessionRead(manager, &session)
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
		session.sessionID = sid
		_ = SessionDestroy(mgr, &session)
	}
	cookie = &http.Cookie{Name: manager.cookieName, Value: "deleted", Path: "/", HttpOnly: true, Expires: time.Unix(0, 0)}
	http.SetCookie(w, cookie)
	return
}
