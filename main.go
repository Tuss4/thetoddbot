package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type slackMsg struct {
	Text  string
	Token string
}

var (
	port     = ":8000"
	commands = map[string]string{
		"ping":     `{"text": "PING FIVE! WHOO-PSH! SNAP!"}`,
		"betrayal": `{"text": "<https://www.youtube.com/watch?v=uw5P2up5bL4>"}`,
		"miracle":  `{"text": "<https://www.youtube.com/watch?v=uehf8e43Vtk>"}`,
		"meow":     `{"text": "<https://www.youtube.com/watch?v=xRhATB9NelU>"}`,
	}
	notFound     = `{"text": "That command was bad and you should feel bad. #ZoidbergLevelRoast"}`
	hookURL      = os.Getenv("SLACK_WEBHOOK_URL")
	slackToken   = os.Getenv("SLACK_TOKEN")
	trigger      = "thetodd: "
	invalidToken = `{"status": "Invalid Token"}`
	notAllowed   = `{"status": "Method Not Allowed"}`
)

func hello(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Contet-Type", "application/json")
	io.WriteString(res, `{"status": "Sup, fam?"}`)
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Fatal(err)
		}
		command := strings.TrimPrefix(r.Form.Get("text"), trigger)
		msg := slackMsg{command, r.Form.Get("token")}
		if msg.Token == slackToken {
			if commands[msg.Text] != "" {
				io.WriteString(w, commands[msg.Text])
			} else {
				io.WriteString(w, notFound)
			}
		} else {
			io.WriteString(w, invalidToken)
		}
	} else {
		io.WriteString(w, notAllowed)
	}
}

func main() {
	fmt.Println("Server running listening on port", port)
	http.HandleFunc("/command", handlePost)
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(port, nil)
}
