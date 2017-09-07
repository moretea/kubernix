{ pkgs ? (import <nixpkgs> {}) }:
with pkgs;
with (import ./dep.nix { inherit pkgs; });
let
  crictl = buildGoPackage {
    name = "crictl";

    src = fetchFromGitHub {
      owner = "kubernetes-incubator";
      repo = "cri-tools";
      rev = "e03736e429bcd0dba6f46cf6d6f7ccf0f5c70cc3";
      sha256= "00zy269k3y3q664gm6lxvr8v2ky8zdgkbxi9banx08irwxcrg22p";
    };

    goPackagePath = "github.com/kubernetes-incubator/cri-tools";
    subPackages = ["cmd/crictl"];
    deps = null;
  };


in stdenv.mkDerivation rec {
  name = "kubernix-dev";
  goPackagePath = "github.com/moretea/kubernix";
  buildInputs = [ go crictl jq dep ];

  CRI_RUNTIME_ENDPOINT = "/tmp/kubernix.sock";

  shellHook = ''
    projectGoPath=$(mktemp -d)
    mkdir -p $projectGoPath/src/$(dirname ${goPackagePath})
    ln -s $(pwd) $projectGoPath/src/${goPackagePath}
    export GOPATH=$projectGoPath:$GOPATH
  '';
}
