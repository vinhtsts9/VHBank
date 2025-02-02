package worker

import (
	"Golang-Masterclass/simplebank/internal/database"
	"context"

	"github.com/hibiken/asynq"
)

const (
	QueueCritical = "critical"
	QueueDefault  = "default"
)

type Taskprocessor interface {
	Start() error
	ProcessTaskSendVerifyEmail(ctx context.Context, task *asynq.Task) error
}

type RedisTaskProcessor struct {
	server *asynq.Server
	db     *database.Queries
}

func NewRedisTaskProcessor(redisOpt asynq.RedisClientOpt, db *database.Queries) Taskprocessor {
	server := asynq.NewServer(redisOpt, asynq.Config{
		Queues: map[string]int{
			QueueCritical: 6,
			QueueDefault:  3,
		},
	})
	return &RedisTaskProcessor{server: server, db: db}
}

func (r *RedisTaskProcessor) Start() error {
	mux := asynq.NewServeMux()
	mux.HandleFunc(TaskSendVerifyEmail, r.ProcessTaskSendVerifyEmail)
	return r.server.Start(mux)
}
