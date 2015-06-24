package handlers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/zenazn/goji/web"
)

type notification struct {
	CollectionType string `json:"collectionType"`
	Date           string `json:"date"`
	OwnerId        string `json:"ownerId"`
	OwnerType      string `json:"ownerType"`
	SubscriberId   string `json:"subscriberId"`
}

func Subscriber(c web.C, w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	r.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	var n []notification
	if err := json.Unmarshal(body, &n); err != nil {
		panic(err)
	}

	log.Printf("Incoming notification: %+v", n)
	w.WriteHeader(http.StatusNoContent)
}
