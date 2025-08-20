package handler

import (
	"fmt"
	"net/http"
	"github.com/skip2/go-qrcode"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	switch path {
	case "/qrcodegen":
		qrCodeHandler(w, r)
	case "/qrcodegen/image":
		qrCodeImageHandler(w, r)
	default:
		homeHandler(w, r)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
	<html>
	<head>
		<meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Go MultiTool</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body>
		<div class="container mt-5 mb-5">
			<div class="card shadow-lg p-4">
				<h1 class="text-center mb-4 text-primary"> Go MultiTool 🌟 </h1>
				<div class="d-flex flex-wrap justify-content-center">
					<a href="/qrcodegen" class="btn btn-outline-warning tool-btn">📷 QR Code Generator</a>
				</div>
			</div>
		</div>
	</body>
	</html>
	`)
}

func qrCodeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		if r.FormValue("download") != "" {
			png, _ := qrcode.Encode(text, qrcode.Medium, 256)
			w.Header().Set("Content-Type", "application/octet-stream")
			w.Header().Set("Content-Disposition", "attachment; filename=qrcode.png")
			w.Write(png)
			return
		}
		http.Redirect(w, r, "/qrcodegen?preview="+text, http.StatusSeeOther)
		return
	}

	preview := r.URL.Query().Get("preview")

	fmt.Fprint(w, `
	<html>
	<head>
		<meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>QR Code Generator</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.2/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body class="bg-light">
		<div class="container mt-5">
			<div class="card shadow p-4">
				<h2 class="text-center mb-4 text-success">📷 QR Code Generator</h2>
				<form method="POST" class="d-flex gap-2">
					<input type="text" name="text" class="form-control" placeholder="Masukkan teks atau URL" required>
					<button type="submit" class="btn btn-primary">Generate</button>
				</form>
	`)

	if preview != "" {
		fmt.Fprintf(w, `
			<div class="text-center mt-4">
				<img src="/qrcodegen/image?text=%s" class="img-thumbnail" alt="QR Code">
				<form method="POST" class="mt-3">
					<input type="hidden" name="text" value="%s">
					<button type="submit" name="download" value="1" class="btn btn-success">⬇️ Download</button>
				</form>
			</div>
		`, preview, preview)
	}

	fmt.Fprint(w, `
			</div>
		</div>
	</body>
	</html>
	`)
}

func qrCodeImageHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	png, _ := qrcode.Encode(text, qrcode.Medium, 256)
	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
