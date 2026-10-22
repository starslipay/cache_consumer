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

// UpdateRelation 回调 cache_mgr 的 UpdateRelation RPC, 由其执行实际的 redis relation 写入
func (c *CacheMgrClient) UpdateRelation(ctx context.Context, userId string, uid int64) (*cache_mgr_pb.UpdateRelationRsp, error) {
	client := cache_mgr_pb.NewCacheMgrClient(c.cli.Conn())
	return client.UpdateRelation(ctx, &cache_mgr_pb.UpdateRelationReq{
		UserId: userId,
		Uid:    uid,
	})
}

// DeleteRelation 回调 cache_mgr 的 DeleteRelation RPC, 由其执行实际的 redis relation 删除(注销用户场景)
func (c *CacheMgrClient) DeleteRelation(ctx context.Context, userId string) (*cache_mgr_pb.DeleteRelationRsp, error) {
	client := cache_mgr_pb.NewCacheMgrClient(c.cli.Conn())
	return client.DeleteRelation(ctx, &cache_mgr_pb.DeleteRelationReq{
		UserId: userId,
	})
}
