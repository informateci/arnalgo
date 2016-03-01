package main

import "github.com/thoj/go-ircevent"

const (
	nickname string = "arnalgo"
	server   string = "irc.freenode.net:6667"
)

var Dialogo = []botCommand{}
var channels = []string{"##mariolotting"}

func main() {
	ircconn := irc.IRC(nickname, nickname)
	ircconn.AddCallback("PRIVMSG", func(event *irc.Event) {
        var replyto string
        var msgs []string
        for _, commando := range Dialogo {
            if commando.matchino.MatchString(event.Message()) {
                msgs, replyto = commando.process(event)
                for _, msg := range msgs {
                    ircconn.Privmsgf(replyto, msg)
                }
                break
            }
        }
	})

	ircconn.Connect(server)
	for _, channel := range channels {
		ircconn.Join(channel)
	}
	ircconn.Loop()
}
