package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/benschw/satis-go/satis"
	"gopkg.in/yaml.v2"
)

// Config struct used by satis-go
type Config struct {
	Dbpath      string
	Bind        string
	RepoUIPath  string
	AdminUIPath string
	Reponame    string
	Repohost    string
}

func getConfig(path string) (Config, error) {
	config := Config{}

	if _, err := os.Stat(path); err != nil {
		return config, errors.New("config path not valid")
	}

	ymlData, err := ioutil.ReadFile(path)
	if err != nil {
		return config, err
	}

	err = yaml.Unmarshal([]byte(ymlData), &config)
	return config, err
}

func main() {
	cfgPath := flag.String("config", "/opt/satis-go/config.yaml", "Path to Config File")
	generateSatisWeb := flag.Bool("generate", false, "Generate Satis Web")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [arguments] \n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()

	// Load Config
	cfg, err := getConfig(*cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	// Make Data Dir
	if err := os.MkdirAll(cfg.Dbpath, 0744); err != nil {
		log.Fatalf("Unable to create path: %v", err)
	}

	// Configure Server
	s := &satis.Server{
		DbPath:      cfg.Dbpath,
		AdminUIPath: cfg.AdminUIPath,
		WebPath:     cfg.RepoUIPath,
		Bind:        cfg.Bind,
		Name:        cfg.Reponame,
		Homepage:    cfg.Repohost,
	}

	log.Println("Satis is starting on " + cfg.Bind)

	// Start Server
	if err := s.Run(*generateSatisWeb); err != nil {
		log.Fatal(err)
	}
}
