package handler

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
        fmt.Fprintf(w, `
        <html>
        <head><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet"><title>Guess Number</title></head>
        <body class="bg-light">
            <div class="container mt-5">
                <div class="card shadow p-4 text-center">
                    <h2>Guess the Number</h2>
                    <div class="alert alert-info">%s</div>
                    <a href="/games" class="btn btn-primary">Try Again</a>
                    <a href="/" class="btn btn-secondary">Menu</a>
                </div>
            </div>
        </body>
        </html>
        `, msg)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet"><title>Guess Number</title></head>
    <body class="bg-light">
        <div class="container mt-5">
            <div class="card shadow p-4">
                <h2 class="text-center mb-3">Guess the Number (1-100)</h2>
                <form method="POST" class="text-center">
                    <input type="number" name="guess" class="form-control w-50 mx-auto mb-3" placeholder="Enter your guess" required>
                    <button type="submit" class="btn btn-success">Guess</button>
                </form>
                <div class="mt-3 text-center">
                    <a href="/" class="btn btn-secondary">Back to menu</a>
                </div>
            </div>
        </div>
    </body>
    </html>
    `)
}
