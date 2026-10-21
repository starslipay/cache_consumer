package main

import (
	"context"

	"github.com/starslipay/cache_mgr/cache_mgr_pb"
	"github.com/zeromicro/go-zero/zrpc"
)

type CacheMgrClient struct {
	cli zrpc.Client
}

func NewCacheMgrClient(conf zrpc.RpcClientConf) *CacheMgrClient {
	conn := zrpc.MustNewClient(conf)
	return &CacheMgrClient{
		cli: conn,
	}
}

// UpdateBalance 回调 cache_mgr 的 UpdateBalance RPC, 由其执行实际的 Redis 余额 CAS 更新
func (c *CacheMgrClient) UpdateBalance(ctx context.Context, uid, curType, balance, modifyTimeS, modifyTimeMs int64) (*cache_mgr_pb.UpdateBalanceRsp, error) {
	client := cache_mgr_pb.NewCacheMgrClient(c.cli.Conn())
	return client.UpdateBalance(ctx, &cache_mgr_pb.UpdateBalanceReq{
		Uid:                 uid,
		CurType:             curType,
		Balance:             balance,
		BalanceModifyTimeS:  modifyTimeS,
		BalanceModifyTimeMs: modifyTimeMs,
	})
}
