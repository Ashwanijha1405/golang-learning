package main

import (
	"fmt"
	"os"
	"text/template"
)

type EmailData struct {
	ReciepientName string
	SenderName     string
	Subject        string
	Body           string
	Items          []string // demo a loop
	UnreadCount    int
}

func main() {

	fmt.Println("-------------- Text template example ---------------")

	emailTemplate := `
Subject: {{ .Subject }}

{{.Body}}

{{if .Items}}
   Related Items:
{{range .Items}}
    - {{.}}
{{end}}
{{end}}

{{if gt .UnreadCount 0}}
You have {{.UnreadCount}} unreads.
{{else}}
You have no messages 
{{end}}


- Thanks
{{.SenderName}}
`

	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	data := EmailData{
		ReciepientName: "Alice",
		SenderName:     "Bob",
		Subject:        "Weekly Update",
		Body:           "Here are the items we worked on this week:",
		Items:          []string{"Task 1: Bug fixes", "Task 2: Feature development", "Task 3: Code review"},
		UnreadCount:    2,
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
