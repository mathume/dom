package dom

import (
	"encoding/xml"
	"io"
	"strings"
)

type dombuilder struct {
	d      DOM
	reader io.Reader
	store  DOMStore
}

// creates and initializes IDOMBuilder structure
// only implements building dom from xml file
func NewDOMBuilder(reader io.Reader, store DOMStore) DOMBuilder {
	db := new(dombuilder)
	db.reader = reader
	db.store = store
	return db
}

func (db *dombuilder) Reader() (reader io.Reader) {
	return db.reader
}

func (db *dombuilder) Build() (dom DOM, err error) {

	dom, err = db.convertDecoderToDOM()
	return
}

func (db *dombuilder) convertDecoderToDOM() (dom DOM, err error) {

	dc := xml.NewDecoder(db.reader)
	dom = db.store.NewDOM()
	err = build_subtree(dc, dom)
	return
}

func build_subtree(dc *xml.Decoder, n Node) (err error) {

	tok, err := dc.Token()
	if err != nil {
		if err == io.EOF {
			err = nil
		}
		return
	}

	switch typ := tok.(type) {
	case xml.StartElement:
		el := n.Store().CreateElement(typ.Name.Space, typ.Name.Local, convert_Attr(typ.Attr, n.Store()))
		err = linkAndContinueDown(dc, n, el)
	case xml.CharData:
		txt := n.Store().CreateText(string(typ))
		err = linkAndContinueUp(dc, n, txt)
	case xml.Comment:
		a := n.Store().CreateComment(string(typ))
		err = linkAndContinueUp(dc, n, a)
	case xml.Directive:
		d := n.Store().CreateDirective(string(typ))
		err = linkAndContinueUp(dc, n, d)
	case xml.ProcInst:
		pi := n.Store().CreateProcInst(typ.Target, string(typ.Inst))
		err = linkAndContinueUp(dc, n, pi)
	case xml.EndElement:
		err = build_subtree(dc, n.Parent())
	}
	return
}

func linkAndContinueUp(dc *xml.Decoder, parent Node, child Node)(err error){
	err = linkAndContinueWith(parent, dc, parent, child)
	return
}

func linkAndContinueDown(dc *xml.Decoder, parent Node, child Node)(err error){
	err = linkAndContinueWith(child, dc, parent, child)
	return
}

func linkAndContinueWith(with Node, dc *xml.Decoder, parent Node, child Node)(err error){
	child.SetParent(parent)
	parent.AppendChildNode(child)
	return build_subtree(dc, with)
}

func convert_Attr(a []xml.Attr, store DOMStore) []Attribute {
	as := make([]Attribute, len(a))
	for i := 0; i < len(a); i++ {
		as[i] = store.CreateAttribute(a[i].Name.Space, a[i].Name.Local, a[i].Value)
	}
	return as
}

func isWellformed(d DOM) (b bool) {
	noOfRoots := 0

	for _, node := range d.ChildNodes() {
		switch node.Kind() {
		case ElementKind:
			noOfRoots++
		case TextKind:
			if !containsOnlyWhiteSpace(
				node.(Text).Data()) {
				return
			}
		}
	}

	if noOfRoots == 1 {
		b = true
	}
	return
}

func containsOnlyWhiteSpace(s string)(b bool){
	ws := " \t\n\r"
	if strings.Trim(s, ws) == "" { b = true }
	return
}
