package client

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestFilestationInfo(t *testing.T) {
	srv := serverMock()
	defer srv.Close()

	apiInfo := Info(srv.URL, "all")

	sid := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie").Sid
	for _, c := range []struct {
		address  string
		expected FileStationInfo
	}{
		{srv.URL, FileStationInfo{Is_manager: true, Support_sharing: true, Hostname: "TEST-DS216J"}},
	} {

		got := GetFileStationInfo(apiInfo, c.address, sid)

		if got != c.expected {
			t.Errorf("Info(%q)", c.address)
		}
	}

}

func TestFilestationCreateFolder(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	apiInfo := Info(srv.URL, "all")

	sid := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie").Sid
	for _, c := range []struct {
		address  string
		expected CreateFolderResponse
	}{
		{srv.URL, CreateFolderResponse{Folders: []FileStationFile{{Path: "/home/test_create_folder", Name: "test_create_folder", Isdir: true}}}},
	} {

		got := CreateFolder(apiInfo, c.address, sid, "/home", "test_create_folder", true, "")
		// TODO: check array length
		// check == in position by position
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
	apiInfo := Info(srv.URL, "all")

	sid := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie").Sid
	for _, c := range []struct {
		address  string
		expected int
	}{
		{srv.URL, 200},
	} {
		log.Print(c)

		fileContents, err := ioutil.ReadFile("./test-fixtures/cat.jpg")
		if err != nil {
			t.Error(err)
		}
		got := Upload(apiInfo, srv.URL, sid, "/home/downloaded", true, true, "cat.jpg", fileContents)

		if got != 200 {
			t.Error(got)
		}
	}

}

func TestFilestationDownload(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	apiInfo := Info(srv.URL, "all")

	sid := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie").Sid
	for _, c := range []struct {
		address  string
		expected int
	}{
		{srv.URL, 200},
	} {

		_, got := Download(apiInfo, c.address, sid, "/home/downloaded/empty.json")

		if len(got) != 2 {
			t.Error("Length differs")
		}

	}

}

func TestFilestationDelete(t *testing.T) {
	srv := serverMock()
	defer srv.Close()
	apiInfo := Info(srv.URL, "all")

	sid := Login(apiInfo, srv.URL, "testuser", "password1234", "TestAuth", "cookie").Sid
	for _, c := range []struct {
		address  string
		expected int
	}{
		{srv.URL, 200},
	} {

		got := Delete(apiInfo, c.address, sid, "/home/downloaded/cat.jpg", false)

		if got != 200 {
			t.Error("Wrong status code")
		}

	}

}
