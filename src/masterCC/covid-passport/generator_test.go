package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestGenerateTestFacilityKeys(t *testing.T) {
	varCodeBlock := `
	var (
	%s
	)
	`
	tfKeyLine := `	tfKey%d string = "%s"
	`

	var tfKeyCodeBlock string
	for i := 1; i <= 3; i++ {
		privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
		if err != nil {
			t.Errorf("Error generating ECDSA keypair: %s", err)
		}
		pkStr, err := MarshalPrivateKey(privateKey)
		if err != nil {
			t.Error(err)
		}

		tfKeyCodeBlock += fmt.Sprintf(tfKeyLine, i, pkStr)
	}
	t.Logf(varCodeBlock, tfKeyCodeBlock)
}

func TestUnmarshalPregeneratedKeys(t *testing.T) {
	for key, tf := range TfKeys {
		if _, err := UnmarshalPrivateKey(tf); err != nil {
			t.Errorf("Error unmarshaling pregenerated ECDSA keypair for Test Facility: %s: %w", key, err)
		}
	}
}

func TestDhpSignVerify(t *testing.T) {
	// TF-1 key
	tf1PrivKey, err := UnmarshalPrivateKey(TfKeys["TF-1-Theresienwiese"])
	if err != nil {
		t.Errorf("Error unmarshaling pregenerated ECDSA keypair for Test Facility TF-1: %s", err)
	}
	tf1PubKey := tf1PrivKey.PublicKey

	// Test Patient: Milan
	dhp1, err := GenerateDhp("001", "TF-1-Theresienwiese", tf1PrivKey, "Milan", "PCR", true, time.Time{}, time.Time{})
	if err != nil {
		t.Errorf("Error generating dhp1: %s", err)
	}

	// DEBUG
	t.Logf("Date: %s", dhp1.Data.Date)
	t.Logf("ExpiryDate: %s", dhp1.Data.ExpiryDate)

	// Simulate PubKey Put/Set on channel
	tmp1, err := MarshalPublicKey(&tf1PubKey)
	if err != nil {
		t.Errorf("Error marshaling public key: %s", err)
	}
	tmp2 := []byte(tmp1)
	tmp3 := string(tmp2)
	issCrt, err := UnmarshalPublicKey(tmp3)
	if err != nil {
		t.Errorf("Error unmarshaling public key: %s", err)
	}

	// Validate signature
	data, err := json.Marshal(&dhp1.Data)
	if err != nil {
		t.Errorf("Error marshaling TestResult data inside DHP: %s", err)
	}
	if !ValidateSignature(data, dhp1.Signature, issCrt) {
		t.Errorf("Signature validation failed! \n Issuer: %s \n Signature: %#v \n TestResult: %#v", *issCrt, dhp1.Signature, data)
	}

	// DEBUG
	t.Logf("Issuer: %s \n Signature: %#v \n TestResult: %#v", *issCrt, dhp1.Signature, data)

}
