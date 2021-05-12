# IPFS
```
 Libp2p badger implementation:
    
    client: main.go
    server: grpc/grpcserver.go
    
    client side: 
       two flags: 1. name = "need a file or file name to store in db"
       2. Id = "takes id of file which is recognize as content identifier"
  
  result : when we pass name: return id of file
           when we pass id : return file from db after checking in peer in network.
           
  
