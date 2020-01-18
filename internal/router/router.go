package router

import (
	"github.com/go-chi/chi"
	"github.com/gyaan/short-urls/internal/handler"
	"github.com/gyaan/short-urls/internal/middleware"
)

type router struct {
	h handler.Handler
	r *chi.Mux
}

func RegisterRoutes(h handler.Handler, r *chi.Mux) {
	routes := router{h: h, r: r}

	//home router
	routes.r.Get("/", routes.h.HomeHandler)

	//user
	routes.r.Post("/register", routes.h.RegisterUser)

	//protect user routes
	userRouter := chi.NewRouter()
	userRouter.Use(middleware.Authenticate)
	userRouter.Put("/{user_id}", routes.h.UpdateUser)
	userRouter.Get("/", routes.h.GetUser)

	//authentication
	routes.r.Post("/access-token", routes.h.GetAccessToken)

	//short url manipulation
	shortUrlRouter := chi.NewRouter()
	shortUrlRouter.Use(middleware.Authenticate)
	shortUrlRouter.Post("/", routes.h.CreateShortUrl)
	shortUrlRouter.Get("/{short_url_id}", routes.h.GetAShortUrl)
	shortUrlRouter.Get("/", routes.h.GetAllShortUrl)
	shortUrlRouter.Put("/{short_url_id}", routes.h.UpdateShortUrl)
	shortUrlRouter.Delete("/{short_url_id}", routes.h.DeleteShortUrl)

	//short url redirection handler
	routes.r.Get("/{short_url}", routes.h.RedirectToActualUrl)

	//add short url auth protect routs
	routes.r.Mount("/short-urls", shortUrlRouter)
	routes.r.Mount("/users", userRouter)
}
