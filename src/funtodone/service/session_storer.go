package service

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
	"gopkg.in/authboss.v1"
)

const sessionCookieName = "ab_blog"

var sessionStore *sessions.CookieStore

//SessionStorer stores authorization in a session cookie
type SessionStorer struct {
	w http.ResponseWriter
	r *http.Request
}

//NewSessionStorer returns an authboss ClientStorer
func NewSessionStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &SessionStorer{w, r}
}

//Get returns the value corresponding to a key and a boolean that says whether it was found
func (s SessionStorer) Get(key string) (string, bool) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	strInf, ok := session.Values[key]
	if !ok {
		return "", false
	}

	str, ok := strInf.(string)
	if !ok {
		return "", false
	}

	return str, true
}

//Put stores a value under a key
func (s SessionStorer) Put(key, value string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	session.Values[key] = value
	session.Save(s.r, s.w) // TODO: check error
}

//Del deletes a value if found, but doesn't return an error if it wasn't
func (s SessionStorer) Del(key string) {
	session, err := sessionStore.Get(s.r, sessionCookieName)
	if err != nil {
		fmt.Println(err)
		return
	}

	delete(session.Values, key)
	session.Save(s.r, s.w) // TODO: check error
}
