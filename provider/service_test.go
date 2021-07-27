package provider

import (
	"testing"

	"github.com/sergief/terraform-provider-synology/client"
)

type SynologyClientMock struct{}

func (mockClient SynologyClientMock) Connect(host string, username string, password string) {

}

func (mockClient SynologyClientMock) Disconnect() {

}

func (mockClient SynologyClientMock) CreateFolder(folderPath string, name string, forceParent bool, additional string) client.CreateFolderResponse {
	return client.CreateFolderResponse{}
}

func (mockClient SynologyClientMock) Download(path string) []byte {
	return []byte("Test response")
}

func (mockClient SynologyClientMock) Delete(path string, recursive bool) {

}

func (clmockClientient SynologyClientMock) Upload(path string, createParents bool, overwrite bool, fileName string, fileContents []byte) {
}

func TestServiceConnect(t *testing.T) {
	service := FileItemService{synologyClient: SynologyClientMock{}}

	service.Create("/test/filename.txt", []byte("file contents"))
}
