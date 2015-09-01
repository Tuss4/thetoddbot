package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

// Shout out to http://www.markjberger.com/testing-web-apps-in-golang/

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

func TestMethodNotAllowed(t *testing.T) {
	req, err := http.NewRequest("GET", "/command", nil)
	w := httptest.NewRecorder()
	if err != nil {
		t.Fatal("Creating 'GET /command' request failed!")
	}
	commandHandler().ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Command page didn't return %v", http.StatusOK)
	}
	if w.Body.String() != notAllowed {
		t.Errorf("Response body didn't return %v", notAllowed)
	}
}

func TestCommandFound(t *testing.T) {
	// Test "oneq"
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
	// Test "betrayal"
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
	// Test "ping"
	data.Del("text")
	data.Add("text", "ping")
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
	if w.Body.String() != commands["ping"] {
		t.Errorf("Response body didn't return %v", commands["ping"])
	}
	// Test "miracle"
	data.Del("text")
	data.Add("text", "miracle")
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
	if w.Body.String() != commands["miracle"] {
		t.Errorf("Response body didn't return %v", commands["miracle"])
	}
	// Test "meow"
	data.Del("text")
	data.Add("text", "meow")
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
	if w.Body.String() != commands["meow"] {
		t.Errorf("Response body didn't return %v", commands["meow"])
	}
}
