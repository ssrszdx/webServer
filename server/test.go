package main

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"testing"
)

func genSecret() string {
	var uname = "abc"
	var password = "abc123"

	var data = uname + ":" + password
	var secret = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println(secret)
	return secret
}

func genSecret2() string {
	var password = "demo4123"
	passhash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}
	fmt.Println(string(passhash))
	return string(passhash)
}

func rtnError(t *testing.T) {
	t.Errorf(genSecret2())
}
