package main

import(
	//"encoding/json"
	"fmt"
	//"io/ioutil"
	"log"
	"net/http"
	//"strconv"
  "github.com/gorilla/mux"
  "github.com/gorilla/websocket"

)

var upgrade = websocket.Upgrader{
ReadBufferSize: 1024,
WriteBufferSize: 1024,
}

/*func index(w http.ResponseWriter, r *http.Request){
  file,err := ioutil.ReadFile("./views/index.html")
  if(err != nil){
    fmt.Fprintf(w,"Error al cargar el HTML")
    return
  }
    
  w.Write(file)
}
*/
func reader(conn *websocket.Conn){
  for{
    messageType, p ,err := conn.ReadMessage()
    if err != nil{
        log.Println(err)
        return
    }

    log.Println(string(p))

    if err := conn.WriteMessage(messageType, p); err != nil{
      log.Println(err)
      return
    }
  }
}


func sendMessage(w http.ResponseWriter, r *http.Request){
 upgrade.CheckOrigin = func(r *http.Request) bool {return true}

 ws, err := upgrade.Upgrade(w,r,nil)

 if(err != nil){
   log.Println(err)
   return
 }

 log.Println("client Connected...")

 reader(ws)
}

func main(){
  router := mux.NewRouter().StrictSlash(true)
  router.PathPrefix("/views/").Handler(http.StripPrefix("/views/", http.FileServer(http.Dir("./views"))))
 // router.HandleFunc("/",index).Methods("GET")
  router.HandleFunc("/ws",sendMessage)
  log.Fatal(http.ListenAndServe(":3000",router)) //Activamos el servidor
  fmt.Println("Run Server")

}
