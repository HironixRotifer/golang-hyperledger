package main

import (
	"log"

	cn "github.com/HironixRotifer/golang-hyperledger/chaincode/contract"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	contract, err := contractapi.NewChaincode(&cn.Contract{})
	if err != nil {
		log.Panicf("error creating chaincode: %v", err)
	}

	if err := contract.Start(); err != nil {
		log.Panicf("error starting chaincode: %v", err)
	}
}
