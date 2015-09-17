package server

import (
	"net/http"
	"net/url"

	"github.com/kr/pretty"
)

func searchHandler(w http.ResponseWriter, r *http.Request, q url.Values) {
	pretty.Print(q)
	return
}
