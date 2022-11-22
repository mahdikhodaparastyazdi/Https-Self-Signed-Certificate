package main

//using for writing cert.pem & key.pem
// go run generate_cert.go --host=localhost
import (
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This is an example \n"))
}
func redir(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://127.0.0.1:10443/"+r.RequestURI, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", handler)
	log.Printf("About to")

	go http.ListenAndServe(":9999", http.HandlerFunc(redir))
	err := http.ListenAndServeTLS(":10443", "cert.pem", "key.pem", nil)
	if err != nil {
		log.Fatal(err)
	}
}
