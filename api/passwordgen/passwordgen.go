package handler

import (
    "crypto/rand"
    "fmt"
    "math/big"
    "net/http"
    "strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"

func generatePassword(length int) string {
    result := ""
    for i := 0; i < length; i++ {
        n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
        result += string(charset[n.Int64()])
    }
    return result
}

func randomInsert(base, insert string) string {
    posN, _ := rand.Int(rand.Reader, big.NewInt(int64(len(base)+1)))
    pos := int(posN.Int64())
    return base[:pos] + insert + base[pos:]
}

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.Method == http.MethodPost {
        birth := r.FormValue("birth") // format yyyy-mm-dd
        birthClean := strings.ReplaceAll(birth, "-", "") // yyyymmdd

        randomPart := generatePassword(8) // 8 random chars
        password := randomInsert(randomPart, birthClean)

        fmt.Fprintf(w, `<p>Generated Password: %s</p><a href="/passwordgen">Back</a>`, password)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><title>Password Generator</title></head>
    <body>
        <h1>Password Generator</h1>
        <form method="POST">
            <label>Birth Date:</label>
            <input name="birth" type="date" required>
            <button type="submit">Generate</button>
        </form>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `)
}
