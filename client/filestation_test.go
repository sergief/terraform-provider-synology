package client

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestFilestationInfo(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	apiInfo, err := Info(srv.URL, "all")
	if err != nil {
		t.Error(err)
	}
	response, errLogin := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie")
	if errLogin != nil {
		t.Error(errLogin)
	}
	sid := response.Sid

	for _, c := range []struct {
		address  string
		expected FileStationInfo
	}{
		{srv.URL, FileStationInfo{Is_manager: true, Support_sharing: true, Hostname: "TEST-DS216J"}},
	} {

		got, errInfo := GetFileStationInfo(apiInfo, c.address, sid)
		if errInfo != nil {
			t.Error(errInfo)
		}
		if got != c.expected {
			t.Errorf("Info(%q)", c.address)
		}
	}

}

func TestFilestationCreateFolder(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	apiInfo, err := Info(srv.URL, "all")
	if err != nil {
		t.Error(err)
	}

	loginResponse, errLogin := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie")
	sid := loginResponse.Sid
	if errLogin != nil {
		t.Error(errLogin)
	}

	for _, c := range []struct {
		address  string
		expected CreateFolderResponse
	}{
		{srv.URL, CreateFolderResponse{Folders: []FileStationFile{{Path: "/home/test_create_folder", Name: "test_create_folder", Isdir: true}}}},
	} {

		got, errCreateFolder := CreateFolder(apiInfo, c.address, sid, "/home", "test_create_folder", true, "")
		if errCreateFolder != nil {
			t.Error(errCreateFolder)
		}

		if len(got.Folders) != len(c.expected.Folders) {
			t.Error("Length differs")
		}

		for i, expectedElement := range c.expected.Folders {
			if expectedElement != got.Folders[i] {
				t.Error("Different values for ", i)
			}
		}
	}

}

func TestFilestationUpload(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	apiInfo, errInfo := Info(srv.URL, "all")
	if errInfo != nil {
		t.Error(errInfo)
	}

	loginResponse, errLogin := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie")
	sid := loginResponse.Sid
	if errLogin != nil {
		t.Error(errLogin)
	}

	for _, c := range []struct {
		address  string
		expected int
	}{
		{srv.URL, 200},
	} {
		log.Print(c)

		fileContents, err := ioutil.ReadFile("./test-fixtures/cat.txt")
		if err != nil {
			t.Error(err)
		}
		got, err := Upload(apiInfo, srv.URL, sid, "/home/downloaded", true, true, "cat.txt", fileContents)
		if err != nil {
			t.Error(err)
		}
		if got != 200 {
			t.Error(got)
		}
	}

}

func TestFilestationDownload(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	apiInfo, errInfo := Info(srv.URL, "all")
	if errInfo != nil {
		t.Error(errInfo)
	}

	loginResponse, errLogin := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie")
	sid := loginResponse.Sid
	if errLogin != nil {
		t.Error(errLogin)
	}

	for _, c := range []struct {
		address  string
		expected int
	}{
		{srv.URL, 200},
	} {

		_, got, downloadError := Download(apiInfo, c.address, sid, "/home/downloaded/empty.json")
		if downloadError != nil {
			t.Error(downloadError)
		}

		if len(got) != 2 {
			t.Error("Length differs")
		}

	}

}

func TestFilestationDelete(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	apiInfo, errInfo := Info(srv.URL, "all")
	if errInfo != nil {
		t.Error(errInfo)
	}

	loginResponse, errLogin := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie")
	sid := loginResponse.Sid
	if errLogin != nil {
		t.Error(errLogin)
	}
	for _, c := range []struct {
		address  string
		expected int
	}{
		{srv.URL, 200},
	} {

		got, errDelete := Delete(apiInfo, c.address, sid, "/home/downloaded/cat.txt", false)
		if errDelete != nil {
			t.Error(errDelete)
		}

		if got != 200 {
			t.Error("Wrong status code")
		}

	}

}
