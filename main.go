package main

import (
	"fmt"
	"net/http"
)

// define the server handler function : 

func basicHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Go Server under runing now ..."))
}



func  main()  {
	// setup an http server :
	server := &http.Server{
		// define the http server port :
		Addr: ":3000",
		// define the http server handler :
		Handler: http.HandlerFunc(basicHandler),
	}
	// run the server:
	err := server.ListenAndServe()
	// check if there was an error :
	if err != nil {
		fmt.Println("Failed to Runing Server: ", err)
	}
}


