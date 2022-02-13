package main

import (
	"fmt"
	"github.com/zolinz/protobuf-test/simplepb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	fmt.Println("hello Zoli")
	myMessage := doSimple()
	writeToFile(myMessage, "testproto.data")

	sm2 := &simplepb.Simple{}
	readFromFile("testproto.data", sm2)

	fmt.Println("here are the values ")
	fmt.Println(sm2.GetName())
	fmt.Println(sm2.GetAge())
}

//marshal - serialize
func writeToFile(pb proto.Message, fileName string) error {
	fmt.Printf("hello from file \n %s", fileName)
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	ioutil.WriteFile(fileName, out, 0644)

	return nil
}

//unmarshal - deserialize
func readFromFile(fileName string, pb proto.Message) {

	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("error reading the file")
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("couldnot read file")
	}
}

func doSimple() *simplepb.Simple {
	sm := simplepb.Simple{
		Name: "Zolika",
		Age:  46,
	}

	fmt.Print(sm)
	return &sm
}
