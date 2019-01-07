package node

type Config struct {
	Server   string
	Contract string
}
type Node struct {
	Config *Config
}

var DefaultConfig = Config{
	Server:   "http://localhost",
	Contract: "",
}

func New(conf *Config) (*Node, error) {
	return &Node{
		Config: conf,
	}, nil
}
