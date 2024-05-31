# Emailist

This library providers a higher level interface to sending messages via SMTP. It makes it easy to send text, html only, or mixed part emails.

## Table of contents

* [Installation](#installation)
* [Quickstart](#quickstart)
* [Advanced configuration](#advanced-configuration)
* [Contributing](#contributing)
* [License](#license)

## Installation

Use go get:

```bash

go get github.com/fluynt/emailist
```

## Quickstart

1. Configure your smtp server:

```go
c := NewConfig()
```

1. Create an identity, the person from whom an email is going to be sent:

```go
i := NewIdentity("Justa Person", "justa@no.response")
```

1. Create a message, that has TEXT and HTML parts.

```go
m := NewMessage("I am a subject", "I am a body", "<b>I am an html body</b>")
```

1. Send the message

```go
c.SendMail(i, ["to@person.com", "toanother@person.com"], m)
```

## Advanced Configuration

All configuration variables are defined in the [config}(./config.go) object. Configuration occurs by setting (optionally) any or all of the following environment variables.

* MAIL_HOST - The hostname of the SMTP server, defaults to localhost

* MAIL_PORT - The port associated with the smtp server, defaults to 25.

* MAIL_USER - If set, the username for SMTP authentication

* MAIL_PASSWORD - If set, the password for SMTP authentication

* MAIL_SSL - If set, use SSL

* MAIL_TLS - If set, use TLS

* DEBUG - If set, output messages, rather then send them.

These variables, can all be set with a custom prefix (i.e `foo_MAIL_HOST`). This allows for the library to be used repeteatedly, with multiple configurations. To do so, try the following - note, the underscore is assumed, and doesn't have to be passed in with the `prefix`.

```go
c := NewConfigWithPrefix("foo")
```

## Contributing

1. Open an issue

2. Open a pull request and link to the issue

## License

[The MIT License](./LICENSE)
