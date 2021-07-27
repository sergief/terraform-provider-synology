package client

import (
	"log"
)

type SynologyClient interface {
	Connect(host string, username string, password string)
	Disconnect()
	CreateFolder(folderPath string, name string, forceParent bool, additional string) CreateFolderResponse
	Download(path string) []byte
	Delete(path string, recursive bool)
	Upload(path string, createParents bool, overwrite bool, fileName string, fileContents []byte)
}

type synologyClient struct {
	host    string
	apiInfo map[string]InfoData
	sid     string
}

func (client *synologyClient) Connect(host string, username string, password string) {
	log.Println("Connect")

	client.host = host
	client.apiInfo = Info(client.host, "all")
	client.sid = Login(client.apiInfo, client.host, username, password, "sergief-synology-client", "cookie").Sid
}

func (client synologyClient) Disconnect() {
	log.Println("Disconnect")

	Logout(client.apiInfo, client.host, client.sid)
}

func (client synologyClient) CreateFolder(folderPath string, name string, forceParent bool, additional string) CreateFolderResponse {
	log.Println("CreateFolder")

	return CreateFolder(client.apiInfo, client.host, client.sid, folderPath, name, forceParent, additional)
}

func (client synologyClient) Download(path string) []byte {
	log.Println("Download")

	statusCode, body := Download(client.apiInfo, client.host, client.sid, path)
	if statusCode != 200 {
		return []byte("")
	}

	return body
}

func (client synologyClient) Delete(path string, recursive bool) {
	log.Println("Delete")

	Delete(client.apiInfo, client.host, client.sid, path, recursive)
}

func (client synologyClient) Upload(path string, createParents bool, overwrite bool, fileName string, fileContents []byte) {
	log.Println("Upload path=" + path + " filename=" + fileName)

	statusCode := Upload(client.apiInfo, client.host, client.sid, path, createParents, overwrite, fileName, fileContents)
	log.Println(statusCode)
}

func NewClient() SynologyClient {
	return &synologyClient{}
}
