package dom

import (
	"errors"
	"strconv"
)

type node struct {
	parent Node
	child  []Node
	store  DOMStore
}

func (n *node) ChildNodes() []Node {
	return n.child
}

func (n *node) ChildAt(index int) (child Node, err error) {
	if len(n.ChildNodes()) < index+1 {
		err = errors.New("No child node at index " + strconv.Itoa(index) + ".")
	} else {
		child = n.ChildNodes()[index]
	}
	return
}

func (n *node) AppendChildNode(m Node) {
	n.child = append(n.ChildNodes(), m)
	return
}

func (n *node) Parent() Node {
	return n.parent
}

func (n *node) SetParent(m Node) {
	n.parent = m
	return
}

func (n *node) Store() DOMStore {
	return n.store
}

func (n *node) String() string {
	s := ""
	for _, c := range n.ChildNodes() {
		s += c.String()
	}
	return s
}

type element struct {
	name   string
	attr   []Attribute
	prefix string
	node
}

func (e *element) String() string {
	s := "<"
	if e.Prefix() != "" {
		s += e.Prefix() + ":"
	}
	s += e.Name()
	for _, a := range e.Attr() {
		s += a.String()
	}
	
	if len(e.ChildNodes())==0 {
		s += "/>"
		return s
	}
	
	s += ">"
	for _, e := range e.ChildNodes() {
		s += e.String()
	}
	s += "</" + e.Name() + ">"
	return s
}

func (e *element) Name() string {
	return e.name
}

func (e *element) Attr() []Attribute {
	return e.attr
}

func (e *element) Prefix() string {
	return e.prefix
}

type attribute struct {
	name   string
	value  string
	prefix string
	node
}

func (a *attribute) String() string {
	return " " + a.name + "=\"" + a.value + "\""
}

func (a *attribute) Name() string {
	return a.name
}

func (a *attribute) Value() string {
	return a.value
}

func (a *attribute) Prefix() string {
	return a.prefix
}

type procInst struct {
	target string
	data   string
	node
}

func (pi *procInst) String() string {
	return "<?" + pi.target + " " + pi.data + "?>"
}

func (pi *procInst) Target() string {
	return pi.target
}

func (pi *procInst) Data() string {
	return pi.data
}

type comment struct {
	data string
	node
}

func (c *comment) String() string {
	return "<!--" + c.data + "-->"
}

func (c *comment) Data() string {
	return c.data
}

type text struct {
	data string
	node
}

func (t *text) String() string {
	return t.data
}

func (t *text) Data() string {
	return t.data
}

type declaration struct {
	procInst
}

type directive struct {
	data string
	node
}

func (d *directive) String() string {
	return "<!" + d.data + ">"
}

func (d *directive) Data() string {
	return d.data
}
