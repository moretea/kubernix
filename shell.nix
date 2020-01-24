{ pkgs ? (import <nixpkgs> {}) }:
with pkgs;
with (import ./dep.nix { inherit pkgs; });
let
  crictl = buildGoPackage rec {
    name = "crictl";
    version = "1.17.0";

    src = fetchFromGitHub {
      owner = "kubernetes-sigs";
      repo = "cri-tools";
      rev = "v${version}";
      sha256= "0h9gry56graif761lmcy91q9fzwvmwb15wcx8245927yfg5j0zgh";
    };

    goPackagePath = "github.com/kubernetes-sigs/cri-tools";
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
    alias build='go build github.com/moretea/kubernix/cmd/kubernix'
  '';
}
