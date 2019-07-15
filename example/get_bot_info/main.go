package get_bot_info

import (
	"fmt"

	"github.com/nlopes/slack"
)

func main() {
	//myToken := "xoxp-689344750534-687199142064-694479525669-e1291c91686c60b709400122218f7885"
	botToken := "xoxb-689344750534-694391105429-0u2oOX6kzvXRKxH3ZVSOjlgR"
	api := slack.New(botToken, slack.OptionDebug(true))
	// If you set debugging, it will log all requests to the console
	// slack.New("YOUR_TOKEN_HERE", slack.OptionDebug(true))
	// 查不到在url浏览器里面有， 最后一排字母就是例如
	// https://bamboorat.slack.com/services/BLDSRGCFN
	bot, err := api.GetBotInfo("BLDSRGCFN")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, Name: %s\n", bot.ID, bot.Name)

}
