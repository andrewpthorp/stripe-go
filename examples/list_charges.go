// List all Charges in your account. To run this on your system:
//
// STRIPE_SECRET_KEY=sk_your_key go run examples/list_charges.go
package main

import (
  "fmt"
  "os"
  "github.com/stripe/stripe-go/stripe"
)

func main() {
  client := stripe.NewClient(os.Getenv("STRIPE_SECRET_KEY"))
  charges, err := client.Charges.List()

  if err != nil {
    fmt.Println("Error listhing charges: ", err)
  } else {

    for _, v := range charges.Data {
      fmt.Println("Charge Id: ", v.Id)
    }
  }
}

