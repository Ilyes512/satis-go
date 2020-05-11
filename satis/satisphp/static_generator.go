package satisphp

import (
	"log"
	"os/exec"

	"github.com/Ilyes512/satis-go/satis/satisphp/db"
)

type Generator interface {
	Generate() error
}

type StaticWebGenerator struct {
	DbPath  string
	WebPath string
}

func (s *StaticWebGenerator) Generate() error {
	log.Print("Generating...")
	out, err := exec.
		Command("satis", "--no-interaction", "build", s.DbPath+db.StagingFile, s.WebPath).
		CombinedOutput()
	if err != nil {
		log.Printf("Satis Generation Error: %s", string(out[:]))
	}
	return err
}
