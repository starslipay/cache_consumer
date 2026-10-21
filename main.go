package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/config.yaml", "the config file")

func main() {
	flag.Parse()

	var c Config
	conf.MustLoad(*configFile, &c)

	cacheMgrClient := NewCacheMgrClient(c.CacheMgrRpc)

	updateBalanceConsumer := NewUpdateBalanceConsumer(cacheMgrClient)
	updateRelationConsumer := NewUpdateRelationConsumer(cacheMgrClient)

	p := kq.MustNewQueue(c.KqConf, kq.WithHandle(func(ctx context.Context, key, value string) error {
		return updateBalanceConsumer.Consume(ctx, []byte(value))
	}))

	pRelation := kq.MustNewQueue(c.KqConfRelation, kq.WithHandle(func(ctx context.Context, key, value string) error {
		return updateRelationConsumer.Consume(ctx, []byte(value))
	}))

	fmt.Printf("Starting update balance cache consumer...\n")
	go p.Start()

	fmt.Printf("Starting update relation cache consumer...\n")
	pRelation.Start()
}
