package golt

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func NewTLSConfig(certFile, keyFile, caFile string, noVerify bool) (*tls.Config, error) {
	c := tls.Config{}

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	c.Certificates = []tls.Certificate{cert}

	if caFile != "" {
		caCert, err := ioutil.ReadFile(caFile)
		if err != nil {
			return nil, err
		}

		pool := x509.NewCertPool()
		pool.AppendCertsFromPEM(caCert)

		c.RootCAs = pool
	}

	if noVerify {
		c.InsecureSkipVerify = true
	}

	return &c, nil
}
