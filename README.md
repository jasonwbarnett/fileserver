# jsonfs

## What is this?

There was a feature request ([golang/go#23898][3]) posted on golang to serve
directory listings in JSON instead of HTML. This would mimick
similar behavior found in NGINX's [ngx_http_autoindex_module][1].

## How do I use it?

This is a direct drop-in replacement for `func FileServer(root FileSystem) Handler`.
Just import this package as `jsonfs_http` and use the `FileServer` function.

Any request sent to the file server with the `Accept` header set to `application/json`
will respond with a json response.

### Example

```golang
// main.go
package main

import (
	"log"
	"net/http"

	jsonfs_http "github.com/jasonwbarnett/jsonfs/net/http"
)

func main() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", jsonfs_http.FileServer(http.Dir("/usr/share/doc"))))
}
```

```text
go run main.go
```

List directory contents as JSON

```text
$ curl 'http://127.0.0.1:8080/' -H 'Accept: application/json' | jq '.'
[
  {
    "name": "bash",
    "type": "directory",
    "mtime": "2017-07-15T14:09:01-07:00"
  },
  {
    "name": "cups",
    "type": "directory",
    "mtime": "2017-12-10T16:16:59-08:00"
  },
  {
    "name": "groff",
    "type": "directory",
    "mtime": "2017-07-15T14:07:46-07:00"
  },
  {
    "name": "ntp",
    "type": "directory",
    "mtime": "2017-05-04T17:21:03-07:00"
  },
  {
    "name": "postfix",
    "type": "directory",
    "mtime": "2017-07-25T18:47:14-07:00"
  }
]
```

List directory contents as HTML

```text
$ curl 'http://127.0.0.1:8080/'
<pre>
<a href="bash/">bash/</a>
<a href="cups/">cups/</a>
<a href="groff/">groff/</a>
<a href="ntp/">ntp/</a>
<a href="postfix/">postfix/</a>
</pre>
```

[1]: http://nginx.org/en/docs/http/ngx_http_autoindex_module.html
[2]: https://golang.org/pkg/net/http/#FileServer
[3]: https://github.com/golang/go/issues/23898
