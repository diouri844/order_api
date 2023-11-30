package main

import (
	"fmt"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// define the server handler function : 

func basicHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Go Server under runing now ..."))
}



func  main()  {
	// seetup new router by chi :
	router := chi.NewRouter()
	// setup logger middleware for the router:
	router.Use(middleware.Logger)
	// use defined router to setup endpoint:
	router.Get("/Ping", basicHandler)
	// setup an http server :
	server := &http.Server{
		// define the http server port :
		Addr: ":3000",
		// define the http server handler :
		Handler: router,
	}
	// run the server:
	err := server.ListenAndServe()
	// check if there was an error :
	if err != nil {
		fmt.Println("Failed to Runing Server: ", err)
	}
}


