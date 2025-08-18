package handler

import (
    "fmt"
    "net/http"
    "strconv"
)

var numbers = []string{
    "nol", "satu", "dua", "tiga", "empat",
    "lima", "enam", "tujuh", "delapan", "sembilan",
}

func numToWords(n int) string {
    if n < 10 {
        return numbers[n]
    }
    if n < 20 {
        return numbers[n-10] + " belas"
    }
    if n < 100 {
        return numbers[n/10] + " puluh " + numbers[n%10]
    }
    return "Angka terlalu besar"
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        num, _ := strconv.Atoi(r.FormValue("number"))
        result := numToWords(num)
        fmt.Fprintf(w, `<p>%d = %s</p><a href="/num2words">Back</a>`, num, result)
        return
    }

    fmt.Fprint(w, `
    <html><head><link rel="stylesheet" href="/static/style.css"><title>Number to Words</title></head>
    <body>
        <h1>Number to Words</h1>
        <form method="POST">
            <input name="number" type="number" placeholder="Enter number (0-99)" required>
            <button type="submit">Convert</button>
        </form>
        <a href="/">Back to menu</a>
    </body></html>
    `)
}
