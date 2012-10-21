package dom

type domstore struct {
}

func NewDOMStore() DOMStore {
	return new(domstore)
}

func (s *domstore) CreateNode() Node {
	n := new(node)
	n.store = s
	return n
}

func (s *domstore) CreateDirective(data string) Directive {
	d := new(directive)
	d.data = data
	d.store = s
	return d
}

func (s *domstore) CreateDeclaration(data string) Declaration {
	return s.CreateProcInst("xml", data)
}

func (s *domstore) CreateText(data string) Text {
	t := new(text)
	t.data = data
	t.store = s
	return t
}

func (s *domstore) CreateComment(data string) Comment {
	c := new(comment)
	c.data = data
	c.store = s
	return c
}

func (s *domstore) CreateProcInst(target string, data string) ProcInst {
	pi := new(procInst)
	pi.target = target
	pi.data = data
	pi.store = s
	return pi
}

func (s *domstore) CreateAttribute(prefix string, name string, value string) Attribute {
	a := new(attribute)
	a.prefix = prefix
	a.name = name
	a.value = value
	return a
}

func (s *domstore) CreateElement(prefix string, name string, attr []Attribute) Element {
	e := new(element)
	e.prefix = prefix
	e.name = name
	e.attr = attr
	e.store = s
	return e
}

func (s *domstore) NewDOM() DOM {
	d := new(dom)
	d.self = s.CreateNode()
	return d
}
