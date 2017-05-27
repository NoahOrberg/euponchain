package model

import (
	"fmt"
	"testing"
)

func TestNodeToString(t *testing.T) {
	node := Node{
		Host: "localhost",
		Port: 8070,
	}
	str := node.ToString()
	fmt.Printf("node.ToString() = %v \n", str)
}

func TestNodeFromString(t *testing.T) {
	node := Node{}
	str := "localhost:8070"
	err := node.FromString(str)
	if err != nil {
		t.Fatal("Missing translate.")
	}
	fmt.Printf("node.Host = %v \n", node.Host)
	fmt.Printf("node.Port = %v \n", node.Port)
}
