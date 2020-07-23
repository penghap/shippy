package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/micro/go-micro/v2"
	pb "github.com/penghap/shippy/service-consignment/proto/consignment"
)

const (
	defaultFilename = "consignment.json"
)

func parseFile(file string) (*pb.Consignment, error) {
	var consignment *pb.Consignment
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &consignment)
	return consignment, err
}

const (
	srvName            = "shippy.service.cli"
	srvVersion         = "latest"
	consignmentSrvName = "shippy.service.consignment"
)

func main() {
	// NewShippingServiceClient
	// Set up a connection to the server.
	srv := micro.NewService(
		micro.Name(srvName),
		micro.Version(srvVersion),
	)
	srv.Init()

	// Create new greeter client
	client := pb.NewShippingService(consignmentSrvName, srv.Client())

	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	consignment, err := parseFile(file)
	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := client.CreateConsignment(ctx, consignment)
	if err != nil {
		log.Fatalf("Could not create: %v", err)
	}

	log.Printf("Created: %t", r.Created)

	resp, err := client.GetConsignments(ctx, &pb.Consignment{})
	if err != nil {
		log.Fatalf("Failed to list consignments: %v", err)
	}

	for _, consignment := range resp.Consignments {
		log.Printf("%+v", consignment)
	}
}
