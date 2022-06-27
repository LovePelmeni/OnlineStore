package kafka 


import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"fmt"
	"os"
	"socket"
	"RealTLocation/location_parser/loggers"
	"sync"
)

var (
	kafka_host = os.Getenv("KAFKA_HOST")
	kafka_port = os.Getenv("KAFKA_PORT")
)

type KafkaOptions struct {
	sync.Mutex
	client_id string 
	acks string 
}

type KafkaClient struct {
	sync.Mutex 
	kafka_host string 
	kafka_port string 
	options KafkaOptions
}

type KafkaEventMessage struct {
	sync.Mutex
	message string 
}

func (this *KafkaClient) createProducer() (*kafka.NewProducer, error){
	producer_opts := KafkaOptions{client_id: socket.gethostname(), acks: "all"}
	producer, error := kafka.NewProducer(&kafka.ConfigMaps{
		"bootstrap.servers":
		 fmt.Sprintf("%s:%s", this.kafka_host, this.kafka_port),
		"client.id": producer_opts.client_id,
		"acks": producer_opts.acks,
	})
	if error != nil {
		loggers.ErrorLogger.Println(
		fmt.Sprintf("Error while running kafka producer."))
		return nil, error 
	}
	return producer, nil  
}

func eventResponse(channel chan kafka.Event){
	sended := <- channel
	switch sended {
		default:
			if responded := sended.(kafka.Event); responded != nil {
				loggers.DebugLogger.Println(fmt.Sprintf("Event has been delivered successfully."))
			} else {
				loggers.ErrorLogger.Println(
				fmt.Sprintf("Event has been missed.. Responded with None"))
		}
	}
}

func (this *kafka.NewProducer) ProduceEvent(topic string, Partition string, message string) (bool, error){
	delivery_response_channel := make(chan kafka.Event, 10000)
	go eventResponse(delivery_response_channel)
	this.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: topic, Partition: Partition},
		Value: []byte(message),
	}, delivery_response_channel)
	return false, nil
}

