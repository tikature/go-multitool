package handler

import (
    "bytes"
    "encoding/base64"
    "fmt"
    "image/color"
    "image/png"
    "net/http"

    "github.com/skip2/go-qrcode"
)

func QRCodeHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        text := r.FormValue("text")
        fgColor := r.FormValue("fg")
        bgColor := r.FormValue("bg")

        qr, err := qrcode.New(text, qrcode.Medium)
        if err != nil {
            http.Error(w, "Failed to generate QR", http.StatusInternalServerError)
            return
        }

        fg := parseHexColor(fgColor, color.Black)
        bg := parseHexColor(bgColor, color.White)
        qr.BackgroundColor = bg
        qr.ForegroundColor = fg

        var buf bytes.Buffer
        _ = png.Encode(&buf, qr.Image(200))

        base64Img := base64.StdEncoding.EncodeToString(buf.Bytes())

        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, `
        <html>
        <head>
            <title>QR Code Generator</title>
            <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
        </head>
        <body class="bg-light">
        <div class="container py-5">
            <div class="card shadow p-4 text-center">
                <h1 class="mb-4">QR Code Result</h1>
                <p><b>Input:</b> %s</p>
                <img src="data:image/png;base64,%s" class="img-thumbnail mb-3" style="max-width:200px;"/>
                <div class="mt-3">
                    <a href="/qrcodegen" class="btn btn-primary me-2">Generate Again</a>
                    <a href="/" class="btn btn-secondary">Menu</a>
                </div>
            </div>
        </div>
        </body>
        </html>
        `, text, base64Img, text, fgColor, bgColor)
        return
    }

    fmt.Fprint(w, `
    <html>
    <head>
        <title>QR Code Generator</title>
        <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.3/dist/css/bootstrap.min.css" rel="stylesheet">
    </head>
    <body class="bg-light">
        <div class="container py-5">
            <div class="card shadow p-4">
                <h1 class="mb-4">QR Code Generator</h1>
                <form method="POST">
                    <div class="mb-3">
                        <label class="form-label">Text / URL</label>
                        <input type="text" name="text" class="form-control" placeholder="Enter text or link" required>
                    </div>
                    <div class="row">
                        <div class="col-md-6 mb-3">
                            <label class="form-label">Foreground</label>
                            <input type="color" name="fg" class="form-control form-control-color" value="#000000">
                        </div>
                        <div class="col-md-6 mb-3">
                            <label class="form-label">Background</label>
                            <input type="color" name="bg" class="form-control form-control-color" value="#ffffff">
                        </div>
                    </div>
                    <button type="submit" class="btn btn-success">Generate</button>
                </form>
                <a href="/" class="btn btn-secondary mt-3">Back to Menu</a>
            </div>
        </div>
    </body>
    </html>
    `)
}

// endpoint untuk download
func QRDownloadHandler(w http.ResponseWriter, r *http.Request) {
    text := r.URL.Query().Get("text")
    fgColor := r.URL.Query().Get("fg")
    bgColor := r.URL.Query().Get("bg")

    qr, err := qrcode.New(text, qrcode.Medium)
    if err != nil {
        http.Error(w, "Failed to generate QR", http.StatusInternalServerError)
        return
    }

    fg := parseHexColor(fgColor, color.Black)
    bg := parseHexColor(bgColor, color.White)
    qr.BackgroundColor = bg
    qr.ForegroundColor = fg

    var buf bytes.Buffer
    _ = png.Encode(&buf, qr.Image(300)) // versi download lebih besar

    w.Header().Set("Content-Type", "application/octet-stream")
    w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")
    w.Write(buf.Bytes())
}

func parseHexColor(s string, fallback color.Color) color.Color {
    var r, g, b uint8
    if len(s) == 7 && s[0] == '#' {
        fmt.Sscanf(s[1:], "%02x%02x%02x", &r, &g, &b)
        return color.RGBA{R: r, G: g, B: b, A: 255}
    }
    return fallback
}
