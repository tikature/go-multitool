package handler

import (
	"fmt"
	"net/http"
	"strings"
)

func TextFormatterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	if r.Method == http.MethodPost {
		text := r.FormValue("text")
		action := r.FormValue("action")
		var result string

		switch action {
		case "upper":
			result = strings.ToUpper(text)
		case "lower":
			result = strings.ToLower(text)
		case "count":
			result = fmt.Sprintf("Word count: %d", len(strings.Fields(text)))
		case "nospace":
			result = strings.ReplaceAll(text, " ", "")
		default:
			result = text
		}

		fmt.Fprintf(w, `
		<html>
		<head>
			<title>Text Formatter</title>
			<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
		</head>
		<body class="bg-light">
			<div class="container py-5">
				<div class="card shadow p-4">
					<h1 class="mb-4 text-primary">Text Formatter</h1>
					<div class="alert alert-info"><b>Result:</b> %s</div>
					<a href="/textformatter" class="btn btn-outline-primary">Back</a>
					<a href="/" class="btn btn-secondary">Menu</a>
				</div>
			</div>
		</body>
		</html>
		`, result)
		return
	}

	fmt.Fprint(w, `
	<html>
	<head>
		<title>Text Formatter</title>
		<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
	</head>
	<body class="bg-light">
		<div class="container py-5">
			<div class="card shadow p-4">
				<h1 class="mb-4 text-primary">Text Formatter</h1>
				<form method="POST" class="mb-3">
					<div class="mb-3">
						<textarea class="form-control" name="text" rows="4" placeholder="Enter text"></textarea>
					</div>
					<div class="mb-3">
						<select class="form-select" name="action">
							<option value="upper">Uppercase</option>
							<option value="lower">Lowercase</option>
							<option value="count">Word Count</option>
							<option value="nospace">Remove Spaces</option>
						</select>
					</div>
					<button type="submit" class="btn btn-primary">Format</button>
				</form>
				<a href="/" class="btn btn-secondary">Back to menu</a>
			</div>
		</div>
	</body>
	</html>
	`)
}
