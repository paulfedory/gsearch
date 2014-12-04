package main

import (
  "fmt"
  "golang.org/x/blog/content/context/google"
  "golang.org/x/net/context"
)

func main() {
  var (
    ctx    context.Context
  )

  ctx = context.Background()

  results, err := google.Search(ctx, "golang")

  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(results)
  }
}
