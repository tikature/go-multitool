package handler

import (
    "fmt"
    "net/http"

    "github.com/skip2/go-qrcode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")

    if r.Method == http.MethodPost {
        text := r.FormValue("text")
        png, err := qrcode.Encode(text, qrcode.Medium, 256)
        if err != nil {
            http.Error(w, "Failed to generate QR", http.StatusInternalServerError)
            return
        }

        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, `
        <html>
        <head><title>QR Code Generator</title></head>
        <body>
            <h1>QR Code Result</h1>
            <p>Input: %s</p>
            <img src="data:image/png;base64,%s"/>
            <br><a href="/qrcodegen">Back</a>
        </body>
        </html>
        `, text, encodeBase64(png))
        return
    }

    fmt.Fprint(w, `
    <html>
    <head><title>QR Code Generator</title></head>
    <body>
        <h1>QR Code Generator</h1>
        <form method="POST">
            <input type="text" name="text" placeholder="Enter text or link" required>
            <button type="submit">Generate</button>
        </form>
        <a href="/">Back to menu</a>
    </body>
    </html>
    `)
}

// helper: encode bytes → base64 string
func encodeBase64(data []byte) string {
    return fmt.Sprintf("%s", encodeToBase64(data))
}

func encodeToBase64(data []byte) string {
    const base64Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
    var result string
    n := len(data)
    for i := 0; i < n; i += 3 {
        var v uint32
        remain := n - i
        switch {
        case remain >= 3:
            v = uint32(data[i])<<16 | uint32(data[i+1])<<8 | uint32(data[i+2])
            result += string(base64Table[v>>18&0x3F]) +
                string(base64Table[v>>12&0x3F]) +
                string(base64Table[v>>6&0x3F]) +
                string(base64Table[v&0x3F])
        case remain == 2:
            v = uint32(data[i])<<16 | uint32(data[i+1])<<8
            result += string(base64Table[v>>18&0x3F]) +
                string(base64Table[v>>12&0x3F]) +
                string(base64Table[v>>6&0x3F]) + "="
        case remain == 1:
            v = uint32(data[i]) << 16
            result += string(base64Table[v>>18&0x3F]) +
                string(base64Table[v>>12&0x3F]) + "=="
        }
    }
    return result
}
