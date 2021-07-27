package client

import (
	"encoding/json"
)

type InfoData struct {
	Path          string
	MinVersion    int
	MaxVersion    int
	RequestFormat string
}

type LoginResponse struct {
	Sid string
}

func Info(host string, query string) map[string]InfoData {
	url := host + "/webapi/query.cgi"

	queryParams := make(map[string]string)
	queryParams["query"] = query
	queryParams["api"] = "SYNO.API.Info"
	queryParams["version"] = "1"
	queryParams["method"] = "query"

	_, response := Call(url, queryParams)
	var info map[string]InfoData
	json.Unmarshal(response.Data, &info)

	return info

}

func Login(apiInfo map[string]InfoData, host string, account string, passwd string, session string, format string) LoginResponse {
	queryString := make(map[string]string)
	queryString["account"] = account
	queryString["passwd"] = passwd
	queryString["session"] = session
	queryString["format"] = format
	queryString["method"] = "login"

	_, apiResponse := CallAPI(host, "SYNO.API.Auth", apiInfo, queryString)

	var loginResponse LoginResponse

	json.Unmarshal(apiResponse.Data, &loginResponse)

	return loginResponse

}

func Logout(apiInfo map[string]InfoData, host string, session string) bool {

	queryString := make(map[string]string)
	queryString["session"] = session
	queryString["method"] = "logout"

	statusCode, _ := CallAPI(host, "SYNO.API.Auth", apiInfo, queryString)

	return statusCode == 200
}
