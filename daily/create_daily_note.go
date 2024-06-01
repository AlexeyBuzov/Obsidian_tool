package daily

import (
 "fmt"
 "time"
 "os"
 "strings"
 "path/filepath"
)

const (
  PATH = "/workdir/Obsidian/"
  DAILY_PATH = "Everyday_notes/2024/"
)

func CreateDailyNote() (string, error) {
  monthPath, err := createMonthPath()
  if err != nil {
    fmt.Printf("Month path mkdir error: %s", err)
  }

  return createDailyNoteFile(monthPath)
}


func createMonthPath() (string, error) {
  currentMonth := currentMonth()
  monthPath := filepath.Join("/Users/jbane/" + PATH + DAILY_PATH + currentMonth)
  err := os.Mkdir(monthPath, os.ModePerm)

  return monthPath, err
}


func currentMonth() string {
  r := strings.NewReplacer(
    "January", "Янаварь",
    "February", "Февраль",
    "March", "Март",
    "April", "Арель",
    "May", "Май",
    "June", "Июнь",
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

func createDailyNoteFile(monthPath string) (string, error) {
  fileName := currentDayFileNameWithoutFormat() + ".md"

  f, err := os.Create(os.ExpandEnv(monthPath + "/" + fileName))

  if err != nil {
    return "", err
  }
  defer f.Close()

  err = pushFileNameToList(monthPath)

  return fileName, err
}

func currentDayFileNameWithoutFormat() string {
  return "Daily " + time.Now().Format("02.01.2006")
}

func pushFileNameToList(monthPath string) error {
  currentMonth := currentMonth()
  listFileName := "012 " + currentMonth + ".md"
  f, err := os.OpenFile(monthPath + "/" + listFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

  if err != nil {
	   return err
  }
  defer f.Close()

  value := fmt.Sprintf("\n[[%s]]", currentDayFileNameWithoutFormat())
  if _, err := f.WriteString(value); err != nil {
	   return err
  }

  return nil
}
