package.main

import (
  "fmt"
  "time"
)

func heavy() {
  time.Sleep(time.Second * 3)
}

func main () {
  heavy()
  fmt.Println("fin")
}
