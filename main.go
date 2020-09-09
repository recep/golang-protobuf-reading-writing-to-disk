package main

import (
	"fmt"
	userpb "github.com/golang-protobuf-reading-writing-to-disk/src/user"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	user := createUser()
	user2 := &userpb.User{}

	writeToFile("user.bin", user)
	readFromFile("user.bin",user2)
	fmt.Printf("%+v\n",user2)
}

func createUser() *userpb.User {
	user := userpb.User{
		Id:   1,
		Name: "recep",
		Age:  20,
		City: "Rize",
	}

	return &user
}

func writeToFile(fname string, pb proto.Message) error {
	data, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	if err = ioutil.WriteFile(fname, data, 0644); err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	data, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	err = proto.Unmarshal(data, pb)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}
