package jsonutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type MyJson struct {
	json_map map[string]interface{}
	path     string
}

func New(p string) *MyJson {
	return &MyJson{
		json_map: make(map[string]interface{}, 0),
		path:     p,
	}
}

func (self *MyJson) LoadJson() error {
	buf, err := ioutil.ReadFile(self.path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &self.json_map)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func (self *MyJson) WriteJson() error {
	file, err := os.Create(self.path)

	if err != nil {
		fmt.Println("Open Error")
		return err
	}

	defer file.Close()

	e := json.NewEncoder(file)
	e.SetIndent("", strings.Repeat(" ", 2))
	e.SetEscapeHTML(true)

	err = e.Encode(self.json_map)

	if err != nil {
		fmt.Println("Write Error")
		return err
	}

	return err
}

func (self *MyJson) AddItem(key string, data int64) *MyJson {
	self.json_map[key] = data // map[string]int{"Data": int(data)}
	return self
}

func (self *MyJson) DeleteItem(key string) *MyJson {
	delete(self.json_map, key)
	return self
}

func (self *MyJson) GetMap() map[string]interface{} {
	return self.json_map
}
