package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8") // UTF-8 untuk emoji
    fmt.Fprint(w, `
    <html>
    <head>	
    <meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
        <title>Go MultiTool</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
        <style>
            body {
                background: linear-gradient(to right, #f8f9fa, #fff0f5);
            }
            .tool-btn {
                min-width: 200px;
                margin: 8px;
            }
            footer {
                background-color: #ffb6c1; /* Pink blossom */
                color: #333;
                font-weight: 500;
                letter-spacing: 1px;
            }
        </style>
    </head>
    <body>
        <div class="container mt-5 mb-5">
            <div class="card shadow-lg p-4">
                <h1 class="text-center mb-4 text-primary"> Go MultiTool 🌟 </h1>
                <div class="d-flex flex-wrap justify-content-center">
                    <a href="/calculator" class="btn btn-outline-primary tool-btn">🧮 Calculator</a>
                    <a href="/games" class="btn btn-outline-success tool-btn">🎮 Guess Number</a>
                    <a href="/textformatter" class="btn btn-outline-warning tool-btn">📝 Text Formatter</a>
                    <a href="/unitconverter" class="btn btn-outline-info tool-btn">📏 Unit Converter</a>
                    <a href="/bmi" class="btn btn-outline-danger tool-btn">⚖️ BMI Calculator</a>
                    <a href="/agecalc" class="btn btn-outline-secondary tool-btn">📅 Age Calculator</a>
                    <a href="/palindrome" class="btn btn-outline-dark tool-btn">🔄 Palindrome Checker</a>
                    <a href="/prime" class="btn btn-outline-primary tool-btn">🔢 Prime Checker</a>
                    <a href="/passwordgen" class="btn btn-outline-success tool-btn">🔑 Password Generator</a>
                    <a href="/qrcodegen" class="btn btn-outline-warning tool-btn">📷 QR Code Generator</a>
                </div>
            </div>
        </div>

        <!-- Footer -->
        <footer class="text-center py-3 fixed-bottom shadow">
            <p class="mb-0">Made with ❤️ by <a href="github.com/tikature" target="_blank">Tikature</a></p>
        </footer>
    </body>
    </html>
    `)
}
