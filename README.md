# Stripe Go Bindings

You can sign up for a stripe account at https://stripe.com

Installation
============

Import the library:

    import "github.com/stripe/stripe-go/stripe"


Usage
=====

    package main

    import (
      "fmt"
      "github.com/stripe/stripe-go/stripe"
    )

    func main() {

      client := stripe.NewClient("sk_your_secret_key")

      params := stripe.CustomerParams{
        Email: "apt@stripe.com",
        CardParams: &stripe.CardParams{
          Name: "4242424242424242",
          ExpMonth: 01,
          ExpYear: 2020,
          CVC: "111",
        },
        Metadata: stripe.Metadata{
          "twitter": "@andrewpthorp"
        },
      }

      customer, err := client.Customers.Create(&params)

      if err != nil {
        fmt.Println("Error creating customer: ", err)
      } else {
        fmt.Println("Created customer: ", customer.id)
      }

    }

Unfinished
==========

Most of the API is covered, but this is missing some major things:

* Tests
* Auth/Capture flow
* Filters (passing 'recipient' to Transfers.List(), etc)
* Cards on Customer.Retrieve()

