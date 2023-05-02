package service

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/go-stomp/stomp/v3"
	"github.com/google/uuid"
	"go.uber.org/ratelimit"
)

type StompConService struct {
	addr  string
	Conn  *stomp.Conn
	Limit ratelimit.Limiter
}

func NewStompConService(addr string) *StompConService {
	conn, err := stomp.Dial("tcp", addr, options...)
	if err != nil {
		fmt.Println("Failed to connect:", err.Error())
	}

	return &StompConService{addr: addr, Conn: conn, Limit: ratelimit.New(50, ratelimit.Per(1*time.Minute))}
}

var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	stomp.ConnOpt.Login("admin", "admin"),
}

func (s *StompConService) Reconnect() error {
	conn, err := stomp.Dial("tcp", s.addr, options...)
	if err != nil {
		return err
	}
	s.Conn = conn

	return nil
}

// func (s *StompConService) Send(destination string, msg string) error {
// 	conn, err := s.Connect()
// 	if err != nil {
// 		return err
// 	}
// 	defer conn.Disconnect()
// 	return conn.Send(
// 		destination,  // destination
// 		"text/plain", // content-type
// 		[]byte(msg))  // body
// }

func catch() {
	if r := recover(); r != nil {
		fmt.Println("Error occured", r)
	} else {
		fmt.Println("Application running perfectly")
	}
}

func (s *StompConService) Subscribe(thread int, destination string, handler func(rate ratelimit.Limiter, err error, msg string, thread int, m *stomp.Message)) error {
	rate := ratelimit.New(7, ratelimit.Per(1*time.Minute))
	for {
		err := s.Reconnect()
		if err != nil {
			fmt.Println("Failed to connect:", err.Error())
			fmt.Println("Trying reset the connection...")
			time.Sleep(time.Second * time.Duration(5))
		} else {
			//fmt.Println("Success to connect")
			fmt.Printf("Connecting to Destination : %v, Thread : %v\n", destination, thread)
			for {
				sub, err := s.Conn.Subscribe(destination, stomp.AckClient, stomp.SubscribeOpt.Header("subscription-type", "ANYCAST"))
				if err != nil {
					break
				}

				defer sub.Unsubscribe()
				defer catch()

				for {
					m := <-sub.C
					if m.Err != nil {
						break
					}
					//s.Limit.Take()
					handler(rate, m.Err, string(m.Body), thread, m)
				}
			}
		}
	}
}

func (s *StompConService) Thread(thread int, destination string) {

	//limit := ratelimit.New(5, ratelimit.Per(1*time.Minute))

	fmt.Println("Listen to : ", destination)

	handler := func(rate ratelimit.Limiter, err error, msg string, thread int, m *stomp.Message) {
		guid := uuid.New()

		t := time.Now()
		//fmt.Printf("%v | Thread : %d, Date : %s\n", guid, thread, t.Local())

		if err != nil {
			fmt.Println("Error caused by ", err)
		} else {
			fmt.Println("Receive Message : \n", msg)
			hit(guid, rate, thread)
			m.Conn.Ack(m)

		}
		fmt.Println("Duration : ", time.Since(t))
	}

	for i := 1; i <= thread; i++ {
		go s.Subscribe(i, destination, handler)

	}
}

func hit(guid uuid.UUID, rate ratelimit.Limiter, thread int) {
	rate.Take()
	var err error
	var client = &http.Client{}

	request, err := http.NewRequest("POST", "http://localhost:8085/dummy", nil)
	//http://localhost:8085/dummy
	if err != nil {
		return
	}

	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	b, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	t := time.Now()
	fmt.Printf("%v |Request HIT, Thread : %d, Date : %s\n", guid, thread, t.Local())
	fmt.Println("Response Hit:	", string(b))
}
