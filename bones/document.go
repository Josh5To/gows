package bones

import (
	_ "embed"
	"html/template"
	"net/http"
)

const templateDir = "templates"

var (
	//go:embed templates/base.tmpl
	baseTempString string
	//go:embed templates/body.tmpl
	bodyTempString string
	//go:embed templates/footer.tmpl
	footerTempString string
	//go:embed templates/head.tmpl
	headTempString string
	//go:embed templates/header.tmpl
	headerTempString string
)

type Page struct {
	Name string
	Lang string

	Head     Head
	Header   Header
	Body     Body
	Footer   Footer
	funcMap  template.FuncMap
	template *template.Template
}

func CreatePage(pageData *Page) error {
	if pageData.funcMap == nil {
		pageData.funcMap = template.FuncMap{}
	}
	//docBase := filepath.Join(templateDir, "base.tmpl")
	//docHead := filepath.Join(templateDir, "head.tmpl")
	//docBody := filepath.Join(templateDir, "body.tmpl")
	// Load our funcMap from various sections
	pageData.Body.AddFuncMap(pageData.funcMap)

	temp, err := template.New("html-doc").Funcs(pageData.funcMap).Parse(baseTempString)
	if err != nil {
		return err
	}

	temp, err = temp.Parse(headTempString)
	if err != nil {
		return err
	}

	temp, err = temp.Parse(headerTempString)
	if err != nil {
		return err
	}

	temp, err = temp.Parse(bodyTempString)
	if err != nil {
		return err
	}

	temp, err = temp.Parse(footerTempString)
	if err != nil {
		return err
	}

	pageData.template = temp

	return nil
}

func (p *Page) HandlerFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := p.template.ExecuteTemplate(w, "base", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
