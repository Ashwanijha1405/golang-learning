package main

import (
	"fmt"
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

{{if gt. UnreadCount 0}}
You have {{.UnreadCount}} unreads.
{{else}}
You have no messages 
{{end}}


- Thanks
{{.SenderName}}
`
}
