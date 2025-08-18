package handler

import (
    "fmt"
    "net/http"
    "strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        text := r.FormValue("text")
        action := r.FormValue("action")
        var result string

        switch action {
        case "upper":
            result = strings.ToUpper(text)
        case "lower":
            result = strings.ToLower(text)
        case "count":
            result = fmt.Sprintf("Word count: %d", len(strings.Fields(text)))
        case "nospace":
            result = strings.ReplaceAll(text, " ", "")
        default:
            result = text
        }

        fmt.Fprintf(w, `<p>Result: %s</p><a href="/textformatter">Back</a>`, result)
        return
    }

    fmt.Fprint(w, `
    <html><head><link rel="stylesheet" href="/static/style.css"><title>Text Formatter</title></head>
    <body>
        <h1>Text Formatter</h1>
        <form method="POST">
            <textarea name="text" rows="4" cols="50" placeholder="Enter text"></textarea><br>
            <select name="action">
                <option value="upper">Uppercase</option>
                <option value="lower">Lowercase</option>
                <option value="count">Word Count</option>
                <option value="nospace">Remove Spaces</option>
            </select>
            <button type="submit">Format</button>
        </form>
        <a href="/">Back to menu</a>
    </body></html>
    `)
}
