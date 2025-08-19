package handler

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    fmt.Fprint(w, `
    <html>
    <head>
    <link rel="stylesheet" href="/static/style.css">
    <title>Go MultiTool</title>
    </head>
    <body>
        <h1>Go MultiTool</h1>
        <ul>
            <li><a href="/calculator">Calculator</a></li>
            <li><a href="/games">Guess Number Game</a></li>
            <li><a href="/textutils">Text Utilities</a></li>
            <li><a href="/textformatter">Text Formatter</a></li>
            <li><a href="/unitconverter">Unit Converter</a></li>
            <li><a href="/quotes">Random Quote Generator</a></li>
            <li><a href="/bmi">BMI Calculator</a></li>
            <li><a href="/agecalc">Age Calculator</a></li>
            <li><a href="/palindrome">Palindrome Checker</a></li>
            <li><a href="/prime">Prime Checker</a></li>
            <li><a href="/passwordgen">Password Generator</a></li>
            <li><a href="/qrcodegen">QR Code Generator</a></li>
        </ul>
    </body>
    </html>
    `)
}
