
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
   $.ajax({
      mode: "cors",
      method: "POST",
      url: 'http://localhost:9000/data',
      data:{"name":cmdObj.innerHTML},
      async:true,
      contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
      dataType : 'json',   //you may use jsonp for cross origin request
      crossDomain:true,
      success: function(data, status, xhr) {
        alert(xhr.getResponseHeader('Location'));
      }
    });

    cmdObj.innerHTML = "";

 
  }
}
