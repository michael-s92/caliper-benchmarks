package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"path"
	"strconv"
	"time"

	"gopkg.in/yaml.v2"
)

var (
	TestTypes   []string = []string{"PCR", "Antibody"}
	PatientDids []string = []string{"Agustin", "Elane", "Randal", "Noel", "Shanel", "Piper", "Towanda", "Tameka", "Fidela", "Brittney", "Conception", "Lan", "Fawn", "Garry", "Patrick", "Arnold", "Delcie", "Kim", "Dacia", "Octavio"}
	TfKeys               = map[string]string{
		"TF-1-Theresienwiese":  "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgN33Tn1iVORYJKbMlyWK5erHetMRmbNqTbFiAzwE+vA6hRANCAASxhkksdofM3NoKHrZoFl79RMoh+tMjWIiu61mlBTvml5jcdP3X2XHxtYuCXmkjK7dT0fQw72kyojwm6vbd+clC",
		"TF-2-Sonnenstraße":    "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQgRDsFbT75TZ1mP7EIMGGfdt83VmoVyypPZVJkn9QJxWuhRANCAATwTCATI0CP2oMcA0u0swtzFDEmEQIUgd++CTF9JIBKNxqbe92g3X7saHK2puc+OSDKvVsO3TO6o5kNRZ0727o+",
		"TF-3-DeutschesMuseum": "MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg0AQREEYO75R0Dmtp1/DPa3nFWSTk/7QWpLjefY46SR6hRANCAAQTq7kswYZ6q+A6J1rw42YRGC+WACZvoQq9CnuWP3pMA58bcXsLi0eB3BwEdwtTV9zKD6qgCrxM99Qd0qBkNecA",
	}
)
var seeds *Seeds = nil

type Seeds struct {
	ValidDhps []Dhp `json:"validDHPs"`
}

type SeedParameters struct {
	AllValidDhp int `yaml:"allValidDHP,omitempty"`
}

func SeedRandomValidDhp() (Dhp, error) {
	if err := loadSeeds(); err != nil {
		return Dhp{}, err
	}
	return seeds.ValidDhps[rand.Intn(len(seeds.ValidDhps))], nil
}

func loadSeeds() error {
	if seeds == nil {
		seeds = &Seeds{}
		// seedsB, err := ioutil.ReadFile(path.Join("hack", "seed", "seeds.json"))
		// seedsB, err := ReadFileFromGCS("milan-thesis-21", "seeds.json")
		seedsB, err := loadFromUrl("https://storage.googleapis.com/milan-thesis-21/seeds.json")
		if err != nil {
			return fmt.Errorf("Error loading seeds.json: %w", err)
		}
		if err := json.Unmarshal(seedsB, seeds); err != nil {
			return fmt.Errorf("Error unmarshaling seeds: %w", err)
		}
	}
	return nil
}

func loadFromUrl(url string) ([]byte, error) {
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error loading URL: %w", err)
	}
	return ioutil.ReadAll(r.Body)
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
