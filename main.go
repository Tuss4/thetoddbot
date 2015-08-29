package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type slackResp struct {
	Text string
}

var (
	port     = ":8000"
	commands = map[string]slackResp{
		"ping":     slackResp{"PING FIVE! WHOO-PSH! SNAP!"},
		"betrayal": slackResp{"<https://www.youtube.com/watch?v=uw5P2up5bL4>"},
		"miracle":  slackResp{"<https://www.youtube.com/watch?v=uehf8e43Vtk>"},
		"meow":     slackResp{"<https://www.youtube.com/watch?v=xRhATB9NelU>"},
	}
	notFound = slackResp{"That command was bad and you should feel bad."}
)

func handleQuery(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	qMap := req.URL.Query()
	qVal := qMap["q"][0]
	if commands[qVal] != (slackResp{}) {
		data, err := json.Marshal(commands[qVal])
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(res, string(data))
	} else {
		data, err := json.Marshal(notFound)
		if err != nil {
			log.Fatal(err)
		}
		io.WriteString(res, string(data))
	}
}

func main() {
	fmt.Println("Server running listening on port", port)
	http.HandleFunc("/command", handleQuery)
	http.ListenAndServe(port, nil)
}
