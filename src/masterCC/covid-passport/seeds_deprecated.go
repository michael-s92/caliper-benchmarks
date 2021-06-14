package main

import (
	"fmt"
	"math/rand"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func GenerateTestFacility(stub shim.ChaincodeStubInterface, id, privateKey string) error {
	privKey, err := UnmarshalPrivateKey(privateKey)
	if err != nil {
		return fmt.Errorf("Error unmarshaling pregenerated ECDSA keypair for Test Facility: %s: %w", id, err)
	}
	publicKeyStr, err := MarshalPublicKey(&privKey.PublicKey)
	if err != nil {
		return fmt.Errorf("Error marshaling ECDSA public key for Test Facility: %s: %w", id, err)
	}
	if err := stub.PutState(id, []byte(publicKeyStr)); err != nil {
		return fmt.Errorf("Error persisting Test Facility %s on ledger: %w", id, err)
	}
	return nil
}

func SeedRandomValidDhp() (Dhp, error) {
	if err := loadSeeds(); err != nil {
		return Dhp{}, err
	}
	return seeds.ValidDhps[rand.Intn(len(seeds.ValidDhps))], nil
}

func SeedFirstValidDhp() (Dhp, error) {
	if err := loadSeeds(); err != nil {
		return Dhp{}, err
	}
	return seeds.ValidDhps[0], nil
}
