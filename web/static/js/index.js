
terminal = document.querySelector(".card-content");
terminal.scrollTop = terminal.scrollHeight;
window.SpeechRecognition = window.webkitSpeechRecognition || window.SpeechRecognition;

document.querySelector("#send").addEventListener("click", sendMsg);

async function sendMsg() {
  termObj = document.querySelector("#terminal");
  cmdObj = document.querySelector("#cmd")

  if (cmdObj.innerHTML !== "") {
    termObj.innerHTML = termObj.innerHTML + "\n<p>$: " + cmdObj.innerHTML + "</p>" 
    cmdObj.innerHTML = "";
  }
}


function sleep(ms) {
  return new Promise(resolve => setTimeout(resolve, ms));
}