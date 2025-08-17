package handler

import (
    "fmt"
    "net/http"
    "strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.Method == http.MethodPost {
        a, _ := strconv.ParseFloat(r.FormValue("a"), 64)
        b, _ := strconv.ParseFloat(r.FormValue("b"), 64)
        op := r.FormValue("op")
        var result float64
        switch op {
        case "+": result = a + b
        case "-": result = a - b
        case "*": result = a * b
        case "/": result = a / b
        default: result = 0
        }
        fmt.Fprintf(w, `<p>Result: %.2f</p><a href="/calculator">Back</a>`, result)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><link rel="stylesheet" href="/static/style.css"><title>Calculator</title></head>
    <body>
        <h1>Calculator</h1>
        <form method="POST">
            <input name="a" type="number" step="any" placeholder="First number" required>
            <select name="op">
                <option value="+">+</option>
                <option value="-">-</option>
                <option value="*">*</option>
                <option value="/">/</option>
            </select>
            <input name="b" type="number" step="any" placeholder="Second number" required>
            <button type="submit">Calculate</button>
        </form>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `)
}
