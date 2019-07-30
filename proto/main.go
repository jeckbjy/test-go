package main

import (
	proto "./proto"
	pb "./example"
	"log"
)

func main() {
	test := &pb.Test{
		Label: "hello",
		Type:  17,
		Reps:  []int64{1, 2, 3},
		// Optionalgroup: &pb.Test_OptionalGroup{
		// 	RequiredField: proto.String("good bye"),
		// },
		Union: &pb.Test_Name{"fred"},
	}
	data, err := proto.Marshal(test)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	} else {
		log.Printf("marshal:%+s\n", data)
	}
	newTest := &pb.Test{}
	err = proto.Unmarshal(data, newTest)
	if err != nil {
		log.Fatal("unmarshaling error: ", err)
	} else {
		log.Printf("unmarshal:%+v", newTest)
	}
	// Now test and newTest contain the same data.
	if test.GetLabel() != newTest.GetLabel() {
		log.Fatalf("data mismatch %q != %q", test.GetLabel(), newTest.GetLabel())
	}
	// Use a type switch to determine which oneof was set.
	switch test.Union.(type) {
	case *pb.Test_Number: // u.Number contains the number.
		log.Printf("number:%+v\n", test.GetNumber())
	case *pb.Test_Name: // u.Name contains the string.
		log.Printf("name:%+v\n", test.GetName())
	}
	// etc.
}