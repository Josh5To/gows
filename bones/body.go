package bones

import (
	"html"
	"html/template"
)

type Body struct {
	Divs []string
}

func (b *Body) AddFuncMap(fm template.FuncMap) {
	fm["body_unescapeDiv"] = func(s string) template.HTML {
		return template.HTML(html.UnescapeString(s))
	}
}
