package main

import (
  "crypto/sha1"
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

  username := strings.ToUpper(os.Args[1])
  fmt.Printf("Password: ")
  pw1, _ := terminal.ReadPassword(int(syscall.Stdin))
  fmt.Printf("\nRetype Password: ")
  pw2, _ := terminal.ReadPassword(int(syscall.Stdin))

  pw := string(pw1)
  retype := string(pw2)

  if strings.Compare(pw, retype) != 0 {
    fmt.Printf("\nSorry, passwords do not match!\n")
    os.Exit(2)
  }

  password := strings.ToUpper(pw)

  sha := sha1.New()
  sha.Write([]byte(fmt.Sprintf("%s:%s", username, password)))
  identifier := fmt.Sprintf("%X", sha.Sum(nil))

  auth := srp.New()
  auth.GenerateSalt()
  auth.ComputeVerifier(identifier)

  fmt.Printf("\ns: %v\nv: %v\n", auth.GetSalt(), auth.GetVerifier())
}
