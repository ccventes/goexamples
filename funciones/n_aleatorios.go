package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func sayRandomName() string {

	//firstName := "pepe"
	//lastName := "lujanes"
	response, err := http.Get("https://randomuser.me/api/")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	type Result struct {
		Gender string `json:"gender"`
		Name   struct {
			Title string `json:"title"`
			First string `json:"first"`
			Last  string `json:"last"`
		} `json:"name"`
		// Otros campos que no necesitamos imprimir
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(string(responseData))
	var data map[string][]Result //por alguna razón que no entiendo el json siempre debe meterse en un slice (array)

	json.Unmarshal(responseData, &data)
	//fmt.Println(data["results"][0], "namae wa ") //imprimiendo el unico que hay
	//fmt.Println(data["results"][0].Name.First, data["results"][0].Name.Last, " es el nombre")

	result := strings.Join([]string{data["results"][0].Name.First, " ", data["results"][0].Name.Last}, "")

	return result
}
