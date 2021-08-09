# Terraform Provider Synology

Provider for managing Synology Resources.

## Build and install
### Using the local environment
If you have the `Go` environment installed, you can simply run the makefile:
```bash
make clean build test release
```
### Using the docker-compose image
You don't need to install the golang development environment if you already have a working `docker` environment, simply run:
```bash
docker-compose run app make clean build test release
```

After this, pick up the version specifically compiled for your OS and architecture from `./bin` and put it in `$HOME/.terraform.d/plugins/github.com/sergief/synology/0.1/$OS_$ARCH/terraform-provider-synology`


## Acceptance Tests
Run the following command setting the required environment variables (no docker support):
```bash
SYNOLOGY_ADDRESS=http://aaa.bbb.ccc.dddd:5000 SYNOLOGY_USERNAME=test_user SYNOLOGY_PASSWORD=test_password make testacc
```

## Terraform Resources

### File

This resource creates a text file in a Synology Filestation.
Example:
```terraform
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
    username = "testuser"
    password = "testpass"
    # these variables can be set as env vars in SYNOLOGY_ADDRESS SYNOLOGY_USERNAME and SYNOLOGY_PASSWORD
}

resource "synology_file" "hello-world" {
  filename = "/home/downloaded/hello-world.txt"
  content = "Hello World"
}
```

### Folder

This resource creates a folder in a Synology Filestation.
Example:
```terraform
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
    username = "test"
    password = "test"
    # these variables can be set as env vars in SYNOLOGY_ADDRESS SYNOLOGY_USERNAME and SYNOLOGY_PASSWORD
}

resource "synology_folder" "my-folder" {
  path = "/home/downloaded/sample-folder"
}
```
