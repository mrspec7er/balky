package auth

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/google"
)

type AuthController struct{}

var Store *sessions.CookieStore

func (*AuthController) Config() {

	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	maxAge := 86400 * 1 // 1 days
	isProd := false     // Set to true when serving over https

	Store.MaxAge(maxAge)
	Store.Options.Path = "/"
	Store.Options.HttpOnly = true
	Store.Options.Secure = isProd

	gothic.Store = Store

	goth.UseProviders(
		google.New(os.Getenv("GOOGLE_AUTH_KEY"), os.Getenv("GOOGLE_AUTH_SECRET"), os.Getenv("API_URL")+"/auth/callback?provider=google"),
	)
}

func (c *AuthController) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, false)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (c *AuthController) Callback(w http.ResponseWriter, r *http.Request) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		fmt.Fprintln(w, err)
		return
	}
	session, _ := Store.Get(r, "auth")
	session.Values["email"] = user.Email
	session.Save(r, w)

	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(w, user)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := Store.Get(r, "auth")
	session.Values["email"] = nil
	session.Save(r, w)
	gothic.Logout(w, r)
	w.Header().Set("Location", "/api/auth/index")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
