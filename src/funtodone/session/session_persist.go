package session

import (
	"errors"
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//TODO: move database refs out of config struct to its own struct

/* Note: see article on escape analysis: references are passed into functions, not out.
http://www.agardner.me/golang/garbage/collection/gc/escape/analysis/2015/10/18/go-escape-analysis.html
*/

//Config -- keep track of all the config data
type Config struct {
	mongoSession        *mgo.Session
	mongoCollection     *mgo.Collection
	mongoHost           string
	mongoDatabase       string
	mongoCollectionName string
}

//DbConn -- return the database connection
func (manager *Manager) DbConn() *mgo.Collection {
	return manager.sessionConfig.mongoCollection
}

//GetSessionConfig -- return the config data for the session
//TODO: read the host and database from a config file
func GetSessionConfig(mgr *Manager) {
	sessionConfig := Config{
		mongoHost:           "127.0.0.1",
		mongoDatabase:       "test", // change to funtodone
		mongoCollectionName: "sessions",
	}
	mgr.sessionConfig = sessionConfig
}

//GetDatabaseConnection - open a Mongo database for storing sessions
func GetDatabaseConnection(mgr *Manager) (err error) {
	//Note: pass globalSessionMgr as the argument to this function
	if mgr == nil {
		err = errors.New("GetDatabaseConnection called with nil Manager")
		return
	}
	mongoSession, err := mgo.Dial(mgr.sessionConfig.mongoHost)
	if err != nil {
		fmt.Printf("Could not open mongo database session: %s", err.Error())
		return err
	}
	mgr.sessionConfig.mongoSession = mongoSession
	mgr.sessionConfig.mongoSession.SetMode(mgo.Monotonic, true)
	// Error check on every access
	mgr.sessionConfig.mongoSession.SetSafe(&mgo.Safe{})

	mongoCollection := mgr.sessionConfig.mongoSession.DB(mgr.sessionConfig.mongoDatabase).C(mgr.sessionConfig.mongoCollectionName)
	mgr.sessionConfig.mongoCollection = mongoCollection
	return
}

//SessionInit - create a new session record in MongoDB
func (mgr *Manager) SessionInit(session *Session) (err error) {
	err = mgr.checkMgrAndSession(session, "SessionInit")
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// session.ID = bson.NewObjectId()
	//c := mgr.sessionConfig.mongoCollection // TODO: figure out why this doesn't work
	c := mgr.sessionConfig.mongoSession.DB("test").C("sessions")
	err = c.Insert(session)
	if err != nil {
		return err
	}
	return nil
}

func (mgr *Manager) checkMgrAndSession(session *Session, fn string) (err error) {
	if session == nil {
		err = errors.New(fn + " called with nil Session")
	}
	if mgr.sessionConfig.mongoSession == nil {
		err = errors.New(fn + " called with nil Manager mongoSession")
	}
	return
}

//SessionRead -- get the session out of mongodb
func (mgr *Manager) SessionRead(session *Session) (err error) {
	err = mgr.checkMgrAndSession(session, "SessionRead")
	if err != nil {
		return err
	}
	// id := bson.ObjectIdHex(session.sessionID)
	//c := mgr.sessionConfig.mongoCollection
	c := mgr.sessionConfig.mongoSession.DB("test").C("sessions")

	err = c.Find(bson.M{"sessionid": session.SessionID}).One(session)
	return err // err is nil if it found it
}

//SessionDestroy -- delete a session record from mongodb
func (mgr *Manager) SessionDestroy(session *Session) (err error) {
	err = mgr.checkMgrAndSession(session, "SessionDestroy")
	if err != nil {
		return err
	}
	c := mgr.sessionConfig.mongoSession.DB("test").C("sessions")

	//c := mgr.sessionConfig.mongoCollection
	err = c.Remove(bson.M{"sessionID": session.SessionID})
	return err // err is nil if it found it
}

//SessionUpdate -- update a session in mongodb, and update the last access time
func (mgr *Manager) SessionUpdate(session *Session) (err error) {
	err = mgr.checkMgrAndSession(session, "SessionUpdate")
	if err != nil {
		return err
	}
	c := mgr.sessionConfig.mongoSession.DB("test").C("sessions")

	//c := mgr.sessionConfig.mongoCollection
	err = c.Update(bson.M{"sessionID": session.SessionID}, session)
	return err // err is nil if it found it
}

func init() {

}
