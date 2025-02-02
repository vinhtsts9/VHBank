package worker

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
	"github.com/pingcap/log"
	"go.uber.org/zap"
)

const TaskSendVerifyEmail = "task:send_verify_email"

type PayloadSendVerifyEmail struct {
	Username string `json:"username"`
}

func (r *RedisTaskDistributor) DistributeTaskSendVerifyEmail(
	ctx context.Context,
	payload *PayloadSendVerifyEmail,
	opts ...asynq.Option,
) error {
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("Failed to marshal payload: %w", err)
	}
	task := asynq.NewTask(TaskSendVerifyEmail, jsonPayload, opts...)
	info, err := r.client.EnqueueContext(ctx, task)
	if err != nil {
		return fmt.Errorf("Failed to enqueue task: %w", err)
	}

	log.Info("task enqueued",
		zap.String("type", task.Type()),
		zap.ByteString("payload", task.Payload()),
		zap.String("queue", info.Queue),
		zap.Int("max_retry", info.MaxRetry),
	)
	return nil
}

func (r *RedisTaskProcessor) ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error {
	var payload PayloadSendVerifyEmail
	if err := json.Unmarshal(task.Payload(), &payload); err != nil {
		return fmt.Errorf("Failed to unmarshal payload: %w", err)
	}
	user, err := r.db.GetUser(ctx, payload.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info("processing task",
				zap.String("type", task.Type()),
				zap.ByteString("payload", task.Payload()),
				zap.String("email", user.Email),
			)
			return fmt.Errorf("Failed to get user: %w", err)
		}
		log.Info("processing task",
			zap.String("type", task.Type()),
			zap.ByteString("payload", task.Payload()),
			zap.String("email", user.Email),
		)
	}
	return nil
}
