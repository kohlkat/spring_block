package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)


type OK struct {
	Welcome string
}

func connect(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/connect" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	switch r.Method {

	case "GET":
		ack := OK{Welcome: "Connect to liquidOptimizer"}
		err := json.NewEncoder(w).Encode(ack)
		if err != nil {
			fmt.Println("error encoding peers", err)
		}
	}
}



func LaunchServer() {
	fs := http.FileServer(http.Dir("frontend"))
	http.Handle("/", fs)
	fmt.Println("Server up and running")

	//fmt.Printf("Starting server for GUI\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		//fmt.Println(err)
	}

	http.HandleFunc("/connect", connect)


}
