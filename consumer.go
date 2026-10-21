package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

// UpdateBalanceMessage 余额缓存更新消息, 与 account_mgr 中 sendUpdateBalanceMessage 投递的结构一致
type UpdateBalanceMessage struct {
	Uid                 int64 `json:"uid"`
	CurType             int64 `json:"cur_type"`
	Balance             int64 `json:"balance"`
	BalanceModifyTimeS  int64 `json:"balance_modify_time_s"`  // 余额变更时间(秒)
	BalanceModifyTimeMs int64 `json:"balance_modify_time_ms"` // 余额变更时间的毫秒部分(0-999)
}

type UpdateBalanceConsumer struct {
	cacheMgrClient *CacheMgrClient
	logger         logx.Logger
}

func NewUpdateBalanceConsumer(cacheMgrClient *CacheMgrClient) *UpdateBalanceConsumer {
	return &UpdateBalanceConsumer{
		cacheMgrClient: cacheMgrClient,
		logger:         logx.WithContext(context.Background()),
	}
}

// Consume 消费 t_update_balance 消息, 回调 cache_mgr.UpdateBalance 执行 Redis 余额 CAS 更新
func (c *UpdateBalanceConsumer) Consume(ctx context.Context, message []byte) error {
	var msg UpdateBalanceMessage
	if err := json.Unmarshal(message, &msg); err != nil {
		c.logger.Errorf("unmarshal update balance message failed: %v", err)
		return err
	}

	c.logger.Infof("received update balance message: uid=%d, cur_type=%d, balance=%d, modify_time=%d/%d",
		msg.Uid, msg.CurType, msg.Balance, msg.BalanceModifyTimeS, msg.BalanceModifyTimeMs)

	maxRetries := 3
	var lastErr error
	for i := 0; i < maxRetries; i++ {
		_, err := c.cacheMgrClient.UpdateBalance(ctx,
			msg.Uid, msg.CurType, msg.Balance, msg.BalanceModifyTimeS, msg.BalanceModifyTimeMs)
		if err != nil {
			lastErr = err
			c.logger.Errorf("call UpdateBalance failed (retry %d/%d): %v", i+1, maxRetries, err)
			time.Sleep(time.Duration(i+1) * time.Second)
			continue
		}

		c.logger.Infof("update balance cache success: uid=%d, balance=%d", msg.Uid, msg.Balance)
		return nil
	}

	c.logger.Errorf("call UpdateBalance failed after %d retries: %v", maxRetries, lastErr)
	return fmt.Errorf("failed after %d retries: %w", maxRetries, lastErr)
}
