package idauth

import (
	"encoding/json"
	"testing"
)

type MyAuth struct {
	Id string
}

func (this MyAuth) Authenticate() (bool, string) {
	return true, this.Id
}

func TestLogin(t *testing.T) {
	auth := MyAuth{Id: "1"}
	token, err := Login(auth)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}
	t.Log(token)
	t.Log("Test Login Passed")
}

func TestLogout(t *testing.T) {
	auth := MyAuth{Id: "2"}
	rst, err := Login(auth)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	un := AuthResp{}
	json.Unmarshal([]byte(rst), &un)
	rsto, err := Logout(un.Response)
	if err != nil {
		t.Fatalf("Logout failed: %v", err)
	}
	t.Log(rsto)
	t.Log("Test Logout Passed")
}

func TestCheck(t *testing.T) {
	auth := MyAuth{Id: "3"}
	rst, err := Login(auth)
	if err != nil {
		t.Fatalf("Login failed: %v", err)
	}

	un := AuthResp{}
	json.Unmarshal([]byte(rst), &un)
	rsto, _ := Check(un.Response)
	if rsto != true {
		t.Fatal("Check failed.")
	}

	t.Log("Test Check Passed")
}
