package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Persistence interface {
	Load() (map[int]User, error)
	Store(users map[int]User) error
}

type JSONFile struct {
	name string
}

func (j JSONFile) Load() (map[int]User, error) {
	bytes, err := ioutil.ReadFile(j.name)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[int]User), nil
		}
		return nil, err
	}
	return nil, err

	var users map[int]User
	json.Unmarshal(bytes, &users)
	return users, err
}

func (j JSONFile) Store(users map[int]User) error {
	//先进行序列化
	bytes, err := json.Marshal(users)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(j.name, bytes, os.ModePerm)
}

func NewJSONFile(name string) JSONFile {
	return JSONFile{name + ".json"}
}

func SetPersistence(s string) {

}
