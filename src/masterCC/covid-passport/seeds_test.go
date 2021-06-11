package main

import (
	"testing"
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
