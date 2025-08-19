package handler

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"net/http"

	"github.com/skip2/go-qrcode"
)

// Handler untuk QR generator
func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		style := r.FormValue("style")
		fgColor := r.FormValue("fg")
		bgColor := r.FormValue("bg")

		// Generate QR dengan warna default dulu
		qr, err := qrcode.New(text, qrcode.Medium)
		if err != nil {
			http.Error(w, "Failed to generate QR", http.StatusInternalServerError)
			return
		}

		// Parse warna dari input user
		fg := parseHexColor(fgColor, color.Black)
		bg := parseHexColor(bgColor, color.White)

		qr.BackgroundColor = bg
		qr.ForegroundColor = fg

		var buf bytes.Buffer
		err = png.Encode(&buf, qr.Image(256))
		if err != nil {
			http.Error(w, "Failed to encode QR", http.StatusInternalServerError)
			return
		}

		// Kalau style circle, konversi pixel hitam jadi bulatan
		var img image.Image
		if style == "circle" {
			img = toCircles(qr.Image(256), fg, bg)
		} else {
			img, _ = png.Decode(&buf)
		}

		// Encode ke base64
		var finalBuf bytes.Buffer
		png.Encode(&finalBuf, img)

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
		`, text, base64.StdEncoding.EncodeToString(finalBuf.Bytes()))
		return
	}

	// Form input
	fmt.Fprint(w, `
	<html>
	<head><title>QR Code Generator</title></head>
	<body>
		<h1>QR Code Generator</h1>
		<form method="POST">
			<input type="text" name="text" placeholder="Enter text or link" required><br><br>
			
			<label>Style:</label>
			<select name="style">
				<option value="square">Square</option>
				<option value="circle">Circle</option>
			</select><br><br>

			<label>Foreground:</label>
			<input type="color" name="fg" value="#000000"><br><br>

			<label>Background:</label>
			<input type="color" name="bg" value="#ffffff"><br><br>

			<button type="submit">Generate</button>
		</form>
		<a href="/">Back to menu</a>
	</body>
	</html>
	`)
}

// Ubah QR pixel jadi bulatan
func toCircles(src image.Image, fg, bg color.Color) image.Image {
	bounds := src.Bounds()
	dest := image.NewRGBA(bounds)
	draw.Draw(dest, bounds, &image.Uniform{bg}, image.Point{}, draw.Src)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 4 {
		for x := bounds.Min.X; x < bounds.Max.X; x += 4 {
			if src.At(x, y) == color.Black {
				// Gambar lingkaran kecil (radius 2px)
				for dy := -2; dy <= 2; dy++ {
					for dx := -2; dx <= 2; dx++ {
						if dx*dx+dy*dy <= 4 {
							dest.Set(x+dx, y+dy, fg)
						}
					}
				}
			}
		}
	}
	return dest
}

// Convert hex → color.Color
func parseHexColor(s string, fallback color.Color) color.Color {
	var r, g, b uint8
	if len(s) == 7 && s[0] == '#' {
		fmt.Sscanf(s[1:], "%02x%02x%02x", &r, &g, &b)
		return color.RGBA{R: r, G: g, B: b, A: 255}
	}
	return fallback
}
