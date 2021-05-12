package main

import (
	"context"
	"flag"
	"fmt"
	serv "ipfs/m/v2/grpc/pb"

	logger "github.com/ipfs/go-log/v2"
	"google.golang.org/grpc"
)

const (
	port = "8080"
)

var log = logger.Logger("client")

func main() {
	ctx := context.Background()
	name := flag.String("name", " ", "data name")
	id := flag.String("id", "", "Id of a data")
	flag.Parse()
	conn, err := grpc.Dial(fmt.Sprintf("localHost:%s", port), grpc.WithInsecure())
	if err != nil {
		log.Errorf("Failed to connect server %s", err.Error())
		return
	}
	client := serv.NewAddContentClient(conn)
	if *id != "" {
		peer := &serv.Peer{
			Cid: *id,
		}
		result, err := client.FindPeer(ctx, peer)
		if err != nil {
			log.Errorf("Failed to find peer and fetch value %s", err.Error())
			return
		}
		fmt.Printf("File is retrived %s", result.Id)
		return 
	}

	if *name != "" {
		f := &serv.Fileinfo{
			Filename: *name,
		}
		id, err := client.AddFile(ctx, f)
		if err != nil {
			log.Errorf("Failed to add file in data base %s", err.Error())
			return
		}
		fmt.Printf("Id of file is %s", id.Id)
		return
	}

	

}
