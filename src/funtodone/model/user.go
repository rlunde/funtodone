package model

import "golang.org/x/crypto/bcrypt"

/*User -- a placeholder for when we build in auth
So far, at least, we're thinking of just having collections tied to
a user, and only having tasks as part of collections. A user might move
a task from one collection to another, or copy a task from one to another,
but it would always be the same user.*/
type User struct {
	FirstName    string `json:"firstname"`
	LastName     string `json:"lastname"`
	Email        string `json:"email"`
	PasswordHash []byte `json:"-"`
	Password     string `json:"password,omitempty"`
	//TODO: indicate whether to use password hash or OAuth2, and if OAuth2, which provider
}

func clear(b []byte) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

//Crypt use bcrypt to create the password hash
func Crypt(password []byte) ([]byte, error) {
	defer clear(password)
	return bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
}
