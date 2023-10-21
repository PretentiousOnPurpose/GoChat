
// terminal = document.querySelector(".card-content");
// terminal.scrollTop = terminal.scrollHeight;

document.querySelector("#loginBtn").addEventListener("click", sendLoginDetails);

function loginFailed() {
  statusObj = document.querySelector("status-indicator");
  statusHdrObj = document.querySelector(".card-header p");

  statusObj.setAttribute('negative', '');
  statusObj.removeAttribute('pulse');
  statusHdrObj.innerHTML = "Connection failed to My Lord";
}

function loginPassed() {
  statusObj = document.querySelector("status-indicator");
  statusHdrObj = document.querySelector(".card-header p");

  statusObj.setAttribute('positive', '');
  // statusObj.removeAttribute('pulse');
  statusHdrObj.innerHTML = "The Lord is with you";
}

function sendLoginDetails() {
  userNameObj = document.querySelector("#username");
  passCodeObj = document.querySelector("#passcode");

  statusObj = document.querySelector("status-indicator");
  statusHdrObj = document.querySelector(".card-header p");

  statusObj.setAttribute('active', '');
  statusObj.setAttribute('pulse', '');

  statusHdrObj.innerHTML = "Connecting to My Lord...";

   $.ajax({
      mode: "cors",
      method: "POST",
      url: 'http://127.0.0.1:8080/login',
      data:{"username":userNameObj.value, "passcode": passCodeObj.value},
      async:true,
      contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
      dataType : 'json',
      crossDomain:true,
      success: function(data, status, xhr) {
        alert(xhr.getResponseHeader('Location'));
      },
      error: function(data, status, xhr) {
        console.log(data);
      }
    });

    console.log("Done");
    setTimeout(loginPassed, 2000);
}
