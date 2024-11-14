package main

import (
	"context"
	"log"
	"time"

	"fsedano.net/act1/api"
	dapr "github.com/dapr/go-sdk/client"
)

func main() {

	client, err := dapr.NewClient()
	if err != nil {
		log.Printf("error create client: %s", err)
		return
	}
	ctx := context.Background()
	inv(ctx, client)

	log.Printf("Done")
}
func inv(ctx context.Context, client dapr.Client) {

	m0 := &api.ClientStub{}

	m0.SetPlate("CAR3")
	client.ImplActorClientStub(m0)
	m0.EnterLane(ctx, "pl1")
	for {
		time.Sleep(1 * time.Second)
		//log.Printf("Sleep...")
	}
	// r, err := m0.ExitLane(ctx, "pl1")
	// if err != nil {
	// 	log.Printf("exitLane0: %s", err)
	// }
	// log.Printf("m0: %s", r)

}
