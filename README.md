# Golang email reply remover

## Installation

     go get github.com/Liberkeys/replyremover

## Usage

```Go
package main

import (
    "fmt"

    "github.com/Liberkeys/replyremover"
)

func main() {
    email := `Awesome! I haven't had another problem with it.

On Aug 22, 2011, at 7:37 PM, defunkt<reply@reply.github.com> wrote:

> Loader seems to be working well.`

    fmt.Println(replyremover.RemoveReplies(email))
}
```
