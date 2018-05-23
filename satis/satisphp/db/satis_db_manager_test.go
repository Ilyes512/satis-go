package db

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
)

const dbPath = "/tmp/satis-test-data"

func ARandomDbMgr() SatisDbManager {
	// Make Data Dir
	if err := os.MkdirAll(dbPath, 0744); err != nil {
		log.Fatalf("Unable to create path: %v", err)
	}

	mgr := SatisDbManager{Path: dbPath}
	mgr.Db.Name = "My Repo"
	mgr.Db.Homepage = "http://repo.com"
	mgr.Db.RequireAll = true
	mgr.Db.Repositories = []SatisRepository{
		SatisRepository{Type: "vcs", URL: "http://package.com"},
	}

	mgr.Path = dbPath
	mgr.Write()
	mgr.WriteStaging()

	return mgr
}

func TestDbLoad(t *testing.T) {

	// given
	mgr := ARandomDbMgr()
	r := SatisDbManager{Path: mgr.Path}

	// when
	err := r.Load()

	// then
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(r.Db, mgr.Db) {
		t.Error("loaded config doesn't match original")
	}
}
func TestDbWrite(t *testing.T) {
	// given
	r := ARandomDbMgr()
	oldName := r.Db.Name
	// when
	r.Db.Name = "foo"
	modifiedDb := r.Db

	err := r.Write()

	// then
	if err != nil {
		t.Error(err)
	}

	err = r.Load()
	if err != nil {
		t.Error(err)
	}

	if oldName == r.Db.Name {
		t.Errorf("config should have changed: %s / %s", oldName, r.Db.Name)
	}
	if !reflect.DeepEqual(r.Db, modifiedDb) {
		t.Errorf("config didn't persist changes when written: %s / %s", r.Db.Name, modifiedDb.Name)
	}
}

func TestDbWriteStaging(t *testing.T) {
	// given
	r := ARandomDbMgr()
	oldName := r.Db.Name
	// when
	r.Db.Name = "foo"
	modifiedDb := r.Db

	err := r.WriteStaging()

	// then
	if err != nil {
		t.Error(err)
	}

	content, err := ioutil.ReadFile(dbPath + StagingFile)

	if err != nil {
		t.Error(err)
	}

	var c SatisDbManager
	if err = json.Unmarshal(content, &c.Db); err != nil {
		t.Error(err)
	}

	if oldName == r.Db.Name {
		t.Errorf("staging config should have changed: %s / %s", oldName, r.Db.Name)
	}
	if !reflect.DeepEqual(r.Db, modifiedDb) {
		t.Errorf("staging config didn't persist changes when written: %s / %s", r.Db.Name, modifiedDb.Name)
	}
}
