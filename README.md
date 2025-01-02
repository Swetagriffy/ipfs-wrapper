IPFS:
A decentralized file storage system where files are stored and referenced by their content hash (CID).


1: Install Go:
Download and install Go from the official Go website.
Verify the installation by running:
go version

2: Install IPFS:
Download and install IPFS from the official IPFS website.
Verify the installation by running:
ipfs --version

3: Start the IPFS Daemon:
Initialize IPFS (if not already initialized):
ipfs init
Start the IPFS daemon to enable the HTTP API:
ipfs daemon
By default, the API is accessible at http://localhost:5001.

4: Create a new directory for the project:
   mkdir ipfs-client && cd ipfs-client
5: Create a Go file:
   touch main.go

-----Run the Program
-Initialize the Go Module:
go mod init ipfs-client

----Run the Program:
go run main.go
