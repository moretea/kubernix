![Kubernix Logo](./logo/kubernix.png)

# Kubernix

Kubernix is a [Container Runtime Interface (CRI)](http://blog.kubernetes.io/2016/12/container-runtime-interface-cri-in-kubernetes.html) implementation that allows you to orchestrate native [NixOS containers](http://nixos.org/nixos/manual/#ch-containers) with Kubernetes, without depending on any other containerizer technology.

*_warning_*: It's just a proof of concept for now! Don't use it for anything serious

## Development
I'll nixify all of this some day. For now you'll have to run the following steps:

- Setup Go source code
  ```
  export GOPATH="SOMEDIR"
  mkdir -p $GOPATH/src/github.com/moretea/
  cd $GOPATH/src/github.com/moretea/
  git clone git@github.com:moretea/kubernix.git
  cd kubernix
  ```
- Enter a nix-shell (or use direnv)
- Install godep (note: `nix-env -i godep` is _not_ the same package)
  This should probably be packaged in nixpkgs.
  ```
    go get -u github.com/golang/dep/cmd/dep 
  ```
- Install dependencies
  `$GOPATH/bin/dep ensure`
- Building of kubernix
  ```
    cd cmd/kubernix
    go build
    rm -f kubernix.sock; ./kubernix -socket ./kubernix.sock
  ```
- Use the [cri-tools](https://github.com/kubernetes-incubator/cri-tools) that are made available in the nix-shell to play with kubernix.
- Alternatively, run one of the tests in the [tests](./tests) directory.

## Demo
Nope, it doesn't work yet!
