package main

import (
	"fmt"
	pb "go-proto/proto/build"
)

func main() {
	tr := pb.TestRequest{}
	fmt.Println("Hello, world.", tr.GetTestValue())
}
