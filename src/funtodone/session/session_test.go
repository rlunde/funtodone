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
	// add something to the map
	session.Set("email", "al@pa.ca")
	email := session.Get("email")
	if email != "al@pa.ca" {
		t.Errorf("session should have email: %s but has email: %s", "al@pa.ca", email)
	}
	err = mgr.SessionUpdate(session)
	if err != nil {
		t.Errorf("SessionUpdate failed: %s\n", err.Error())
	}
	returnedSession := Session{SessionID: session.SessionID}
	err = mgr.SessionRead(&returnedSession)
	if err != nil {
		t.Errorf("SessionRead failed: %s\n", err.Error())
	}
	if session.SessionID != returnedSession.SessionID {
		t.Errorf("SessionInit has SessionID: %s but returnedSession has SessionID: %s", session.SessionID, returnedSession.SessionID)
	}
	returnedEmail := returnedSession.Get("email")
	if returnedEmail != email {
		t.Errorf("loaded session has email: %s but should have email: %s", returnedEmail, email)
	}
}
