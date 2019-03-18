package kmip

/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at http://mozilla.org/MPL/2.0/. */

import (
	"context"
	"crypto/tls"
	"log"
	"net"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

type ServerSuite struct {
	suite.Suite

	certs  CertificateSet
	server Server
	client Client

	listenCh chan error
}

func (s *ServerSuite) SetupSuite() {
	s.Require().NoError(s.certs.Generate([]string{"localhost"}, []net.IP{net.IPv4(127, 0, 0, 1)}))

	s.server.Addr = "localhost:"
	s.server.TLSConfig = &tls.Config{}
	DefaultServerTLSConfig(s.server.TLSConfig)
	s.server.TLSConfig.ClientCAs = s.certs.CAPool
	s.server.TLSConfig.Certificates = []tls.Certificate{
		{
			Certificate: [][]byte{s.certs.ServerCert.Raw},
			PrivateKey:  s.certs.ServerKey,
		},
	}

	s.server.ReadTimeout = time.Second
	s.server.WriteTimeout = time.Second

	s.server.Log = log.New(os.Stderr, "[kmip] ", log.LstdFlags)

	s.listenCh = make(chan error, 1)
	initializedCh := make(chan struct{})

	go func() {
		s.listenCh <- s.server.ListenAndServe(initializedCh)
	}()

	<-initializedCh
}

func (s *ServerSuite) SetupTest() {
	s.server.mu.Lock()
	addr := s.server.l.Addr().String()
	s.server.mu.Unlock()

	_, port, err := net.SplitHostPort(addr)
	s.Require().NoError(err)

	s.client.Endpoint = "localhost:" + port
	s.client.TLSConfig = &tls.Config{}
	DefaultClientTLSConfig(s.client.TLSConfig)
	s.client.TLSConfig.RootCAs = s.certs.CAPool
	s.client.TLSConfig.Certificates = []tls.Certificate{
		{
			Certificate: [][]byte{s.certs.ClientCert.Raw},
			PrivateKey:  s.certs.ClientKey,
		},
	}

	s.client.ReadTimeout = time.Second
	s.client.WriteTimeout = time.Second

	s.Require().NoError(s.client.Connect())
}

func (s *ServerSuite) TearDownTest() {
	s.Require().NoError(s.client.Close())
}

func (s *ServerSuite) TearDownSuite() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second)
	defer ctxCancel()

	s.Require().NoError(s.server.Shutdown(ctx))
	s.Require().NoError(<-s.listenCh)
}

func (s *ServerSuite) TestDiscoverVersions() {
	versions, err := s.client.DiscoverVersions(DefaultSupportedVersions)
	s.Require().NoError(err)
	s.Require().Equal(DefaultSupportedVersions, versions)

	versions, err = s.client.DiscoverVersions(nil)
	s.Require().NoError(err)
	s.Require().Equal(DefaultSupportedVersions, versions)

	versions, err = s.client.DiscoverVersions([]ProtocolVersion{{Major: 1, Minor: 2}})
	s.Require().NoError(err)
	s.Require().Equal([]ProtocolVersion{{Major: 1, Minor: 2}}, versions)

	versions, err = s.client.DiscoverVersions([]ProtocolVersion{{Major: 2, Minor: 0}})
	s.Require().NoError(err)
	s.Require().Equal([]ProtocolVersion(nil), versions)
}

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}
