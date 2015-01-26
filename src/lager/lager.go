package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "strings"
  "time"
)

var (
  linePrefix  = flag.String("prefix", "INFO", "Prefix for each printed line")
  useDate = flag.Bool("useDate", true, "Insert the date at the beginning of the log line")
)


type Printer struct {
  prefix  string
  useDate bool
}

func now() string {
  return time.Now().Format("Jan 2 15:04:05 MST")
}

func (p *Printer) Print(words string) {
  var prefix string
  if (p.useDate) {
    prefix = now() + " " + p.prefix
  } else {
    prefix = p.prefix
  }
  fmt.Println(prefix + ":", words)
}

func main() {
  flag.Parse()

  printer := &Printer{prefix: *linePrefix, useDate: *useDate}

  bio := bufio.NewReader(os.Stdin)
  for {
    line, _ := bio.ReadString('\n')
    if (len(line) == 0) {
      continue
    }
    printer.Print(strings.Trim(line, "\n"))
  }
}
