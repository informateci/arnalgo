package main

import (
    "bufio"
    "github.com/thoj/go-ircevent"
    "log"
    "os"
    "regexp"
    "strings"
)

var Icseh map [string] []string

func stampaIcse (icsamelo string) []string {
    lines := []string{"", "", "", "", "", "", ""}
    for _, letter := range icsamelo {
        for idx, line := range Icseh[string(letter)] {
            lines[idx] = lines[idx] + line
        }
    }

    return lines
}

func caricaIcseh() {
    leggino, err := os.Open("./dati/icse.txt")

    if err != nil {
        log.Fatal(err)
        return
    }

    var icsehLines []string
    r := bufio.NewReader(leggino)
    line, _, erre := r.ReadLine()
    for erre == nil {
        icsehLines = append(icsehLines, string(line))
        line, _, erre = r.ReadLine()
    }

    Icseh = map [string] []string {}

    for _, icsina := range icsehLines {
        var splitti []string
        splitti = strings.Split(icsina, "|")
        if len(splitti) == 10 {
            splitti = splitti[:len(splitti)-3]
            splitti = append(splitti, "|", "")
        }
        Icseh[splitti[len(splitti)-2]] = splitti[:len(splitti)-2]
    }
}

func init() {
    caricaIcseh()
    Dialogo = append(Dialogo,
        botCommand {
            regexp.MustCompile("^icsah (.+)"),
            func(event *irc.Event, matches []string) ([]string, string) {
                return stampaIcse(matches[1]), whomToReply(event)
            },
        },
        botCommand {
            regexp.MustCompile("^bamba$"),
            func(event *irc.Event, c []string) ([]string, string) {
                return stampaIcse("ROSA"), whomToReply(event)
            },
        },
    )
}
