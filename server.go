package main

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"

// 	"sync"
// )

// var wg sync.WaitGroup
// var mutex sync.Mutex

// func uploadImage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("File upload endpoint Hit")

// 	// maximum of 10 MB files.
// 	if err := r.ParseMultipartForm(10 << 20); err != nil {
// 		fmt.Fprintln(w, err)
// 		return
// 	}

// 	formData := r.MultipartForm // read form data

// 	files := formData.File["images"] // grab the file names

// 	for _, file := range files {
// 		multiPartFile, err := file.Open()
// 		defer multiPartFile.Close()

// 		if err != nil {
// 			fmt.Fprintln(w, err)
// 			return
// 		}

// 		// create temporary file
// 		tempFile, err := ioutil.TempFile("temp-images", "upload-*.jpg")
// 		defer tempFile.Close()
// 		if err != nil {
// 			fmt.Fprintln(w, "Unable to create file for writing.")
// 			return
// 		}

// 		// copy content:
// 		_, err = io.Copy(tempFile, multiPartFile)
// 		if err != nil {
// 			fmt.Fprintln(w, err)
// 			return
// 		}
// 	}

// 	fmt.Fprintln(w, "Uploaded successfully")
// }

// func main() {
// 	http.HandleFunc("/upload", uploadImage)
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
