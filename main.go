package main

import (
 "flag"
 "fmt"
 "obsidian_tips/daily"
)

const (
  CREATE_FLAG = "c"
  CREATE_DAILY_NOTE = "d"

  HELP_FLAG = "h"
  HELP_FLAG_CREATE = "c"
)

func main() {
  var createArg string
  flag.StringVar(&createArg, CREATE_FLAG, "NONE", "Creates notes. To see posible values: -h=c")

  var helpArg string
  flag.StringVar(&helpArg, HELP_FLAG, "NONE", "Help command")

  flag.Parse()

  switch createArg {
  case CREATE_DAILY_NOTE:
    fmt.Println("Creating daily note...")
    fileName, err := daily.CreateDailyNote()

    if err != nil {
      fmt.Println(err)
      return
    }

    fmt.Println("Successfuly created file:", fileName)
  }

  switch helpArg {
  case HELP_FLAG_CREATE:
    fmt.Println(`To create note run -c={type}\nPosible types:\nd - Daily notes`)
  }
}
