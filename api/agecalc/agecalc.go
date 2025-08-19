package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {
		birthdate := r.FormValue("birthdate")

		// pakai layout yang fix sesuai HTML date input (YYYY-MM-DD)
		layout := "2006-01-02"
		bd, err := time.Parse(layout, birthdate)
		if err != nil {
			fmt.Fprintf(w, "<p style='color:red'>Invalid date format. Use YYYY-MM-DD.</p><a href='/agecalc'>Back</a>")
			return
		}

		today := time.Now()
		age := today.Year() - bd.Year()
		if today.YearDay() < bd.YearDay() {
			age--
		}

		fmt.Fprintf(w, `
			<html>
			<head><title>Age Calculator</title></head>
			<body>
				<h1>Age Calculator</h1>
				<p>Birthdate: %s</p>
				<p>Your age: %d years</p>
				<a href="/agecalc">Back</a>
			</body>
			</html>
		`, birthdate, age)
		return
	}

	fmt.Fprint(w, `
	<html>
	<head><title>Age Calculator</title></head>
	<body>
		<h1>Age Calculator</h1>
		<form method="POST">
			<input type="date" name="birthdate" required>
			<button type="submit">Calculate</button>
		</form>
		<a href="/">Back to menu</a>
	</body>
	</html>
	`)
}
