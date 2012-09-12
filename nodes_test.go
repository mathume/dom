package dom

import (
	"testing"
	"fmt"
	"reflect"
)

func check_not_nil(t *testing.T, a ...interface{}){
	for i:=0; i<len(a); i++{
		switch a[i]{
		case nil:
			t.Fatal(fmt.Sprint(a[i]) + " was nil!")
		}
	}
}

func check_nil(t *testing.T, a ...interface{}){
	for i:=0; i<len(a); i++{
		switch a[i]{
		case nil:
		default:
			t.Fatal(fmt.Sprint(a[i]) + " is of kind " + reflect.TypeOf(a[i]).Kind().String() + " , not nil!")
		}
	}
}

func check_type(t *testing.T, typ interface{}, e Node){
	if reflect.TypeOf(typ) != reflect.TypeOf(e) {
		t.Fatal("Wrong Type().")
	}
}
