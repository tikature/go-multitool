package handler

import (
	"fmt"
	"net/http"
	"time"
)

func AgeCalcHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {
		birthdate := r.FormValue("birthdate")
		layout := "2006-01-02"
		bd, err := time.Parse(layout, birthdate)
		if err != nil {
			fmt.Fprintf(w, `<div class="alert alert-danger">Invalid date format. Use YYYY-MM-DD.</div><a href='/agecalc'>Back</a>`)
			return
		}

		today := time.Now()
		age := today.Year() - bd.Year()
		if today.YearDay() < bd.YearDay() {
			age--
		}

		fmt.Fprintf(w, `
		<html>
		<head>
			<title>Age Calculator</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
		</head>
		<body class="bg-light">
			<div class="container py-5">
				<div class="card shadow p-4">
					<h1 class="mb-4 text-warning">Age Calculator</h1>
					<p><b>Birthdate:</b> %s</p>
					<div class="alert alert-info"><b>Your age:</b> %d years</div>
					<a href="/agecalc" class="btn btn-outline-warning">Back</a>
					<a href="/" class="btn btn-secondary">Menu</a>
				</div>
			</div>
		</body>
		</html>
		`, birthdate, age)
		return
	}

	fmt.Fprint(w, `
	<html>
	<head>
		<title>Age Calculator</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body class="bg-light">
		<div class="container py-5">
			<div class="card shadow p-4">
				<h1 class="mb-4 text-warning">Age Calculator</h1>
				<form method="POST" class="mb-3">
					<div class="mb-3">
						<input type="date" class="form-control" name="birthdate" required>
					</div>
					<button type="submit" class="btn btn-warning">Calculate</button>
				</form>
				<a href="/" class="btn btn-secondary">Back to menu</a>
			</div>
		</div>
	</body>
	</html>
	`)
}
