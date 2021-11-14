package rabbitmq

import (
	"fmt"
	"spectra/commom"
	appErrors "spectra/errors"

	"github.com/streadway/amqp"
)

var RabbitMQChannel *amqp.Channel

func CreateConnection() (*amqp.Connection, *amqp.Channel, appErrors.ErrorResponse) {
	connectRabbitMQ, err := amqp.Dial(commom.Envs.RabbitMQHost)
	if err != nil {
		return nil, nil, appErrors.InternalServerError(fmt.Sprintf("Error to connect in rabbitmq host - %s", err.Error()))
	}

	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		return nil, nil, appErrors.InternalServerError("Error to create channel with rabbit mq")
	}

	_, err = channelRabbitMQ.QueueDeclare(
		commom.Envs.RabbitMQQueue, // queue name
		true,                      // durable
		false,                     // auto delete
		false,                     // exclusive
		false,                     // no wait
		nil,                       // arguments
	)

	if err != nil {
		return nil, nil, appErrors.InternalServerError("Error to create queue with rabbit mq")
	}

	return connectRabbitMQ, channelRabbitMQ, appErrors.ErrorResponse{}
}

func SendMessage(channel *amqp.Channel, input amqp.Publishing, queueName string) appErrors.ErrorResponse {
	if err := channel.Publish(
		"",        // exchange
		queueName, // queue name
		false,     // mandatory
		false,     // immediate
		input,     // message to publish
	); err != nil {
		return appErrors.InternalServerError(fmt.Sprintf("Error to send message to queue - %s", queueName))
	}
	return appErrors.ErrorResponse{}
}
