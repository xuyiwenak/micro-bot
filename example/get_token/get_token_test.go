package get_token

import (
	"fmt"
	"github.com/nlopes/slack"
	"net/http"
	"testing"
)

func getToken(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response := []byte(`{"ok": true, "team": {
			"id": "evan",
			"name": "notalar",
          }
		}}`)
	rw.Write(response)
}

func TestGetTokenInfo(t *testing.T) {
	http.HandleFunc("/oauth.access", getToken)
	authCode := "111111111"
	clientID := "xxxx"
	clientSecret := "xxxxxx"
	redirectURI := ""

	oAuthRes, err := slack.GetOAuthResponse(http.DefaultClient, clientID, clientSecret, authCode, redirectURI)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	fmt.Printf("auth result : %v", &oAuthRes)
}
