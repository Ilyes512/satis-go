package db

import (
	"encoding/json"
	"io/ioutil"
)

const (
	// DbFile const
	DbFile = "/db.json"
	// StagingFile const
	StagingFile = "/stage.json"
)

// SatisDbManager struct
type SatisDbManager struct {
	Path string
	Db   SatisDb
}

// Load method
func (c *SatisDbManager) Load() error {
	content, err := ioutil.ReadFile(c.Path + DbFile)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(content, &c.Db); err != nil {
		return err
	}
	return nil
}

// Write method
func (c *SatisDbManager) Write() error {
	return c.doWrite(c.Path + DbFile)
}

// WriteStaging method
func (c *SatisDbManager) WriteStaging() error {
	return c.doWrite(c.Path + StagingFile)
}

func (c *SatisDbManager) doWrite(path string) error {
	b, err := json.MarshalIndent(c.Db, "", "    ") // pretty print
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(path, b, 0644); err != nil {
		return err
	}
	return nil
}

// SaveRepo method
func (c *SatisDbManager) SaveRepo(repo SatisRepository) error {
	return nil
}
