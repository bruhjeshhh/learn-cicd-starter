package auth

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	req, err := http.NewRequest("GET", "https://someting.com", nil)
	if err != nil {
		fmt.Print("error thrown while  making a req: ", err)
	}
	key := "abc123"
	req.Header.Set("Authorization", "ApiKey "+key)
	resp, _ := GetAPIKey(req.Header)

	if !reflect.DeepEqual(key, resp) {
		t.Fatalf("expected: %v, got: %v", key, resp)
	}
}
