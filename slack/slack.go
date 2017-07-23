package main

import (
  "fmt"
  "os/exec"
  "time"
)

func run(api *slack.Client) int {
  rtm := api.NewRTM()
  go rtm.ManageConnection()

  for {
      select {
      case msg := <-rtm.IncomingEvents:
          switch ev := msg.Data.(type) {
          case *slack.HelloEvent:
              log.Print("Hello Event")

          case *slack.MessageEvent:
              log.Printf("Message: %v\n", ev)
              rtm.SendMessage(rtm.NewOutgoingMessage("Hello world", ev.Channel))

          case *slack.InvalidAuthEvent:
              log.Print("Invalid credentials")
              return 1

          }
      }
  }
}
