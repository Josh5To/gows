package components

import (
	"fmt"
	"html/template"
	"math/rand"
	"strings"
)

const (
	H1 Heading = "h1"
	H2 Heading = "h2"
	H3 Heading = "h3"
	H4 Heading = "h4"
	H5 Heading = "h5"
	H6 Heading = "h6"
)

const HeadingTemplate = `{{if ne .HeadingContent ""}}{{printHeading}}{{end}}
<p>{{.ParagraphContent}}</p>
`

type Heading string

type TextBox struct {
	HeadingSize      Heading
	HeadingContent   string
	ParagraphContent string

	parsedHeading template.HTML
}

func NewTextBox(tb TextBox) (string, error) {
	//TODO: Find a better way to make these unique?
	uid := rand.Int()

	tbFuncMap := template.FuncMap{
		"printHeading": func() string {
			return fmt.Sprintf("<%s>%s</%s>", tb.HeadingSize, tb.HeadingContent, tb.HeadingSize)
		},
	}

	sb := strings.Builder{}
	tbTemp, err := template.New(fmt.Sprintf("%s-%d", "text-box", uid)).Funcs(tbFuncMap).Parse(HeadingTemplate)
	if err != nil {
		return "", err
	}

	if err := tbTemp.Execute(&sb, tb); err != nil {
		return "", err
	}

	return sb.String(), nil
}
