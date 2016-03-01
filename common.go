package main

import (
    "encoding/json"
    "github.com/thoj/go-ircevent"
    "io/ioutil"
    "net/http"
    "net/url"
    "regexp"
)

func whomToReply(e *irc.Event) string {
    var replyto string
    if e.Arguments[0] == nickname {
        replyto = e.Nick
    } else {
        replyto = e.Arguments[0]
    }
    return replyto
}

type NoembedJSON struct {
    Html string
    ProviderName string `json:"provider_name"`
    Title string
    Type string
}

func noembedRequest(urlstr string) *NoembedJSON {
    var val url.Values
    val = url.Values{}
    val.Add("url", urlstr)
    resp, _ := http.Get("http://noembed.com/embed?" + val.Encode())
    body, _ := ioutil.ReadAll(resp.Body)
    ret := NoembedJSON{}
    json.Unmarshal(body, &ret)
    return &ret
}

type botCommand struct {
    matchino *regexp.Regexp
    handler func(*irc.Event, []string) ([]string, string)
}

func (cmd *botCommand) process(e *irc.Event) ([]string, string) {
    return cmd.handler(e, cmd.matchino.FindStringSubmatch(e.Message()))
}

func init() {
    /* uuido, _ := uuid.NewV4()
    mario := noembedRequest("http://it.wikipedia.org/wiki/Speciale:PaginaCasuale#" + uuido.String())

    fmt.Printf("%#v\n", sanitize.HTML(doc.Find("p").Text())) */
}
