package main

import (
    "fmt"
    "math/rand"
    "net/http"
    "strconv"
    "time"
)

var secret int

func init() {
    rand.Seed(time.Now().UnixNano())
    secret = rand.Intn(100) + 1
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        guess, _ := strconv.Atoi(r.FormValue("guess"))
        msg := ""
        if guess < secret {
            msg = "Too low!"
        } else if guess > secret {
            msg = "Too high!"
        } else {
            msg = "Correct! You win! New number generated."
            secret = rand.Intn(100) + 1
        }
        fmt.Fprintf(w, `<p>%s</p><a href="/games">Back</a>`, msg)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><link rel="stylesheet" href="/static/style.css"><title>Guess Number</title></head>
    <body>
        <h1>Guess the Number (1-100)</h1>
        <form method="POST">
            <input type="number" name="guess" required>
            <button type="submit">Guess</button>
        </form>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `)
}

func main() {}
