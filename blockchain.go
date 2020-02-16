package main

import(
      "bytes"
      "fmt"
      "crypto/sha256"
)

type Block struct{
  Hash       int
  Data       int
  prevhash   int
}
