
terminal = document.querySelector(".card-content");
terminal.scrollTop = terminal.scrollHeight;

document.querySelector("#send").addEventListener("click", sendMsg);



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
    // fetch("http://127.0.0.1:8080/data", {
    //   method: "POST",
    //   body: JSON.stringify({
    //     userId: 1,
    //     title: "Fix my bugs",
    //     completed: false
    //   }),
    //   headers: {
    //     'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8',
    //     "Access-Control-Allow-Origin" : "*"
    //   }
    // });

    $.post({
      // type: "POST",
      url: 'http://127.0.0.1:8080/data',
      data:{"name":"Yash"},
      async:true,
      contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
      dataType : 'jsonp',   //you may use jsonp for cross origin request
      crossDomain:true,
      success: function(data, status, xhr) {
        alert(xhr.getResponseHeader('Location'));
      }
    });

    console.log("Worked");

  }
}
