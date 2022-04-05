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

	content := `#comment line
    #coment line
default=default
application=application
environment=default
env-instance=default
local=default
local-development=default
application-profile-1=default
application-profile-2=default
newline=application

step1.step2.step3=application
multiline1=1-\
  2-\
  3\
multiline2=r+\
  g+\
  b\  
l2.multiline3=1-\
  2-\
  3\
multiline4=application\
`
	config := Parse(content)
	fmt.Println(config)
	t.Log("Meh")
}
