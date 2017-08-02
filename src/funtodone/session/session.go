package session

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

var globalSessionMgr *Manager

//GetMgr - we only have a single global session manager -- do we need any?
func GetMgr() *Manager {
	return globalSessionMgr
}

//  initialize the session manager (init is run automatically)
func init() {
	var err error
	globalSessionMgr = &Manager{
		cookieName:  "gosessionid",
		maxlifetime: 3600,
	}
	if err != nil {
		fmt.Printf("Error creating session manager: %s", err.Error())
	}
}

//Session -- keep track of web session
type Session struct {
	// ID         bson.ObjectId               `json:"_id" bson:"_id,omitempty"`
	SessionID  string                 `json:"sessionid" bson:"sessionid"`
	LastAccess int64                  `json:"lastaccess" bson:"lastaccess"` // unix time of last access
	M          map[string]interface{} `json:"values" bson:"values"`         // holds a map of any key to any value
}

//Set -- store a value of any type in a session
func (session *Session) Set(key string, value interface{}) error {
	session.M[key] = value
	return nil
}

//Get -- get a value of any type from a session
func (session *Session) Get(key string) interface{} {
	return session.M[key]
}

//Delete -- delete a key/value pair from a session
func (session *Session) Delete(key string, value interface{}) error {
	delete(session.M, key) // do we need to return an error if it isn't there?
	return nil
}

//NewSession return a new session with the map and lastAccess initialized
func NewSession(sid string) (session Session) {
	session = Session{SessionID: sid,
		M:          make(map[string]interface{}),
		LastAccess: time.Now().Unix()}
	return
}

//Manager -- I don't know if we still need a lock?
type Manager struct {
	cookieName    string
	lock          sync.Mutex // protects session
	maxlifetime   int64
	sessionConfig Config
}

//sessionID -- make an ID as a 32 byte random number
func (manager *Manager) sessionID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

//SessionStart -- get the session cookie (if it exists) or make a new sessionID,
//then return the session.
func (manager *Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session Session, err error) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	// if this is the first session, open a database connection
	if manager.sessionConfig.mongoSession == nil {
		err = GetDatabaseConnection(manager)
		if err != nil {
			return Session{}, err
		}
	}
	cookie, err := r.Cookie(manager.cookieName)

	if err != nil || cookie.Value == "" {
		sid := manager.sessionID()
		session = NewSession(sid)
		err = manager.SessionInit(&session)
		if err != nil {
			return session, err
		}
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session = NewSession(sid)
		err = manager.SessionRead(&session)
	}
	return
}

//SessionEnd -- delete the session from the server, then delete the cookie.
func (manager *Manager) SessionEnd(mgr *Manager, w http.ResponseWriter, r *http.Request) (session Session) {
	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err == nil && cookie.Value != "" {
		sid, _ := url.QueryUnescape(cookie.Value)
		session.SessionID = sid
		_ = mgr.SessionDestroy(&session)
	}
	cookie = &http.Cookie{Name: manager.cookieName, Value: "deleted", Path: "/", HttpOnly: true, Expires: time.Unix(0, 0)}
	http.SetCookie(w, cookie)
	return
}
