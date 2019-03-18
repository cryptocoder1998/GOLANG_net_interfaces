package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"
)

//IntInfo contains main information about a network interface
type IntInfo struct {
	Name      string   `json:"name,omitempty"`
	MAC       string   `json:"MAC,omitempty"`
	Addresses []string `json:"addr,omitempty"`
	MTU       int      `json:"MTU,omitempty"`
}

var VersionNum string = "v1.0"
var InterfacesNames []string

func VersionEndpoint(w http.ResponseWriter, req *http.Request) {
	if VersionNum == "" {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error: version error")
		return
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(VersionNum)
}

func IntListEndpoint(w http.ResponseWriter, req *http.Request) {

	list, err := net.Interfaces()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error: interfaces error")
		return
	}
	for _, buffer := range list {
		InterfacesNames = append(InterfacesNames, buffer.Name)
	}
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(InterfacesNames)
	InterfacesNames = nil
}

func IntInfoEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	int_name := params["int-name"]

	list, err := net.Interfaces()
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Error: interfaces error")
		return
	}
	for _, buffer := range list {
		if buffer.Name == int_name {
			var SendInfo IntInfo
			SendInfo.Name = buffer.Name
			SendInfo.MAC = buffer.HardwareAddr.String()

			addr_buf, err := buffer.Addrs()
			if err != nil {
				w.WriteHeader(500)
				json.NewEncoder(w).Encode("Error: interface " + int_name + "error")
				return
			}

			for _, addr := range addr_buf {
				SendInfo.Addresses = append(SendInfo.Addresses, addr.String())
			}

			SendInfo.MTU = buffer.MTU
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(SendInfo)
			return
		}
	}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode("Such interface doesn't exist")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/service/version", VersionEndpoint).Methods("GET")
	router.HandleFunc("/service/{api-version}/interfaces", IntListEndpoint).Methods("GET")
	router.HandleFunc("/service/{api-version}/interface/{int-name}", IntInfoEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
