package main
import "fmt"

func queryDatabase(query string) string {
  var result string
  connectDatabase()
  // Add deferred call to disconnectDatabase() here
  defer disconnectDatabase()
  if query == "SELECT * FROM coolTable;" {
    result = "NAME|DOB\nVincent Van Gogh|March 30, 1853"
  }  
  fmt.Println(result)
  return result
}

func connectDatabase() {
  fmt.Println("Connecting to the database.")
}

func disconnectDatabase() {
  fmt.Println("Disconnecting from the database.")
}

func main() {
  queryDatabase("SELECT * FROM coolTable;")
}