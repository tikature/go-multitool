package handler

import (
    "fmt"
    "net/http"
    "strconv"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        value, _ := strconv.ParseFloat(r.FormValue("value"), 64)
        unit := r.FormValue("unit")
        var result string

        switch unit {
        case "ctof":
            result = fmt.Sprintf("%.2f °C = %.2f °F", value, (value*9/5)+32)
        case "ftoc":
            result = fmt.Sprintf("%.2f °F = %.2f °C", value, (value-32)*5/9)
        case "mtokg":
            result = fmt.Sprintf("%.2f g = %.2f kg", value, value/1000)
        case "ktom":
            result = fmt.Sprintf("%.2f km = %.2f m", value, value*1000)
        default:
            result = "Unknown conversion"
        }

        fmt.Fprintf(w, `<p>%s</p><a href="/unitconverter">Back</a>`, result)
        return
    }

    fmt.Fprint(w, `
    <html><head><link rel="stylesheet" href="/static/style.css"><title>Unit Converter</title></head>
    <body>
        <h1>Unit Converter</h1>
        <form method="POST">
            <input name="value" type="number" step="any" placeholder="Enter value" required>
            <select name="unit">
                <option value="ctof">Celsius → Fahrenheit</option>
                <option value="ftoc">Fahrenheit → Celsius</option>
                <option value="mtokg">Gram → Kilogram</option>
                <option value="ktom">Kilometer → Meter</option>
            </select>
            <button type="submit">Convert</button>
        </form>
        <a href="/">Back to menu</a>
    </body></html>
    `)
}
