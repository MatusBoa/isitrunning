package pusher

import "github.com/pusher/pusher-http-go/v5"

func CreateWebsocketClient(id string, secret string, key string, host string) PusherWebsocketClient {
	return PusherWebsocketClient{
		client: pusher.Client{
			AppID:  id,
			Secret: secret,
			Key:    key,
			Host:   host,
		},
	}
}

type PusherWebsocketClient struct {
	client pusher.Client
}

func (ws *PusherWebsocketClient) Emit(channel string, event string, message string) error {
	return ws.client.Trigger(channel, event, message)
}
