package tasks

import (
	"context"
	"demo/internal/jobs/client"
	"encoding/json"
	"log"

	"github.com/hibiken/asynq"
)

type LogPayload struct {
	UserID string `json:"userId"`
	Event  string `json:"event"`
}

func HandleLogEvent(_ context.Context, t *asynq.Task) error {
	var p LogPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	log.Printf("📝 User %s triggered event: %s", p.UserID, p.Event)
	return nil
}

func EnqueueLogEvent(userID, event string) error {
	jobClient := client.GetJobsClient()
	payload, _ := json.Marshal(LogPayload{
		UserID: userID,
		Event:  event,
	})
	task := asynq.NewTask(TaskLogEvent, payload)
	_, err := jobClient.Enqueue(task)
	return err
}
