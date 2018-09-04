package admin

import (
	"github.com/caicloud/nirvana"
	"github.com/caicloud/nirvana/config"
	"github.com/caicloud/nirvana/log"

	cfg "github.com/caicloud/simple-object-storage/pkg/config"
	"github.com/caicloud/simple-object-storage/pkg/constants"
	"github.com/caicloud/simple-object-storage/pkg/metadata"
	"github.com/caicloud/simple-object-storage/pkg/metadata/boltdb"
	"github.com/caicloud/simple-object-storage/pkg/storage"
	"github.com/caicloud/simple-object-storage/pkg/storage/local"
)

type Server struct {
	cfg cfg.Config
	cmd config.NirvanaCommand

	bucket metadata.Bucket
	object metadata.Object
	store  storage.Interface
}

func NewServer() (*Server, error) {
	s := &Server{
		cfg: *cfg.NewDefaultConfig(),
		cmd: config.NewNirvanaCommand(&config.Option{
			Port: uint16(constants.DefaultListenPort),
		}),
	}
	s.cmd.AddOption("env", &s.cfg)
	s.cmd.SetHook(&config.NirvanaCommandHookFunc{
		PreConfigureFunc: s.init,
	})
	return s, nil
}

func (s *Server) Run() error {
	return s.cmd.Execute()
}

func (s *Server) init(config *nirvana.Config) error {
	e := s.cfg.Validate()
	if e != nil {
		log.Errorf("config Validate failed, %v", e)
		return e
	}
	log.Infof("init for config:%v", s.cfg.String())

	s.bucket, e = boltdb.NewBucket(s.cfg.DatabaseString, s.cfg.RootAllowed)
	if e != nil {
		log.Errorf("NewBucket failed, %v", e)
		return e
	}
	s.object, e = boltdb.NewObject(s.cfg.DatabaseString, s.cfg.RootAllowed)
	if e != nil {
		log.Errorf("NewObject failed, %v", e)
		return e
	}
	s.store, e = local.NewStorage(s.cfg.StorageString, s.cfg.RootAllowed)
	if e != nil {
		log.Errorf("NewStorage failed, %v", e)
		return e
	}
	config.Configure(
		nirvana.Descriptor(s.getNirvanaDescriptors()...),
	)
	return nil
}
