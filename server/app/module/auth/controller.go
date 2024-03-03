package auth

import (
	"html/template"
	"net/http"

	"github.com/markbates/goth/gothic"
	"github.com/mrspec7er/balky/app/utils"
)

type AuthController struct {
	service  AuthService
	response utils.Response
}

func (c *AuthController) Index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, false)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	gothic.BeginAuthHandler(w, r)
}

func (c *AuthController) Callback(w http.ResponseWriter, r *http.Request) {
	user, err := c.service.SaveUserSessions(w, r)
	if err != nil {
		c.response.UnauthorizeUser(w)
		return
	}
	t, _ := template.ParseFiles("templates/success.html")
	t.Execute(w, user)
}

func (c *AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	c.service.RemoveUserSessions(w, r)
	w.Header().Set("Location", "/api/auth/index")
	w.WriteHeader(http.StatusTemporaryRedirect)
}
