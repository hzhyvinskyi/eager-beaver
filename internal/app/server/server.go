package server

// Server ...
type Server struct {
	config *Config
}

// New ...
func New(config *Config) *Server {
	return &Server{
		config: config,
	}
}

// Start ...
func (s *Server) Start() error {
	return nil
}
