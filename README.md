# hyperledger-fabric-400
hyperledger fabric in less than 400 lines of Go

# tutorial
https://hamait.tistory.com/1012

# run 
leader peer : go run peer.go -name 28000 -port 28000 -leader
oleader peer : go run peer.go -name 28001 -port 28001 -bootstrap 28000
