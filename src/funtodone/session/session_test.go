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
	err = SessionInit(mgr, session)
	if err != nil {
		t.Errorf("SessionInit failed: %s\n", err.Error())
	}
	returnedSession := Session{SessionID: session.SessionID}
	err = SessionRead(mgr, &returnedSession)
	if err != nil {
		t.Errorf("SessionRead failed: %s\n", err.Error())
	}
	// TODO: check contents
}
