package notification

import (
	"console/src/config"
	"log"

	"github.com/streadway/amqp"
)

func SendStockMessage(message []byte) error {
	ConnectRabbit(message, config.RabbitStockRoutingKey)
	return nil
}

func SendFiatMessage(message []byte) error {
	ConnectRabbit(message, config.RabbitFiatRountingKey)
	return nil
}

func ConnectRabbit(message []byte, routingKey string) error {

	// Conecta com o RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		return err
	}

	// Cria um canal
	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	defer conn.Close()
	defer ch.Close()

	// Configura a exchange
	err = ch.ExchangeDeclare(
		config.RabbitExchange, // exchange name
		"direct",              // exchange type
		true,                  // durable
		false,                 // auto-deleted
		false,                 // internal
		false,                 // no-wait
		nil,                   // arguments
	)
	if err != nil {
		return err
	}

	// Configura a fila de STOCK com DLQ e TTL
	q, err := ch.QueueDeclare(
		config.RabbitStockQueue, // name
		true,                    // durable
		false,                   // delete when unused
		false,                   // exclusive
		false,                   // no-wait
		amqp.Table{
			"x-dead-letter-exchange":    config.RabbitExchangeDLQ,
			"x-dead-letter-routing-key": config.RabbitStockRoutingKey, //config.RabbitQueueDLQ,
			"x-message-ttl":             config.RabbitTTL,             // TTL de 5 segundos
		}, // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Liga a fila de STOCK à exchange
	err = ch.QueueBind(
		q.Name,                       // queue name
		config.RabbitStockRoutingKey, // routing key
		config.RabbitExchange,        // exchange
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Configura a fila de FIAT com DLQ e TTL
	q2, err := ch.QueueDeclare(
		config.RabbitFiatQueue, // name
		true,                   // durable
		false,                  // delete when unused
		false,                  // exclusive
		false,                  // no-wait
		amqp.Table{
			"x-dead-letter-exchange":    config.RabbitExchangeDLQ,
			"x-dead-letter-routing-key": config.RabbitFiatRountingKey, //config.RabbitQueueDLQ,
			"x-message-ttl":             config.RabbitTTL,             // TTL de 5 segundos
		}, // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Liga a fila de FIAT à exchange
	err = ch.QueueBind(
		q2.Name,                      // queue name
		config.RabbitFiatRountingKey, // routing key
		config.RabbitExchange,        // exchange
		false,                        // no-wait
		nil,                          // arguments
	)
	if err != nil {
		log.Fatal(err)
	}

	// Configura a exchange de DLQ
	err = ch.ExchangeDeclare(
		config.RabbitExchangeDLQ, // exchange name
		"fanout",                 // exchange type
		true,                     // durable
		false,                    // auto-deleted
		false,                    // internal
		false,                    // no-wait
		nil,                      // arguments
	)
	if err != nil {
		return err
	}

	// Configura a fila de DLQ
	_, err = ch.QueueDeclare(
		config.RabbitQueueDLQ, // Nome da fila DLQ
		true,                  // Durable
		false,                 // Auto-delete
		false,                 // Exclusive
		false,                 // No-wait
		nil,                   // Argumentos extras
	)
	if err != nil {
		return err
	}

	// Liga a fila de DLQ à exchange DLQ
	err = ch.QueueBind(
		config.RabbitQueueDLQ,    // Nome da fila DLQ
		config.RabbitQueueDLQ,    // Routing key para a fila DLQ
		config.RabbitExchangeDLQ, // Nome da exchange DLQ
		false,                    // No-wait
		nil,                      // Argumentos extras
	)
	if err != nil {
		return err
	}

	msg := amqp.Publishing{
		ContentType:  "text/plain",
		Body:         message,         //[]byte(message),
		DeliveryMode: amqp.Persistent, // make message persistent
	}

	err = ch.Publish(
		config.RabbitExchange, // exchange
		routingKey,            // routing key
		false,                 // mandatory
		false,                 // immediate
		msg,
	)
	if err != nil {
		return err
	}

	return nil
}
