package handler

import (
    "fmt"
    "net/http"
    "strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.Method == http.MethodPost {
        input := strings.ToLower(strings.ReplaceAll(r.FormValue("text"), " ", ""))
        rev := ""
        for i := len(input) - 1; i >= 0; i-- {
            rev += string(input[i])
        }
        result := "Not Palindrome"
        if input == rev {
            result = "Palindrome"
        }
        fmt.Fprintf(w, `<p>Result: %s</p><a href="/palindrome">Back</a>`, result)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><title>Palindrome Checker</title></head>
    <body>
        <h1>Palindrome Checker</h1>
        <form method="POST">
            <input name="text" type="text" placeholder="Enter text" required>
            <button type="submit">Check</button>
        </form>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `)
}
