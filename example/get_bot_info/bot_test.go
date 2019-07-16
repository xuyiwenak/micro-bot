package get_bot_info

import (
	"fmt"
	"github.com/nlopes/slack"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

var (
	serverAddr string
	once       sync.Once
)

func getBotInfo(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	response := []byte(`{"ok": true, "bot": {
			"id":"BLDSRGCFN",
			"deleted":false,
			"name":"github",
			"icons": {
              "image_36":"https:\/\/a.slack-edge.com\/2fac\/plugins\/github\/assets\/service_36.png",
              "image_48":"https:\/\/a.slack-edge.com\/2fac\/plugins\/github\/assets\/service_48.png",
              "image_72":"https:\/\/a.slack-edge.com\/2fac\/plugins\/github\/assets\/service_72.png"
            }
        }}`)
	rw.Write(response)
}

func startServer() {
	server := httptest.NewServer(nil)
	serverAddr = server.Listener.Addr().String()
	log.Print("Test WebSocket server listening on ", serverAddr)
}

func TestGetBotInfo(t *testing.T) {
	http.HandleFunc("/bots.info", getBotInfo)
	botToken := "xoxb-689344750534-694391105429-Um9qmK7dDUhYJjHvjd8PxnOO"
	once.Do(startServer)

	api := slack.New(botToken, slack.OptionDebug(true))

	bot, err := api.GetBotInfo("BLDSRGCFN")
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
		return
	}

	fmt.Printf("%v", bot)
}
