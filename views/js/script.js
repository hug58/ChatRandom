const form = document.getElementById("form-text")
const socket = new WebSocket("ws://localhost:3000/ws")
console.log("Try connect to WebSocket server")

socket.addEventListener('message', function (event) {
    console.log('Message from server ', event.data);
});


socket.onopen = (e) =>{
console.log("Connected!",e)
}

socket.onmessage = (e) =>{
  document.querySelector('.displayChat').innerHTML += `
   <div class="chat-bubble">
        <p>${e.data}</p>
      </div>

  `
  console.log(e.data)
}

socket.onclose = (e)=>{
console.log("Connection finish",e)
}

socket.onerror = (err)=>{
  console.log("Error:",err)
}


form.addEventListener("submit", (e) => {
  e.preventDefault()
  const message = form['textBox'].value
  socket.send(message)
  form.reset()
})
