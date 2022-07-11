package jsonutil

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DM_JSON struct {
	Dm struct {
		No   int `json:"no"`
		Data int `json:"data"`
	} `json:"DM"`
}

type MyJson struct {
	json_map map[string]interface{}
}

func New() *MyJson {
	return &MyJson{}
}

func (js *MyJson) LoadJson(path string) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(buf, &js.json_map)

	if err != nil {
		panic(err)
	}

	fmt.Println(js.json_map)

	for k, v := range js.json_map {
		fmt.Println(k, v)
	}
}

func (js *MyJson) WriteJson(obj interface{}, outpath string) {
	file, err := os.Create(outpath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	aa := map[string]interface{}{
		"DM130": struct { key string val int } {"Data", 100}
	}
	js.json_map = append(js.json_map, aa)

	err = json.NewEncoder(file).Encode(js.json_map)

}
