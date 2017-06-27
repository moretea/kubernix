{ pkgs ? (import <nixpkgs> {}) }:
with pkgs;
let 
  crictl = buildGoPackage {
    name = "crictl";

    src = fetchFromGitHub {
      owner = "kubernetes-incubator";
      repo = "cri-tools";
      rev = "e03736e429bcd0dba6f46cf6d6f7ccf0f5c70cc3";
      sha256= "00zy269k3y3q664gm6lxvr8v2ky8zdgkbxi9banx08irwxcrg22p";
    };

    subPackages = ["cmd/crictl"]; 

    goPackagePath = "github.com/kubernetes-incubator/cri-tools";
    deps = null;
  };
in
stdenv.mkDerivation {
  name = "kubernix";
  buildInputs = [ crictl go jq];
}
