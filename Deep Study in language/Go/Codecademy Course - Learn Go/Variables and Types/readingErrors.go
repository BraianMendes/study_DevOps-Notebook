package main

import "fmt"

func main() {
  var jellybeanCounter int8
  
  // Go will raise errors both for
  // jellybeanCounter being unused
  // and for "fmt" being imported
  // and unused.
  
  // Uncomment the following line
  // and watch the program run
  // successfully!
  
  fmt.Println(jellybeanCounter)
}