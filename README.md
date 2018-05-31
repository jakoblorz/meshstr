# brokerutil
brokerutil provides a common interface to message-brokers for pub-sub applications.

[![GoDoc](https://godoc.org/github.com/jakoblorz/brokerutil?status.svg)](https://godoc.org/github.com/jakoblorz/brokerutil)

Use brokerutil to be able to build pub-sub applications which are not
highly dependent on the message-brokers driver implementation quirks.
brokerutil provides a common interface which enables developers to switch
the message broker without having to rewrite major parts of the applications
pub-sub logic.

## Installation

`go get github.com/jakoblorz/brokerutil`

Use the native drivers in the driver package or implement your own. Currently provided drivers:
- [redis](https://redis.io/) `go get github.com/jakoblorz/brokerutil/driver/redis`
- loopback (for testing) `go get github.com/jakoblorz/brokerutil/driver/loopback`

## Example
This example will use redis as message-broker. It uses the brokerutil native redis driver which
relies on the [github.com/go-redis/redis](http://github.com/go-redis/redis) driver.

```go
package main

import (
    "github.com/jakoblorz/brokerutil"
    "github.com/jakoblorz/brokerutil/stream"
    "github.com/jakoblorz/brokerutil/driver/redis"
    r "github.com/go-redis/redis"
)

func main() {

    // create redis driver to support pub-sub
    driver, err := redis.NewDriver(&redis.DriverOptions{
        channel: "brokerutil-chan",
        driver:  &r.Options{
            Addr:     "localhost:6379",
            Password: "",
            DB:       0
        },
    })

    if err != nil {
        log.Fatalf("could not create brokerutil redis driver: %v", err)
        return
    }

    // create new pub sub using the initialized redis driver
    ps, err := brokerutil.NewPubSubFromDriver(driver)
    if err != nil {
        log.Fatalf("could not create brokerutil pub sub: %v", err)
        return
    }

    // run blocking subscribe as go routine
    go ps.SubscribeSync(func(msg stream.Message) error {
        fmt.Printf("%v", msg)

        return nil
    })

    // start controller routine which blocks execution
    if err := controller.Listen(); err != nil {
        log.Fatalf("could not run controller routine: %v", err)
        return
    }
}
```
