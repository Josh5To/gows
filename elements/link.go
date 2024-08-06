package elements

import (
	"html/template"
)

var link = ElementAnatomy{
	TagName: "l",
	Void:    true,
}

type Link struct {
	Href  Attribute `html:"href"`
	Rel   Attribute `html:"rel"`
	Title Attribute `html:"title"`
	Type  Attribute `html:"type"`
}

func (l *Link) Nest(elements ...Element) Element {
	//Cannot be nested
	return l
}

func (l *Link) nested() []Element {
	return nil
}

func (l *Link) AddFuncMap(funcMap *template.FuncMap) error {
	//TODO implement me
	panic("implement me")
}

func (l *Link) Attributes() AttributesMap {
	return GetAttributes(l)
}

func (l *Link) Content() []template.HTML {
	return []template.HTML{}
}

func (l *Link) IsVoid() bool {
	return link.Void
}

func (l *Link) TagName() string {
	return link.TagName
}
