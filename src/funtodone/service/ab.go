package service

import ( // plugin package
	// register authboss register module
	authboss "gopkg.in/authboss.v1"
	// register authboss login module
	// to lock user after N authentication failures
)

func initAuthBossPolicy(ab *authboss.Authboss) {
	ab.Policies = []authboss.Validator{
		authboss.Rules{
			FieldName:       "email",
			Required:        true,
			AllowWhitespace: false,
		},
		authboss.Rules{
			FieldName:       "password",
			Required:        true,
			MinLength:       4,
			MaxLength:       8,
			AllowWhitespace: false,
		},
	}
}
