
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

  statusObj.removeAttribute('negative', '');
  statusObj.removeAttribute('positive', '');
  statusObj.setAttribute('active', '');
  statusObj.setAttribute('pulse', '');

  statusHdrObj.innerHTML = "Connecting to My Lord...";

   $.ajax({
      // mode: "cors",
      method: "POST",
      url: 'http://localhost:9000/login',
      data:{"username":userNameObj.value, "passcode": passCodeObj.value},
      async:true,
      contentType: 'application/x-www-form-urlencoded;charset=UTF-8',
      dataType : 'html or json',
      // crossDomain:false,
      success: function(data, status, xhr) {        
        setTimeout(loginPassed, 2000);
      },
      error: function(data, status, xhr) {
        console.log(data);
        setTimeout(loginFailed, 2000);
      }
    });

    console.log("Done");
}
