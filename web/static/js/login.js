
// terminal = document.querySelector(".card-content");
// terminal.scrollTop = terminal.scrollHeight;

document.querySelector("#loginBtn").addEventListener("click", sendLoginDetails);

function sendLoginDetails() {
  userNameObj = document.querySelector("#username");
  passCodeObj = document.querySelector("#passcode")

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
}
