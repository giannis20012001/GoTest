package stompserver

/**
 * Created by John Tsantilis
 * (i [dot] tsantilis [at] yahoo [dot] com A.K.A lumi) on 8/11/2017.
 */

func main() {
	stompServerAddress := "127.0.0.1" //TODO: Delete this line
	stompServerPort := "61613" //ActiveMQ STOMP port: 61613
	stompServerPath := "/topic/eu.arcadia.monitoring" //ActiveMQ path: "/queue/test" or "/topic/test"

	con, err := ConnectToStompServerWithOutCred(stompServerAddress, stompServerPort)
	if err == nil {

		msg := "Test message #1"
		SendMessageToStompServer(con, stompServerPath, "text/plain", msg)
		DisconnectFromStompServer(con)

	}

}