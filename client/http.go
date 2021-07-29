package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	Success bool
	Data    json.RawMessage
}

func HttpCall(url string, queryParams map[string]string) (int, []byte, error) {
	httpClient := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, nil, err
	}
	q := req.URL.Query()
	for k, v := range queryParams {
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()

	log.Print("Call url=" + req.URL.String())

	resp, err := httpClient.Do(req)
	if err != nil {
		return 0, nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, nil, err
	}

	return resp.StatusCode, body, nil
}

func Call(url string, queryParams map[string]string) (int, ApiResponse, error) {
	statusCode, body, err := HttpCall(url, queryParams)
	if err != nil {
		return 0, ApiResponse{}, err
	}

	log.Println(string(body))

	var result ApiResponse
	json.Unmarshal(body, &result)
	if !result.Success {
		return statusCode, result, fmt.Errorf("Error retrieving server data: %s", string(body))
	}

	return statusCode, result, nil
}

type keyValuePair struct {
	key      string
	value    io.Reader
	isFile   bool
	fileName string
}

func HttpPostMultiFormCall(url string, formParams []keyValuePair) (int, ApiResponse, error) {
	tr := &http.Transport{
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}

	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for _, keyvalue := range formParams {
		var fw io.Writer
		var err error
		key := keyvalue.key
		r := keyvalue.value
		var buf []byte
		isFile := keyvalue.isFile
		if x, ok := r.(io.Closer); ok {
			x.Close()
		}
		if isFile {
			if fw, err = w.CreateFormFile(key, keyvalue.fileName); err != nil {
				return 0, ApiResponse{}, err
			}
			if buf, err = ioutil.ReadAll(r); err != nil {
				return 0, ApiResponse{}, err
			}
			fw.Write(buf)

			log.Println(string(buf))
			w.Close()

		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return 0, ApiResponse{}, err

			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return 0, ApiResponse{}, err
		}

	}

	req, err := http.NewRequest("POST", url, &b)
	if err != nil {
		return 0, ApiResponse{}, err
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	log.Print("Call url=" + req.URL.String())

	// Submit the request
	res, err := client.Do(req)
	if err != nil {
		return 0, ApiResponse{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, ApiResponse{}, err
	}

	w.Close()

	var result ApiResponse
	json.Unmarshal(body, &result)
	if !result.Success {
		return 0, ApiResponse{}, fmt.Errorf("Error retrieving server data: %s", string(body))
	}
	return res.StatusCode, result, nil
}

func CallAPI(url string, apiName string, apiInfo map[string]InfoData, queryParams map[string]string) (int, ApiResponse, error) {
	info := apiInfo[apiName]
	queryParams["api"] = apiName
	queryParams["version"] = strconv.Itoa(info.MaxVersion)
	return Call(url+"/webapi/"+info.Path, queryParams)

}
