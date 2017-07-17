package service

import (
	"container/list"
	"fmt"
	"sync"
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

var pder = &Provider{list: list.New()}

//SessionStore - for each session, keep a map of name/value pairs and a last accessed time
type SessionStore struct {
	sid          string                      // unique session id
	timeAccessed time.Time                   // last access time
	value        map[interface{}]interface{} // session value stored inside
}

//Set - save a key/value pair in a session in the database
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

//Get - retrieve a value for a key, for a session, from the database
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
	//return nil
}

func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

func (st *SessionStore) SessionID() string {
	return st.sid
}

type Provider struct {
	lock     sync.Mutex               // lock
	sessions map[string]*list.Element // save in memory // TODO: replace with mongodb
	list     *list.List               // gc
}

func (pder *Provider) SessionInit(sid string) (Session, error) {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := pder.list.PushBack(newsess)
	pder.sessions[sid] = element // TODO: store in mongodb
	return newsess, nil
}

func (pder *Provider) SessionRead(sid string) (Session, error) {
	if element, ok := pder.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil // TODO: retrieve from mongodb
	} else {
		sess, err := pder.SessionInit(sid)
		return sess, err
	}
	//return nil, nil
}

func (pder *Provider) SessionDestroy(sid string) error {
	if element, ok := pder.sessions[sid]; ok {
		delete(pder.sessions, sid) // TODO: delete from mongodb
		pder.list.Remove(element)
		return nil
	}
	return nil
}

func (pder *Provider) SessionGC(maxlifetime int64) {
	pder.lock.Lock()
	defer pder.lock.Unlock()

	for {
		element := pder.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxlifetime) < time.Now().Unix() {
			pder.list.Remove(element) // TODO: delete from mongodb
			delete(pder.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

func (pder *Provider) SessionUpdate(sid string) error {
	pder.lock.Lock()
	defer pder.lock.Unlock()
	if element, ok := pder.sessions[sid]; ok { // TODO: update in mongodb
		element.Value.(*SessionStore).timeAccessed = time.Now()
		pder.list.MoveToFront(element)
		return nil
	}
	return nil
}

func init() {
	pder.sessions = make(map[string]*list.Element, 0) // TODO: connect to mongodb
	Register("mongodb", pder)
}
