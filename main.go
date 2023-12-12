package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"path"
	"strconv"

	"git.mrcyjanek.net/p3pch4t/p3pgo/lib/core"
	"github.com/joho/godotenv"
)

var botPi *core.PrivateInfoS

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("I2P_HTTP_PROXY") == "" {
		log.Fatalln("I2P_HTTP_PROXY is \"\"")
	}
	if os.Getenv("PRIVATEINFO_ROOT_ENDPOINT") == "" {
		log.Fatalln("PRIVATEINFO_ROOT_ENDPOINT is \"\"")
	}

	core.I2P_HTTP_PROXY = os.Getenv("I2P_HTTP_PROXY")
	core.LOCAL_SERVER_PORT, err = strconv.Atoi(os.Getenv("LOCAL_SERVER_PORT"))
	if err != nil {
		log.Fatalln("LOCAL_SERVER_PORT", err)
	}
	botPi = core.OpenPrivateInfo(path.Join(os.Getenv("HOME"), ".config", ".p3pchucknorris"), "Group Host", "")
	botPi.Endpoint = core.Endpoint(os.Getenv("PRIVATEINFO_ROOT_ENDPOINT"))
	botPi.MessageCallback = append(botPi.MessageCallback, botMsgHandler)
	botPi.IntroduceCallback = append(botPi.IntroduceCallback, botIntroduceHandler)
	if !botPi.IsAccountReady() {
		botPi.Create("Chuck Norris", "chuchnorris-jokes@mrcyjanek.net", 4096)
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	log.Println("p3p is running...", botPi.Endpoint)
	for sig := range c {
		log.Println("Closing [", sig, "] ...")
		return
	}
}

//go:embed welcome.md
var welcomeMessage string

func botIntroduceHandler(pi *core.PrivateInfoS, ui *core.UserInfo, evt *core.Event) {
	pi.SendMessage(ui, core.MessageTypeText, welcomeMessage)
}
func getJoke() string {
	r, err := http.Get("https://api.chucknorris.io/jokes/random")
	if err != nil {
		return err.Error()
	}
	defer r.Body.Close()
	b, err := io.ReadAll(r.Body)
	if err != nil {
		return err.Error()
	}
	var joke Joke
	err = json.Unmarshal(b, &joke)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%s [source](%s)", joke.Value, joke.URL)
}

type Joke struct {
	Categories []interface{} `json:"categories"`
	CreatedAt  string        `json:"created_at"`
	IconURL    string        `json:"icon_url"`
	ID         string        `json:"id"`
	UpdatedAt  string        `json:"updated_at"`
	URL        string        `json:"url"`
	Value      string        `json:"value"`
}

func botMsgHandler(pi *core.PrivateInfoS, ui *core.UserInfo, evt *core.Event, msg *core.Message) {
	if msg.Body != "!joke" {
		pi.SendMessage(ui, core.MessageTypeText, welcomeMessage)
		return
	}
	pi.SendMessage(ui, core.MessageTypeText, getJoke())
}
