package server

import (
	"fmt"
	"net/http"

	"github.com/inhies/newznab"
)

// Handles all API requests. Should be registered to /api as per convention.
func APIhandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if len(query["t"]) > 1 {
		//fmt.Fprintf(w, errorString, 201, "Incorrect parameter")
		fmt.Println(newznab.ErrBadParam.String())
		sendError(w, newznab.ErrBadParam)
		return
	}
	switch query["t"][0] {
	case "search":
		searchHandler(w, r, query)
	case "caps", "register", "tvsearch", "movie", "music", "book",
		"details", "getnfo", "get", "cartadd", "cartdel", "comments",
		"commentadd", "user":
		fmt.Println(newznab.ErrFuncUnavail.String())
		sendError(w, newznab.ErrFuncUnavail)
		//fmt.Fprintf(w, errorString, 203, "Function not available")
	default:
		//fmt.Fprintf(w, errorString, 202, "No such function")
		fmt.Println(newznab.ErrNoSuchFunc.String())
		sendError(w, newznab.ErrNoSuchFunc)
	}
	return
}

func sendError(w http.ResponseWriter, e newznab.Error) {
	output, err := e.AsXML()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(output)
}
