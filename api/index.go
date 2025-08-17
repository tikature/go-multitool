package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, `
    <html>
    <head>
    <link rel="stylesheet" href="/static/style.css">
    <title>Go MultiTool</title>
    </head>
    <body>
        <h1>Go MultiTool</h1>
        <ul>
            <li><a href="/calculator">Calculator</a></li>
            <li><a href="/games">Guess Number Game</a></li>
            <li><a href="/textutils">Text Utilities</a></li>
        </ul>
    </body>
    </html>
    `)
}
