package api

import (
	"net/http"
	"os"
	"path/filepath"
)

func APIhost() {
	filePath := filepath.Join(os.Getenv("HOME"), "webdev/basedlibrary/API/dbOUT.json")

	http.HandleFunc("/dbOUT.json", func(w http.ResponseWriter, r *http.Request) {
		// Set the content type to application/json
		w.Header().Set("Content-Type", "application/json")

		// Read the file and write it to the response
		http.ServeFile(w, r, filePath)
	})

}
