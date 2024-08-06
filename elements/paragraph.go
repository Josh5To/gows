package elements

import "html/template"

var paragraph = ElementAnatomy{
	TagName: "p",
	Void:    false,
}

type Paragraph struct {
	Text string

	nestedElements []Element
}

func (p *Paragraph) nested() []Element {
	return nil
}

func (p *Paragraph) Nest(element ...Element) Element {
	p.nestedElements = append(p.nestedElements, element...)
	return p
}

func (p *Paragraph) AddFuncMap(funcMap *template.FuncMap) error {
	return nil
}

func (p *Paragraph) Attributes() AttributesMap {
	return nil
}

func (p *Paragraph) Content() []template.HTML {
	var result = []template.HTML{template.HTML(p.Text)}
	th, err := parseNested(p.nestedElements)
	if err != nil {
		return nil
	}
	return append(result, th...)
}

func (p *Paragraph) IsVoid() bool {
	return paragraph.Void
}

func (p *Paragraph) TagName() string {
	return paragraph.TagName
}
