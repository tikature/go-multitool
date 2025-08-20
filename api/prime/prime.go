package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if r.Method == http.MethodPost {
		n, _ := strconv.Atoi(r.FormValue("number"))
		result := "❌ Not Prime"
		if isPrime(n) {
			result = "✅ Prime"
		}
		fmt.Fprintf(w, `
		<html>
		<head>
			<meta charset="UTF-8">
        	<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Prime Checker</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
		</head>
		<body class="bg-light">
			<div class="container mt-5">
				<div class="card shadow p-4">
					<h2 class="mb-3 text-success">🔢 Prime Number Checker</h2>
					<p class="alert alert-info">Result: <b>%d</b> is <b>%s</b></p>
					<a href="/prime" class="btn btn-outline-success">🔄 Try Again</a>
					<a href="/" class="btn btn-secondary ms-2">🏠 Menu</a>
				</div>
			</div>
		</body>
		</html>
		`, n, result)
		return
	}

	fmt.Fprint(w, `
	<html>
	<head>
		<meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>Prime Checker</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body class="bg-light">
		<div class="container mt-5">
			<div class="card shadow p-4">
				<h2 class="mb-3 text-success">🔢 Prime Number Checker</h2>
				<form method="POST" class="row g-3">
					<div class="col-12">
						<input name="number" type="number" class="form-control" placeholder="Enter number" required>
					</div>
					<div class="col-12">
						<button type="submit" class="btn btn-success">Check</button>
						<a href="/" class="btn btn-secondary ms-2">Back to Menu</a>
					</div>
				</form>
			</div>
		</div>
	</body>
	</html>
	`)
}
