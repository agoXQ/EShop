package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	// "os"
	// "strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

var ch *amqp.Channel

func init() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	// defer conn.Close()

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
	// defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	failOnError(err, "Failed to declare an exchange")

	//声明固定响应队列

    _, err = ch.QueueDeclare(
        "response_queue", // 队列名称
        true,            // 持久化
        false,           // 不自动删除
        false,           // 不独占
        false,           // 不等待
        nil,            // 参数
    )
    failOnError(err, "Failed to declare response queue")
}


func VerifyToken(token string) (bool, error) {
    // 生成唯一的 request_id
    requestID := uuid.New().String()

    // 构造请求消息
    request := map[string]interface{}{
        "request_id": requestID,
        "token":      token,
        "reply_to":   "response_queue", // 使用固定的响应队列名称
    }
    body, err := json.Marshal(request)
    if err != nil {
        return false, fmt.Errorf("failed to marshal request: %v", err)
    }

    // 发送请求消息
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    err = ch.PublishWithContext(ctx,
        "logs_topic",          // exchange
        "server.token.verify", // routing key
        false,                 // mandatory
        false,                 // immediate
        amqp.Publishing{
            ContentType: "application/json",
            Body:        body,
        },
    )
    if err != nil {
        return false, fmt.Errorf("failed to publish request: %v", err)
    }

    // 创建一个独立的消费者
    msgs, err := ch.Consume(
        "response_queue", // 固定的响应队列名称
        "",              // 消费者标签
        false,           // 手动确认
        false,           // 不独占
        false,           // 不阻塞
        false,           // 不等待
        nil,             // 参数
    )
    if err != nil {
        return false, fmt.Errorf("failed to register consumer: %v", err)
    }

    // 设置超时机制
    for {
        select {
        case msg := <-msgs:
            var response map[string]interface{}
            if err := json.Unmarshal(msg.Body, &response); err != nil {
                msg.Nack(false, true) // 拒绝消息并重新入队
                return false, fmt.Errorf("failed to unmarshal response: %v", err)
            }
            if response["request_id"] == requestID {
                msg.Ack(false) // 确认消息
                if response["status"].(string) == "success" {
                    return response["valid"].(bool), nil
                } else {
                    return false, fmt.Errorf("token verification failed: %v", response["error"])
                }
            } else {
                msg.Nack(false, true) // 拒绝消息并重新入队
            }
        case <-time.After(5 * time.Second):
            return false, fmt.Errorf("timeout waiting for response")
        }
    }
}

func verifyHandler(ctx context.Context, rctx *app.RequestContext) {
    token := rctx.Query("token")
    result, err := VerifyToken(token)
    if err != nil {
        rctx.JSON(http.StatusBadRequest, utils.H{"error": err.Error()})
        return
    }
    rctx.JSON(http.StatusOK, utils.H{"valid": result})
}