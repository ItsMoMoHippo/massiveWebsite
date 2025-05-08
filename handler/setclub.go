package handler

import (
	"fmt"
	"maWeb/state"
	"net/http"
)

func SetClubFoo(w http.ResponseWriter, r *http.Request) {
	state.SetClub("foo")
	fmt.Println("Set club to foo")
	w.WriteHeader(http.StatusNoContent)
}
