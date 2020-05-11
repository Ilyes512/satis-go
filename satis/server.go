package satis

import (
	"net/http"
	"os"

	"github.com/Ilyes512/satis-go/satis/satisphp"
	"github.com/Ilyes512/satis-go/satis/satisphp/db"
	"github.com/Ilyes512/satis-go/satis/satisphp/job"
	"github.com/gorilla/mux"
)

// Server struct used by satis-go
type Server struct {
	DbPath       string
	AdminUIPath  string
	WebPath      string
	Bind         string
	Name         string
	Homepage     string
	Username     string
	APIToken     string
	jobProcessor satisphp.SatisJobProcessor
	jobClient    satisphp.SatisClient
}

// Run the satis-go server
func (s *Server) Run(generate bool) error {
	// sync config to db
	if err := s.initDb(); err != nil {
		return err
	}

	// Shared Jobs Channel to queue/process db modifications and generation task
	jobs := make(chan job.SatisJob)

	// Job Processor responsible for interacting with db & static web docs
	gen := &satisphp.StaticWebGenerator{
		DbPath:  s.DbPath,
		WebPath: s.WebPath,
	}

	s.jobProcessor = satisphp.SatisJobProcessor{
		DbPath:    s.DbPath,
		Jobs:      jobs,
		Generator: gen,
	}

	// Client to Job Processor
	jobClient := satisphp.SatisClient{
		DbPath: s.DbPath,
		Jobs:   jobs,
	}

	// route handlers
	resource := &SatisResource{
		Host:           s.Homepage,
		SatisPhpClient: jobClient,
		Username:       s.Username,
		APIToken:       s.APIToken,
	}

	// Configure Routes
	r := mux.NewRouter()

	r.HandleFunc("/api/repo", resource.addRepo).Methods("POST")
	r.HandleFunc("/api/repo/{id}", resource.saveRepo).Methods("PUT")
	r.HandleFunc("/api/repo/{id}", resource.findRepo).Methods("GET")
	r.HandleFunc("/api/repo", resource.findAllRepos).Methods("GET")
	r.HandleFunc("/api/repo/{id}", resource.deleteRepo).Methods("DELETE")
	r.HandleFunc("/api/generate-web-job", resource.generateStaticWeb).Methods("POST")
	r.HandleFunc("/api/update-package", resource.updatePackage).Queries("username", "", "apiToken", "").Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(s.WebPath)))

	// r.Handle("/dist/{rest}", http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist/"))))
	// r.PathPrefix("/dist/").Handler(http.StripPrefix("/dist/", http.FileServer(http.Dir("./dist"))))

	http.Handle("/", r)
	http.Handle("/admin/", http.StripPrefix("/admin/", http.FileServer(http.Dir(s.AdminUIPath))))

	// Start update processor
	go s.jobProcessor.ProcessUpdates()

	if generate {
		err := resource.generateStaticWebNow()
		if err != nil {
			return err
		}
	}

	// Start HTTP Server
	return http.ListenAndServe(s.Bind, nil)
}

// Sync configured values to satis repository meta data
func (s *Server) initDb() error {
	dbMgr := &db.SatisDbManager{Path: s.DbPath}

	// create empty db if it doesn't exist
	if _, err := os.Stat(s.DbPath + db.DbFile); os.IsNotExist(err) {
		if err := dbMgr.Write(); err != nil {
			return err
		}
	}

	if err := dbMgr.Load(); err != nil {
		return err
	}

	boolPointer := func(b bool) *bool { return &b }

	dbMgr.Db.Name = s.Name
	dbMgr.Db.Homepage = s.Homepage
	dbMgr.Db.RequireAll = boolPointer(true)
	return dbMgr.Write()
}
