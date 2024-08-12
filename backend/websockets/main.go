package websockets

type WebsocketClient interface {
	Emit(channel string, event string, message string) error
}