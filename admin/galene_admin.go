package main

import (
	"net/http"
	"log"
	"galene/admin/serverAdmin"
)


func main() {
	serverAdmin.Handle()

	log.Fatal(http.ListenAndServe(":8444", nil))

}
