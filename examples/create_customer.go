// Create a Customer with a default Card and some Metadata. To run this on your
// system:
//
// STRIPE_SECRET_KEY=sk_your_key go run examples/create_customer.go
package main

import (
  "fmt"
  "os"
  "github.com/stripe/stripe-go/stripe"
)

func main() {
  client := stripe.NewClient(os.Getenv("STRIPE_SECRET_KEY"))

  params := stripe.CustomerParams{
    Description: "A pretty awesome customer",
    Email: "apt@stripe.com",
    CardParams: &stripe.CardParams{
      Name: "Andrew Thorp",
      Number: "4242424242424242",
      ExpMonth: 01,
      ExpYear: 2020,
    },
    Metadata: stripe.Metadata{
      "awesome": "yes",
      "twitter": "@andrewpthorp",
    },
  }

  customer, err := client.Customers.Create(&params)

  if err != nil {
    fmt.Println("Error creating customer: ", err)
  } else {
    fmt.Println("Created customer: ", customer.Id)
  }
}
