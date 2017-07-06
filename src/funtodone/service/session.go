package service

import (
	"fmt"
	"sync"
)

/*
 * https://astaxie.gitbooks.io/build-web-application-with-golang/en/06.2.html
 */
//Session -- keep track of web session
type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}
type SessionProvider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}

var provides = make(map[string]SessionProvider)

type SessionManager struct {
	cookieName  string
	lock        sync.Mutex // protects session
	provider    SessionProvider
	maxlifetime int64
}

func NewManager(providerName, cookieName string, maxlifetime int64) (*SessionManager, error) {
	provider, ok := provides[providerName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", providerName)
	}
	return &SessionManager{provider: provider, cookieName: cookieName, maxlifetime: maxlifetime}, nil
}
