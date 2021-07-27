package client

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
)

func serverMock() *httptest.Server {
	handler := http.NewServeMux()
	handler.HandleFunc("/webapi/query.cgi", queryMock)
	handler.HandleFunc("/webapi/auth.cgi", authMock)
	handler.HandleFunc("/webapi/entry.cgi", entryMock)

	srv := httptest.NewServer(handler)

	return srv
}

func authMock(w http.ResponseWriter, r *http.Request) {

	fixtureContent, err := ioutil.ReadFile("test-fixtures/login.json")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write([]byte(fixtureContent))
}

func entryMock(w http.ResponseWriter, r *http.Request) {

	keys := r.URL.Query()["api"]
	api := keys[0]

	var filesByApi = map[string]string{
		"SYNO.FileStation.Info":         "test-fixtures/filestation_info.json",
		"SYNO.FileStation.CreateFolder": "test-fixtures/filestation_create_folder.json",
		"SYNO.FileStation.Upload":       "test-fixtures/upload-success.json",
		"SYNO.FileStation.Download":     "test-fixtures/empty.json",
		"SYNO.FileStation.Delete":       "test-fixtures/empty.json",
	}

	fixtureContent, err := ioutil.ReadFile(filesByApi[api])
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write([]byte(fixtureContent))
}

func queryMock(w http.ResponseWriter, r *http.Request) {

	fixtureContent, err := ioutil.ReadFile("test-fixtures/info.json")
	if err != nil {
		log.Fatal(err)
	}
	_, _ = w.Write([]byte(fixtureContent))
}
