package elements

import (
	"html/template"
	"testing"
)

func TestParse(t *testing.T) {
	var nestedDiv = Div{
		ClassName: "nested-div",
	}.Nest(&Paragraph{
		Text: "YO!",
	})
	//nestedDiv

	type args struct {
		el Element
	}
	tests := []struct {
		name    string
		args    args
		want    template.HTML
		wantErr bool
	}{
		{
			name: "Verify correct parsing of a Link",
			args: args{
				el: &Link{
					Href: "test/style.css",
					Rel:  "stylesheet",
				},
			},
			want:    template.HTML(`<l href="test/style.css" rel="stylesheet"/>`),
			wantErr: false,
		},
		{
			name: "Verify correct parsing of a Div",
			args: args{
				el: nestedDiv,
			},
			want:    template.HTML(`<div class="nested-div"><p>YO!</p></div>`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.args.el)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
