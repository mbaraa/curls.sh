package main

import (
	_ "embed"
	"encoding/json"
	"log"
	"net/http"

	"github.com/mbaraa/curls.sh/scripts"
)

func main() {
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		resp := make(map[string]string)
		scripts := scripts.All()
		for _, script := range scripts {
			resp[script.Id()] = script.Description()
		}

		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(resp)
	})

	http.HandleFunc("GET /{script_id}", func(w http.ResponseWriter, r *http.Request) {
		scripts := scripts.AllById()
		script, ok := scripts[r.PathValue("script_id")]
		if !ok {
			w.Write([]byte("refer to https://github.com/mbaraa/curls.sh for more info, or simply check the index endpoint curls.sh/"))
			w.WriteHeader(http.StatusNotFound)
			return
		}

		args := r.URL.Query()["args"]

		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(script.RunString(args)))
	})

	log.Println("Starting http server at port 3000")
	log.Fatalln(http.ListenAndServe(":3000", nil))
}
