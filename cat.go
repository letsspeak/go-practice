package main

import (
  "./file"
  "flag"
  "fmt"
  "os"
)

func cat(f *file.File) {
  const NBUF = 512
  var buf [NBUF]byte
  for {
    switch nr, er := f.Read(buf[:]); true {
      case nr < 0:
        fmt.Fprintf(os.Stderr, "cat: error reading from %s: %s\n", f.String(), er.Error())
      case nr == 0: // EOF
        return
      case nr > 0:
        if nw, ew := file.Stdout.Write(buf[0:nr]); nw != nr {
          fmt.Fprintf(os.Stderr, "cat: error writing from %s: %s\n", f.String(), ew.Error())
          os.Exit(1)
        }
    }
  }
}

func main() {
  flag.Parse()
  if flag.NArg() == 0 {
    cat(file.Stdin)
  }
  for i := 0; i < flag.NArg(); i++ {
    f, err := file.Open(flag.Arg(i))
    if f == nil {
      fmt.Fprintf(os.Stderr, "cat: can't open %s error %s\n", flag.Arg(i), err.Error())
      os.Exit(1)
    }
    cat(f)
    f.Close()
  }
}
