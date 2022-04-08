# DiscordGo command manager

[![](https://img.shields.io/github/v/tag/thechampagne/discordgomanager?label=version)](https://github.com/thechampagne/discordgomanager/releases/latest) [![](https://img.shields.io/github/license/thechampagne/discordgomanager)](https://github.com/thechampagne/discordgomanager/blob/main/LICENSE)

DiscordGo command manager.

### Download

```
go get github.com/thechampagne/discordgomanager
```

### Example

```go
package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/thechampagne/discordgomanager"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type hello struct{}

func (hello) GetCommand() string {
	return "hello"
}

func (hello) Run(args []string,s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Hi")
}

func main(){

	manager := discordgomanager.New("!")
	manager.AddCommand(hello{})

	client, err := discordgo.New("Bot " + "token")

	if err != nil {
		log.Fatal(err)
	}
	client.AddHandler(manager.Handler)

	err = client.Open()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	client.Close()
}
```

### License

This repo is released under the [MIT License](https://github.com/thechampagne/discordgomanager/blob/main/LICENSE).