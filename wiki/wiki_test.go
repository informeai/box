package main


import (

	"testings"

)


func TestWiki(t *testings.T){
	
	err := verifyArgs(os.Args)
	if err != nil{
		t.Fataln(err)
	}



}
