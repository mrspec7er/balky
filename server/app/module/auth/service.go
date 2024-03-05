package auth

import (
	"fmt"
	"net/http"

	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/mrspec7er/balky/app/utility"
)

type AuthService struct{}

func (a AuthService) GetUserEmail(r *http.Request) (string, error) {
	session, err := utility.Store.Get(r, "auth")
	if err != nil {
		return "", err
	}
	email, ok := session.Values["email"].(string)

	if !ok || email == "" {
		return "", fmt.Errorf("unauthorize user")
	}

	return email, nil
}

func (a AuthService) SaveUserSessions(w http.ResponseWriter, r *http.Request) (*goth.User, error) {
	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil {
		return nil, err
	}

	session, err := utility.Store.Get(r, "auth")
	if err != nil {
		return nil, err
	}
	session.Values["email"] = user.Email
	session.Save(r, w)

	return &user, nil
}

func (a AuthService) RemoveUserSessions(w http.ResponseWriter, r *http.Request) {
	session, _ := utility.Store.Get(r, "auth")
	session.Values["email"] = nil
	session.Save(r, w)
	gothic.Logout(w, r)
}
