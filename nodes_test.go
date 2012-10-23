package dom

import (
	"strings"
	"testing"
)

func TestTextOnElement(t *testing.T) {
	xmlSampleWithText := "<root>a\n<child><!-- c --><?proc inst ?>b</child>c<child/>d</root>"
	db := NewDOMBuilder(strings.NewReader(xmlSampleWithText),
		NewDOMStore())
	d, _ := db.Build()
	expectedTextOnRoot := "a\nbcd"
	if d.Root().Text() != expectedTextOnRoot {
		t.Fatal(d.Root().Text())
	}
}
