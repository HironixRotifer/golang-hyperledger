./network.sh down
./network.sh up createChannel -c test -ca
./network.sh deployCC -ccn basic -ccp ../chaincode/ -ccv 1 -ccl go -c test