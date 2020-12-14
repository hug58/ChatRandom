const form = document.getElementById("form-text")
const socket = new WebSocket("ws://localhost:3000/ws")
console.log("Try connect to WebSocket server")

socket.onopen = () =>{
console.log("Connected!")
  socket.send("hello from cliente xd")
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
  console.log(message)

})

