package main

import (
	"codeguardian/parser"
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/mux"
	amqp "github.com/rabbitmq/amqp091-go"
)

func FindHash() string {
	f, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("%x", h.Sum(nil))
}

var requestChannel chan []byte

var counter int = 0

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

type RabbitConn struct {
	conn *amqp.Connection
}

func AllGroceries(d RabbitConn) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		d.sendTo()

		var data parser.FinalReport

		fmt.Println("working")
		msg := <-requestChannel
		json.Unmarshal(msg, &data)
		fmt.Println(data)
		json.NewEncoder(w).Encode(data)

	}

}

func (r RabbitConn) sendTo() {

	ch, err := r.conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"one", // name
		false, // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	src := []byte(`<?php
	$host = "localhost";
	$username  = "db_user";
	$passwd = ".mypwd";
	$dbname = "my_db";
 
	//Creating a connection
	$con = mysqli_connect($host, $username, $passwd, $dbname,"3307");
 
	if($con){
	   print("Connection Established Successfully");
	}else{
	   print("Connection Failed ");
	}
	$sql = "SELECT name FROM user";
 $result = mysqli_query($con,$sql);
 if ($result->num_rows > 0) {
  // output data of each row
  while($row = $result->fetch_assoc()) {
   echo $row['name']."<br>";
  }
 } else {
  echo "0 results";
 }
 $con->close();
 ?>
	`)

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        src,
		})
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Sent %s\n", src)

}

func main() {

	con, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer con.Close()

	requestChannel = make(chan []byte)

	d := RabbitConn{
		conn: con,
	}

	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/analyze", AllGroceries(d))

	// <-- Need to run as separate process -->
	// fileName := "./layer1/file_create.go"
	// cmd := exec.Command("go", "run", fileName)
	// cmd.CombinedOutput()

	// fileName = "./layer2/diffmain.go"
	// cmd = exec.Command("go", "run", fileName)
	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println(string(output))

	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:

				if !ok {
					return
				}
				if event.Has(fsnotify.Write) && event.Name == "parser\\example.go" {

					// the file is changed - because it call two time - restrict to one
					fmt.Println("changed")
					if counter%2 == 0 {
						fileName := "./layer2/diffmain.go"
						cmd := exec.Command("go", "run", fileName)
						output, err := cmd.CombinedOutput()
						if err != nil {
							fmt.Println("Error:", err)
							continue
						}

						requestChannel <- output

					}

				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
			counter++
		}
	}()

	// Add a path.
	err = watcher.Add("./parser/example.go")
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.

	log.Fatal(http.ListenAndServe(":3000", r))

}
