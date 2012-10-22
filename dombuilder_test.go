package dom

import (
	"encoding/xml"
	"strings"
	"testing"
)

func TestConvertDecoderToDOM_ReadsDeclaration(t *testing.T) {
	text, expected_declaration := declaration_and_empty_root()
	db := NewDOMBuilder(strings.NewReader(text), NewDOMStore())
	dom, _ := db.Build()
	d := dom.Declaration()

	if d.String() != expected_declaration {
		t.Fail()
	}
}

func TestConvertDecoderToDOM_ReadsEmptyRoot(t *testing.T) {
	db := NewDOMBuilder(strings.NewReader(empty_root()), NewDOMStore())
	expected := empty_root()
	dom, _ := db.Build()
	r := dom.Root()

	if r.String() != expected {
		t.Fail()
	}
}

func TestConvertDecoderToDOM_Build_ReturnsError_WithNonXml(t *testing.T) {
	nonXml := "abc"
	db := NewDOMBuilder(strings.NewReader(nonXml), NewDOMStore())
	_, err := db.Build()
	if err == nil {
		t.Fail()
	}
}

func TestConvertDecoderToDOM_MissingEndTag_ReturnsError(t *testing.T) {
	xmlMissingEndTag := "<root>"
	db := NewDOMBuilder(strings.NewReader(xmlMissingEndTag), NewDOMStore())
	_, err := db.Build()
	if err == nil {
		t.Fail()
	}
}

func TestConvertDecoderToDOM_OnSampleDocument(t *testing.T) {
	sampleXml := getSampleXml()
	db := NewDOMBuilder(strings.NewReader(sampleXml), NewDOMStore())
	d, err := db.Build()
	if err != nil {
		t.Log("db.Build() returned error")
		t.Fatal(err)
	}
	ds := d.String()
	de := sampleXml
	if ds != de {
		t.Fail()
	}
}

func TestDecoderAtMissingEndTag(t *testing.T) {
	reader := strings.NewReader("<rootwithoutendtag>")
	dec := xml.NewDecoder(reader)
	t.Log(dec.Strict)
	_, err := dec.Token()
	_, err = dec.Token()
	if err == nil {
		t.Fail()
	}
}

func declaration_and_empty_root() (text, expected_declaration string) {
	text = xml.Header + empty_root()
	expected_declaration = strings.Split(xml.Header, "\n")[0]
	return
}

func empty_root() string {
	return "<root/>"
}

func getSampleXml() (s string) {
	s = "<?xml version=\"1.0\" standalone=\"yes\" encoding=\"UTF-8\"?><!-- comment --><root><?proc processing instruction ?><firstchild name=\"value\"><grandchild/></firstchild><secondchild id=\"5\"/><!-- comment --></root>"
	return
}
