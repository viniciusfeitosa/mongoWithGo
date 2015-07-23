package middlewares

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httputil"

	"github.com/viniciusfeitosa/mongo/models"
)

// CmaMiddleware is a middleware to cma
func CmaMiddleware(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" && r.URL.Path == "/ping/cma/create" {
			dumpHead, _ := httputil.DumpRequest(r, false)
			dumpBody, _ := httputil.DumpRequest(r, true)
			dumpBody = dumpBody[len(dumpHead):]

			cma := models.CmaJSON{}
			if err := json.NewDecoder(bytes.NewReader(dumpBody)).Decode(&cma); err != nil {
				log.Println(err)
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if cma.Identifier != "publicacao" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusBadGateway)
				return
			}
		}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
