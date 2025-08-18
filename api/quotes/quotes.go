package handler

import (
    "fmt"
    "math/rand"
    "net/http"
    "time"
)

var quotes = []string{
    "Code is like humor. When you have to explain it, it’s bad.",
    "Fix the cause, not the symptom.",
    "Simplicity is the soul of efficiency.",
    "Before software can be reusable it first has to be usable.",
    "In Go we trust.",
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    rand.Seed(time.Now().UnixNano())
    q := quotes[rand.Intn(len(quotes))]

    fmt.Fprintf(w, `
    <html><head><link rel="stylesheet" href="/static/style.css"><title>Quotes</title></head>
    <body>
        <h1>Random Quote</h1>
        <p>%s</p>
        <a href="/quotes">Another one</a><br>
        <a href="/">Back to menu</a>
    </body></html>
    `, q)
}
