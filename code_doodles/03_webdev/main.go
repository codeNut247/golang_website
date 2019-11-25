package main

func headers(w *http.ResponseWriter, r http.Request) {

}

func main() {
  server := http.Server {
      Addr: "127.0.0.1:8080"
  }
  mux := http.DefaultServeMux

  

}
