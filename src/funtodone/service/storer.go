package service

import (
	"fmt"
	"time"

	authboss "gopkg.in/authboss.v1"

	"github.com/davecgh/go-spew/spew"
)

var nextUserID int

//User is the user struct needed by authboss
type User struct {
	ID   int
	Name string

	// Auth
	Email    string
	Password string

	// OAuth2
	Oauth2Uid      string
	Oauth2Provider string
	Oauth2Token    string
	Oauth2Refresh  string
	Oauth2Expiry   time.Time

	// Confirm
	ConfirmToken string
	Confirmed    bool

	// Lock
	AttemptNumber int64
	AttemptTime   time.Time
	Locked        time.Time

	// Recover
	RecoverToken       string
	RecoverTokenExpiry time.Time

	// Remember is in another table
}

//MemStorer is a temp storer struct to use until I get mongodb version working
type MemStorer struct {
	Users  map[string]User
	Tokens map[string][]string
}

//NewMemStorer creates a dummy MemStorer struct
func NewMemStorer() *MemStorer {
	return &MemStorer{
		Users: map[string]User{
			"zeratul@heroes.com": User{
				ID:        1,
				Name:      "Zeratul",
				Password:  "$2a$10$XtW/BrS5HeYIuOCXYe8DFuInetDMdaarMUJEOg/VA/JAIDgw3l4aG", // pass = 1234
				Email:     "zeratul@heroes.com",
				Confirmed: true,
			},
		},
		Tokens: make(map[string][]string),
	}
}

//Create adds a new user and an attribute to the MemStorer
func (s MemStorer) Create(key string, attr authboss.Attributes) error {
	var user User
	if err := attr.Bind(&user, true); err != nil {
		return err
	}

	user.ID = nextUserID
	nextUserID++

	s.Users[key] = user
	fmt.Println("Create")
	spew.Dump(s.Users)
	return nil
}

//Put adds an attribute to the MemStorer
func (s MemStorer) Put(key string, attr authboss.Attributes) error {
	return s.Create(key, attr)
}

//Get returns an attribute from the MemStorer
func (s MemStorer) Get(key string) (result interface{}, err error) {
	user, ok := s.Users[key]
	if !ok {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

//PutOAuth adds an OAuth attribute to the MemStorer
func (s MemStorer) PutOAuth(uid, provider string, attr authboss.Attributes) error {
	return s.Create(uid+provider, attr)
}

//GetOAuth returns an OAuth attribute from the MemStorer
func (s MemStorer) GetOAuth(uid, provider string) (result interface{}, err error) {
	user, ok := s.Users[uid+provider]
	if !ok {
		return nil, authboss.ErrUserNotFound
	}

	return &user, nil
}

//AddToken appends a key/value to the tokens in a Storer
func (s MemStorer) AddToken(key, token string) error {
	s.Tokens[key] = append(s.Tokens[key], token)
	fmt.Println("AddToken")
	spew.Dump(s.Tokens)
	return nil
}

//DelTokens deletes all tokens with the given key
//it says it should return an error, but it doesn't
func (s MemStorer) DelTokens(key string) error {
	delete(s.Tokens, key)
	fmt.Println("DelTokens")
	spew.Dump(s.Tokens)
	return nil
}

//UseToken looks for a token in the tokens store and removes it.
//If there isn't one, it's an error.
func (s MemStorer) UseToken(givenKey, token string) error {
	toks, ok := s.Tokens[givenKey]
	if !ok {
		return authboss.ErrTokenNotFound
	}

	for i, tok := range toks {
		if tok == token {
			toks[i], toks[len(toks)-1] = toks[len(toks)-1], toks[i]
			s.Tokens[givenKey] = toks[:len(toks)-1]
			return nil
		}
	}

	return authboss.ErrTokenNotFound
}

//ConfirmUser looks for and returns a user with the given token
func (s MemStorer) ConfirmUser(tok string) (result interface{}, err error) {
	fmt.Println("==============", tok)

	for _, u := range s.Users {
		if u.ConfirmToken == tok {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}

//RecoverUser looks for and returns the user with the given recoverToken
func (s MemStorer) RecoverUser(rec string) (result interface{}, err error) {
	for _, u := range s.Users {
		if u.RecoverToken == rec {
			return &u, nil
		}
	}

	return nil, authboss.ErrUserNotFound
}
