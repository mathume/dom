package dom

import(
	"io"
)

// represents an arbitrary xml node
// xml declaration and root element
// have Parent() == null
type Node interface{
	Parent() Node
	SetParent(n Node)
	ChildNodes() []Node
	ChildAt(index int) (Node, error)
	AppendChildNode(n Node)
	Store() DOMStore
	String() string
}

// represents an xml element
// ChildNodes() also returns Data as Text
type Element interface{
	Node

	Prefix() string
	Name() string
	Attr() []Attribute
}

// represents an xml attribute
// ChildNodes() always returns nil
type Attribute interface{
	Prefix() string
	Name() string
	Value() string
	String() string
}

// represents an xml comment (<!-- $text -->).
// ChildNodes() will return slice with $text
type Comment interface{
	Node
	Data
}

// represents an xml processing instruction (<?--$target $text-->)
// ChildNodes() will return slice with $text
type ProcInst interface{
	Node
	Data

	Target() string
}

// represents Data as a node
// Data() returns the text contents
type Text interface{
	Node
	Data
}

// represents the top level xml declaration
// it is a processing instruction with Target()=="xml"
type Declaration interface{
	ProcInst

}

// represents an xml directive (<!text ...>)
type Directive interface{
	Node
	Data
}

// returns the text contents of text nodes
type Data interface{
	Data() string
}

// represents a DOM
// by returning Root() you can navigate deeper
// other dom functions like GetElementsById etc
// might be added later
type DOM interface{
	Node
	Declaration() Declaration
	Root() Element
}

// builds an DOM
// by default the DOM is read from a file
type DOMBuilder interface{
	Build() (DOM, error)
	Reader()(reader io.Reader)
}

type DOMStore interface{
	NewDOM() DOM
	CreateElement(prefix string, name string, attr []Attribute) Element
	CreateAttribute(prefix string, name string, value string) Attribute
	CreateProcInst(target string, data string) ProcInst
	CreateComment(data string) Comment
	CreateText(data string) Text
	CreateDeclaration(data string) Declaration
	CreateDirective(data string) Directive
}
