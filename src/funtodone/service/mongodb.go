package service

import (
	"fmt"
	"time"

	mgo "gopkg.in/mgo.v2"
)

type sessionConfig struct {
	mongoSession        *mgo.Session
	mongoCollection     *mgo.Collection
	mongoHost           string
	mongoDatabase       string
	mongoCollectionName string
}

//SessionConfig -- global config for session manager configuration
var SessionConfig sessionConfig

//GetDatabaseConnection - open a Mongo database for storing sessions
func GetDatabaseConnection() (*sessionConfig, error) {
	//TODO: read the host and database from a config file
	SessionConfig = sessionConfig{
		mongoHost:           "127.0.0.1",
		mongoDatabase:       "funtodone",
		mongoCollectionName: "sessions",
	}
	mongoSession, err := mgo.Dial(SessionConfig.mongoHost)
	if err != nil {
		fmt.Printf("Could not open mongo database session: %s", err.Error())
		return nil, err
	}
	SessionConfig.mongoSession = mongoSession
	SessionConfig.mongoSession.SetMode(mgo.Monotonic, true)

	mongoCollection := SessionConfig.mongoSession.DB(SessionConfig.mongoDatabase).C(SessionConfig.mongoCollectionName)
	SessionConfig.mongoCollection = mongoCollection
	return &SessionConfig, nil
}

//var pder = &Provider{list: list.New()}

//SessionStore - for each session, keep a map of name/value pairs and a last accessed time
type SessionStore struct {
	sid          string                      // unique session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session value stored inside
}

//Set - save a key/value pair in a session in the database
func (st *SessionStore) Set(mgr *SessionManager, key, value interface{}) error {
	st.value[key] = value
	SessionUpdate(mgr, st.sid)
	return nil
}

//Get - retrieve a value for a key, for a session, from the database
func (st *SessionStore) Get(mgr *SessionManager, key interface{}) interface{} {
	SessionUpdate(mgr, st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	//return nil
}

func (st *SessionStore) Delete(mgr *SessionManager, key interface{}) error {
	delete(st.value, key)
	SessionUpdate(mgr, st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

// type Provider struct {
// 	lock     sync.Mutex               // lock
// 	sessions map[string]*list.Element // save in memory // TODO: replace with mongodb
// 	list     *list.List               // gc
// }

// func (pder *Provider) SessionInit(sid string) (Session, error) {
//TODO: rethink this whole thing
func SessionInit(mgr *SessionManager, sid string) (Session, error) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// v := make(map[interface{}]interface{}, 0)

	//TODO: store the session in mongodb
	// element := pder.list.PushBack(newsess)
	// pder.sessions[sid] = element // TODO: store in mongodb
	return nil, nil // TODO: return a session
}

// func (pder *Provider) SessionRead(sid string) (Session, error) {
func SessionRead(mgr *SessionManager, sid string) (Session, error) {
	//TODO: retrieve the session from mongodb if it's there
	element := mgr.store.Get(mgr, sid) // TODO: change this to FindObjectByID
	if element != nil {
		return element.(Session), nil
	} else {
		sess, err := SessionInit(mgr, sid)
		return sess, err
	}
	//return nil, nil
}

func SessionDestroy(mgr *SessionManager, sid string) error {
	// TODO: delete from mongodb
	// if element, ok := pder.sessions[sid]; ok {
	// 	delete(pder.sessions, sid)
	// 	pder.list.Remove(element)
	// 	return nil
	// }
	return nil
}

func SessionGC(mgr SessionManager, maxlifetime int64) {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// TODO: delete from mongodb
	// for {
	// 	element := pder.list.Back()
	// 	if element == nil {
	// 		break
	// 	}
	// 	if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
	// 		pder.list.Remove(element)
	// 		delete(pder.sessions, element.Value.(*SessionStore).sid)
	// 	} else {
	// 		break
	// 	}
	// }
}

func SessionUpdate(mgr *SessionManager, sid string) error {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// TODO: update in mongodb
	// if element, ok := pder.sessions[sid]; ok {
	// 	element.Value.(*SessionStore).timeAccessed = time.Now()
	// 	pder.list.MoveToFront(element)
	// 	return nil
	// }
	return nil
}

func init() {
	//pder.sessions = make(map[string]*list.Element, 0) // TODO: connect to mongodb
	//Register("mongodb", pder)
}
