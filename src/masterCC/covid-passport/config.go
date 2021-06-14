package main

import "path"

var Config = Cfg{
	DBType: LevelDB,
	// DBType: CouchDB,
	Seeds: SeedsCfg{
		InputParametersPath: path.Join("hack", "seed", "seedParameters.yaml"),
		OutputJsonPath:      path.Join("hack", "seed", "seeds.json"),
		HostedUrl:           "https://storage.googleapis.com/milan-thesis-21/covid-passport/seeds-3x30.json",
		LoadSeedsFromUrl:    true,
	},
}

type Cfg struct {
	DBType DBType
	Seeds  SeedsCfg
}

type DBType int

const (
	LevelDB DBType = iota
	CouchDB
)

type SeedsCfg struct {
	InputParametersPath string
	OutputJsonPath      string
	HostedUrl           string
	LoadSeedsFromUrl    bool
}
