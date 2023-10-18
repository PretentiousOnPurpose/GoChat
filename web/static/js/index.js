
terminal = document.querySelector(".card-content");
terminal.scrollTop = terminal.scrollHeight;
window.SpeechRecognition = window.webkitSpeechRecognition || window.SpeechRecognition;
// const WebSocket = require('ws');

document.querySelector("#send").addEventListener("click", sendMsg);

// const socket = new WebSocket("ws://localhost:8080");

connect();

function connect() {
  ws = new WebSocket('ws://localhost:8080');
  setTimeout(bindEvents, 1000);
  setReadyState();
}

function bindEvents() {
  ws.onopen = function() {
    console.log('onopen called');
    setReadyState();
  };
}

function setReadyState() {
  console.log('ws.readyState: ' + ws.readyState);
}

function sendMsg() {
  termObj = document.querySelector("#terminal");
  cmdObj = document.querySelector("#cmd")

  if (cmdObj.innerHTML !== "") {
    if (cmdObj.innerHTML === "clear") {
      termObj.innerHTML = "";  
    } else {
      termObj.innerHTML = termObj.innerHTML + "\n<p>You: " + cmdObj.innerHTML + "</p>" 
    }

    cmdObj.innerHTML = "";
    //sendMsgToLord(msg)
    console.log("Send 1");
    
    ws.onopen = () => ws.send("Message");
    
    console.log("Send 2");

  }
}
