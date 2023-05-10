package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/z7zmey/php-parser/php5"
	"github.com/z7zmey/php-parser/visitor"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitConn struct {
	conn *amqp.Connection

	text []byte
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func FileCreate(src []byte) {
	// src := r.text

	var b bytes.Buffer

	par := php5.NewParser(src, "")
	par.Parse()

	visitor := visitor.GoDumper{
		Writer: &b,
	}

	//
	rootNode := par.GetRootNode()
	rootNode.Walk(&visitor)

	f, err := os.Create("../parser/example.go")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// Write the package declaration to the file
	_, err = fmt.Fprintf(f, "package parser\n\n")
	if err != nil {
		panic(err)
	}

	// Write the import statements to the file
	_, err = fmt.Fprintf(f, "type program *Root\n\n")
	if err != nil {
		panic(err)
	}

	// Write the code to the file
	_, err = fmt.Fprintf(f, "func Test() program {\n    p:= "+b.String()+"\n return p \n}\n")
	if err != nil {
		panic(err)
	}

	fmt.Println("Written")

}

// func (r *RabbitConn) ReceiveFrom() {
// 	ch, err := r.conn.Channel()
// 	failOnError(err, "Failed to open a channel")
// 	defer ch.Close()

// 	q, err := ch.QueueDeclare(
// 		"one", // name
// 		false, // durable
// 		false, // delete when unused
// 		false, // exclusive
// 		false, // no-wait
// 		nil,   // arguments
// 	)
// 	failOnError(err, "Failed to declare a queue")

// 	msgs, err := ch.Consume(
// 		q.Name, // queue
// 		"",     // consumer
// 		true,   // auto-ack
// 		false,  // exclusive
// 		false,  // no-local
// 		false,  // no-wait
// 		nil,    // args
// 	)
// 	failOnError(err, "Failed to register a consumer")

// 	var forever chan struct{}

// 	go func() {
// 		for d := range msgs {
// 			log.Printf("Received a message: %s", d.Body)
// 			r.text = d.Body
// 			r.FileCreate()
// 		}
// 	}()

// 	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
// 	<-forever
// }

func main() {
	// con, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	// failOnError(err, "Failed to connect to RabbitMQ")
	// defer con.Close()

	// d := RabbitConn{
	// 	conn: con,
	// }
	// d.ReceiveFrom()

	// src := []byte(`<?php
	// <?php
	// $a = $_GET['name'];
	// echo "Name is".$a;
	// `)

	src, _ := ioutil.ReadFile("index.php")
	FileCreate(src)
}
