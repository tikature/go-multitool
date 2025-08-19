package handler

import (
    "fmt"
    "net/http"
    "strconv"
)

func isPrime(n int) bool {
    if n < 2 {
        return false
    }
    for i := 2; i*i <= n; i++ {
        if n%i == 0 {
            return false
        }
    }
    return true
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.Method == http.MethodPost {
        n, _ := strconv.Atoi(r.FormValue("number"))
        result := "Not Prime"
        if isPrime(n) {
            result = "Prime"
        }
        fmt.Fprintf(w, `<p>Result: %d is %s</p><a href="/prime">Back</a>`, n, result)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><title>Prime Checker</title></head>
    <body>
        <h1>Prime Number Checker</h1>
        <form method="POST">
            <input name="number" type="number" placeholder="Enter number" required>
            <button type="submit">Check</button>
        </form>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `)
}
