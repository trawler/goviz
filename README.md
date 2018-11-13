goviz
=====

a visualization tool for golang project dependency
![](https://raw.githubusercontent.com/trawler/goviz/master/images/own.png)


This tool is for helping source code reading.
The dependency of the whole code can be visualized quickly.

## Installation

```
go get github.com/trawler/goviz
```

and if not being installed [graphviz](http://www.graphviz.org), install it :

```
brew install graphviz
```

## Usage

```
goviz -i $GOPATH/src/github.com/hashicorp/serf | dot -Tpng -o serf.png
```

### Option

```
Usage:
  goviz [flags]

Flags:
  -d, --depth int       max plot depth of the dependency tree (default 128)
  -f, --focus string    focus on the specific module
  -h, --help            help for goviz
  -i, --input string    input project name
  -l, --leaf            whether leaf nodes are plotted
  -m, --metrics         display module metrics
  -o, --output string   output file (default "STDOUT")
  -s, --search string   top directory of searching
```

## Samples

### [anko](https://github.com/mattn/anko)


```
goviz -i $GOPATH/src/github.com/mattn/anko | dot -Tpng -o anko.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/anko.png)


### [serf](https://github.com/hashicorp/serf)


```
goviz -i $GOPATH/src/github.com/hashicorp/serf | dot -Tpng -o serf.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/serf.png)


### [go-xslate](https://github.com/lestrrat/go-xslate)


```
goviz -i $GOPATH/src/github.com/lestrrat/go-xslate | dot -Tpng -o xslate.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/xslate.png)


### [vegeta](https://github.com/tsenart/vegeta)

#### plot depth 1
```
goviz -i $GOPATH/src/github.com/tsenart/vegeta -l -d 1 | dot -Tpng -o vegeta1.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/vegeta1.png)
#### plot depth 2
```
goviz -i $GOPATH/src/github.com/tsenart/vegeta -l -d 2 | dot -Tpng -o vegeta2.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/vegeta2.png)
#### plot depth 3
```
goviz -i $GOPATH/src/github.com/tsenart/vegeta -l -d 3 | dot -Tpng -o vegeta3.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/vegeta3.png)
#### plot depth 4
```
goviz -i $GOPATH/src/github.com/tsenart/vegeta -l -d 4 | dot -Tpng -o vegeta4.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/vegeta4.png)


### [packer](https://github.com/mitchellh/packer)


```
goviz -i $GOPATH/src/github.com/mitchellh/packer --search SELF| dot -Tpng -o packer.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/packer.png)


### [docker plot depth 1](https://github.com/dotcloud/docker/docker)


```
goviz -i $GOPATH/src/github.com/dotcloud/docker/docker -s github.com/dotcloud/docker -d 1| dot -Tpng -o docker1.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/docker1.png)


### [docker plot depth 2](https://github.com/dotcloud/docker/docker)


```
goviz -i $GOPATH/src/github.com/dotcloud/docker/docker -s github.com/dotcloud/docker -d 2| dot -Tpng -o docker2.png
```
![](https://raw.githubusercontent.com/trawler/goviz/master/images/docker2.png)


### docker's metrics
goviz has a function which outputs the metrics (instability) of go project.

```
goviz -i $GOPATH/src/github.com/dotcloud/docker/docker -m
```
Instability is a value of 0 to 1.
It suggests that it is such an unstable module that this value is high.
It becomes easy to distinguish whether it is a module nearer to  application layer, and whether it is a module near a common library.


```
Inst:1.000 Ca(  0) Ce(  9)	github.com/dotcloud/docker/docker
Inst:0.960 Ca(  1) Ce( 24)	github.com/dotcloud/docker/pkg/libcontainer/nsinit
Inst:0.956 Ca(  2) Ce( 43)	github.com/dotcloud/docker/runtime
Inst:0.950 Ca(  1) Ce( 19)	github.com/dotcloud/docker/api/client
Inst:0.950 Ca(  1) Ce( 19)	github.com/dotcloud/docker/server
Inst:0.909 Ca(  1) Ce( 10)	github.com/dotcloud/docker/api/server
Inst:0.867 Ca(  2) Ce( 13)	github.com/dotcloud/docker/runtime/execdriver/native
Inst:0.857 Ca(  1) Ce(  6)	github.com/dotcloud/docker/runtime/graphdriver/devmapper
Inst:0.833 Ca(  1) Ce(  5)	github.com/dotcloud/docker/runtime/graphdriver/aufs
Inst:0.800 Ca(  1) Ce(  4)	github.com/dotcloud/docker/builtins
Inst:0.800 Ca(  2) Ce(  8)	github.com/dotcloud/docker/runtime/networkdriver/bridge
Inst:0.800 Ca(  1) Ce(  4)	github.com/dotcloud/docker/runtime/execdriver/execdrivers
Inst:0.778 Ca(  2) Ce(  7)	github.com/dotcloud/docker/pkg/libcontainer/network
Inst:0.750 Ca(  1) Ce(  3)	github.com/dotcloud/docker/sysinit
Inst:0.750 Ca(  3) Ce(  9)	github.com/dotcloud/docker/runtime/execdriver/lxc
Inst:0.750 Ca(  1) Ce(  3)	github.com/dotcloud/docker/runtime/execdriver/native/template
Inst:0.727 Ca(  3) Ce(  8)	github.com/dotcloud/docker/graph
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/runtime/execdriver/native/configuration
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/runtime/networkdriver/portmapper
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/runtime/networkdriver/ipallocator
Inst:0.667 Ca(  1) Ce(  2)	github.com/dotcloud/docker/links
Inst:0.571 Ca(  9) Ce( 12)	github.com/dotcloud/docker/runconfig
Inst:0.500 Ca(  2) Ce(  2)	github.com/dotcloud/docker/pkg/selinux
Inst:0.500 Ca(  5) Ce(  5)	github.com/dotcloud/docker/api
Inst:0.500 Ca(  2) Ce(  2)	github.com/dotcloud/docker/daemonconfig
Inst:0.500 Ca(  5) Ce(  5)	github.com/dotcloud/docker/image
Inst:0.500 Ca(  1) Ce(  1)	github.com/dotcloud/docker/pkg/libcontainer/capabilities
Inst:0.500 Ca(  1) Ce(  1)	github.com/gorilla/mux
Inst:0.500 Ca(  1) Ce(  1)	github.com/dotcloud/docker/runtime/graphdriver/btrfs
Inst:0.500 Ca(  1) Ce(  1)	github.com/dotcloud/docker/runtime/graphdriver/vfs
Inst:0.444 Ca( 10) Ce(  8)	github.com/dotcloud/docker/archive
Inst:0.333 Ca(  2) Ce(  1)	github.com/dotcloud/docker/opts
Inst:0.333 Ca(  2) Ce(  1)	github.com/dotcloud/docker/runtime/networkdriver/portallocator
Inst:0.250 Ca(  6) Ce(  2)	github.com/dotcloud/docker/registry
Inst:0.250 Ca(  6) Ce(  2)	github.com/dotcloud/docker/pkg/cgroups
Inst:0.250 Ca(  3) Ce(  1)	github.com/dotcloud/docker/pkg/sysinfo
Inst:0.250 Ca(  3) Ce(  1)	github.com/dotcloud/docker/runtime/networkdriver
Inst:0.154 Ca( 11) Ce(  2)	github.com/dotcloud/docker/runtime/graphdriver
Inst:0.125 Ca(  7) Ce(  1)	github.com/dotcloud/docker/pkg/label
Inst:0.091 Ca( 10) Ce(  1)	github.com/dotcloud/docker/nat
Inst:0.083 Ca( 11) Ce(  1)	github.com/dotcloud/docker/runtime/execdriver
Inst:0.077 Ca( 36) Ce(  3)	github.com/dotcloud/docker/utils
Inst:0.067 Ca( 14) Ce(  1)	github.com/dotcloud/docker/engine
Inst:0.056 Ca( 17) Ce(  1)	github.com/dotcloud/docker/pkg/libcontainer
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/collections
Inst:0.000 Ca(  1) Ce(  0)	github.com/gorilla/context
Inst:0.000 Ca(  1) Ce(  0)	code.google.com/p/go.net/websocket
Inst:0.000 Ca(  7) Ce(  0)	github.com/dotcloud/docker/dockerversion
Inst:0.000 Ca(  3) Ce(  0)	github.com/dotcloud/docker/pkg/mflag
Inst:0.000 Ca(  4) Ce(  0)	github.com/dotcloud/docker/pkg/mount
Inst:0.000 Ca(  1) Ce(  0)	github.com/dotcloud/docker/pkg/namesgenerator
Inst:0.000 Ca(  4) Ce(  0)	github.com/dotcloud/docker/pkg/netlink
Inst:0.000 Ca(  1) Ce(  0)	github.com/dotcloud/docker/pkg/proxy
Inst:0.000 Ca(  1) Ce(  0)	github.com/dotcloud/docker/pkg/listenbuffer
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/signal
Inst:0.000 Ca( 10) Ce(  0)	github.com/dotcloud/docker/pkg/system
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/systemd
Inst:0.000 Ca(  6) Ce(  0)	github.com/dotcloud/docker/pkg/term
Inst:0.000 Ca(  3) Ce(  0)	github.com/dotcloud/docker/pkg/user
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/version
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/libcontainer/utils
Inst:0.000 Ca(  4) Ce(  0)	github.com/dotcloud/docker/pkg/libcontainer/apparmor
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/iptables
Inst:0.000 Ca(  2) Ce(  0)	github.com/dotcloud/docker/pkg/graphdb
Inst:0.000 Ca(  5) Ce(  0)	github.com/dotcloud/docker/vendor/src/code.google.com/p/go/src/pkg/archive/tar

```
## License

MIT

## Author

hirokidaichi [at] gmail.com
