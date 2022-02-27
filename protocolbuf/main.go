package main

import (
	"fmt"
	"io/ioutil"
	"log"
	simplepb "protocolbuf/src/simple"

	"google.golang.org/protobuf/proto"
)

func main(){
	sam := MySimple()

	writeToFile("simple.bin", sam)

	sm2 := &simplepb.SimpleMessage{}

	readFromFile("simple.bin", sm2)
}	

func writeToFile(fname string, pb proto.Message) error{
	out, err := proto.Marshal(pb)
	if (err != nil){
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname,out, 0644); err != nil{

		log.Fatalln("Can't Write to file ",err)
		return err
	}
	fmt.Println("Data Has been Written")
	return nil

}

func readFromFile(fname string, pb proto.Message) error{
	in, err := ioutil.ReadFile(fname)

	if(err != nil){
		log.Fatalln("Something went wrong in read file",err)
		return err
	}
	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("COuld'nt put the bytes  into the protocol buffers  struct", err2)
		return err2
	}
	return nil

}

func MySimple() *simplepb.SimpleMessage{

	sm := simplepb.SimpleMessage{
		Id: 1234,
		IsSimple: true,
		Name: "ABid",
		SimpleList: []int32{1,4,5,9},
	}

	fmt.Println(sm)

	sm.Name = "I rename you"

	fmt.Println(sm)

	fmt.Println("The id is ",sm.GetId())
	return &sm

}