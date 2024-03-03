package auth

import "github.com/go-chi/chi/v5"

func RouteConfig(router chi.Router) {
	controller := &AuthController{}

	controller.Config()

	router.Get("/index", controller.Index)
	router.Get("/login", controller.Login)
	router.Get("/callback", controller.Callback)
	router.Get("/logout", controller.Logout)
}
