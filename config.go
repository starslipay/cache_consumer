package main

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	CacheMgrRpc zrpc.RpcClientConf
	KqConf      kq.KqConf
}
