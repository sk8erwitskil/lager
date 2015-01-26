package main

import (
  "bufio"
  "flag"
  "fmt"
  "os"
  "time"
)

var (
  linePrefix = flag.String("prefix", "INFO", "Prefix for each printed line")
  outputWriter = flag.String("output", "stderr", "Where to output to. Either stdout or stderr")
  useDate = flag.Bool("useDate", true, "Insert the date at the beginning of the log line")
)

type Printer struct {
  prefix  string
  useDate bool
  writer *os.File
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
  fmt.Fprintf(p.writer, prefix + ": " + words)
}

func main() {
  flag.Parse()

  var writer *os.File
  switch *outputWriter {
    case "stdout":
      writer = os.Stdout
    case "stderr":
      writer = os.Stderr
    default:
      panic("Invalid output: " + *outputWriter)
  }

  printer := &Printer{
      prefix: *linePrefix,
      useDate: *useDate,
      writer: writer,
  }

  bio := bufio.NewReader(os.Stdin)
  for {
    line, _ := bio.ReadString('\n')
    if (len(line) == 0) {
      continue
    }
    printer.Print(line)
  }
}
