# Probe

Probe is a command line tool that does one thing:
interrogates a service to determine if it is alive/ready.

Whether the probe result corresponds to aliveness or readiness depends on how the service handles connections and/or requests on the specified address.


### Probe Schemes

Probe supports three schemes of probing: tcp, http, &amp; https.

TCP probing simply opens a connection to the supplied address and port and then closes it.

HTTP probing opens a connection to the supplied address, port, and path, makes an HTTP GET request, reads the response, and closes the connection.

HTTPS probing acts like HTTP probing, except with TLS/SSL certification.


### Example Usage

TCP probing:

```
probe tcp://example.com:80
```

HTTP probing, with timeout:

```
probe --timeout 5s http://example.com/
```

HTTPS probing, with timeout shortcut:

```
probe -t 1s https://example.com/
```


### Exit Codes

- `0` - Success
- `1` - Runtime Error
- `2` - Input Validation Error

The error description will be printed to STDERR.


### Install

#### From Source

Install the bleeding edge revision (HEAD):

```
go get github.com/karlkfi/probe
```

(Requires [Go (Golang)](https://golang.org/doc/install).)


#### With Homebrew

[Homebrew formula](https://raw.githubusercontent.com/karlkfi/homebrew/probe/Library/Formula/probe.rb) available on [a branch](https://github.com/karlkfi/homebrew/tree/probe), pending upstreaming.

```
wget https://raw.githubusercontent.com/karlkfi/homebrew/probe/Library/Formula/probe.rb \
  -O /usr/local/Library/Formula/probe.rb
brew install probe
```

(Requires [Homebrew](http://brew.sh/) and [wget](http://www.gnu.org/software/wget/) (`brew install wget`).)


### Build

Build locally:

```
make
```

(Requires [Go (Golang)](https://golang.org/doc/install) and [Make](https://www.gnu.org/software/make/).)

Build in docker:

```
make build-docker
```

(Requires [Docker](https://docs.docker.com/installation/).)

Build docker builder:

```
make builder
```

(Requires [Docker](https://docs.docker.com/installation/).)


### DNS Resolution

Probe uses Go's DNS resolver, which can be configured by environment variable to force Go-based or C-based resolution.

See http://golang.org/pkg/net/#hdr-Name_Resolution for more details.

If the Go-based resolver is used, there is no built-in DNS caching, which may or may not be desirable.


### TODO

1. Add SSL certificate validation options (currently ignores cert validity).
2. Detect timeouts better
  - `request canceled while waiting for connection`
  - `read tcp 93.184.216.34:443: use of closed network connection` (https://example.com/)
3. Upload cross-platform pre-compiled binaries
4. Add configurable DNS caching

### License

   Copyright 2015 Karl Isenberg

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
