package main

import (
	"fmt"
	"log"
	"net/http"
	"traefik-avahi-helper/traefik"
)

func main() {
	client, err := traefik.CreateApiClient(http.DefaultClient)
	if err != nil {
		log.Fatalln(err)
	}

	routers, err := client.GetHttpRouters()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(routers[0])
}
