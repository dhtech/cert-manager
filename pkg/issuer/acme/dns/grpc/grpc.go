package grpc

import (
	"crypto/tls"
	"time"

	"github.com/golang/glog"
	"github.com/jetstack/cert-manager/pkg/issuer/acme/dns/util"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	//pb "github.com/jetstack/cert-manager/proto"
)

type provider struct {
	creds     *grpc.DialOption
	timeout   time.Duration
	interval  time.Duration
}

func NewDNSProviderTLS(service, serverName string, clientCert *tls.Certificate, timeout, interval time.Duration) (*provider, error) {
	tc := &tls.Config{
		ServerName: serverName,
	}

	if clientCert != nil {
		tc.Certificates = []tls.Certificate{*clientCert}
	}
	creds := grpc.WithTransportCredentials(credentials.NewTLS(tc))
	return &provider{&creds, timeout, interval}, nil
}

func NewDNSProviderInsecure(service string, timeout, interval time.Duration) (*provider, error) {
	creds := grpc.WithInsecure()
	return &provider{&creds, timeout, interval}, nil
}

func (c *provider) Present(domain, token, key string) error {
	fqdn, value, ttl := util.DNS01Record(domain, key)
	glog.Infof("TODO: Insert %v, %v, %v", fqdn, value, ttl)
	_, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	return errors.New("not implemented")
}

func (c *provider) CleanUp(domain, token, key string) error {
	fqdn, _, _ := util.DNS01Record(domain, key)
	glog.Infof("TODO: Remove %s", fqdn)
	return errors.New("not implemented")
}

func (c *provider) Timeout() (timeout, interval time.Duration) {
	return c.timeout, c.interval
}
