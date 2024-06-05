package cmd

import (
	"encoding/json"
	"fmt"
	"ghar/util"
	"io/ioutil"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

func (m *Har) ParseHar(harFilePath string) (exportPath string, err error) {
	harPath, filename := filepath.Split(harFilePath)
	ext := filepath.Ext(filename)

	data, err := ioutil.ReadFile(harFilePath)
	if err != nil {
		util.XWarning(fmt.Sprintf("ParseHar ioutil.ReadFile error : %v\n", err))
		return exportPath, err

	}
	exportPath = filepath.Join(harPath, filename[:len(filename)-len(ext)]+"_js")
	os.MkdirAll(exportPath, os.ModePerm)

	jsonData := map[string]interface{}{}
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		util.XWarning(fmt.Sprintf("ParseHar json.Unmarshal error : %v\n", err))
		return exportPath, err
	}

	staticURLList := []string{}
	if jsonData["log"] == nil {
		util.XWarning(fmt.Sprintf("ParseHar log error : %v\n", err))
		return exportPath, err
	}
	logJsonData := jsonData["log"].(map[string]interface{})

	if logJsonData["entries"] == nil {
		fmt.Printf("%v\n", logJsonData["entries"])
		util.XWarning(fmt.Sprintf("ParseHar jsonData entries error : %v\n", err))
		return exportPath, err
	}
	for _, entrie := range logJsonData["entries"].([]interface{}) {
		if entrie == nil {
			continue
		}
		if entrie.(map[string]interface{})["request"] == nil {
			continue
		}
		request := entrie.(map[string]interface{})["request"].(map[string]interface{})
		if request["url"] == nil {
			continue
		}
		requestURL := request["url"].(string)
		parsedURL, err := url.Parse(requestURL)
		if err != nil {
			continue
		}
		parsedURLPath := strings.ReplaceAll(parsedURL.Path[1:], "/", "_")
		if strings.Contains(ext, ".jsc") || strings.Contains(ext, ".js") || strings.Contains(ext, ".sc") {
			if util.InArray(requestURL, staticURLList) == false {
				staticURLList = append(staticURLList, requestURL)
			}

			if entrie.(map[string]interface{})["response"] == nil {
				continue
			}
			content := entrie.(map[string]interface{})["response"].(map[string]interface{})["content"]
			if content.(map[string]interface{})["text"] == nil {

				continue
			}
			ioutil.WriteFile(filepath.Join(exportPath, parsedURLPath), []byte(content.(map[string]interface{})["text"].(string)), os.ModePerm)

		}

	}

	ioutil.WriteFile(filepath.Join(exportPath, "js_list.txt"), []byte(strings.Join(staticURLList, "\n")), os.ModePerm)
	return exportPath, nil
}
