package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
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
	hookURL  = os.Getenv("SLACK_WEBHOOK_URL")
)

func handleHook(res http.ResponseWriter, req *http.Request) {
	client := &http.Client{}
	the_string := `{"text": "<https://www.youtube.com/watch?v=uehf8e43Vtk>"}`
	print(hookURL)
	req, err := http.NewRequest("POST", hookURL, bytes.NewBuffer([]byte(the_string)))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	print(resp.Status, resp.StatusCode)
}

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
	http.HandleFunc("/hook", handleHook)
	http.HandleFunc("/command", handleQuery)
	http.ListenAndServe(port, nil)
}
