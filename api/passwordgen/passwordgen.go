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
		birth := r.FormValue("birth")          // format yyyy-mm-dd
		name := strings.ToLower(r.FormValue("name"))
		birthClean := strings.ReplaceAll(birth, "-", "") // yyyymmdd

		// ambil 3 huruf pertama nama (kalau ada)
		namePart := ""
		if len(name) >= 3 {
			namePart = name[:3]
		} else {
			namePart = name
		}

		randomPart := generatePassword(8) // 8 random chars
		password := randomInsert(randomPart, birthClean+namePart)

		fmt.Fprintf(w, `
		<html>
		<head>
			<meta charset="UTF-8">
        	<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Password Generator</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
		</head>
		<body class="bg-light">
			<div class="container mt-5">
				<div class="card shadow p-4">
					<h2 class="mb-3 text-primary">🔐 Password Generator</h2>
					<p class="alert alert-success"><b>Generated Password:</b> %s</p>
					<a href="/passwordgen" class="btn btn-outline-primary">🔄 Generate Again</a>
					<a href="/" class="btn btn-secondary ms-2">🏠 Menu</a>
				</div>
			</div>
		</body>
		</html>
		`, password)
		return
	}

	fmt.Fprint(w, `
	<html>
	<head>
		<meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Password Generator</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body class="bg-light">
		<div class="container mt-5">
			<div class="card shadow p-4">
				<h2 class="mb-3 text-primary">🔐 Password Generator</h2>
				<form method="POST" class="row g-3">
					<div class="col-md-6">
						<label class="form-label">Birth Date:</label>
						<input name="birth" type="date" class="form-control" required>
					</div>
					<div class="col-md-6">
						<label class="form-label">Name:</label>
						<input name="name" type="text" class="form-control" placeholder="Enter your name" required>
					</div>
					<div class="col-12">
						<button type="submit" class="btn btn-primary">Generate Password</button>
						<a href="/" class="btn btn-secondary ms-2">Back to Menu</a>
					</div>
				</form>
			</div>
		</div>
	</body>
	</html>
	`)
}
