package main

import (
	"fmt"
	"net/http"
)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/home",func (res http.ResponseWriter, req *http.Request){

		fmt.Fprint(res,"Welcome to home page!")
	})

	mux.HandleFunc("/home/accounts",func (res http.ResponseWriter, req *http.Request){

		fmt.Fprint(res,"Welcome to Accounts page!")
	})

	mux.HandleFunc("/home/subscription",func (res http.ResponseWriter, req *http.Request){

		fmt.Fprint(res,"Welcome to Subscription page!")
	})

	http.ListenAndServe(":9000",mux)
}