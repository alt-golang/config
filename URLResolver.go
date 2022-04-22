package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"
)

type URLResolver struct {
	config   Config
	selector Selector
	Resolver
}

func (URLResolver URLResolver) Callback(value interface{}, path string) interface{} {
	if reflect.TypeOf(value).Kind() == reflect.String &&
		URLResolver.selector.Matches(value.(string)) {

		selectedValue := URLResolver.selector.ResolveValue(value.(string))

		var urlPath = path[0:strings.LastIndex(path, ".")]
		var method, _ = URLResolver.config.GetWithDefault(urlPath+".method", "GET")
		var body, _ = URLResolver.config.Get(urlPath + ".body")
		if body == nil {
			body = ""
		}
		var headers, _ = URLResolver.config.GetWithDefault(urlPath+".headers", "")

		responseData, err := URLResolver.FetchUrlResponseData(selectedValue, method.(string), fmt.Sprint(body), headers.(map[string]interface{}))

		var result map[string]interface{}
		err = json.Unmarshal(responseData, &result)
		if err != nil {
			log.Fatalln(err)
		}

		if err == nil {
			return result
		}
	}
	return value
}

func (URLResolver URLResolver) Resolve(object interface{}, path string) (interface{}, error) {
	return ResolverMapValuesDeep(object, path, URLResolver.Callback), nil
}

func (URLResolver URLResolver) FetchUrlResponseData(url string, method string, body string, headers map[string]interface{}) ([]byte, error) {

	var result []byte
	var err error

	log.Println("Fetching url " + url)
	client := &http.Client{}
	req, err := http.NewRequest(
		strings.ToUpper(method), url,
		strings.NewReader(body))

	for k, v := range headers {
		req.Header.Add(k, v.(string))
	}

	resp, err := client.Do(req)
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	result = out
	return result, err
}

//func main() {
//	apikey := "PH9ppknzxFFzGSNyiaVU"
//	response, err := fetchUrlResponseData(
//		"https://app.cloudability.com/api/2/reporting/compare/enqueue?viewId=33620&end_date=2021-09-30&max_results=50&offset=0&order=desc&relative_period=last_month&sort_by=unblended_cost&start_date=2021-09-01&metrics[]=unblended_cost&metrics[]=total_amortized_cost&dimensions[]=vendor&dimensions[]=vendor_account_name&dimensions[]=vendor_account_identifier",
//		apikey)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	var dat map[string]interface{}
//	err = json.Unmarshal(response, &dat)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	id := int(dat["id"].(float64))
//
//	log.Println("Enqueued report is " + strconv.FormatFloat(dat["id"].(float64), 'f', -1, 64))
//
//	status := ""
//	for status != "finished" {
//		log.Println("Checking report queue state.")
//		response, err := fetchUrlResponseData(
//			"https://app.cloudability.com/api/2/reporting/compare/reports/"+strconv.Itoa(id)+"/state?viewId=33620",
//			apikey)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		s1, _ := prettyJson(response)
//		log.Println("Results:\n" + s1)
//
//		var dat map[string]interface{}
//		err = json.Unmarshal(response, &dat)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		status = dat["status"].(string)
//	}
//	log.Println("Report queue finished")
//
//	response, err = fetchUrlResponseData(
//		"https://app.cloudability.com/api/2/reporting/compare/reports/"+strconv.Itoa(id)+"/results?viewId=33620",
//		apikey)
//
//	s, _ := prettyJson(response)
//	log.Println("Results:\n" + s)
//}
//

//
//func prettyJson(buffer []byte) (string, error) {
//	var result string
//	var err error
//
//	var prettyJSON bytes.Buffer
//	json.Indent(&prettyJSON, buffer, "", "  ")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	result = prettyJSON.String()
//	return result, err
//}
