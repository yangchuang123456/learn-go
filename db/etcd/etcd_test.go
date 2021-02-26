package etcd

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/etcd/clientv3"
	etcdnaming "go.etcd.io/etcd/clientv3/naming"
	"google.golang.org/grpc/naming"
	"testing"
)
func Test_etcd(t *testing.T){
	cli, cerr := clientv3.NewFromURL("http://localhost:2379")
	assert.NoError(t,cerr)
	r := &etcdnaming.GRPCResolver{Client: cli}
/*	b := grpc.RoundRobin(r)
	_, gerr := grpc.Dial("my-service", grpc.WithBalancer(b), grpc.WithBlock())
	assert.NoError(t,gerr)*/

	r.Update(context.TODO(), "my-service", naming.Update{Op: naming.Add, Addr: "1.2.3.4", Metadata: "..."})
}
