package main

import (
	"html/template"
	"net/http"
)

type FormData struct {
	Name    string
	Email   string
	Message string
}

var tpl = `
<!DOCTYPE html>
<html>
<head>
	<title>Form Service</title>
</head>
<body>
	<h1>Contact Form</h1>
	<form method="POST" action="/submit">
		<label for="name">Name:</label><br>
		<input type="text" id="name" name="name"><br>
		<label for="email">Email:</label><br>
		<input type="email" id="email" name="email"><br>
		<label for="message">Message:</label><br>
		<textarea id="message" name="message" rows="4" cols="50"></textarea><br>
		<input type="submit" value="Submit">
	</form>

	{{if .}}
	<h2>Submitted Data:</h2>
	<p><strong>Name:</strong> {{.Name}}</p>
	<p><strong>Email:</strong> {{.Email}}</p>
	<p><strong>Message:</strong> {{.Message}}</p>
	{{end}}
</body>
</html>
`

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("form").Parse(tpl))
		tmpl.Execute(w, nil)
	})

	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Failed to parse form data", http.StatusBadRequest)
				return
			}

			formData := FormData{
				Name:    r.FormValue("name"),
				Email:   r.FormValue("email"),
				Message: r.FormValue("message"),
			}

			tmpl := template.Must(template.New("form").Parse(tpl))
			tmpl.Execute(w, formData)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.ListenAndServe(":8080", nil)
}
