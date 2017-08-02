package session

import (
	"testing"
)

var (
	DropDatabase = true
)

func SetupTestSession() *Session {
	mgr := GetMgr()
	idstr := mgr.sessionID()
	session := NewSession(idstr)
	return &session
}
func TestSessionInitAndRead(t *testing.T) {
	mgr := GetMgr()
	err := GetDatabaseConnection(mgr)
	session := SetupTestSession()
	err = mgr.SessionInit(session)
	if err != nil {
		t.Errorf("SessionInit failed: %s\n", err.Error())
	}
	returnedSession := Session{SessionID: session.SessionID}
	err = mgr.SessionRead(&returnedSession)
	if err != nil {
		t.Errorf("SessionRead failed: %s\n", err.Error())
	}
	// TODO: check contents
}
