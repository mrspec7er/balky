package auth

import (
	"context"
	"net/http"
	"slices"

	"github.com/mrspec7er/balky/app/module/user"
	"github.com/mrspec7er/balky/app/utility"
)

type AuthMiddleware struct {
	service  AuthService
	response utility.Response
	user     user.UserService
}

func (m AuthMiddleware) Authorize(roles ...string) func(http.Handler) http.Handler {
	return (func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			userEmail, err := m.service.GetUserEmail(r)

			if err != nil {
				m.response.UnauthorizeUser(w)
				return
			}

			user, status, err := m.user.FindOne(userEmail)
			if err != nil {
				m.response.InternalServerErrorHandler(w, status, err)
				return
			}

			if !slices.Contains(roles, user.Role) {
				m.response.UnauthorizeUser(w)
				return
			}

			ctx := context.WithValue(r.Context(), utility.UserContextKey, user)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	})
}
