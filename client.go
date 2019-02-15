package golt

type Client struct {
	cfg *Config

	conn Conn
}

func New(cfg *Config) (*Client, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	return &Client{cfg: cfg}, nil
}

func (c *Client) Connect() error {
	return c.conn.Connect(c.cfg.URI, c.cfg.TLSConfig, c.cfg.ConnectionTimeout)
}

// TODO(tserkov): Connection pooling.
// Session duplicates the connection to the server.
// func (c *Client) Session() (*Conn, error) {}
