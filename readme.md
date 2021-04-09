# checkout API

[![Go](https://github.com/mihu8/checkout/actions/workflows/go.yml/badge.svg)](https://github.com/mihu8/checkout/actions/workflows/go.yml)


# Notes

This project is extremely open-ended. It's a giant rabbit hole. There are way too many things to consider,
* I should have used big/decimal for currency for rupiah.
* Current promotion relies on the fact that all promotions are *orthogonal*, meaning they won't interfere with each
  other. If the promotions are not orthogonal,
  * order or sequence of each promotion matters!
  * we may have to repeatedly run through all promotions until the result is stable
  * we can write code to make sure the promotions are orthogonal
* "id" of product is tricky as well.
* How to define _Quantity_ - we can buy one iPhone or two, but we can also buy 1.3kg of Apple.
* "Check Availability / Stock" & "Checkout" should be atomic operation - I would adopt a pessimistic lock with sanity
  check and 5-minutes timeout.
* concurrency, atomic, rollback etc etc
* on and on...

# ci/cd

I decided to use GitHub Actions as it does not rely on self-hosted infrastructures for small projects like this. This
small system should be treated like a library or module, building a binary does not make any sense, but I put a small
`main.go` in `main/` and output it into Actions artifacts, it's just a demo that I understand what's required...

https://github.com/mihu8/checkout/actions

# GraphQL schema


https://github.com/mihu8/checkout/blob/master/checkout.graphql

__Note:__ GraphQL's schema does not match the golang code.

# TODO

Run lint and build on graphql schema.
