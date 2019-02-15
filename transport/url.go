package transport

import (
	"errors"
	"net/url"
)

type URL struct {
	Address  string
	Username string
	Password string
}

func NewURL(uri string) (*URL, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	url := URL{
		Address: u.Host,
	}

	// Validate schema.
	if u.Scheme != "bolt" {
		return nil, errors.New(`URI scheme must be "bolt"`)
	}

	// Validate user + password.
	if u.User != nil {
		p, set := u.User.Password()
		if !set {
			return nil, errors.New(`A password must be provided when providing a username`)
		}

		url.Username = u.User.Username()
		url.Password = p
	}

	return &url, nil
}
