package provider

import (
	"log"
	"path/filepath"

	"github.com/sergief/terraform-provider-synology/client"
)

type FileItemService struct {
	synologyClient client.SynologyClient
}

func (service FileItemService) Create(filename string, contents []byte) {
	log.Println("Create " + string(contents))
	path, filename := getPathAndFilenameFromFullPath(filename)
	service.synologyClient.Upload(path, true, true, filename, contents)
}

func (service FileItemService) Read(filename string) []byte {
	return service.synologyClient.Download(filename)
}

func (service FileItemService) Update(filename string, contents []byte) []byte {
	log.Println("Update " + string(contents))

	path, filename := getPathAndFilenameFromFullPath(filename)
	service.synologyClient.Upload(path, true, true, filename, contents)
	return contents
}

func (service FileItemService) Delete(filename string) {
	service.synologyClient.Delete(filename, false)

}

func getPathAndFilenameFromFullPath(fullPath string) (string, string) {
	return filepath.Dir(fullPath), filepath.Base(fullPath)
}
