package main

import (
    "github.com/kennygrant/sanitize"
    "github.com/nu7hatch/gouuid"
    "github.com/PuerkitoBio/goquery"
    "github.com/thoj/go-ircevent"
    "math"
    "regexp"
    "strings"
)

var summary string = ""

func init() {
    Dialogo = append(Dialogo,
        botCommand {
            regexp.MustCompile("e allora\\?$"),
            func(event *irc.Event, matches []string) ([]string, string) {
                return []string{"e allora le foibe?"}, whomToReply(event)
            },
        },
        botCommand {
            regexp.MustCompile("(^allivello\\?)|(parliamo di)"),
            func(event *irc.Event, matches []string) ([]string, string) {
                uuido, _ := uuid.NewV4()
                wikipedia_url := "http://it.wikipedia.org/wiki/Speciale:PaginaCasuale#" + uuido.String()
                resp := noembedRequest(wikipedia_url)
                doc, _ := goquery.NewDocumentFromReader(strings.NewReader(resp.Html))
                summary = sanitize.HTML(doc.Find("p").Text())
                return []string{resp.Title}, whomToReply(event)
            },
        },
        botCommand {
            regexp.MustCompile("parliamone"),
            func(event *irc.Event, matches []string) ([]string, string) {
                var ret string = ""
                if summary != "" {
                    ret = summary[:int(math.Min(float64(len(summary)), 430.))]
                    summary = ""
                }
                return []string{ret}, whomToReply(event)
            },
        },
        botCommand {
            regexp.MustCompile("anche no"),
            func(event *irc.Event, matches []string) ([]string, string) {
                var ret string = ""
                if summary != "" {
                    ret = "ಥ_ಥ  ockay"
                    summary = "┌∩┐(◕_◕)┌∩┐"
                }
                return []string{ret}, whomToReply(event)
            },
        },
    )
}
