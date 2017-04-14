package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/securecookie"
	"gopkg.in/authboss.v1"
)

var cookieStore *securecookie.SecureCookie

//CookieStorer is used by authboss to save authorization in cookies
type CookieStorer struct {
	w http.ResponseWriter
	r *http.Request
}

//NewCookieStorer creates a CookieStorer
func NewCookieStorer(w http.ResponseWriter, r *http.Request) authboss.ClientStorer {
	return &CookieStorer{w, r}
}

//Get returns the authorization from a cookie
func (s CookieStorer) Get(key string) (string, bool) {
	cookie, err := s.r.Cookie(key)
	if err != nil {
		return "", false
	}

	var value string
	err = cookieStore.Decode(key, cookie.Value, &value)
	if err != nil {
		return "", false
	}

	return value, true
}

//Put stores the authorization in a cookie
func (s CookieStorer) Put(key, value string) {
	encoded, err := cookieStore.Encode(key, value)
	if err != nil {
		fmt.Println(err)
	}

	cookie := &http.Cookie{
		Expires: time.Now().UTC().AddDate(1, 0, 0),
		Name:    key,
		Value:   encoded,
		Path:    "/",
	}
	http.SetCookie(s.w, cookie)
}

//Del deletes a cookie containing an authorization
func (s CookieStorer) Del(key string) {
	cookie := &http.Cookie{
		MaxAge: -1,
		Name:   key,
		Path:   "/",
	}
	http.SetCookie(s.w, cookie)
}
