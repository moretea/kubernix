![Kubernix Logo](./logo/kubernix.png)

# Kubernix

Kubernix is a [Container Runtime Interface (CRI)](http://blog.kubernetes.io/2016/12/container-runtime-interface-cri-in-kubernetes.html) implementation that allows you to orchestrate native [NixOS containers](http://nixos.org/nixos/manual/#ch-containers) with Kubernetes, without depending on any other containerizer technology.

*_warning_*: It's just a proof of concept for now! Don't use it for anything serious

## Development

### Preparation
1. Because we use `nix-shell` magic, there is no need to create a `$GOPATH/src/github.com/moretea` directory.
2. Clone the repository
  ```
    git clone git@github.com:moretea/kubernix.git
  ```
3. Enter a `nix-shell.`
4. Get the dependencies by running
  ```
    dep ensure
  ```
5.  I've made an easy `build` alias available in the `nix-shell`.

### Playing around
- Run the `kubernix` binary. It will open a UNIX socket on /tmp/kubernix.sock.
- Use the [cri-tools](https://github.com/kubernetes-incubator/cri-tools) that are made available in the nix-shell to play with kubernix.
  The `CRI_RUNTIME_ENDPOINT` environmental variable points to the `/tmp/kubernets.sock` UNIX socket,
  so the `cri-tools` can be used without passing any configuration paramaters.
- Alternatively, run one (or all) of the tests in the [tests](./tests) directory.

## Demo
Nope, it doesn't work yet!
