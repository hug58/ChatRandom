const form = document.getElementById("form-text")
const socket = new WebSocket("ws://localhost:3000/ws")

console.log("Try connect to WebSocket server")

socket.onopen = (e) => {
  console.log("Connected!", e)
}

socket.onmessage = (e) => {
  const user = document.getElementById("name").value
  const data = JSON.parse(e.data) 
  document.querySelector('.displayChat').innerHTML += 
`
  <div class=${user === data.username ? 'chat-bubble-transmitter' : 'chat-bubble-receiver'}>
        <p>${data.content}</p>
  </div>

`

}

socket.onclose = (e) => {
  console.log("Connection finish", e)
}

socket.onerror = (err) => {
  console.log("Error:", err)
}


form.addEventListener("submit", (e) => {
  e.preventDefault()
  const message = form['textBox'].value
  const json = JSON.stringify({
    username: document.getElementById("name").value,
    content: message
  })

  socket.send(json)
  form.reset()
})
