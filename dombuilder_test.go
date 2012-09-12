package dom

import (
	"testing"
	"fmt"
)

func TestDOMBuilder(t *testing.T){
	db := NewDOMBuilder()
	db.SetFile("monitors.xml")
	db.Build()
	err := db.DOM().SaveAs("dom.xml")
	if err!=nil {
		t.Fatal(err)
	}
	d := db.DOM()
	fmt.Println(d)
}
