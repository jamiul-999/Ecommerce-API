package handlers

import (
	"log"
	"net/http"
)

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println("This is the handler printed in between")
}
