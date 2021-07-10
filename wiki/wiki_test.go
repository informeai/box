package wiki


import (

	"testing"
	//"os"

)


func TestWiki(t *testing.T){
	
	wiki := NewWiki([]string{"wiki","-l","en","-w","cafe"})
	err := wiki.GetPage()
	if err != nil{
		t.Fatal(err)
	}



}
