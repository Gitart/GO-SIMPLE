package main

import (
    "github.com/mailgun/mailgun-go"
)

func SendSimpleMessage(domain, apiKey string) (string, error) {
  mg := mailgun.NewMailgun("tutorialedge.net", apiKey, "key-12345671234567")
  m := mg.NewMessage(
    "Excited User <elliot@tutorialedge.net>",
    "Hello",
    "Testing some Mailgun!",
    "elliot@tutorialedge.net",
  )
  _, id, err := mg.Send(m)
  return id, err
}

func main(){
    SendSimpleMessage("postmaster@elliotforbes.co.uk", "key-12345671234567")

}
