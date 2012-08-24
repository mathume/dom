package dom

// represents an arbitrary xml node
// xml declaration and root element
// have Parent() == null
type Node interface{
	Parent() Node
	SetParent(n Node)
	ChildNodes() []Node
	AppendChildNode(n Node)
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
	Declaration() Declaration
	Root() Element
	FileName() string
	SaveAs(filename string) error
}

// builds an DOM
// by default the DOM is read from a file
type DOMBuilder interface{
	File() string
	SetFile(filename string)
	DOM() DOM
	Build()
}

type Iterator interface{
	Next() (Node, error)
	Reset()
}
