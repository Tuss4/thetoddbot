package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestHello(t *testing.T) {

	req, err := http.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	if err != nil {
		t.Fatal("Creating 'GET /hello' request failed!")
	}
	helloHandler().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Hello page didn't return %v", http.StatusOK)
	}
	if w.Body.String() != helloMsg {
		t.Errorf("Reponse body didn't return %v", helloMsg)
	}
}

func TestCommandNotFound(t *testing.T) {
	data := url.Values{}
	data.Set("token", getToken())
	data.Add("text", "bruh")
	req, err := http.NewRequest("POST", "/command", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	if err != nil {
		t.Fatal("Creating 'POST /command' request failed!")
	}
	commandHandler().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Command page didn't return %v", http.StatusOK)
	}
	if w.Body.String() != notFound {
		t.Errorf("Response body didn't return %v", notFound)
	}
}

func TestCommandFound(t *testing.T) {
	data := url.Values{}
	data.Set("token", getToken())
	data.Add("text", "oneq")
	req, err := http.NewRequest("POST", "/command", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w := httptest.NewRecorder()
	if err != nil {
		t.Fatal("Creating 'POST /command' request failed!")
	}
	commandHandler().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Command page didn't return %v", http.StatusOK)
	}
	if w.Body.String() != commands["oneq"] {
		t.Errorf("Response body didn't return %v", commands["oneq"])
	}
	data.Del("text")
	data.Add("text", "betrayal")
	req, err = http.NewRequest("POST", "/command", bytes.NewBufferString(data.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	w = httptest.NewRecorder()
	if err != nil {
		t.Fatal("Creating 'POST /command' request failed!")
	}
	commandHandler().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Command page didn't return %v", http.StatusOK)
	}
	if w.Body.String() != commands["betrayal"] {
		t.Errorf("Response body didn't return %v", commands["betrayal"])
	}
}
