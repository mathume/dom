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
	store DOMStore
}

// creates and initializes IDOMBuilder structure
// only implements building dom from xml file
func NewDOMBuilder(reader io.Reader, store DOMStore) DOMBuilder{
	db := new(dombuilder)
	db.reader = reader
	db.store = store
	return db
}

func (db *dombuilder)Reader()(reader io.Reader){
	return db.reader
}

func (db *dombuilder)Build() (dom DOM, err error){
	defer func() {
		if r:=recover(); r!=nil {
			err = errors.New(fmt.Sprint(r))
			return
		}
	}()
	
	dom = db.convertDecoderToDOM()
	return
}

func (db *dombuilder)convertDecoderToDOM() DOM{

	dc := xml.NewDecoder(db.reader)
	dom := db.store.NewDOM()
	build_subtree(dc, dom)
	return dom
}

func build_subtree(dc *xml.Decoder, n Node){
	tok,_ := dc.Token()
	switch typ:=tok.(type){
	case xml.StartElement:
		el := n.Store().CreateElement(typ.Name.Space, typ.Name.Local, convert_Attr(typ.Attr,n.Store()))
		el.SetParent(n)
		n.AppendChildNode(el)
		build_subtree(dc, el)
	case xml.CharData:
		txt := n.Store().CreateText(string(typ))
		txt.SetParent(n)
		n.AppendChildNode(txt)
		build_subtree(dc, n)
	case xml.Comment:
		a := n.Store().CreateComment(string(typ))
		a.SetParent(n)
		n.AppendChildNode(a)
		build_subtree(dc,n)
	case xml.Directive:
		d := n.Store().CreateDirective(string(typ))
		d.SetParent(n)
		n.AppendChildNode(d)
		build_subtree(dc,n)
	case xml.ProcInst:
		pi := n.Store().CreateProcInst(typ.Target, string(typ.Inst))
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

func convert_Attr(a []xml.Attr,store DOMStore) []Attribute{
	as := make([]Attribute,len(a))
	for i:= 0; i<len(a); i++ {
		as[i] = store.CreateAttribute(a[i].Name.Space, a[i].Name.Local, a[i].Value)
	}
	return as
}