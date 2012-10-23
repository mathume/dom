package dom

import (
	"encoding/xml"
	"io"
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

func TestConvertDecoderToDOM_Build_ReturnsNoErrorByDefault_WithNonXml(t *testing.T) {
	nonXml := "abc"
	db := NewDOMBuilder(strings.NewReader(nonXml), NewDOMStore())
	dom, err := db.Build()
	if err != nil {
		t.Fatal(dom)
	}
}

func TestConvertDecoderToDOM_MissingRootEndTag_ReturnsNoError(t *testing.T) {
	xmlMissingEndTag := "<root>"
	db := NewDOMBuilder(strings.NewReader(xmlMissingEndTag), NewDOMStore())
	d, err := db.Build()
	if err != nil {
		t.Fatal(d)
	}
}

func TestConvertDecoderToDOM_MissingChildEndTag_ReturnsError(t *testing.T) {
	xmlMissingEndTagInChild := "<root><child></root>"
	db := NewDOMBuilder(strings.NewReader(xmlMissingEndTagInChild), NewDOMStore())
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

func TestConvertDecoderToDOM_ReturnsNoError_ForTwoRoots(t *testing.T) {
	tworoots := empty_root() + empty_root()
	db := NewDOMBuilder(strings.NewReader(tworoots), NewDOMStore())
	_, err := db.Build()
	if err != nil {
		t.Fail()
	}
}

func TestConvertDecoderToDOM_ReturnsNoError_ForMissingRoot(t *testing.T) {
	onlydeclaration := xml.Header
	db := NewDOMBuilder(strings.NewReader(onlydeclaration), NewDOMStore())
	_, err := db.Build()
	if err != nil {
		t.Fail()
	}
}

func TestContainsOnlyWhiteSpace_True(t *testing.T) {
	if !containsOnlyWhiteSpace(" \n \t \r \n") {
		t.Fatal("Recognised non-whitespace where there is none.")
	}
}

func TestContainsOnlyWhiteSpace_False(t *testing.T) {
	if containsOnlyWhiteSpace(" \n \r j \r \t \n ") {
		t.Fatal("Recognised whitespace where there is non-whitespace.")
	}
}

func TestIsWellformed_AtTwoRoots(t *testing.T) {
	data := notWellformedBecauseOfTwoRoots()
	db := NewDOMBuilder(strings.NewReader(data), NewDOMStore())
	d, _ := db.Build()
	if isWellformed(d) {
		t.Fail()
	}
}

func TestIsWellformed_AtCharacterDataNextToRoot(t *testing.T) {
	data := notWellformedBecauseOfCharacterData()
	db := NewDOMBuilder(strings.NewReader(data), NewDOMStore())
	d, _ := db.Build()
	if isWellformed(d) {
		t.Fail()
	}
}

func TestIsWellformed_AtEmptyDocument(t *testing.T) {
	data := ""
	db := NewDOMBuilder(strings.NewReader(data), NewDOMStore())
	d, _ := db.Build()
	if isWellformed(d) {
		t.Fail()
	}
}

func TestIsWellformed_AtSampleDocument(t *testing.T) {
	sampleXml := getSampleXml()
	db := NewDOMBuilder(strings.NewReader(sampleXml), NewDOMStore())
	d, _ := db.Build()
	if !isWellformed(d) {
		t.Fail()
	}
}

func TestXmlDecoder_AtMissingEndTag(t *testing.T) {
	reader := strings.NewReader("<root><child></root>")
	dec := xml.NewDecoder(reader)
	dec.Strict = true
	_, err := dec.Token()
	_, err = dec.Token()
	_, err = dec.Token()
	if err == io.EOF || err == nil {
		t.Fatal(err)
	}

	reader = strings.NewReader("<root>")
	dec = xml.NewDecoder(reader)
	dec.Strict = true
	_, err = dec.Token()
	_, err = dec.Token()
	if err != io.EOF {
		t.Fatal(err)
	}
}

func notWellformedBecauseOfTwoRoots() (s string) {
	s = xml.Header + empty_root() + empty_root()
	return
}

func notWellformedBecauseOfCharacterData() (s string) {
	ds := NewDOMStore()
	s += xml.Header
	s += ds.CreateProcInst("target", "data").String()
	s += empty_root()
	s += "characterData"
	return
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
