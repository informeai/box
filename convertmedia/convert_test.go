package convertmedia



import(
	"testing"


)

func TestConvertMedia(t *testing.T){

	conv := NewConvert([]string{"convert","inputFile","outputDir"})
	if len(conv.Args) < 3{
		t.Fatal("error: args not completed")
	}

}
