package auth

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetAPIKey1(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey 1234567890")

	key, err := GetAPIKey(header)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	want := "1234567890"

	if key != want {
		t.Fatalf("%v != %v", key, want)
	}
}

func TestGetAPIKey2(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "1234567890")

	key, err := GetAPIKey(header)
	if err == nil {
		fmt.Println(key)
		t.Fail()
	}

	if key != "" {
		t.Fatalf("Expected no value for error, got %v", key)
	}
}

func TestGetAPIKey3(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "")

	key, err := GetAPIKey(header)
	if err == nil {
		fmt.Println(key)
		t.Fail()
	}

	if key != "" {
		t.Fatalf("Expected no value for error, got %v", key)
	}
}
