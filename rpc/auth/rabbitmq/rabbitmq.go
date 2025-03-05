package rabbitmq

import (
	"encoding/json"
	"log"
	"time"

	"eshop/rpc/auth/service"

	amqp "github.com/rabbitmq/amqp091-go"
)

func StartTokenVerificationConsumer() {
	for {
		err := consumeMessages()
		if err != nil {
			log.Printf("Error in consumer: %v", err)
			time.Sleep(5 * time.Second) // 等待一段时间后重试
		}
	}
}

func consumeMessages() error {
	println("received message")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		"server.token.verify",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	// 绑定队列到交换机
	err = ch.QueueBind(
		"server.token.verify", // 队列名称
		"server.token.verify", // 路由键
		"logs_topic",          // 交换机名称
		false,                 // 不等待
		nil,                   // 参数
	)
	if err != nil {
		log.Fatalf("Failed to bind queue: %v", err)
	}

	msgs, err := ch.Consume(
		"server.token.verify",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	s := service.NewServiceStruct("localhost:6379")

	for msg := range msgs {
		go handleMsg(msg, ch, s)
	}

	return nil
}

func handleMsg(msg amqp.Delivery, ch *amqp.Channel, s *service.ServiceStruct) {
	var request map[string]interface{}
	if err := json.Unmarshal(msg.Body, &request); err != nil {
		log.Printf("Failed to unmarshal request: %v", err)
		return
	}

	valid, err := s.Verify(request["token"].(string))
	if err != nil {
		log.Printf("Token verification failed: %v", err)
	}

	response := map[string]interface{}{
		"request_id": request["request_id"],
		"status":     "success",
		"valid":      valid,
	}
	if err != nil {
		response["status"] = "error"
		response["error"] = err.Error()
	}

	body, _ := json.Marshal(response)
	err = ch.Publish(
		"",
		request["reply_to"].(string),
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("Failed to publish response: %v", err)
	}
}

// package rabbitmq

// import (
// 	"encoding/json"
// 	"log"
// 	amqp "github.com/rabbitmq/amqp091-go"
// 	"eshop/rpc/auth/service"
// )

// func StartTokenVerificationConsumer() {
//     // 连接到 RabbitMQ
//     conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//     if err != nil {
//         log.Fatalf("Failed to connect to RabbitMQ: %v", err)
//     }
//     defer conn.Close()
// 	log.Printf("start in rabbitmq")

//     ch, err := conn.Channel()
//     if err != nil {
//         log.Fatalf("Failed to open a channel: %v", err)
//     }
//     defer ch.Close()

//     // 声明请求队列
//     _, err = ch.QueueDeclare(
//         "server.token.verify", // 队列名称
//         false,                // 非持久化
//         false,                // 不自动删除
//         false,                // 不独占
//         false,                // 不等待
//         nil,                  // 参数
//     )
//     if err != nil {
//         log.Fatalf("Failed to declare queue: %v", err)
//     }

//     // 消费请求消息
//     msgs, err := ch.Consume(
//         "server.token.verify", // 队列名称
//         "",                   // 消费者标签
//         true,                 // 自动确认
//         false,                // 不独占
//         false,                // 不阻塞
//         false,                // 不等待
//         nil,                  // 参数
//     )
//     if err != nil {
//         log.Fatalf("Failed to register consumer: %v", err)
//     }

// 	s := service.NewServiceStruct("localhost:6379")

//     for msg := range msgs {
//         var request map[string]interface{}
//         if err := json.Unmarshal(msg.Body, &request); err != nil {
//             log.Printf("Failed to unmarshal request: %v", err)
//             continue
//         }

//         // 校验 token
//         valid, err := s.Verify(request["token"].(string))
//         if err != nil {
//             log.Printf("Token verification failed: %v", err)
//         }

//         // 构造响应消息
//         response := map[string]interface{}{
//             "request_id": request["request_id"],
//             "status":     "success",
//             "valid":      valid,
//         }
//         if err != nil {
//             response["status"] = "error"
//             response["error"] = err.Error()
//         }
//         body, _ := json.Marshal(response)
// 		// var forever chan struct{}
//         // 发送响应消息
//         err = ch.Publish(
//             "",          // exchange
//             request["reply_to"].(string), // routing key
//             false,      // mandatory
//             false,      // immediate
//             amqp.Publishing{
//                 ContentType: "application/json",
//                 Body:        body,
//             },
//         )
//         if err != nil {
//             log.Printf("Failed to publish response: %v", err)
//         }
//     }
// }
