package main

import (
	"encoding/json"
	"fmt"
	"strconv"

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
	_ = c.BenchmarkUploadDhpValid(stub, []string{"0"})
	return c.BenchmarkVerifyResult(stub, []string{"0"})
}

func (c *CovidPassportChaincode) BenchmarkUploadDhpValid(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Input Validation
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	ix, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error(fmt.Sprintf("Error parsing DHP seed index: %s", err))
	}

	if err := loadSeeds(); err != nil {
		return shim.Error(fmt.Sprintf("Error seeding DHP: %s", err))
	}

	dhp1 := seeds.ValidDhps[ix%len(seeds.ValidDhps)]
	dhp1B, err := json.Marshal(&dhp1)
	if err != nil {
		return shim.Error(fmt.Sprintf("Error marshaling dhp1: %s", err))
	}

	return c.UploadDhp(stub, []string{string(dhp1B)})
}

func (c *CovidPassportChaincode) BenchmarkVerifyResult(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// Input Validation
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	ix, err := strconv.Atoi(args[0])
	if err != nil {
		return shim.Error(fmt.Sprintf("Error parsing DHP seed index: %s", err))
	}

	if err := loadSeeds(); err != nil {
		return shim.Error(fmt.Sprintf("Error seeding DHP: %s", err))
	}

	dhp1 := seeds.ValidDhps[ix%len(seeds.ValidDhps)]

	return c.VerifyResult(stub, []string{string(dhp1.Data.Patient), string(dhp1.Data.Method)})
}
