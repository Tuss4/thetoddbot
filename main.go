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
		"oneq":     `{"text": "*Points at your shoe game*\nWHAT ARE THOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOSE?!\n<http://i.imgur.com/EnxQLii.gif>"}`,
	}
	notFound     = `{"text": "That command was bad and you should feel bad. #ZoidbergLevelRoast\n<http://i.imgur.com/UnXhfOJ.gif>"}`
	slackToken   = os.Getenv("SLACK_TOKEN")
	trigger      = "thetodd: "
	invalidToken = `{"status": "Invalid Token"}`
	notAllowed   = `{"status": "Method Not Allowed"}`
	helloMsg     = `{"status": "Sup, fam?"}`
)

func getToken() string {
	if slackToken != "" {
		return slackToken
	} else {
		return ""
	}
}

func helloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, helloMsg)
	})
}

func commandHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				log.Fatal(err)
			}
			t := getToken()
			command := strings.TrimPrefix(r.Form.Get("text"), trigger)
			msg := slackMsg{command, r.Form.Get("token")}
			if msg.Token == t {
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
	})
}

func main() {
	fmt.Println("Server running and listening on port", port)
	http.HandleFunc("/command", commandHandler())
	http.HandleFunc("/hello", helloHandler())
	http.ListenAndServe(port, nil)
}
