package main

import (
	"fmt"
	"net/http"
"log"
)

type User struct{

}
type Account struct{

}

func handleGetUser(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/user" {
		http.Error(w, "Not found.", 404)
		return 
	}
	if r.Method != "GET" {
		http.Error(w, "Request method not supported.", 404)
        return
	}
	fmt.Fprintf(w, "handling get user function.")
}

func main(){
	//listen here;
	serveLandingPage := http.FileServer(http.Dir("./public"))
	http.Handle("/",serveLandingPage)
	http.HandleFunc("/user", handleGetUser)

	
    fmt.Println("Listening on port 8080")
	if error := http.ListenAndServe(":8080", nil);
	error != nil {
		log.Fatal(error)

	}
}