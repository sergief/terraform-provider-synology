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

## Terraform Resources

### File-item

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
    username = "test"
    password = "test"
    #alternatively you can also set SYNOLOGY_USERNAME and SYNOLOGY_PASSWORD as environment variables
}

resource "synology_file" "hello-world" {
  filename = "/home/downloaded/hello-world.txt"
  content = "Hello World"
}
```
