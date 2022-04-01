package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	var jsonConfig map[string]interface{}
	err := decoder.Decode(&jsonConfig)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(jsonConfig["fruits"])

	yamlString, err := os.ReadFile("config.yaml")
	var yamlConfig map[string]interface{}
	yaml.Unmarshal([]byte(string(yamlString)), &yamlConfig)
	fmt.Println(yamlConfig)

	if _, err := os.Stat("file-exists2.file"); os.IsNotExist(err) {
		fmt.Printf("File does not exist\n")
	}
	// continue program

	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}

	var configLoader = new(ConfigLoader)
	configLoader.LoadConfig()

}
