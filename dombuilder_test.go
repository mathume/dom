package dom

import (
	"testing"
	"encoding/xml"
	"strings"
)

func TestConvertDecoderToDOM_ReadsDeclaration(t *testing.T){
	text,expected_declaration := declaration_and_empty_root()
	db := NewDOMBuilder(strings.NewReader(text))
	db.Build()
	d := db.DOM().Declaration()

	if(d.String() != expected_declaration){
		t.Fail()
	}
}

func TestConvertDecoderToDOM_ReadsEmptyRoot(t *testing.T){
	db := NewDOMBuilder(strings.NewReader(empty_root()))
	expected := split_empty_root()
	db.Build()
	r := db.DOM().Root()

	if(r.String() != expected){
		t.Fail()
	}
}

func TestConvertDecoderToDOM_Build_RaisesError_WithNonXml(t *testing.T){
	t.Fatal("Test not implemented yet.")
}

func split_empty_root()(root string){
	return "<root></root>"
}

func declaration_and_empty_root() (text, expected_declaration string){
	text = xml.Header + empty_root()
	expected_declaration = strings.Split(xml.Header,"\n")[0]
	return
}

func empty_root() string{
	return "<root/>"
}
