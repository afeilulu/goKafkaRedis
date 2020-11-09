package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"afeilulu.com/goKafkaRedis/config"
	"afeilulu.com/goKafkaRedis/operator"
	"gopkg.in/confluentinc/confluent-kafka-go.v1/kafka"
)

func main() {

	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "Usage: %s <broker> <group> <topics..>\n",
			os.Args[0])
		os.Exit(1)
	}

	broker := os.Args[1]
	group := os.Args[2]
	topics := os.Args[3:]

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":               broker,
		"group.id":                        group,
		"session.timeout.ms":              6000,
		"go.events.channel.enable":        true,
		"go.application.rebalance.enable": true,
		"security.protocol":               "SASL_PLAINTEXT",
		"sasl.mechanism":                  "PLAIN",
		"sasl.username":                   "nroad",
		"sasl.password":                   "naps@2019@nroad",
		// Enable generation of PartitionEOF when the
		// end of a partition is reached.
		"enable.partition.eof": false,
		"auto.offset.reset":    "earliest"})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Created Consumer %v\n", c)

	err = c.SubscribeTopics(topics, nil)

	offset := ""

	run := true

	for run == true {
		select {
		case sig := <-sigchan:
			fmt.Printf("offset end: %s \n", offset)
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false

		case ev := <-c.Events():
			switch e := ev.(type) {
			case kafka.AssignedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.Assign(e.Partitions)
			case kafka.RevokedPartitions:
				fmt.Fprintf(os.Stderr, "%% %v\n", e)
				c.Unassign()
			case *kafka.Message:
				// fmt.Printf("%% Message on %s:\n%s\n",
				// 	*e.TopicPartition.Topic, string(e.Value))

				switch *e.TopicPartition.Topic {
				case "http":
					// fmt.Printf("%% Message on %s: offset %s\n",
					// 	*e.TopicPartition.Topic, e.TopicPartition.Offset)
					if (len(offset) < 1 ) {
						fmt.Printf("offset start: %s\n", e.TopicPartition.Offset)
					}
					offset = e.TopicPartition.Offset.String()
					
					operator.Handle(e.Value)
					
				}
			case kafka.PartitionEOF:
				fmt.Printf("%% Reached %v\n", e)
			case kafka.Error:
				// Errors should generally be considered as informational, the client will try to automatically recover
				fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			}
		}
	}

	fmt.Printf("Closing consumer\n")
	c.Close()

	fmt.Printf("Closing redis\n")
	config.RedisClient.Close()
}