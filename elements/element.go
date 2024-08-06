package elements

import (
	"fmt"
	"html"
	"html/template"
	"math/rand/v2"
	"reflect"
	"strings"
)

const (
	StructTagName = "gows"
)

type Attribute string

type AttributesMap map[string]string

type ElementAnatomy struct {
	TagName string
	Void    bool
}

type Element interface {
	nested() []Element
	Nest(...Element) Element
	AddFuncMap(*template.FuncMap) error
	Attributes() AttributesMap
	Content() []template.HTML
	IsVoid() bool
	TagName() string
}

func GetAttributes(el Element) AttributesMap {
	//TODO: Make this be able to handle pointers and non-pointers
	values := reflect.ValueOf(el)
	if !values.CanAddr() {
		values = reflect.Indirect(values)
	}

	var attrs = make(AttributesMap)
	for i, v := range reflect.VisibleFields(values.Type()) {
		if v.Type == reflect.TypeOf(Attribute("")) {
			if values.Field(i).String() != "" {
				// Set our field from struct tag if found.
				if val, set := v.Tag.Lookup("html"); set {
					attrs[val] = values.Field(i).String()
					continue
				}

				// Fallback, use lowercase field name as key
				attrs[strings.ToLower(v.Name)] = values.Field(i).String()
			}
		}
	}

	return attrs
}

func Parse(el Element) (template.HTML, error) {
	//TODO: Got to be a better way to make the templates unique
	uid := rand.Int()
	var elementTemplate = `<{{.TagName}}{{ range $akey, $aval := .Attributes}} {{$akey -}}="{{$aval}}"{{- end -}}
{{if .IsVoid}}/>{{else}}>{{range .Content}}{{.}}{{end}}</{{- .TagName}}>{{end}}`

	t, e := template.New(fmt.Sprintf("%s-%10d", el.TagName(), uid)).Parse(elementTemplate)
	if e != nil {
		return "", e
	}

	sb := new(strings.Builder)
	if err := t.Execute(sb, el); err != nil {
		return "", err
	}
	fmt.Println(template.HTML(html.UnescapeString(sb.String())))
	return template.HTML(html.UnescapeString(sb.String())), nil
}

func parseNested(els []Element) ([]template.HTML, error) {
	var NestedElements []template.HTML
	for _, e := range els {
		ph, err := Parse(e)
		if err != nil {
			fmt.Printf("Error parsing nested %s content: %v\n", e.TagName(), err)
			continue
		}
		NestedElements = append(NestedElements, ph)
	}
	return NestedElements, nil
}
