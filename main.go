package main

import (
	"file_server/config"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {
	location := config.Config.FileLocation
	fs := http.FileServer(http.Dir(location))
	http.Handle("/ms-file-server/", http.StripPrefix("/ms-file-server/", fs))
	http.HandleFunc("/", indexHandler)
	fmt.Printf("Starting File Server at port :" + config.Config.ServicePort)
	err := http.ListenAndServe(config.Config.ServiceHost+":"+config.Config.ServicePort, nil)
	if err != nil {
		fmt.Printf("Listen And Serve : ", err)
		return
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	location := config.Config.FileLocation
	if err := r.ParseForm(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	f, err := os.Open(location)
	if f != nil {
		defer f.Close()
		return
	}

	if err != nil {
		fmt.Println(err)
	}

	file, _ := f.Stat()

	FileSize := strconv.FormatInt(file.Size(), 10)

	contentDisposition := fmt.Sprintf("attachment; filename=%s", f.Name())
	w.Header().Set("Content-Disposition", contentDisposition)
	w.Header().Set("Content-Length", FileSize)

	if _, err := io.Copy(w, f); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
