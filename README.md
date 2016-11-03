# Go GRPC Client Quick Start Guide

This guide will walk you through deploying a Go [grpc][grpc] client on [Deis Workflow][]. The client also runs a Go http server which is routable outside of the cluster and proxies the user requests to a Go [grpc server][grpcserver]. The [Go server][serverinstall] to which the requests needs to be sent can be configured using `deis config:set SERVER_NAME=<server app name>`.

## Usage

```console
$ git clone https://github.com/deis/example-grpc-client.git
$ cd example-grpc-client
$ deis create
Creating Application... done, created breezy-playroom
Git remote deis added for app breezy-playroom
$ deis config:set SERVER_NAME=finest-rabbitry
Creating config... done
=== breezy-playroom Config
SERVER_NAME      finest-rabbitry
$ git push deis master
Counting objects: 8, done.
Delta compression using up to 4 threads.
Compressing objects: 100% (6/6), done.
Writing objects: 100% (8/8), 4.00 KiB | 0 bytes/s, done.
Total 8 (delta 0), reused 0 (delta 0)
Starting build... but first, coffee!
-----> Restoring cache...
       No cache file found. If this is the first deploy, it will be created now.
-----> Go app detected
        !!    
        !!    'GOVERSION' isn't set, defaulting to 'go1.7.1'
        !!    
        !!    Run 'heroku config:set GOVERSION=goX.Y' to set the Go version to use
        !!    for future builds
        !!    
-----> Installing go1.7.1... done
-----> Installing glide v0.12.2... done
-----> Installing hg 3.9... done
        !!    Installing package '.' (default)
        !!    
        !!    To install a different package spec for the next build run:
        !!    
        !!    'heroku config:set GO_INSTALL_PACKAGE_SPEC="<pkg spec>"'
        !!    
        !!    For more details see: https://devcenter.heroku.com/articles/go-dependencies-via-glide
        !!    
-----> Fetching any unsaved dependencies (glide install)
       [INFO]	Downloading dependencies. Please wait...
       [INFO]	--> Fetching google.golang.org/grpc.
       [INFO]	--> Fetching github.com/deis/example-grpc-server.
       [INFO]	--> Fetching golang.org/x/net.
       [INFO]	--> Fetching github.com/golang/protobuf.
       [INFO]	Setting references.
       [INFO]	--> Setting version for github.com/golang/protobuf to 8616e8ee5e20a1704615e6c8d7afcdac06087a67.
       [INFO]	--> Setting version for github.com/deis/example-grpc-server to 683c669e94944e2da3716668396dd557ec3c4093.
       [INFO]	--> Setting version for google.golang.org/grpc to b7f1379d3cbbbeb2ca3405852012e237aa05459e.
       [INFO]	--> Setting version for golang.org/x/net to fb93926129b8ec0056f2f458b1f519654814edf0.
       [INFO]	Exporting resolved dependencies...
       [INFO]	--> Exporting github.com/golang/protobuf
       [INFO]	--> Exporting github.com/deis/example-grpc-server
       [INFO]	--> Exporting golang.org/x/net
       [INFO]	--> Exporting google.golang.org/grpc
       [INFO]	Replacing existing vendor dependencies
-----> Running: go install -v -tags heroku .
       github.com/deis/example-grpc-client/vendor/github.com/golang/protobuf/proto
       github.com/deis/example-grpc-client/vendor/golang.org/x/net/context
       github.com/deis/example-grpc-client/vendor/golang.org/x/net/http2/hpack
       github.com/deis/example-grpc-client/vendor/golang.org/x/net/http2
       github.com/deis/example-grpc-client/vendor/golang.org/x/net/internal/timeseries
       github.com/deis/example-grpc-client/vendor/golang.org/x/net/trace
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/codes
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/credentials
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/grpclog
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/internal
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/metadata
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/naming
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/peer
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc/transport
       github.com/deis/example-grpc-client/vendor/google.golang.org/grpc
       github.com/deis/example-grpc-client/vendor/github.com/deis/example-grpc-server/_proto
       github.com/deis/example-grpc-client
-----> Discovering process types
       Procfile declares types -> web
-----> Checking for changes inside the cache directory...
       Files inside cache folder changed, uploading new cache...
       Done: Uploaded cache (85M)
-----> Compiled slug size is 3.0M
Build complete.
Launching App...
Done, breezy-playroom:v3 deployed to Workflow

Use 'deis open' to view this application in your browser

To learn more, use 'deis help' or visit https://deis.com/

To ssh://git@deis-builder.deis.rocks:2222/breezy-playroom.git
 * [new branch]      master -> master
$ curl http://breezy-playroom.deis.rocks
Powered by Deis
```

## Additional Resources

* [GitHub Project](https://github.com/deis/workflow)
* [Documentation](https://deis.com/docs/workflow/)
* [Blog](https://deis.com/blog/)

[Deis Workflow]: https://github.com/deis/workflow#readme
[grpc]: http://www.grpc.io/docs/quickstart/go.html
[grpcserver]: https://github.com/deis/example-grpc-server
[serverinstall]: https://github.com/deis/example-grpc-server/#example-grpc-server
