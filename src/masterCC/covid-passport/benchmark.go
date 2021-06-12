package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func (c *CovidPassportChaincode) InitLedger(stub shim.ChaincodeStubInterface) pb.Response {
	for key, tf := range TfKeys {
		if err := GenerateTestFacility(stub, key, tf); err != nil {
			return shim.Error(err.Error())
		}
	}

	return shim.Success(nil)
}

func (c *CovidPassportChaincode) DoNothing(stub shim.ChaincodeStubInterface) pb.Response {
	return c.TestUploadDhpValid(stub)
}

func (c *CovidPassportChaincode) TestUploadDhpValid(stub shim.ChaincodeStubInterface) pb.Response {
	dhp1, err := SeedFirstValidDhp()
	if err != nil {
		return shim.Error(fmt.Sprintf("Error seeding random valid DHP: %s", err))
	}

	dhp1B, err := json.Marshal(&dhp1)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error marshaling dhp1: %s", err))
	}

	return c.UploadDhp(stub, []string{string(dhp1B)})
}
