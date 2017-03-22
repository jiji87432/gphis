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
	resp := Login(auth)

	j, _ := json.MarshalIndent(resp, "", "  ")

	t.Log(string(j))
	t.Log("Test Login Passed")
}

func TestLogout(t *testing.T) {
	auth := MyAuth{Id: "2"}
	resp := Login(auth)

	rst := Logout(resp.Response)

	j, _ := json.MarshalIndent(rst, "", "  ")

	t.Log(string(j))
	t.Log("Test Logout Passed")
}

func TestCheck(t *testing.T) {
	auth := MyAuth{Id: "3"}
	resp := Login(auth)

	rsto, _ := Check(resp.Response)
	if rsto != true {
		t.Fatal("Check failed.")
	}

	t.Log("Test Check Passed")
}
