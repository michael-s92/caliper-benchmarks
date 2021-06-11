package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"path"
	"strconv"
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

func Test_LoadSeeds(t *testing.T) {
	if err := loadSeeds(); err != nil {
		t.Fatal(err)
	}

	if len(seeds.ValidDhps) < 1 {
		t.Fatalf("No dhps loaded from seeds: %#v", *seeds)
	}
}

func Test_GenerateSeeds(t *testing.T) {
	if err := generateSeeds(); err != nil {
		t.Fatal(err)
	}
}

func generateSeeds() error {
	// Seed random gen
	rand.Seed(time.Now().UnixNano())

	// Read seed params
	seedParamsB, err := ioutil.ReadFile(path.Join("hack", "seed", "seedParameters.yaml"))
	if err != nil {
		return fmt.Errorf("Error reading seedParameters.yaml: %w", err)
	}
	var seedParams SeedParameters
	if err := yaml.Unmarshal(seedParamsB, &seedParams); err != nil {
		return fmt.Errorf("Error unmarshaling seedParameters.yaml: %w", err)
	}

	// Initialize Seeds
	seeds := Seeds{}

	// Generate Valid DHPs
	for i := 0; i < seedParams.AllValidDhp; i++ {
		dhp, err := generateValidDhp(i)
		if err != nil {
			return fmt.Errorf("Error generating Valid DHP #%d: %w", i, err)
		}
		seeds.ValidDhps = append(seeds.ValidDhps, *dhp)
	}

	// Store seeds.json
	seedsJson, err := json.Marshal(&seeds)
	if err != nil {
		return fmt.Errorf("Error marshaling generated seeds: %w", err)
	}
	if err := ioutil.WriteFile(path.Join("hack", "seed", "seeds.json"), seedsJson, 0644); err != nil {
		return fmt.Errorf("Error writing seeds.json: %w", err)
	}

	return nil
}

func generateValidDhp(ix int) (*Dhp, error) {
	var tfId string
	r := rand.Intn(len(TfKeys))
	for k := range TfKeys {
		if r == 0 {
			tfId = k
		}
		r--
	}
	tfKey, err := UnmarshalPrivateKey(TfKeys[tfId])
	if err != nil {
		return nil, fmt.Errorf("Error unmarshaling private key: %w", err)
	}
	patientDid := PatientDids[rand.Intn(len(PatientDids))]

	// fmt.Printf("PatientDID: %s", patientDid)
	// fmt.Printf("SHA256(PatientDID: %s", (ccrypto.Hash([]byte(patientDid))))

	return GenerateDhp(strconv.Itoa(ix), tfId, tfKey, patientDid,
		TestTypes[rand.Intn(len(TestTypes))], rand.Intn(1) == 1,
		time.Now().AddDate(0, 0, -1).Truncate(24*time.Hour),
		time.Now().AddDate(0, 0, 2).Truncate(24*time.Hour),
	)
}
