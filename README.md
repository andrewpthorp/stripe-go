# Stripe Go Bindings

You can sign up for a stripe account at https://stripe.com

API Version
===========

This library pins the Stripe API Version. To see the current version, view
`stripe/stripe.go`. You can view more about Stripes API versioning at the
[stripe documentation](https://stripe.com/docs/api#versioning).

Installation
============

Import the library:

    import "github.com/stripe/stripe-go/stripe"


Usage
=====

```go
package main

import (
  "fmt"
  "github.com/stripe/stripe-go/stripe"
)

func main() {

  // use the DefaultClient
  client := stripe.NewClient(nil, "sk_your_secret_key")

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
```

Testing
=======

Tests are all stubbed out. You can view the fixture responses in `fixtures/`.
Ideally, tests will never actually hit the Stripe API. To run the tests:

Install Dependencies

    script/bootstrap

Run Tests

    script/test

License
=======

stripe-go is released under the MIT license. See
[LICENSE](https://github.com/stripe/stripe-go/blob/master/LICENSE).
