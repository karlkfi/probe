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


### Building

Build locally:

```
make
```

Build in docker:

```
make build-docker
```

Build docker builder:

```
make builder
```


### TODO

1. Add SSL certificate validation options (currently ignores cert validity).
2. Detect timeouts better
  - `request canceled while waiting for connection`
  - `read tcp 93.184.216.34:443: use of closed network connection` (https://example.com/)


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
