package interfaces

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

var baseURL = "http://127.0.0.1:8000"

func TestRegister(t *testing.T) {
	url := baseURL + "/api/signin/v1/register"
	data := map[string]interface{}{
		"name":     "Jerry",
		"password": "666777",
		"phone":    "1234567890",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Error(err.Error())
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err.Error())
	}
	response := map[string]interface{}{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("Register User reply: %v", response)
}

func TestLogin(t *testing.T) {
	url := baseURL + "/api/signin/v1/login"
	data := map[string]interface{} {
		"username": "Jerry",
		"password": "666777",
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		t.Error(err.Error())
	}
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		t.Error(err.Error())
	}
	response := map[string]interface{}{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("User login reply: %v", response)
}

func TestGetSignInInfo(t *testing.T) {
	url := baseURL + "/api/signin/v1/signininfo"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err.Error())
	}
	// 需要设置一个有效的token
	req.Header.Add("Authorization", "NnxKZXJyeQ==")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	response := map[string]interface{}{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("User get sign in info reply: %v", response)
}

func TestSignIn(t *testing.T) {
	url := baseURL + "/api/signin/v1/signin"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err.Error())
	}
	// 一个有效的token
	req.Header.Add("Authorization", "NnxKZXJyeQ==")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err.Error())
	}
	response := map[string]interface{}{}
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Error(err.Error())
	}
	fmt.Printf("User Sign in the day reply, %v", response)
}