package dom

import (
	"encoding/xml"
	"io"
	"errors"
	"fmt"
)

type dombuilder struct{
	d DOM
	reader io.Reader
}

// creates and initializes IDOMBuilder structure
// only implements building dom from xml file
func NewDOMBuilder(reader io.Reader) DOMBuilder{
	db := new(dombuilder)
	db.reader = reader
	return db
}

func (db *dombuilder)SetReader(reader io.Reader){
	db.reader = reader
	return
}

func (db *dombuilder)Reader()(reader io.Reader){
	return db.reader
}

func (db *dombuilder)DOM() DOM{
	return db.d
}

func (db *dombuilder)Build() (err error){
	defer func() {
		if r:=recover(); r!=nil {
			err = errors.New(fmt.Sprint(r))
		}
	}()
	db.convertDecoderToDOM(db.reader)
	return
}

func (db *dombuilder)convertDecoderToDOM(reader io.Reader){

	dc := xml.NewDecoder(reader)
	decl := NewDeclaration()
	root := NewElement()

	tok,_ := dc.Token()
	switch typ := tok.(type){
	case xml.ProcInst:
		decl = CreateDeclaration(string(typ.Inst))
		tok, _ = dc.Token()
	}

	for {
		switch typ := tok.(type){
		case xml.StartElement:
			root = CreateElement(typ.Name.Space, typ.Name.Local, convert_Attr(typ.Attr))
			goto outer
		case nil:
			return
		}
		tok,_ = dc.Token()
	}
outer:
	if root.Name() == "" {
		panic("root element not found")
	}
	build_subtree(dc, root)

	db.d = CreateDOM(decl, root)
}

func build_subtree(dc *xml.Decoder, n Node){
	tok,_ := dc.Token()
	switch typ:=tok.(type){
	case xml.StartElement:
		el := CreateElement(typ.Name.Space, typ.Name.Local, convert_Attr(typ.Attr))
		el.SetParent(n)
		n.AppendChildNode(el)
		build_subtree(dc, el)
	case xml.CharData:
		txt := CreateText(string(typ))
		txt.SetParent(n)
		n.AppendChildNode(txt)
		build_subtree(dc, n)
	case xml.Comment:
		a := CreateComment(string(typ))
		a.SetParent(n)
		n.AppendChildNode(a)
		build_subtree(dc,n)
	case xml.Directive:
		d := CreateDirective(string(typ))
		d.SetParent(n)
		n.AppendChildNode(d)
		build_subtree(dc,n)
	case xml.ProcInst:
		pi := CreateProcInst(typ.Target, string(typ.Inst))
		pi.SetParent(n)
		n.AppendChildNode(pi)
		build_subtree(dc,n)
	case xml.EndElement:
		switch pTyp := n.Parent().(type){
		case Element:
			build_subtree(dc, pTyp)
		default:
			return
		}
	}
	return
}

func convert_Attr(a []xml.Attr) []Attribute{
	as := make([]Attribute,len(a))
	for i:= 0; i<len(a); i++ {
		as[i] = CreateAttribute(a[i].Name.Space, a[i].Name.Local, a[i].Value)
	}
	return as
}
