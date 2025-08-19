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
        case "+":
            result = a + b
        case "-":
            result = a - b
        case "*":
            result = a * b
        case "/":
            result = a / b
        default:
            result = 0
        }
        fmt.Fprintf(w, `
        <html>
        <head><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet"><title>Calculator</title></head>
        <body class="bg-light">
            <div class="container mt-5">
                <div class="card shadow p-4">
                    <h2 class="text-center mb-3">Calculator</h2>
                    <div class="alert alert-success text-center">Result: %.2f</div>
                    <a href="/calculator" class="btn btn-primary">Back</a>
                    <a href="/" class="btn btn-secondary">Menu</a>
                </div>
            </div>
        </body>
        </html>
        `, result)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet"><title>Calculator</title></head>
    <body class="bg-light">
        <div class="container mt-5">
            <div class="card shadow p-4">
                <h2 class="text-center mb-3">Calculator</h2>
                <form method="POST" class="row g-3">
                    <div class="col-md-5">
                        <input name="a" type="number" step="any" class="form-control" placeholder="First number" required>
                    </div>
                    <div class="col-md-2">
                        <select name="op" class="form-select">
                            <option value="+">+</option>
                            <option value="-">-</option>
                            <option value="*">*</option>
                            <option value="/">/</option>
                        </select>
                    </div>
                    <div class="col-md-5">
                        <input name="b" type="number" step="any" class="form-control" placeholder="Second number" required>
                    </div>
                    <div class="col-12 text-center">
                        <button type="submit" class="btn btn-success">Calculate</button>
                    </div>
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
