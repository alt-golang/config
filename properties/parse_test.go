package properties

import (
	"fmt"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	//fmt.Println("Hello, World!")
	//
	//file, _ := os.Open("config.json")
	//defer file.Close()
	//decoder := json.NewDecoder(file)
	//var jsonConfig map[string]interface{}
	//err := decoder.Decode(&jsonConfig)
	//if err != nil {
	//	fmt.Println("error:", err)
	//}
	//fmt.Println(jsonConfig["fruits"])
	//
	//yamlString, err := os.ReadFile("config.yaml")

	//yaml.Unmarshal([]byte(string(yamlString)), &yamlConfig)
	//fmt.Println(yamlConfig)
	//
	//if _, err := os.Stat("file-exists2.file"); os.IsNotExist(err) {
	//	fmt.Printf("File does not exist\n")
	//}
	//continue program

	if _, err := os.Stat("file-exists.go"); err == nil {
		fmt.Printf("File exists\n")
	} else {
		fmt.Printf("File does not exist\n")
	}

	//os.Setenv("GO_ENV", "environment")
	//os.Setenv("GO_APP_INSTANCE", "instance")
	//os.Setenv("GO_PROFILES_ACTIVE", "1,2")
	//var config map[string]interface{}
	//config = loader.LoadConfigWithDir("test/config/properties")
	//fmt.Println(config)
	fmt.Printf("The end\n")
	t.Log("Meh")
}
