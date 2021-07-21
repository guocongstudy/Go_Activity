package main

import (
	"encoding/json"
	"fmt"
)

const (
	Large  = iota //large
	Medium        //medium
	Small         //small
)

type Size int

func (s Size) MarshalText() ([]byte, error) {
	switch s {
	case Large:
		return []byte("large"), nil
	case Medium:
		return []byte("medium"), nil
	case Small:
		return []byte("small"), nil
	default:
		return []byte("unknown"), nil
	}
}

func (s *Size) UnmarshalText(bytes []byte) error {
	switch string(bytes) {
	case "Large":
		*s = Large
	case "Medium":
		*s = Medium
	case "Small":
		*s = Small
	default:
		*s = Small
	}
	return nil
}
func main() {
	var size Size = Medium
	bytes, _ := json.Marshal(size)
	fmt.Println(string(bytes))
}
