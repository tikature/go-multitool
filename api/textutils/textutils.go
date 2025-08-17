package handler

import (
    "fmt"
    "net/http"
    "strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    result := ""
    if r.Method == http.MethodPost {
        text := r.FormValue("text")
        action := r.FormValue("action")
        switch action {
        case "upper":
            result = strings.ToUpper(text)
        case "lower":
            result = strings.ToLower(text)
        case "reverse":
            runes := []rune(text)
            for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
                runes[i], runes[j] = runes[j], runes[i]
            }
            result = string(runes)
        }
    }

    fmt.Fprintf(w, `
    <html>
    <head><link rel="stylesheet" href="/static/style.css"><title>Text Utilities</title></head>
    <body>
        <h1>Text Utilities</h1>
        <form method="POST">
            <input type="text" name="text" placeholder="Enter text" required>
            <select name="action">
                <option value="upper">UPPERCASE</option>
                <option value="lower">lowercase</option>
                <option value="reverse">Reverse</option>
            </select>
            <button type="submit">Transform</button>
        </form>
        <p>%s</p>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `, result)
}
