package service

import ( // plugin package
	// register authboss register module
	"encoding/base64"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin/render"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/justinas/nosurf"
	authboss "gopkg.in/authboss.v1"
	"gopkg.in/gin-gonic/gin.v1"
	// register authboss login module
	// to lock user after N authentication failures
)

// //User is the user struct needed by authboss
// type User struct {
// 	ID   int
// 	Name string
//
// 	// Auth
// 	Email    string
// 	Password string
//
// 	// OAuth2
// 	Oauth2Uid      string
// 	Oauth2Provider string
// 	Oauth2Token    string
// 	Oauth2Refresh  string
// 	Oauth2Expiry   time.Time
//
// 	// Confirm
// 	ConfirmToken string
// 	Confirmed    bool
//
// 	// Lock
// 	AttemptNumber int64
// 	AttemptTime   time.Time
// 	Locked        time.Time
//
// 	// Recover
// 	RecoverToken       string
// 	RecoverTokenExpiry time.Time
// }

var funcs = template.FuncMap{
	"formatDate": func(date time.Time) string {
		return date.Format("2017/01/02 12:00pm")
	},
	"yield": func() string { return "" },
}

//AuBo -- global pointer
var AuBo *authboss.Authboss

func layoutData(w http.ResponseWriter, r *http.Request) authboss.HTMLData {
	currentUserName := ""
	userInter, err := AuBo.CurrentUser(w, r)
	if userInter != nil && err == nil {
		currentUserName = userInter.(*User).Name
	}

	return authboss.HTMLData{
		"loggedin":               userInter != nil,
		"username":               "username",
		authboss.FlashSuccessKey: AuBo.FlashSuccess(w, r),
		authboss.FlashErrorKey:   AuBo.FlashError(w, r),
		"current_user_name":      currentUserName,
	}
}

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
			MaxLength:       40,
			AllowWhitespace: false,
		},
	}
}

func initAuthBossLayout(ab *authboss.Authboss, r *gin.Engine) {
	if os.Getenv(gin.ENV_GIN_MODE) == gin.ReleaseMode {
		ab.Layout = r.HTMLRender.(render.HTMLProduction).Template.Funcs(funcs).Lookup("authboss.tmpl")
	} else {
		html := r.HTMLRender.(render.HTMLDebug).Instance("authboss.tmpl", nil).(render.HTML)
		ab.Layout = html.Template.Funcs(template.FuncMap(funcs)).Lookup("authboss.tmpl")
		// ab.Layout.ExecuteTemplate(os.Stdout, "layout.html.tpl", nil)
	}
}

var database = NewMemStorer()

func initAuthBossParam(r *gin.Engine) *authboss.Authboss {
	ab := authboss.New()
	ab.Storer = database
	ab.CookieStoreMaker = NewCookieStorer
	ab.SessionStoreMaker = NewSessionStorer
	ab.ViewsPath = filepath.Join("ab_views")
	//ab.RootURL = `http://localhost:5567`

	ab.LayoutDataMaker = layoutData

	ab.MountPath = "/auth"
	ab.LogWriter = os.Stdout

	ab.XSRFName = "csrf_token"
	ab.XSRFMaker = func(_ http.ResponseWriter, r *http.Request) string {
		return nosurf.Token(r)
	}

	initAuthBossLayout(ab, r)
	ab.Mailer = authboss.LogMailer(os.Stdout)
	initAuthBossPolicy(ab)

	if err := ab.Init(); err != nil {
		// Handle error, don't let program continue to run
		log.Fatalln(err)
	}
	return ab
}

func initAuthBossRoute(r *gin.Engine) {
	cookieStoreKey, _ := base64.StdEncoding.DecodeString(`NpEPi8pEjKVjLGJ6kYCS+VTCzi6BUuDzU0wrwXyf5uDPArtlofn2AG6aTMiPmN3C909rsEWMNqJqhIVPGP3Exg==`)
	sessionStoreKey, _ := base64.StdEncoding.DecodeString(`AbfYwmmt8UCwUuhd9qvfNA9UCuN1cVcKJN1ofbiky6xCyyBj20whe40rJa3Su0WOWLWcPpO1taqJdsEI/65+JA==`)
	cookieStore = securecookie.New(cookieStoreKey, nil)
	sessionStore = sessions.NewCookieStore(sessionStoreKey)
	AuBo = initAuthBossParam(r)
	r.Any("/auth/*w", gin.WrapH(AuBo.NewRouter()))
}
