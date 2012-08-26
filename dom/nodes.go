package dom
import (
	"strconv"
	"errors"
)

type node struct{
	parent Node
	child []Node
}

func (n *node)ChildNodes() []Node{
	return n.child
}

func (n *node)ChildAt(index int) (child Node, err error){
	if len(n.ChildNodes())<index + 1 {
		child = nil
		err = errors.New("No child node at index " + strconv.Itoa(index) + ".")
	}else{
		child = n.ChildNodes()[index]
		err = nil
	}
	return
}

func (n *node)AppendChildNode(m Node){
	n.child = append(n.ChildNodes(), m)
	return
}

func (n *node)Parent() Node{
	return n.parent
}

func (n *node)SetParent(m Node){
	n.parent = m
	return
}

func (n *node)String() string{
	return "node"
}

type element struct{
	name string
	attr []Attribute
	prefix string
	node
}

func NewElement() Element{
	return new(element)
}

func CreateElement(prefix string, name string, attr []Attribute) Element{
	e := new(element)
	e.prefix = prefix
	e.name = name
	e.attr = attr
	return e
}

func (e *element)String() string{
	s := "<" 
	if e.prefix != "" {
		s += e.prefix + ":"
	}
	s += e.name
	for _, a := range e.attr  {
		s += a.String()
	}
	s += ">"
	for _, e := range e.ChildNodes()  {
		s += e.String()
	}
	s += "</" + e.name + ">"
	return s
}

func (e *element)Name() string{
	return e.name
}

func (e *element)Attr() []Attribute{
	return e.attr
}

func (e *element)Prefix() string{
	return e.prefix
}

type attribute struct{
	name string
	value string
	prefix string
	node
}

func NewAttribute() Attribute{
	return new(attribute)
}

func CreateAttribute(prefix string, name string, value string) Attribute{
	a := new(attribute)
	a.prefix = prefix
	a.name = name
	a.value = value
	return a
}

func (a *attribute)String() string{
	return " " + a.name + "=\"" + a.value + "\""
}

func (a *attribute)Name() string{
	return a.name
}

func (a *attribute)Value() string{
	return a.value
}

func (a *attribute)Prefix() string{
	return a.prefix
}

type procInst struct{
	target string
	data string
	node
}

func NewProcInst() ProcInst{
	return new(procInst)
}

func CreateProcInst(target string, data string) ProcInst{
	pi := new(procInst)
	pi.target = target
	pi.data = data
	return pi
}

func (pi *procInst)String() string{
	return "<?" + pi.target + " " + pi.data + "?>"
}

func (pi *procInst)Target() string{
	return pi.target
}

func (pi *procInst)Data() string{
	return pi.data
}

type comment struct{
	data string
	node
}

func NewComment() Comment{
	return new(comment)
}

func CreateComment(data string) Comment{
	c := new(comment)
	c.data = data
	return c
}

func (c *comment)String() string{
	return "<!--" + c.data + "-->"
}

func (c *comment)Data() string{
	return c.data
}

type text struct{
	data string
	node
}

func NewText() Text{
	return new(text)
}

func CreateText(data string) Text{
	t := new(text)
	t.data = data
	return t
}

func (t *text)String() string{
	return t.data
}

func (t *text)Data() string{
	return t.data
}

type declaration struct{
	procInst
}

func NewDeclaration() Declaration{
	return new(declaration)
}

func CreateDeclaration(data string) Declaration{
	return CreateProcInst("xml", data)
}

type directive struct{
	data string
	node
}

func NewDirective() Directive{
	return new(directive)
}

func CreateDirective(data string) Directive{
	d := new(directive)
	d.data = data
	return d
}

func (d *directive)String() string{
	return "<!" + d.data + ">"
}

func (d *directive)Data() string{
	return d.data
}
