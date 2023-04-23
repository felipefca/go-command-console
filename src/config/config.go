package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	ConnectionString      = ""
	Port                  = 0
	Database              = ""
	StockCollection       = ""
	FiatCollection        = ""
	RabbitUserName        = ""
	RabbitPassword        = ""
	RabbitHost            = ""
	RabbitExchange        = ""
	RabbitStockQueue      = ""
	RabbitFiatQueue       = ""
	RabbitExchangeDLQ     = ""
	RabbitQueueDLQ        = ""
	RabbitTTL             = 0
	RabbitFiatRountingKey = ""
	RabbitStockRoutingKey = ""
)

func Load() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	RabbitTTL, erro = strconv.Atoi(os.Getenv("RABBITMQ_TTL"))
	if erro != nil {
		RabbitTTL = 5000
	}

	ConnectionString = os.Getenv("CONNECTION_STRING")
	StockCollection = os.Getenv("STOCK_COLLECTION")
	FiatCollection = os.Getenv("FIAT_COLLECTION")
	Database = os.Getenv("DATABASE")
	RabbitUserName = os.Getenv("RABBITMQ_USER")
	RabbitPassword = os.Getenv("RABBITMQ_PASS")
	RabbitHost = os.Getenv("RABBITMQ_HOST")
	RabbitExchange = os.Getenv("RABBITMQ_EXCHANGE")
	RabbitFiatQueue = os.Getenv("RABBITMQ_FIAT_QUEUE")
	RabbitStockQueue = os.Getenv("RABBITMQ_STOCK_QUEUE")
	RabbitExchangeDLQ = os.Getenv("RABBITMQ_DLQ_EXCHANGE")
	RabbitQueueDLQ = os.Getenv("RABBITMQ_DLQ_QUEUE")
	RabbitFiatRountingKey = os.Getenv("RABBITMQ_FIAT_ROUTING_KEY")
	RabbitStockRoutingKey = os.Getenv("RABBITMQ_STOCK_ROUTING_KEY")
}
