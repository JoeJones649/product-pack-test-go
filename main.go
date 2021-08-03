package main
import "encoding/json"
import "log"
import "net/http"

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		log.Println("Hello, world!")
		js, err := json.Marshal(map[string]string{
			"test": "test",
		})

		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Running")
	log.Fatal(http.ListenAndServe(":8000", nil))
}