package main

import (
  "fmt"
  "os"
  "strings"
  "syscall"

  "golang.org/x/crypto/ssh/terminal"

  "metagit.org/blizzlike/wowpasswd/srp"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Printf("USAGE: %v <username>\n", os.Args[0])
    os.Exit(1)
  }

  username := os.Args[1]
  fmt.Printf("Password: ")
  pw1, _ := terminal.ReadPassword(int(syscall.Stdin))
  fmt.Printf("\nRetype Password: ")
  pw2, _ := terminal.ReadPassword(int(syscall.Stdin))

  password := string(pw1)
  retype := string(pw2)

  if strings.Compare(password, retype) != 0 {
    fmt.Printf("\nSorry, passwords do not match!\n")
    os.Exit(2)
  }

  identifier := srp.Hash(username, password)

  auth := srp.New()
  auth.GenerateSalt()
  auth.ComputeVerifier(identifier)

  fmt.Printf("\ns: %v\nv: %v\n", auth.GetSalt(), auth.GetVerifier())
}
