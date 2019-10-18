package modules

import (
	"github.com/coreos/etcd/clientv3"
	"time"
)

const (
	dialTimeout = 3 * time.Second
	readTimeout = 3 * time.Second
	writeTimeout = 3 * time.Second
	bakDataTTL = 1800
)

const (
	hgwPrefix = "/hgw-gateway/"
)




var cli *clientv3.Client

func ConnectStore(endPoints []string, username, password string) error {
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   endPoints,
		DialTimeout: dialTimeout,
		Username: username,
		Password: password,
	})
	if err != nil {
		return err
	}
	cli = c
	return nil
}

