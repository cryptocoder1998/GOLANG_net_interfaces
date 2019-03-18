package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const BaseURL string = "http://"
const help string = `AVAILABLE OPTIONS:
help(-h) - shows helpful information 
show(-s) [interface_name] - shows information about specified network interface
list(-l) - shows all names of all available network interfaces 
--version - shows API version of service
USAGE:
./client [command] [command_args] --server [ip_address] --port [port_value]`

//IntInfo contains main information about a network interface
type IntInfo struct {
	Name    string   `json:"name,omitempty"`
	MAC     string   `json:"MAC,omitempty"`
	Address []string `json:"addr,omitempty"`
	MTU     int      `json:"MTU,omitempty"`
}

var ErrorString string

func SendVersionRequest(ip string, port string) (string, error) {
	var URL string = BaseURL + ip + ":" + port + "/service/version"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "Couldn't create an appropriate request", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Server is unavailable", err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading data", err
	}

	if resp.StatusCode == 500 {
		err = json.Unmarshal(body, &ErrorString)
		return ErrorString, err
	}
	var ServerVersion string
	err = json.Unmarshal(body, &ServerVersion)
	if err != nil {
		return "Error while unmarshalling data", err
	}
	return ServerVersion, nil
}
func SendEnumerateRequest(ip string, port string) (string, error, []string) {
	var URL string = BaseURL + ip + ":" + port + "/service/v1/interfaces"
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "Couldn't create an appropriate request", err, nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Server is unavailable", err, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading data", err, nil
	}
	if resp.StatusCode == 500 {
		err = json.Unmarshal(body, &ErrorString)
		return ErrorString, nil, nil
	}
	var InterfacesNames []string
	err = json.Unmarshal(body, &InterfacesNames)
	if err != nil {
		return "Error unmarshalling data", err, nil
	} else {
		return "", nil, InterfacesNames
	}
}
func SendIntRequest(ip string, port string, int_name string) (string, error, *IntInfo) {
	var URL string = BaseURL + ip + ":" + port + "/service/v1/interface/" + int_name
	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return "Couldn't create an appropriate request", err, nil
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "Server is unavailable", err, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error reading data", err, nil
	}

	if resp.StatusCode == 500 || resp.StatusCode == 404 {
		err = json.Unmarshal(body, &ErrorString)
		return ErrorString, nil, nil
	}
	var InterBuf IntInfo
	err = json.Unmarshal(body, &InterBuf)
	if err != nil {
		return "Error unmarshalling data", err, nil
	} else {
		return "", nil, &InterBuf
	}

}
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println(help)
		return
	}
	if len(args) == 1 && (args[0] == "help" || args[0] == "-h") {
		fmt.Println(help)
		return
	} else if len(args) == 5 {
		if args[0] == "list" || args[0] == "-l" {
			str, err, arr := SendEnumerateRequest(args[2], args[4])
			if err != nil || str != "" {
				fmt.Println(str)
				return
			}
			fmt.Println("INTERFACES: ", arr)
		} else if args[0] == "--version" {
			str, err := SendVersionRequest(args[2], args[4])
			if err != nil {
				fmt.Println(str)
				return
			}
			fmt.Println("VERSION: " + str)
		} else {
			fmt.Println(help)
			return
		}
	} else if len(args) == 6 {
		if args[0] == "show" || args[0] == "-s" {
			str, err, inf := SendIntRequest(args[3], args[5], args[1])
			if err != nil || str != "" {
				fmt.Println(str)
				return
			}
			fmt.Printf("%s: %+v", args[1], *inf)
		} else {
			fmt.Println(help)
			return
		}
	} else {
		fmt.Println(help)
		return
	}

}
