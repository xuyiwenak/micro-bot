package get_team_info

import (
	"errors"
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var (
	serverAddr           string
	once                 sync.Once
	ErrIncorrectResponse = errors.New("Response is incorrect")
)

func getTeamInfo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response := []byte(`{"ok": true, "team": {
			"id": "TL9A4N2FQ",
			"name": "notalar",
			"domain": "notalar",
			"icon": {
              "image_34": "https://slack.global.ssl.fastly.net/66f9/img/avatars-teams/ava_0002-34.png",
              "image_44": "https://slack.global.ssl.fastly.net/66f9/img/avatars-teams/ava_0002-44.png",
              "image_55": "https://slack.global.ssl.fastly.net/66f9/img/avatars-teams/ava_0002-55.png",
              "image_default": true
          }
		}}`)
	rw.Write(response)
}
func startServer() {
	server := httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
	log.Print("Test WebSocket server listening on ", serverAddr)
}

func TestGetTeamInfo(t *testing.T) {
	http.HandleFunc("/team.info", getTeamInfo)

	once.Do(startServer)
	api := slack.New("xoxp-689344750534-687199142064-684104843490-62a858cec778bb60561865481641ba67")

	teamInfo, err := api.GetTeamInfo()
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}
	fmt.Println(teamInfo.Name, teamInfo.ID, teamInfo.Domain, teamInfo.Icon)
	// t.Fatal refers to -> t.Errorf & return
}
