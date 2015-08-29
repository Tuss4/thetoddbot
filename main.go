package main

import (
	"fmt"
	"io"
	"net/http"
)

var (
	port     = ":8000"
	commands = map[string]string{
		"ping":     "PING FIVE! WHOO-PSH! SNAP!",
		"betrayal": "https://www.youtube.com/watch?v=uw5P2up5bL4",
		"miracle":  "https://www.youtube.com/watch?v=uehf8e43Vtk",
		"meow":     "https://www.youtube.com/watch?v=xRhATB9NelU",
	}
	notFound = "That command was bad and you should feel bad."
)

func handleQuery(res http.ResponseWriter, req *http.Request) {
	qMap := req.URL.Query()
	fmt.Println(qMap)
	res.Header().Set("Content-Type", "application/json")
	io.WriteString(res, `{"foo": "bar"}`)
}

func main() {
	fmt.Println("Server running listening on port", port)
	http.HandleFunc("/command", handleQuery)
	http.ListenAndServe(port, nil)
}
