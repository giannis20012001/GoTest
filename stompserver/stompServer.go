package stompserver

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 5/7/2017.
 */

import (
	"time"
	"github.com/go-stomp/stomp"
)

var options []func(*stomp.Conn) error = []func(*stomp.Conn) error{
	//stomp.ConnOpt.Login("userid", "userpassword"),
	//stomp.ConnOpt.Host("mybroker"),
	stomp.ConnOpt.HeartBeat(360*time.Second, 360*time.Second),
	stomp.ConnOpt.HeartBeatError(360 * time.Second),
}

func ConnectToStompServerWithCred(address string, port string, username string, passwd string) (*stomp.Conn, error) {
	conn, err := stomp.Dial("tcp", address + ":" + port, options...)
	if err != nil {
		return conn, err

	}

	return conn, err

}

func ConnectToStompServerWithOutCred(address string, port string) (*stomp.Conn, error) {
	conn, err := stomp.Dial("tcp", address + ":" + port, options...)
	if err != nil {
		return conn, err

	}

	return conn, err

}

func DisconnectFromStompServer(conn *stomp.Conn) error {
	return conn.Disconnect()

}

func SendMessageToStompServer(conn *stomp.Conn, destination string, contentType string, message string) error {
	tx := conn.Begin()
	err := tx.Send(
		destination, //queue name or topic name e.g. "/queue/test-1" or "/topic/SampleTopic"
		contentType, //e.g. "text/plain"
		[]byte(message)) // message body
	if err != nil {
		return err

	}

	err = tx.Commit()
	if err != nil {
		return err

	}

	return nil

}

func SubscribeToStompServer(conn *stomp.Conn, destination string) (*stomp.Subscription, error) {
	sub, err := conn.Subscribe(
		destination, //queue name or topic name e.g. "/queue/test-1" or "/topic/SampleTopic"
		stomp.AckClient) //Subscribe to queue or topic with client acknowledgement
	if err != nil {
		return sub, err

	}

	return sub, err

}

func GetMessageFromStompServer(sub *stomp.Subscription, conn *stomp.Conn) (*stomp.Message, error) {
	msg := <-sub.C
	if msg.Err != nil {
		return msg, msg.Err

	}

	err := conn.Ack(msg)
	if err != nil {
		return msg, msg.Err

	}

	return msg, msg.Err

}

func UnsubscribeFromStompServer(sub *stomp.Subscription) error {
	err := sub.Unsubscribe()
	if err != nil {
		return err

	}

	return nil

}