# Probe

Probe is a command line tool that does one thing:
interrogates a service to determine if it is alive/ready.

Whether the probe result corresponds to aliveness or readiness depends on how the service handles connections and/or requests on the specified address.


### Schemes

Probe supports http, http, and tcp. The scheme is specified by the fully qualified address.

TCP probing simply opens a connection to the supplied address and port and then closes it.

HTTP probing opens a connection to the supplied address, port, and path, makes an HTTP GET request, reads the response, and closes the connection.

HTTPS probing acts like HTTP probing, except with TLS/SSL certification.

HTTP and HTTPS attempts expect a successful response status code (&gt;= 200 and &lt; 400).


#### Examples

Open and close a TCP connection:

```
probe tcp://example.com:80
```

Make an HTTP GET request:

```
probe http://example.com/
```

Make an HTTPS GET request:

```
probe https://example.com/
```


### Exit Codes

- `0` - Success
- `1` - Runtime Error
- `2` - Input Validation Error

The error description will be printed to STDERR.


### Retries

Probe can optional retry attempts (e.g. `--max-attempts=2`).

Related Options:

- `--max-attempts` - Maximum number of attempts to make (unlimitted: -1) (default 1)
- `--retry-delay` - Delay between attempts. Valid time units: ns, us (or µs), ms, s, m, h. (default 1s)

#### Examples

Retry 5 times with a half second delay between each attempt:

```
probe --max-attempts=5 --retry-delay=0.5s http://example.com/
```


### Timeouts

Probe can optionally time out (e.g. `--timeout=5s`).

Related Options:

- `--timeout` - Time after which the attempt(s) will be interrupted. Valid time units: ns, us (or µs), ms, s, m, h. (default -1ns)
- `--attempt-timeout` - Time after which each individual attempt will be interrupted. Valid time units: ns, us (or µs), ms, s, m, h. (default -1ns)

#### Examples

Retry for 30 seconds with a two second delay between each attempt and five second timeout for each attempt:

```
probe --timeout=30s --max-attempts=-1 --retry-delay=2s --attempt-timeout=5s http://example.com/
```


### DNS Resolution

Probe uses Go's DNS resolver, which can be configured by environment variable to force Go-based or C-based resolution.

See http://golang.org/pkg/net/#hdr-Name_Resolution for more details.

If the Go-based resolver is used, there is no built-in DNS caching, which may or may not be desirable.


### Install

#### Pre-compiled Binaries

Find platform-specific download URL to the [latest release](https://github.com/karlkfi/probe/releases/latest).

Linux Example:

```
curl -sL https://github.com/karlkfi/probe/releases/download/v0.3.0/probe-0.3.0-linux_amd64.tgz | sudo tar zxf - -C /usr/local/bin/
```


#### From Source

[Go get](https://golang.org/doc/install) can be used to install the bleeding edge revision (HEAD):

```
go get github.com/karlkfi/probe
```


#### With Homebrew

[Homebrew](http://brew.sh/) can be used to install the latest released version using [a formula](https://raw.githubusercontent.com/karlkfi/homebrew-terminal/master/probe.rb) from the [karlkfi/homebrew-terminal](https://github.com/karlkfi/homebrew-terminal) tap:

```
brew tap karlkfi/homebrew-terminal
brew install probe
```


### Build

After checking out the Git repo, Probe can be compiled with [Make](https://www.gnu.org/software/make/).

Build locally (Requires [Go (Golang)](https://golang.org/doc/install)):

```
make
```

Build the linux binary in docker (Requires [Docker](https://docs.docker.com/installation/)):

```
make build-docker
```

Build the Linux and Mac binaries in docker (Requires [Docker](https://docs.docker.com/installation/)):

```
make build-docker-cross
```

Build the docker builder (Requires [Docker](https://docs.docker.com/installation/)):

```
make builder
```

Package binaries:

```
for f in pkg/*; do cd $f && tar -zcvf probe-${VERSION}-$(basename $f).tgz probe; cd ../..; done
```


### TODO

1. Add SSL certificate validation options (currently ignores cert validity).
2. Add configurable DNS caching

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
