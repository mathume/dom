package dom

import (
	"encoding/xml"
	"io"
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
	
	dom,err = db.convertDecoderToDOM()
	return
}

func (db *dombuilder) convertDecoderToDOM()(dom DOM, err error) {

	dc := xml.NewDecoder(db.reader)
	dom = db.store.NewDOM()
	err = build_subtree(dc, dom)
	return
}

func build_subtree(dc *xml.Decoder, n Node) (err error) {
	tok, err := dc.Token()
	if err != nil {
		return
	}
	switch typ := tok.(type) {
	case xml.StartElement:
		el := n.Store().CreateElement(typ.Name.Space, typ.Name.Local, convert_Attr(typ.Attr, n.Store()))
		el.SetParent(n)
		n.AppendChildNode(el)
		err = build_subtree(dc, el)
	case xml.CharData:
		txt := n.Store().CreateText(string(typ))
		txt.SetParent(n)
		n.AppendChildNode(txt)
		err = build_subtree(dc, n)
	case xml.Comment:
		a := n.Store().CreateComment(string(typ))
		a.SetParent(n)
		n.AppendChildNode(a)
		err = build_subtree(dc, n)
	case xml.Directive:
		d := n.Store().CreateDirective(string(typ))
		d.SetParent(n)
		n.AppendChildNode(d)
		err = build_subtree(dc, n)
	case xml.ProcInst:
		pi := n.Store().CreateProcInst(typ.Target, string(typ.Inst))
		pi.SetParent(n)
		n.AppendChildNode(pi)
		err = build_subtree(dc, n)
	case xml.EndElement:
		switch pTyp := n.Parent().(type) {
		case Element:
			err = build_subtree(dc, pTyp)
		default:
			return
		}
	}
	return
}

func convert_Attr(a []xml.Attr, store DOMStore) []Attribute {
	as := make([]Attribute, len(a))
	for i := 0; i < len(a); i++ {
		as[i] = store.CreateAttribute(a[i].Name.Space, a[i].Name.Local, a[i].Value)
	}
	return as
}
