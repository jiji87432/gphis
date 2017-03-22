package permissions

import (
	"encoding/json"
	"testing"
)

func TestGetPermit(t *testing.T) {
	// case 1
	resp := GetPermit("aaa:bbb:ccc:ddd", 1)
	if resp.Success != false {
		t.Fatal("case1: failed")
	} else {
		tmp, _ := json.MarshalIndent(resp, "", "  ")
		t.Log(string(tmp))
		t.Log("case1: passed")
	}

	// case 2
	resp = GetPermit("1:2:3:4:5", 1)
	if resp.Success != true {
		t.Fatal("case2: failed")
	} else {
		tmp, _ := json.MarshalIndent(resp, "", "  ")
		t.Log(string(tmp))
		t.Log("case2: passed")
	}

	// case 3
	resp = GetPermit("1:2:3:4:5", 6)
	if resp.Success != false {
		t.Fatal("case3: failed")
	} else {
		tmp, _ := json.MarshalIndent(resp, "", "  ")
		t.Log(string(tmp))
		t.Log("case3: passed")
	}

	// case 4
	resp = GetPermit("1", 1)
	if resp.Success != true {
		t.Fatal("case4: failed")
	} else {
		tmp, _ := json.MarshalIndent(resp, "", "  ")
		t.Log(string(tmp))
		t.Log("case4: passed")
	}

	// case 5
	resp = GetPermit("1", 2)
	if resp.Success != false {
		t.Fatal("case5: falied")
	} else {
		tmp, _ := json.MarshalIndent(resp, "", "  ")
		t.Log(string(tmp))
		t.Log("case5: passed")
	}
}
