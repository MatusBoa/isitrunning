import Pusher from "pusher-js"

const pusher = new Pusher('app-key', {
    wsHost: '127.0.0.1',
    wsPort: 6001,
    forceTLS: false,
    disableStats: true,
    enabledTransports: ['ws', 'wss'],
    cluster: "eu",
})

export default pusher