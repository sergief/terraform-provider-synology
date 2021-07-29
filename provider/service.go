package provider

import (
	"log"
	"path/filepath"

	"github.com/sergief/terraform-provider-synology/client"
)

type FileItemService struct {
	synologyClient client.SynologyClient
}

func (service FileItemService) Create(filename string, contents []byte) error {
	log.Println("Create " + string(contents))
	path, filename := getPathAndFilenameFromFullPath(filename)
	return service.synologyClient.Upload(path, true, true, filename, contents)
}

func (service FileItemService) Read(filename string) ([]byte, error) {
	return service.synologyClient.Download(filename)
}

func (service FileItemService) Update(filename string, contents []byte) ([]byte, error) {
	log.Println("Update " + string(contents))

	path, filename := getPathAndFilenameFromFullPath(filename)
	err := service.synologyClient.Upload(path, true, true, filename, contents)
	return contents, err
}

func (service FileItemService) Delete(filename string) error {
	return service.synologyClient.Delete(filename, false)

}

func getPathAndFilenameFromFullPath(fullPath string) (string, string) {
	return filepath.Dir(fullPath), filepath.Base(fullPath)
}
