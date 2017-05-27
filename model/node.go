package model

import (
	"errors"
	"strconv"
	"strings"
)

type Node struct {
	Host string `json:"host"`
	Port int64  `json:"port"`
}

func (n Node) ToString() string {
	return n.Host + ":" + strconv.FormatInt(n.Port, 10)
}

func (n *Node) FromString(str string) error {
	splitsStr := strings.Split(str, ":")
	if len(splitsStr) != 2 {
		return errors.New("Cannot convert string -> model.Node{}")
	}
	n.Host = splitsStr[0]
	port, err := strconv.ParseInt(splitsStr[1], 10, 64)
	if err != nil {
		return err
	}
	n.Port = port
	return nil
}
