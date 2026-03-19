package auth

import (
	"net/http"
	"testing"
)

func TestAuthorizationHeaderNoKey(t *testing.T) {
	headers := http.Header{"Content-Type": {"application/json"}}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected error, missing key")
	}
}

func TestAuthorizationHeaderEmptyKey(t *testing.T) {
	headers := http.Header{"Authorization": {""}}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected error, empty key")
	}
}

func TestAuthorizationHeaderNoSep(t *testing.T) {
	headers := http.Header{"Authorization": {"ApiKeyblablablabla"}}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected error, no separator")
	}
}

func TestAuthorizationHeaderNoApiKey(t *testing.T) {
	headers := http.Header{"Authorization": {"MapiKey blablablabla"}}
	_, err := GetAPIKey(headers)
	if err == nil {
		t.Fatal("Expected error, ApiKey substring not present")
	}
}

func TestAuthorizationValid(t *testing.T) {
	headers := http.Header{"Authorization": {"ApiKey blablablabla"}}
	key, err := GetAPIKey(headers)
	if err != nil {
		t.Fatal(err)
	}
	if key != "blablablabla" {
		t.Fatal("not parsed correctly")
	}
}
