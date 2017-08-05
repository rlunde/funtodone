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
	session := NewSession(mgr, idstr)
	return &session
}
func TestCreateReadDestroy(t *testing.T) {
	mgr := GetMgr()
	err := GetDatabaseConnection(mgr)
	session := SetupTestSession()
	err = Create(session)
	if err != nil {
		t.Errorf("Create failed: %s\n", err.Error())
	}
	// add something to the map
	session.Set("email", "al@pa.ca")
	email := session.Get("email")
	if email != "al@pa.ca" {
		t.Errorf("session should have email: %s but has email: %s", "al@pa.ca", email)
	}

	returnedSession := NewSession(session.Mgr, session.SessionID)
	err = Read(&returnedSession)
	if err != nil {
		t.Errorf("Read failed: %s\n", err.Error())
	}
	if session.SessionID != returnedSession.SessionID {
		t.Errorf("Create has SessionID: %s but returnedSession has SessionID: %s", session.SessionID, returnedSession.SessionID)
	}
	returnedEmail := returnedSession.Get("email")
	if returnedEmail != email {
		t.Errorf("loaded session has email: %s but should have email: %s", returnedEmail, email)
	}
	err = Destroy(session)
	if err != nil {
		t.Errorf("Destroy failed: %s\n", err.Error())
	}
	goneSession := NewSession(session.Mgr, session.SessionID)
	err = Read(&goneSession)
	if err == nil || err.Error() != "not found" {
		t.Errorf("Read found a deleted session: %s\n", session.SessionID)
	}
}
