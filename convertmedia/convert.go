package convertmedia

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"errors"
	"github.com/xfrr/goffmpeg/transcoder"
)

type MediaConverter interface{
	verifyArgs() error
	verifyFile() error
	transcoderFile() error
}

type Convert struct{
	Args []string


}

func NewConvert(args []string) *Convert{
	return &Convert{Args:args}

}

func(c *Convert) verifyArgs() error{
	if len(c.Args) < 3{
		fmt.Println("Usage: convert [input file] [output directory]")
		return errors.New("error: args not completed") 

	}

	return nil
}
func(c *Convert) verifyFile() error{
	input := c.Args[1]
	output := c.Args[2]
	
	if _, err := os.Stat(input); os.IsNotExist(err){
		return errors.New("error: input file not existed")
	}

	if _, err := os.Stat(filepath.Dir(output)); os.IsNotExist(err){
		err = os.Mkdir(filepath.Dir(output), 0755)
		if err != nil{
			return errors.New("error: output directory not created")
		}
	}
	return nil
}

func(c *Convert) transcoderFile() error{
	trans := new(transcoder.Transcoder)
	err := trans.Initialize(c.Args[1],c.Args[2])
	if err != nil{
		return errors.New("error: transcoder not initialized")
	}
	duration := trans.MediaFile().DurationInput()
	fmt.Println(duration)
	done := trans.Run(true)

	progress := trans.Output()
	fmt.Printf("Initialized\n")
	fmt.Printf("INPUT: %v\n", filepath.Base(c.Args[1]))
	fmt.Printf("OUTPUT: %v\n", filepath.Base(c.Args[2]))
	for p := range progress{
		fmt.Printf("\rPROGRESS: %v\tSPEED: %v", p.CurrentTime, p.Speed[:len(p.Speed)-1])

	}
	err = <- done
	if err != nil && err != io.EOF{
		return errors.New("error: problem in process of transcoder file")

	}
	fmt.Printf("\nFinished\n")
	return nil
}
func(c *Convert) Convert() error{
	err := c.verifyArgs()
	if err != nil{
		return err
	}
	err = c.verifyFile()
	if err != nil{
		return err
	}
	err = c.transcoderFile()
	if err != nil{
		return err
	}
	return nil

}
