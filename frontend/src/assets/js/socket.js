
import sharedData from './data.js';

console.log("state:", sharedData.AllChats);


class MySocket {
    constructor() {
        this.mysocket = null;
        // this.PostsContainer = document.getElementById("posts");
    }

    send(obj) {
        this.mysocket.send(JSON.stringify(obj));
    }



    connectSocket() {
        var socket = new WebSocket("ws://localhost:8080/ws"); //make sure the port matches with your golang code
        this.mysocket = socket;

        socket.onmessage = (e) => {

            let message = JSON.parse(e.data)
            // console.log("received", message);
            // console.log("pour", message.To, "de", message.From, "actual", sharedData.Myaccount.ID, "latest message", sharedData.LatestChat.ID);
            if (message.Payload === "chat") {
                if (message.To === sharedData.Myaccount.ID && sharedData.LatestChat.ID === message.From) {
                    // console.log(sharedData.AllChats, message.Obj);
                    sharedData.AllChats = message.Obj
                }
            } else if (message.Payload === "group") {
                sharedData.AllChatsGroup = message.Obj
                //si le user a ouvert le chat de groupe correspondant latestchatgroup pour avoir le [] chats correspondant
                // if (message.To===sharedData.Myaccount.ID && sharedData.LatestChat.ID===message.From ){
                //     console.log(sharedData.AllChats,message.Obj);
                //     sharedData.AllChats=message.Obj
                // }
            }

        }
        socket.onopen = () => {
            console.log("socket open")
        };
        socket.onclose = (e) => {
            console.log("socket close")
            if (e.wasClean) {
                console.log(`Closed cleanly, code = ${e.code}, reason = ${e.reason}`);
            } else {
                // mysocket.mysocket.send(JSON.stringify({Connected:'yes',Too:"all"}))
                console.error('Connection died', e);
                setTimeout(mysocket.connectSocket(), 1000)
            }
        }
    }
}

window.addEventListener("unload", () => {
    mysocket.mysocket.close(1008, "window reloading");
});


export var mysocket = new MySocket();
mysocket.connectSocket();
