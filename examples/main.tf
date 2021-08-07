terraform {
  required_providers {
    synology = {
      version = "0.1"
      source = "github.com/sergief/synology"
    }
  }
}

provider "synology" {
    url = "http://192.168.1.5:5000"
    #username = "test"
    #password = "test"
    # these variables can be set as env vars in SYNOLOGY_USERNAME and SYNOLOGY_PASSWORD
}

resource "synology_file" "hello-world" {
  filename = "/home/downloaded/hello-world.txt"
  content = "Hello World"
}

resource "synology_file" "hello-world-from-file" {
  filename = "/home/downloaded/hello-world-ref.txt"
  content = file("./hello-world.txt")
}