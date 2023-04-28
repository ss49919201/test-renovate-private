package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	protocol := os.Getenv("PROTOCOL")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dsn := user + ":" + password + "@" + protocol + "(" + dbHost + ":3306)/" + dbName
	fmt.Println("Data source name:", dsn)

	// DB接続
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Failed to open db:", err)
		return err
	}
	defer conn.Close()
	if err := conn.Ping(); err != nil {
		fmt.Println("Failed to connect db:", err)
		return err
	}
	r, err := conn.Query("SHOW DATABASES")
	if err != nil {
		fmt.Println("Failed to excecute query 'SHOW DATABASES':", err)
		return err
	}
	fmt.Println("Sucess excecute query 'SHOW DATABASES' ", r)
	fmt.Println("Rows :", r)

	for _, message := range sqsEvent.Records {
		fmt.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
