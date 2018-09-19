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
	History     []Option
	Additional  []Option
	Links       []ListItem
}

// Option represent the command line switches of the command
// Details takes
// 1. ListParameterDetail for an enumerated list with items that can be
//    plain test or a URL
// 2. TextParameterDetail for a block of text to talk about how dope you are
type Option struct {
	Name       string
	URL        string
	CSS        string
	Leader     string
	Parameters []Parameter
	Details    []interface{}
}

// Parameter defines the specifics of each option
// Details takes
// 1. ListParameterDetail for an enumerated list with items that can be
//    plain test or a URL
// 2. TextParameterDetail for a block of text to talk about how dope you are
type Parameter struct {
	Type    string
	Details []interface{}
}

// ListParameterDetail ... a wrapper for a list parameter
type ListParameterDetail struct {
	Items []ListItem
}

// ListItem ... an item in a list that can represent a URL
type ListItem struct {
	Text        string
	URL         string
	Description string
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
				"",
				"",
				"--",
				[]Parameter{
					Parameter{
						"golang",
						[]interface{}{
							TextParameterDetail{
								"Reasons you love Go and want to see more people write in Go.",
							},
							ListParameterDetail{
								[]ListItem{
									ListItem{"Text", "", "Some nice text"},
									ListItem{"Text URL", "http://test.com", "an additional description about this item"},
								},
							},
						},
					},
					Parameter{
						"dotnet",
						[]interface{}{
							TextParameterDetail{
								"Reasons you came to like .NET and how it led you to Go.",
							},
							ListParameterDetail{
								[]ListItem{
									ListItem{"Text", "", ""},
									ListItem{"Text URL", "http://test.com", ""},
								},
							},
						},
					},
					Parameter{
						"javascript",
						[]interface{}{
							TextParameterDetail{
								"Reasons why you can't stop liking JavaScript and how new front-end frameworks make it difficult to hate JavaScript.",
							},
							ListParameterDetail{
								[]ListItem{
									ListItem{"Text", "", ""},
									ListItem{"Text URL", "http://test.com", ""},
								},
							},
						},
					},
				},
				nil,
			},
			Option{
				"option-1",
				"",
				"",
				"--",
				nil,
				nil,
			},
			Option{
				"option-2",
				"",
				"",
				"--",
				nil,
				nil,
			},
		},
		[]string{
			"Tell a story about yourself. Explain how dope your skills are and the awesome things you've done.",
			"Go ahead, make it multiple paragraphs.",
			"Just make it interesting. Don't bore people.",
		},
		[]Option{
			Option{
				"Excellent Company",
				"http://www.excellent.com",
				"bottomcut",
				"* ",
				nil,
				nil,
			},
		},
		[]Option{
			Option{},
		},
		[]ListItem{
			ListItem{"GitHub", "https://github.com/jeryanders", ""},
			ListItem{"Twitter", "https://twitter.com/atomdata", ""},
			ListItem{"LinkedIn", "https://www.linkedin.com/in/jesse-anderson-99469349/", ""},
		},
	}
}

// GetIndentClass ...
func GetIndentClass(d interface{}) string {
	switch d.(type) {
	case ListParameterDetail:
		return "List"
	default:
		return "Text"
	}
}
