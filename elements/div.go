package elements

import (
	"html/template"
)

var div = ElementAnatomy{
	TagName: "div",
	Void:    false,
}

type Div struct {
	ClassName Attribute `html:"class"`

	nestedElements []Element
}

func (d Div) nested() []Element {
	return d.nestedElements
}

func (d Div) Nest(element ...Element) Element {
	d.nestedElements = append(d.nestedElements, element...)
	return d
}

func (d Div) AddFuncMap(funcMap *template.FuncMap) error {
	//TODO implement me
	panic("implement me")
}

func (d Div) Attributes() AttributesMap {
	return GetAttributes(d)
}

func (d Div) Content() []template.HTML {
	th, err := parseNested(d.nestedElements)
	if err != nil {
		return nil
	}
	return th
}

func (d Div) IsVoid() bool {
	return div.Void
}

func (d Div) TagName() string {
	return div.TagName
}
