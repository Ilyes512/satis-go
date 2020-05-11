package job

import (
	"github.com/Ilyes512/satis-go/satis/satisphp/api"
	"github.com/Ilyes512/satis-go/satis/satisphp/db"
)

// Add or save a repo tp the repo collection
func NewSaveRepoJob(dbPath string, repo api.Repo) *SaveRepoJob {
	return &SaveRepoJob{
		dbPath:     dbPath,
		repository: repo,
		exitChan:   make(chan error, 1),
	}
}

type SaveRepoJob struct {
	dbPath     string
	repository api.Repo
	exitChan   chan error
}

func (j SaveRepoJob) ExitChan() chan error {
	return j.exitChan
}
func (j SaveRepoJob) Run() error {
	dbMgr := db.SatisDbManager{Path: j.dbPath}

	if err := dbMgr.Load(); err != nil {
		return err
	}
	repos, err := j.doSave(j.repository, dbMgr.Db.Repositories)
	if err != nil {
		return err
	}
	dbMgr.Db.Repositories = repos

	if err := dbMgr.Write(); err != nil {
		return err
	}
	return nil
}
func (j SaveRepoJob) doSave(repo api.Repo, repos []db.SatisRepository) ([]db.SatisRepository, error) {
	repoEntity := db.SatisRepository{Type: repo.Type, URL: repo.URL}
	found := false
	for i, r := range repos {
		tmp := api.NewRepo(r.Type, r.URL)
		if tmp.ID == repo.ID {
			repos[i] = repoEntity
			found = true
		}
	}
	if !found {
		return append(repos, repoEntity), nil
	}

	return repos, nil
}
