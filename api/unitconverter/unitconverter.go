package handler

import (
    "fmt"
    "net/http"
    "strconv"
)

func UnitConverterHandler(w http.ResponseWriter, r *http.Request) {
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

        fmt.Fprintf(w, `
        <html>
        <head>
            <title>Unit Converter</title>
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
        </head>
        <body class="bg-light">
        <div class="container py-5">
            <div class="card shadow p-4">
                <h1 class="mb-3">Unit Converter</h1>
                <div class="alert alert-success">%s</div>
                <a href="/unitconverter" class="btn btn-primary">Back</a>
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
        <title>Unit Converter</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
    </head>
    <body class="bg-light">
        <div class="container py-5">
            <div class="card shadow p-4">
                <h1 class="mb-4">Unit Converter</h1>
                <form method="POST">
                    <div class="mb-3">
                        <label class="form-label">Value</label>
                        <input name="value" type="number" step="any" class="form-control" placeholder="Enter value" required>
                    </div>
                    <div class="mb-3">
                        <label class="form-label">Conversion</label>
                        <select name="unit" class="form-select">
                            <option value="ctof">Celsius → Fahrenheit</option>
                            <option value="ftoc">Fahrenheit → Celsius</option>
                            <option value="mtokg">Gram → Kilogram</option>
                            <option value="ktom">Kilometer → Meter</option>
                        </select>
                    </div>
                    <button type="submit" class="btn btn-success">Convert</button>
                </form>
                <a href="/" class="btn btn-secondary mt-3">Back to Menu</a>
            </div>
        </div>
    </body>
    </html>
    `)
}
