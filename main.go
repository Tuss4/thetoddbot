package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	port     = ":8000"
	commands = map[string]string{
		"ping":     `{"text": "PING FIVE! WHOO-PSH! SNAP!"}`,
		"betrayal": `{"text": "<https://www.youtube.com/watch?v=uw5P2up5bL4>"}`,
		"miracle":  `{"text": "<https://www.youtube.com/watch?v=uehf8e43Vtk>"}`,
		"meow":     `{"text": "<https://www.youtube.com/watch?v=xRhATB9NelU>"}`,
	}
	notFound = `{"text": "That command was bad and you should feel bad."}`
	hookURL  = os.Getenv("SLACK_WEBHOOK_URL")
)

func handleHook(res http.ResponseWriter, req *http.Request, val string) {
	client := &http.Client{}
	req, err := http.NewRequest("POST", hookURL, bytes.NewBuffer([]byte(val)))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Status:", resp.Status)
}

func handleQuery(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	qMap := req.URL.Query()
	qVal := qMap["q"][0]
	if commands[qVal] != "" {
		handleHook(res, req, commands[qVal])
		io.WriteString(res, `{"status": "Valid Command"}`)
	} else {
		handleHook(res, req, notFound)
		io.WriteString(res, `{"status": "Invalid Command"}`)
	}
}

func main() {
	fmt.Println("Server running listening on port", port)
	http.HandleFunc("/command", handleQuery)
	http.ListenAndServe(port, nil)
}
