package satisphp

import (
	"errors"

	"github.com/Ilyes512/satis-go/satis/satisphp/api"
	"github.com/Ilyes512/satis-go/satis/satisphp/job"
)

var ErrRepoNotFound = errors.New("Repository Not Found")

type SatisClient struct {
	Jobs   chan job.SatisJob
	DbPath string
}

func (s *SatisClient) FindRepo(id string) (api.Repo, error) {
	var repo api.Repo

	repos, err := s.FindAllRepos()
	if err != nil {
		return repo, err
	}

	found := false
	for _, r := range repos {
		if r.ID == id {
			found = true
			repo = r
		}
	}
	if found {
		return repo, nil
	} else {
		return repo, ErrRepoNotFound
	}
}
func (s *SatisClient) FindAllRepos() ([]api.Repo, error) {
	j := job.NewFindAllJob(s.DbPath)

	err := s.performJob(j)

	repos := <-j.ReposResp

	rs := make([]api.Repo, len(repos), len(repos))
	for i, repo := range repos {
		rs[i] = *api.NewRepo(repo.Type, repo.URL)
	}

	return rs, err
}

func (s *SatisClient) SaveRepo(repo *api.Repo, generate bool) error {
	// repoEntity := db.SatisRepository{
	// 	Type: repo.Type,
	// 	URL:  repo.URL,
	// }
	j := job.NewSaveRepoJob(s.DbPath, *repo)
	if err := s.performJob(j); err != nil {
		return err
	}
	if generate {
		return s.GenerateSatisWeb()
	} else {
		return nil
	}
}

func (s *SatisClient) DeleteRepo(id string, generate bool) error {
	var toDelete api.Repo

	repos, err := s.FindAllRepos()
	if err != nil {
		return err
	}

	found := false
	for _, r := range repos {
		if r.ID == id {
			found = true
			toDelete = r
		}
	}

	if found {
		j := job.NewDeleteRepoJob(s.DbPath, toDelete.URL)
		if err = s.performJob(j); err != nil {
			switch err {
			case job.ErrRepoNotFound:
				return ErrRepoNotFound
			default:
				return err
			}
		}

		if generate {
			return s.GenerateSatisWeb()
		} else {
			return nil
		}
	} else {
		return ErrRepoNotFound
	}
}

func (s *SatisClient) GenerateSatisWeb() error {
	j := job.NewGenerateJob()
	return s.performJob(j)
}

func (s *SatisClient) Shutdown() error {
	j := job.NewExitJob()
	return s.performJob(j)
}

func (s *SatisClient) performJob(j job.SatisJob) error {
	s.Jobs <- j

	return <-j.ExitChan()
}
