package lib

import (
	"context"

	"github.com/ipfs/go-cid"
	logger "github.com/ipfs/go-log/v2"
	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p-core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

const (
	p1 = "/ip4/127.0.0.1/tcp/3000"
	p2 = "/ip4/127.0.0.1/tcp/4000"
	p3 = "/ip4/127.0.0.1/tcp/5000"
)

type LibConfig struct {
	ctx context.Context
	dh  *dht.IpfsDHT
}

var log = logger.Logger("libp2p")

func InitPeer() (*LibConfig, error) {
	ctx := context.Background()
	host, err := libp2p.New(ctx, libp2p.ListenAddrStrings(p1))
	if err != nil {
		log.Errorf("Failed to start new node %s", err.Error())
		return nil, err
	}
	defer host.Close()
	netPeer, err := libp2p.New(ctx, libp2p.ListenAddrStrings(p2))
	if err != nil {
		log.Errorf("Failed to start new node %s", err.Error())
		return nil, err
	}
	defer netPeer.Close()
	peers := &peer.AddrInfo{
		ID:    host.ID(),
		Addrs: host.Addrs(),
	}
	mlt, err := peer.AddrInfoToP2pAddrs(peers)
	if err != nil {
		log.Errorf("Failed to get multi address %S", err.Error())
		return nil, err
	}
	peerAddInfo, err := peer.AddrInfoFromP2pAddr(mlt[0])
	if err != nil {
		log.Errorf("Failed to add peer address info %s", err.Error())
		return nil, err
	}
	err = netPeer.Connect(ctx, *peerAddInfo)
	if err != nil {
		log.Errorf("Failed to connect peers %s", err.Error())
		return nil, err
	}
	host3, err := libp2p.New(ctx, libp2p.ListenAddrStrings(p3))
	if err != nil {
		log.Errorf("Failed to create third node %s", err.Error())
		return nil, err
	}
	defer host3.Close()
	pi := &peer.AddrInfo{
		ID:    host3.ID(),
		Addrs: host3.Addrs(),
	}
	mlt, err = peer.AddrInfoToP2pAddrs(pi)
	if err != nil {
		log.Errorf("Failed to add peer address info %s", err.Error())
		return nil, err
	}
	pinfo, err := peer.AddrInfoFromP2pAddr(mlt[0])
	if err != nil {
		log.Errorf("Failed to add peer address info %s", err.Error())
		return nil, err
	}
	err = netPeer.Connect(ctx, *pinfo)
	if err != nil {
		log.Errorf("Failed to connect third node %s", err.Error())
		return nil, err
	}

	net := netPeer.Network()
	pids := net.Peers()
	log.Errorf("%s", pids)
	dh, err := dht.New(ctx, netPeer)
	if err != nil {
		log.Errorf("Failed to get dht %s", err.Error())
		return nil, err
	}
	return &LibConfig{ctx: ctx, dh: dh}, nil
}

func (lc *LibConfig) SetCid(cid []byte) error {
	err := lc.dh.Host().Peerstore().Put(lc.dh.Host().ID(), string(cid), cid)
	if err != nil {
		log.Errorf("Failed to put cid in dht %s", err)
		return err
	}
	return nil
}

func (lc *LibConfig) FindPeers(cid cid.Cid) error {
	_, err := lc.dh.GetValue(lc.ctx, cid.Hash().B58String())
	if err != nil {
		log.Errorf("failed to get cid from network %s", err.Error())
		return err
	}
	return nil
}
