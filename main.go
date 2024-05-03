package main

import (
 "flag"
 "fmt"
 "time"
 "os"
 "strings"
 "path/filepath"
)

const (
  PATH = "/workdir/Obsidian/"
  DAILY_PATH = "Everyday_notes/2024/"

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
    fileName, err := createDailyNote()

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

func createDailyNote() (string, error) {
  date := time.Now().Format("02.01.2006")
  currentMonthPath := currentMonth()
  fileName := "Daily " + date + ".md"

  monthPath := filepath.Join("/Users/jbane/" + PATH + DAILY_PATH + currentMonthPath)
  err := os.Mkdir(monthPath, os.ModePerm)
  // TODO: Handle error when path already exist
  if err != nil {
    fmt.Println("Month path mkdir error:")
    fmt.Println(err)
  }

  f, err := os.Create(os.ExpandEnv(monthPath + "/" + fileName))
  if err != nil {
    return "", err
  }

  f.Close()

  return fileName, nil
}


func currentMonth() string {
  r := strings.NewReplacer(
    "January", "Янаварь",
    "February", "Февраль",
    "March", "Март",
    "April", "Арель",
    "May", "Май",
    "June", "Июни",
    "July", "Июль",
    "August", "Август",
    "September", "Сентябрь",
    "October", "Октябрь",
    "November", "Ноябрь",
    "December", "Декабрь",
  )
  time := time.Now().Format("January_2006")
  return r.Replace(time)
}
