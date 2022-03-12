package admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

var baseURL = "http://127.0.0.1:8001"

func TestGetSignInRecord(t *testing.T) {
	url := baseURL + "/api/admin/signin/v1/getSignInRecord"
	data := map[string]interface{} {
		"user": []int64{1, 2},
		"day":  []string{"20220311", "20220312"},
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
	fmt.Printf("Admin get user sign in record reply: %v", response)

}