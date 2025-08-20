package handler

import (
    "fmt"
    "net/http"
    "strings"
)

func PalindromeHandler(w http.ResponseWriter, r *http.Request) {
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

        fmt.Fprintf(w, `
        <html>
        <head>
        	<meta charset="UTF-8">
            <meta name="viewport" content="width=device-width, initial-scale=1.0">
            <title>Palindrome Checker</title>
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
        </head>
        <body class="bg-light">
        <div class="container py-5">
            <div class="card shadow p-4">
                <h1 class="mb-3">Palindrome Checker</h1>
                <div class="alert alert-info">Result: %s</div>
                <a href="/palindrome" class="btn btn-primary">Check Again</a><br>
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
    <head>
    	<meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Palindrome Checker</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
    </head>
    <body class="bg-light">
        <div class="container py-5">
            <div class="card shadow p-4">
                <h1 class="mb-4">Palindrome Checker</h1>
                <form method="POST">
                    <div class="mb-3">
                        <label class="form-label">Enter text</label>
                        <input name="text" type="text" class="form-control" placeholder="Enter text" required>
                    </div>
                    <button type="submit" class="btn btn-success">Check</button>
                </form>
                <a href="/" class="btn btn-secondary mt-3">Back to Menu</a>
            </div>
        </div>
    </body>
    </html>
    `)
}
