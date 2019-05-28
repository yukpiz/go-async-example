package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {

	e := gin.Default()

	e.GET("/hello", func(gc *gin.Context) {
		ac := NewAsyncClient()
		sjob := NewSlackNotifyJob("hello!")
		ac.Run(sjob)

		gc.String(http.StatusOK, "success!")
	})

	e.Run(":8888")
}

type AsyncClientInterface interface {
	Run(i AsyncExecutor)
}

type AsyncClient struct{}

func NewAsyncClient() AsyncClientInterface {
	return &AsyncClient{}
}

func (c *AsyncClient) Run(i AsyncExecutor) {
	go func() {
		if err := i.Exec(); err != nil {
			log.Fatal(err)
		}
	}()
}

type AsyncExecutor interface {
	Exec() error
}

type SlackNotifyJob struct {
	Message string
}

func NewSlackNotifyJob(msg string) AsyncExecutor {
	return &SlackNotifyJob{
		Message: msg,
	}
}

func (job *SlackNotifyJob) Exec() error {
	time.Sleep(5 * time.Second)
	log.Printf("POST SLACK: %s\n", job.Message)
	return nil
}
