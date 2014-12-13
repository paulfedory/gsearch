package main

import (
  "fmt"
  "golang.org/x/blog/content/context/google"
  "golang.org/x/net/context"
  "flag"
)

func main() {
  ctx := context.Background()
  flag.Parse()
  term := flag.Args()

  results, err := google.Search(ctx, term[0])

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(results)
  }
}
