package main

import(
	//"encoding/json"
  "fmt"
	"io/ioutil"
	"log"
	"net/http"
	//"strconv"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"

)



var clients = make(map[*websocket.Conn]bool) // connected clients


var upgrader = websocket.Upgrader{
 ReadBufferSize:  1024,
 WriteBufferSize: 1024,
}

func index(w http.ResponseWriter, r *http.Request){
  file,err := ioutil.ReadFile("./views/index.html")

  if(err != nil){
    fmt.Fprintf(w,"Error al cargar el HTML")
    return
  }

  w.Write(file)
}

func reader(conn *websocket.Conn){
for {
   // Capture the message from the client
      messageType, p ,err := conn.ReadMessage()

    if err != nil{ 
          for client := range clients{
          delete(clients,client)
         }
       log.Println(err)
       return
      }

  // Walk through connected clients and display the message
       for client := range clients {
          err := client.WriteMessage(messageType,p)
               if err != nil {
                    log.Printf("error: %v", err)
                    client.Close()
                    delete(clients, client)
                    return 
              }
          }
      } 
}
    


func handleConnections(w http.ResponseWriter, r *http.Request){
   // Upgrade initial GET request to a websocket
 upgrader.CheckOrigin = func(r *http.Request) bool {return true}

 ws, err := upgrader.Upgrade(w,r,nil)

 if(err != nil){
   log.Println(err)
   ws.Close()
   return
 }

 //Register new client
 clients[ws] = true

 log.Println("client Connected...")

 reader(ws)


}

func main(){
  fmt.Println("Run Server")
  router := mux.NewRouter().StrictSlash(true)
  router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir("./views"))))
  router.HandleFunc("/",index).Methods("GET")
  router.HandleFunc("/ws",handleConnections)
  log.Fatal(http.ListenAndServe(":3000",router)) //Activamos el servidor
}
