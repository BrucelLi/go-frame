package tasks

import (
	"context"
	"demo/internal/jobs/client"
	"encoding/json"
	"fmt"
	"github.com/hibiken/asynq"
	"time"
)

type EmailPayload struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func HandleSendEmail(_ context.Context, t *asynq.Task) error {
	var p EmailPayload
	if err := json.Unmarshal(t.Payload(), &p); err != nil {
		return err
	}
	// 这里可以接入实际邮件发送逻辑
	fmt.Printf("📧 Sending email to %s: %s\n", p.To, p.Subject)
	return nil
}

func EnqueueSendEmail(to, subject, body string) error {
	jobClient := client.GetJobsClient()
	payload, _ := json.Marshal(EmailPayload{
		To:      to,
		Subject: subject,
		Body:    body,
	})
	task := asynq.NewTask(TaskSendEmail, payload)
	_, err := jobClient.Enqueue(task, asynq.MaxRetry(3), asynq.Timeout(30*time.Second))
	return err
}
