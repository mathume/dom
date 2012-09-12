/* needs documentation
*/
package dom
import "io/ioutil"

type dom struct{
	decl Declaration
	root Element
	filename string
}

func (d *dom)FileName() string{
	return d.filename
}

func (d *dom)Declaration() Declaration{
	return d.decl
}

func (d *dom)Root() Element{
	return d.root
}

func (d *dom)String() string{
	return d.decl.String() + "\n" + d.root.String()
}

// returns and empty DOM structure
// TODO: implement store mechanism for faster traversing/search
// TODO: implement DOMSearch (define that one too)
func NewDom() DOM{
	return new(dom)
}

func CreateDOM(decl Declaration, root Element, filename string) DOM{
	d := new(dom)
	d.decl = decl
	d.root = root
	d.filename = filename
	return d
}

func (d *dom)SaveAs(filename string) error{
	return ioutil.WriteFile(filename, []byte(d.String() + "\n"), 0666)
}
