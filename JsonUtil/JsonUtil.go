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

func (js *MyJson) LoadJson(path string) error {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(buf, &js.json_map)

	if err != nil {
		fmt.Println(err)
		return err
	}

	for k, v := range js.json_map {
		fmt.Println(k, v)
	}

	return err
}

func (js *MyJson) WriteJson(outpath string) error {
	file, err := os.OpenFile(outpath, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		fmt.Println("Open Error")
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(js.json_map)

	if err != nil {
		fmt.Println("Write Error")
		return err
	}

	return err
}

func (js *MyJson) AddItem(key string, data int64) *MyJson {
	js.json_map[key] = map[string]int{"Data": int(data)}
	return js
}
