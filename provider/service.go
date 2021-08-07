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

type FolderItemService struct {
	synologyClient client.SynologyClient
}

func (service FolderItemService) Create(path string) error {
	log.Println("Create Folder" + string(path))
	basePath, name := getPathAndFilenameFromFullPath(path)
	_, error := service.synologyClient.CreateFolder(basePath, name, true, "")
	return error
}

func (service FolderItemService) Delete(path string) error {
	log.Println("Delete Folder" + string(path))
	return service.synologyClient.Delete(path, true)

}
