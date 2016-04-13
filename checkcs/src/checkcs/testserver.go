package checkcs

// {
//   "code": 0,
//   "status": "OK"
// }

// {
//   "code": 3,
//   "status": "UNKNOWN",
//   "message": "[rc-notfound] No replication controllers found!"
// }


import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

func Startserver() {
	http.HandleFunc("/ok", okhandler)
	http.HandleFunc("/failed", failedhandler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func okhandler(w http.ResponseWriter, r *http.Request) {
	okcheck := Check{
		Code: 0,
		Status: "OK",
	}

	b, err := json.Marshal(okcheck)
	if err != nil {
		fmt.Println("error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func failedhandler(w http.ResponseWriter, r *http.Request) {
	failedcheck := Check{
		Code: 3,
		Status: "UNKNOWN",
		Message: "[rc-notfound] No replication controllers found!",
	}

	b, err := json.Marshal(failedcheck)
	if err != nil {
		fmt.Println("error: ", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}
