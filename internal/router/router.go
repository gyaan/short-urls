package router

import (
	"github.com/go-chi/chi"
	"github.com/gyaan/short-urls/internal/handler"
	"github.com/gyaan/short-urls/internal/middleware"
)

type router struct {
	sh handler.ShortUrlHandler
	uh handler.UserHandler
	ah handler.AuthenticationHandler
	r *chi.Mux
}

func RegisterRoutes(sh handler.ShortUrlHandler, uh handler.UserHandler, ah handler.AuthenticationHandler, r *chi.Mux) {
	routes := router{sh: sh, uh:uh, ah:ah, r: r}

	//home router
	routes.r.Get("/", routes.sh.HomeHandler)

	//user
	routes.r.Post("/register", routes.uh.RegisterUser)

	//protect user routes
	userRouter := chi.NewRouter()
	userRouter.Use(middleware.Authenticate)
	userRouter.Put("/{user_id}", routes.uh.UpdateUser)
	userRouter.Get("/", routes.uh.GetUser)

	//authentication
	routes.r.Post("/access-token", routes.ah.GetAccessToken)

	//short url manipulation
	shortUrlRouter := chi.NewRouter()
	shortUrlRouter.Use(middleware.Authenticate)
	shortUrlRouter.Post("/", routes.sh.CreateShortUrl)
	shortUrlRouter.Get("/{short_url_id}", routes.sh.GetAShortUrl)
	shortUrlRouter.Get("/", routes.sh.GetAllShortUrl)
	shortUrlRouter.Put("/{short_url_id}", routes.sh.UpdateShortUrl)
	shortUrlRouter.Delete("/{short_url_id}", routes.sh.DeleteShortUrl)

	//short url redirection handler
	routes.r.Get("/{short_url}", routes.sh.RedirectToActualUrl)

	//add short url auth protect routs
	routes.r.Mount("/short-urls", shortUrlRouter)
	routes.r.Mount("/users", userRouter)
}
