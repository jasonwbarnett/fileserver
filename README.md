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

[1]: http://nginx.org/en/docs/http/ngx_http_autoindex_module.html
[2]: https://golang.org/pkg/net/http/#FileServer
[3]: https://github.com/golang/go/issues/23898
