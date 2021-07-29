package client

import (
	"bytes"
	"encoding/json"
	"strconv"
	"strings"
)

type FileStationInfo struct {
	Is_manager               bool
	Support_sharing          bool
	Support_virtual_protocol string
	Hostname                 string
}

type CreateFolderResponse struct {
	Folders []FileStationFile
}

type FileStationFile struct {
	Path       string
	Name       string
	Isdir      bool
	Additional interface{}
}

func GetFileStationInfo(apiInfo map[string]InfoData, host string, sid string) (FileStationInfo, error) {

	queryString := make(map[string]string)
	queryString["method"] = "get"
	queryString["_sid"] = sid

	_, apiResponse, err := CallAPI(host, "SYNO.FileStation.Info", apiInfo, queryString)

	if err != nil {
		return FileStationInfo{}, err
	}
	var fileStationInfo FileStationInfo

	json.Unmarshal(apiResponse.Data, &fileStationInfo)

	return fileStationInfo, nil
}

func CreateFolder(apiInfo map[string]InfoData, host string, sid string, folderPath string, name string, forceParent bool, additional string) (CreateFolderResponse, error) {
	queryString := make(map[string]string)
	queryString["method"] = "create"
	queryString["_sid"] = sid
	queryString["folder_path"] = folderPath
	queryString["name"] = name
	queryString["force_parent"] = strconv.FormatBool(forceParent)
	if additional != "" {
		queryString["additional"] = additional
	}

	_, apiResponse, err := CallAPI(host, "SYNO.FileStation.CreateFolder", apiInfo, queryString)
	if err != nil {
		return CreateFolderResponse{}, err
	}

	var createFolderResponse CreateFolderResponse
	json.Unmarshal(apiResponse.Data, &createFolderResponse)

	return createFolderResponse, nil
}

func Download(apiInfo map[string]InfoData, host string, sid string, path string) (int, []byte, error) {
	apiName := "SYNO.FileStation.Download"
	info := apiInfo[apiName]

	queryString := make(map[string]string)

	queryString["method"] = "download"
	queryString["mode"] = "download"
	queryString["_sid"] = sid
	queryString["path"] = path
	queryString["api"] = apiName
	queryString["version"] = strconv.Itoa(info.MaxVersion)
	wsUrl := host + "/webapi/" + info.Path

	statusCode, body, err := HttpCall(wsUrl, queryString)

	return statusCode, body, err
}

func Delete(apiInfo map[string]InfoData, host string, sid string, path string, recursive bool) (int, error) {
	apiName := "SYNO.FileStation.Delete"
	info := apiInfo[apiName]

	queryString := make(map[string]string)

	queryString["method"] = "delete"
	queryString["mode"] = "open"
	queryString["_sid"] = sid
	queryString["path"] = path
	queryString["recursive"] = strconv.FormatBool(recursive)

	queryString["api"] = apiName
	queryString["version"] = strconv.Itoa(info.MaxVersion)
	wsUrl := host + "/webapi/" + info.Path

	statusCode, _, err := HttpCall(wsUrl, queryString)

	return statusCode, err
}

func Upload(apiInfo map[string]InfoData, host string, sid string, path string, createParents bool, overwrite bool, fileName string, fileContents []byte) (int, error) {
	apiName := "SYNO.FileStation.Upload"
	info := apiInfo[apiName]
	formParams := []keyValuePair{
		{key: "api", value: strings.NewReader(apiName), isFile: false},
		{key: "version", value: strings.NewReader(strconv.Itoa(info.MinVersion)), isFile: false},
		{key: "method", value: strings.NewReader("upload"), isFile: false},
		{key: "path", value: strings.NewReader(path), isFile: false},
		{key: "create_parents", value: strings.NewReader(strconv.FormatBool(createParents)), isFile: false},
		{key: "overwrite", value: strings.NewReader(strconv.FormatBool(overwrite)), isFile: false},
		{key: "file", value: bytes.NewReader(fileContents), isFile: true, fileName: fileName},
	}

	wsUrl := host + "/webapi/" + info.Path + "?api=" + apiName + "&version=" + strconv.Itoa(info.MinVersion) + "&method=upload&_sid=" + sid

	statusCode, _, err := HttpPostMultiFormCall(wsUrl, formParams)

	return statusCode, err
}
