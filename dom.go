/* needs documentation
 */
package dom

type dom struct {
	self Node
	decl Declaration
	root Element
}

func (d *dom) Parent() Node {
	return d.self
}
func (d *dom) SetParent(n Node) {
	return
}
func (d *dom) ChildNodes() []Node {
	return d.self.ChildNodes()
}
func (d *dom) ChildAt(index int) (Node, error) {
	return d.self.ChildAt(index)
}
func (d *dom) AppendChildNode(n Node) {
	d.self.AppendChildNode(n)
}

func (d *dom) Store() DOMStore {
	return d.self.Store()
}

func (d *dom) String() string {
	return d.self.String()
}

func (d *dom) Declaration() Declaration {
	if d.declAndRootNotSet() {
		d.getDeclAndRoot()
	}
	return d.decl
}
func (d *dom) Root() Element {
	if d.declAndRootNotSet() {
		d.getDeclAndRoot()
	}
	return d.root
}

func (d *dom) Kind() Kind{
	return NodeKind
}

func (d *dom)declAndRootNotSet() bool{
	return d.root == nil || d.decl == nil
}

func (d *dom) getDeclAndRoot() {
	children := d.self.ChildNodes()
	for i := 0; i < len(children); i++ {
		switch children[i].(type) {
		case ProcInst:
			if children[i].(ProcInst).Target() == "xml" {
				d.decl = children[i].(Declaration)
			}
		case Element:
			d.root = children[i].(Element)
		}
	}
}
