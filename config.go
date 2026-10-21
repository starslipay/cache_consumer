package main

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	CacheMgrRpc    zrpc.RpcClientConf
	KqConf         kq.KqConf // t_update_balance 余额缓存更新
	KqConfRelation kq.KqConf // t_update_relation relation缓存更新
}
