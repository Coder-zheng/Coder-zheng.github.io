package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	//POST takes the uploaded file(s) and saves it to disk.
	case "POST":
		//parse the multipart form in the request
		err := r.ParseMultipartForm(100000)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		//get a ref to the parsed multipart form
		m := r.MultipartForm

		//get the *fileheaders
		files := m.File["uploadfile"]
		for i, _ := range files {
			//for each fileheader, get a handle to the actual file
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				fmt.Println("open err:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//create destination file making sure the path is writeable.
			dst, err := os.Create("./upload/" + files[i].Filename)
			defer dst.Close()
			if err != nil {
				fmt.Println("create err:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			//copy the uploaded file to the destination file
			if _, err := io.Copy(dst, file); err != nil {
				fmt.Println("copy err:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/upload", uploadHandler)

	//static file handler.
	http.Handle("/staticfile/", http.StripPrefix("/staticfile/", http.FileServer(http.Dir("./upload"))))

	//Listen on port 8080
	http.ListenAndServe(":8080", nil)
}
