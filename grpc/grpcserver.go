package main

import (
	"context"
	"errors"
	"fmt"
	"net"

	serv "ipfs/m/v2/grpc/pb"

	"ipfs/m/v2/lib"
	"ipfs/m/v2/store"

	badger "github.com/dgraph-io/badger/v3"
	"github.com/golang/protobuf/proto"
	"github.com/ipfs/go-cid"
	logger "github.com/ipfs/go-log/v2"
	"google.golang.org/grpc"
)

var log = logger.Logger("grpc/server")

const (
	port = "8080"
)

type server struct {
	DB  *badger.DB
	dht *lib.LibConfig
}

func main() {
	// logger.SetLogLevel("*", "debug")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", port))
	if err != nil {
		log.Errorf("Failed to create server %s", err)
		return
	}
	db, err := store.DBIntialization()
	if err != nil {
		log.Errorf("Failed to intialize db %s", err.Error())
		return
	}

	grp := grpc.NewServer()
	dht, err := lib.InitPeer()
	if err != nil {
		log.Errorf("Failed to get dht instance %s", err.Error())
		return
	}
	log.Infof("Server is runnning on port:%s", port)
	serv.RegisterAddContentServer(grp, &server{DB: db, dht: dht})
	grp.Serve(lis)
}

func (sr *server) AddFile(ctx context.Context, f *serv.Fileinfo) (*serv.Id, error) {
	file := &store.File{
		Name: f.Filename,
	}
	sc := &store.StoreConf{
		DB: sr.DB,
	}
	cid, err := sc.Create(file)
	if err != nil {
		log.Errorf("Failed to store in db %s", err.Error())
		return nil, err
	}
	id := &serv.Id{}
	err = proto.Unmarshal(cid, id)
	if err != nil {
		log.Errorf("Failed to unmarshal data %s", err.Error())
		return nil, err
	}
	log.Errorf("string %s", id.Id)
	return id, nil
}

func (sr *server) FindPeer(ctx context.Context, peer *serv.Peer) (*serv.Id, error) {
	file := &store.ContentBased{
		Id: peer.Cid,
	}
	sc := &store.StoreConf{
		DB: sr.DB,
	}
	id, err := cid.Cast([]byte(peer.Cid))
	if err != nil {
		log.Errorf("Failed to convert into cid %s", err.Error())
		return nil, err
	}
	err = sr.dht.FindPeers(id)
	if err != nil {
		log.Errorf("Failed to find cid in network %s", err.Error())
		return nil, err
	}
	item, err := sc.Read(file)
	if err != nil {
		log.Errorf("Failed to read to value from database %s", err.Error())
		return nil, err
	}
	f, ok := item.(*store.File)
	if !ok {
		return nil, errors.New("item is not type of file %s")
	}
	return &serv.Id{Id: f.Name}, nil
}
