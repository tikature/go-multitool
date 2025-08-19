package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8") // penting: UTF-8
    fmt.Fprint(w, `
    <html>
    <head>
        <meta charset="UTF-8">
        <title>Go MultiTool</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
    </head>
    <body class="bg-light">
        <div class="container mt-5 mb-5">
            <div class="card shadow p-4">
                <h1 class="text-center mb-4">Go MultiTool</h1>
                <ul class="list-group">
                    <li class="list-group-item"><a href="/calculator">🧮 Calculator</a></li>
                    <li class="list-group-item"><a href="/games">🎮 Guess Number Game</a></li>
                    <li class="list-group-item"><a href="/textformatter">📝 Text Formatter</a></li>
                    <li class="list-group-item"><a href="/unitconverter">📏 Unit Converter</a></li>
                    <li class="list-group-item"><a href="/bmi">⚖️ BMI Calculator</a></li>
                    <li class="list-group-item"><a href="/agecalc">📅 Age Calculator</a></li>
                    <li class="list-group-item"><a href="/palindrome">🔄 Palindrome Checker</a></li>
                    <li class="list-group-item"><a href="/prime">🔢 Prime Checker</a></li>
                    <li class="list-group-item"><a href="/passwordgen">🔑 Password Generator</a></li>
                    <li class="list-group-item"><a href="/qrcodegen">📷 QR Code Generator</a></li>
                </ul>
            </div>
        </div>

        <!-- Footer tetap -->
        <footer class="bg-dark text-white text-center py-3 fixed-bottom">
            by hilda
        </footer>
    </body>
    </html>
    `)
}
