package data

import "strings"

// ManPage represents the main help page document
type ManPage struct {
	FirstName   string
	LastName    string
	Description string
	Years       int
	Synopsis    []Option
	Summary     []string
	Links       []ListItem
}

// Option represent the command line switches of the command
type Option struct {
	Name       string
	Parameters []Parameter
}

// Parameter defines the specifics of each option
type Parameter struct {
	Type    string
	Details []interface{}
}

// GetIndentClass ...
func (p *Parameter) GetIndentClass(d interface{}) string {
	switch d.(type) {
	case ListParameterDetail:
		return "List"
	default:
		return "Text"
	}
}

// ListParameterDetail ... a wrapper for a list parameter
type ListParameterDetail struct {
	Items []ListItem
}

// ListItem ... an item in a list that can represent a URL
type ListItem struct {
	Text string
	URL  string
}

// TextParameterDetail ... a text paragraph
type TextParameterDetail struct {
	Text string
}

// JoinParams joins the parameters together with a pipe "|"
func (o *Option) JoinParams() string {
	var types []string
	for _, p := range o.Parameters {
		types = append(types, p.Type)
	}
	return strings.Join(types, "|")
}

// GetManPage retrieves the primary help document to display on the page
func GetManPage() *ManPage {
	return &ManPage{
		"atom",
		"Data",
		"software engineer",
		6,
		[]Option{
			Option{
				"developer",
				[]Parameter{
					Parameter{
						"dotnet",
						[]interface{}{
							TextParameterDetail{
								"d quam pellentesque volutpat et vel tellus. Curabitur turpis dolor, efficitur in orci sit amet, ornare mollis tortor. Sed maximus quam nunc, ut molestie lacus placerat et. ",
							},
							ListParameterDetail{
								[]ListItem{
									ListItem{"Text", ""},
									ListItem{"Text URL", "http://test.com"},
								},
							},
						},
					},
					Parameter{
						"javascript",
						nil,
					},
					Parameter{
						"golang",
						nil,
					},
				},
			},
			Option{
				"mentor",
				nil,
			},
			Option{
				"architect",
				nil,
			},
			Option{
				"leader",
				nil,
			},
		},
		[]string{
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce auctor magna sed velit aliquet lacinia. Phasellus tincidunt libero turpis, nec commodo elit volutpat ac. Fusce non enim sed quam pellentesque volutpat et vel tellus. Curabitur turpis dolor, efficitur in orci sit amet, ornare mollis tortor. Sed maximus quam nunc, ut molestie lacus placerat et. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Sed nec ante eu nisi pharetra blandit sed ac ex.",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce auctor magna sed velit aliquet lacinia. Phasellus tincidunt libero turpis, nec commodo elit volutpat ac. Fusce non enim sed quam pellentesque volutpat et vel tellus. Curabitur turpis dolor, efficitur in orci sit amet, ornare mollis tortor. Sed maximus quam nunc, ut molestie lacus placerat et. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Sed nec ante eu nisi pharetra blandit sed ac ex.",
			"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Fusce auctor magna sed velit aliquet lacinia. Phasellus tincidunt libero turpis, nec commodo elit volutpat ac. Fusce non enim sed quam pellentesque volutpat et vel tellus. Curabitur turpis dolor, efficitur in orci sit amet, ornare mollis tortor. Sed maximus quam nunc, ut molestie lacus placerat et. Vestibulum ante ipsum primis in faucibus orci luctus et ultrices posuere cubilia Curae; Sed nec ante eu nisi pharetra blandit sed ac ex.",
		},
		[]ListItem{
			ListItem{"GitHub", "https://github.com/jeryanders"},
			ListItem{"Twitter", "https://twitter.com/atomdata"},
			ListItem{"LinkedIn", "https://www.linkedin.com/in/jesse-anderson-99469349/"},
		},
	}
}
