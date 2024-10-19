package main

import (
	"fmt"
	"hdu/internal/registry_client"
	"hdu/internal/utils"
	"log"
)

func main() {

	client := registry_client.NewRegistryClient("https://localhost:5000")
	catalog, err := client.GetCatalog(1)
	_check_error(err)
	utils.PrintAsJson(catalog)

	if len(catalog.Repositories) > 0 {

		repo, err := client.GetRepository(catalog.Repositories[0])
		_check_error(err)

		fmt.Println(repo)

		if len(repo.Tags) > 0 {

			manifest, err := client.GetManivestV2(repo.Name, repo.Tags[0])
			_check_error(err)

			utils.PrintAsJson(manifest)
		}

	}
}

func _check_error(err error) {
	if err != nil {
		fmt.Println("Error ------------------------------")
		log.Panic(err)
		fmt.Println("Error ------------------------------")
	}
}
