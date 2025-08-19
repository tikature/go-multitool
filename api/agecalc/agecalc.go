package handler

import (
    "fmt"
    "net/http"
    "strconv"
    "time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.Method == http.MethodPost {
        day, _ := strconv.Atoi(r.FormValue("day"))
        month, _ := strconv.Atoi(r.FormValue("month"))
        year, _ := strconv.Atoi(r.FormValue("year"))

        dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
        today := time.Now()
        age := today.Year() - dob.Year()
        if today.YearDay() < dob.YearDay() {
            age--
        }

        fmt.Fprintf(w, `<p>Your Age: %d years</p><a href="/agecalc">Back</a>`, age)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><title>Age Calculator</title></head>
    <body>
        <h1>Age Calculator</h1>
        <form method="POST">
            <input name="day" type="number" placeholder="Day" required>
            <input name="month" type="number" placeholder="Month" required>
            <input name="year" type="number" placeholder="Year" required>
            <button type="submit">Calculate</button>
        </form>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `)
}
