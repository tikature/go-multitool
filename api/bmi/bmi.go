package handler

import (
	"fmt"
	"net/http"
	"strconv"
)

func BMIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {
		weight, _ := strconv.ParseFloat(r.FormValue("weight"), 64)
		height, _ := strconv.ParseFloat(r.FormValue("height"), 64)
		bmi := weight / (height * height)

		category := ""
		switch {
		case bmi < 18.5:
			category = "Underweight"
		case bmi < 24.9:
			category = "Normal"
		case bmi < 29.9:
			category = "Overweight"
		default:
			category = "Obese"
		}

		fmt.Fprintf(w, `
		<html>
		<head>
			<meta charset="UTF-8">
        	<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>BMI Calculator</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
		</head>
		<body class="bg-light">
			<div class="container py-5">
				<div class="card shadow p-4">
					<h1 class="mb-4 text-success">BMI Calculator</h1>
					<div class="alert alert-info">
						<b>BMI:</b> %.2f <br>
						<b>Category:</b> %s
					</div>
					<a href="/bmi" class="btn btn-outline-success">Back</a>
					<a href="/" class="btn btn-secondary">Menu</a>
					<div class="d-flex justify-content-center gap-3 mt-3">
                        <a href="/bmi" class="btn btn-primary px-4">Back</a>
                        <a href="/" class="btn btn-secondary px-4">Menu</a>
                    </div>
				</div>
			</div>
		</body>
		</html>
		`, bmi, category)
		return
	}

	fmt.Fprint(w, `
	<html>
	<head>
		<meta charset="UTF-8">
        <meta name="viewport" content="width=device-width, initial-scale=1.0">
		<title>BMI Calculator</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body class="bg-light">
		<div class="container py-5">
			<div class="card shadow p-4">
				<h1 class="mb-4 text-success">BMI Calculator</h1>
				<form method="POST" class="mb-3">
					<div class="mb-3">
						<input class="form-control" name="weight" type="number" step="any" placeholder="Weight (kg) ex: 52" required>
					</div>
					<div class="mb-3">
						<input class="form-control" name="height" type="number" step="any" placeholder="Height (m) ex: 1.53" required>
					</div>
					<button type="submit" class="btn btn-success">Calculate</button>
				</form>
				<a href="/" class="btn btn-secondary">Back to menu</a>
			</div>
		</div>
	</body>
	</html>
	`)
}
