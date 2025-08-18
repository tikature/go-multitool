package handler

import (
    "fmt"
    "net/http"
    "strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
        height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
        bmi := weight / (height * height)

        category := ""
        switch {
        case bmi < 18.5:
            category = "Underweight"
        case bmi < 24.9:
            category = "Normal"
        case bmi < 29.9:
            category = "Overweight"
        default:
            category = "Obese"
        }

        fmt.Fprintf(w, `<p>BMI: %.2f (%s)</p><a href="/bmi">Back</a>`, bmi, category)
        return
    }

    fmt.Fprint(w, `
    <html><head><link rel="stylesheet" href="/static/style.css"><title>BMI Calculator</title></head>
    <body>
        <h1>BMI Calculator</h1>
        <form method="POST">
            <input name="weight" type="number" step="any" placeholder="Weight (kg)" required>
            <input name="height" type="number" step="any" placeholder="Height (m)" required>
            <button type="submit">Calculate</button>
        </form>
        <a href="/">Back to menu</a>
    </body></html>
    `)
}
