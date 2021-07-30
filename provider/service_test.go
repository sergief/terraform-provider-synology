package provider

import (
	"testing"

	"github.com/sergief/terraform-provider-synology/client"
)

type SynologyClientMock struct{}

func (mockClient SynologyClientMock) Connect(host string, username string, password string) error {
	return nil
}

func (mockClient SynologyClientMock) Disconnect() error {
	return nil
}

func (mockClient SynologyClientMock) CreateFolder(folderPath string, name string, forceParent bool, additional string) (client.CreateFolderResponse, error) {
	return client.CreateFolderResponse{}, nil
}

func (mockClient SynologyClientMock) Download(path string) ([]byte, error) {
	return []byte("Test response"), nil
}

func (mockClient SynologyClientMock) Delete(path string, recursive bool) error {
	return nil
}

func (mockClientient SynologyClientMock) Upload(path string, createParents bool, overwrite bool, fileName string, fileContents []byte) error {
	return nil
}

func TestServiceConnect(t *testing.T) {
	service := FileItemService{synologyClient: SynologyClientMock{}}

	service.Create("/test/filename.txt", []byte("file contents"))
}
