package songkick

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

func get(requestURL string, obj interface{}) error {
	// fmt.Println(requestURL)
	rsp, err := http.Get(requestURL)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()
	bodyByte, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return err
	}

	jsonData := &obj
	err = json.Unmarshal(bodyByte, &jsonData)
	if err != nil {
		return err
	}

	return err
}

func convertPage(pageNumber int) string {
	page := strconv.Itoa(pageNumber)
	if pageNumber <= 0 {
		page = "1"
	}
	return page
}
