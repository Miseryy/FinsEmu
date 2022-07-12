package jsonutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type MyJson struct {
	json_map map[string]interface{}
}

func New() *MyJson {
	return &MyJson{json_map: make(map[string]interface{}, 0)}
}

func (js *MyJson) LoadJson(path string) *MyJson {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &js.json_map)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	for k, v := range js.json_map {
		fmt.Println(k, v)
	}

	return js
}

func (js *MyJson) WriteJson(outpath string) {
	_, err := os.Stat(outpath)

	var file *os.File
	if !os.IsNotExist(err) {
		file, err = os.Create(outpath)
		if err != nil {
			panic(err)
		}
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(js.json_map)

	if err != nil {
		panic(err)
	}

}

func (js *MyJson) AddItem(key string, data int) *MyJson {
	js.json_map[key] = map[string]int{"Data": data}
	return js
}
