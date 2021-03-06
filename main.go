package main

import (
	"log"
	"net/http"

	"github.com/bradford-hamilton/explore-the-chi/api/routes"
	"github.com/bradford-hamilton/explore-the-chi/config"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

// Routes defines middleware, sets version namespace, mounts different routes
// and returns a pointer to a chi multiplexer
func Routes(dbConn *config.DBConfig) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		render.SetContentType(render.ContentTypeJSON), // set content-type headers as application/json
		middleware.Logger,          // log api request calls
		middleware.DefaultCompress, // compress results, mostly gzipping assets and json
		middleware.StripSlashes,    // strip slashes to no slash URL versions
		middleware.Recoverer,       // recover from panics without crashing server
	)

	router.Route("/api", func(r chi.Router) {
		r.Mount("/v1/tx", routes.Routes(dbConn))
	})

	return router
}

func main() {
	dbConn := config.NewDynamo()

	env, err := config.NewEnvConfig()
	if err != nil {
		log.Panicln("Configuration error", err)
	}

	router := Routes(dbConn)

	walkFunc := func(
		method string,
		route string,
		handler http.Handler,
		middlewares ...func(http.Handler) http.Handler,
	) error {
		log.Printf("%s %s\n", method, route) // walk and print all the routes
		return nil
	}

	if err := chi.Walk(router, walkFunc); err != nil {
		log.Panicf("Error Log: %s\n", err.Error()) // panic if there is an error
	}

	log.Println("Serving application on PORT :", env.PORT)
	log.Fatal(http.ListenAndServe(":"+env.PORT, router))
}
