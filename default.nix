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

    goPackagePath = "github.com/kubernetes-incubator/cri-tools";
    subPackages = ["cmd/crictl"];
    deps = null;
  };

  dep = buildGoPackage rec {
    name = "dep-unstable-${version}";
    version = "2017-06-27";
    rev = "4bfa359b53746db53fcf09fc06044689c55e3949";

    src = fetchFromGitHub {
      owner = "golang";
      repo = "dep";
      inherit rev;
      sha256 = "1brsz4kw6gg8cy49641g8ky27s35xwa05vn1blnkhkl1973xwpir";
    };

    goPackagePath = "github.com/golang/dep";
    subPackages = [ "cmd/dep" ];
    goDeps = null;
  };
in buildGoPackage rec {
  name = "kubernix-dev";
  goPackagePath = "github.com/moretea/kubernix";
  buildInputs = [ crictl jq dep ];

  subPackages = ["cmd/kubernix"];

  src = ./.;
  installPhase = ''
    mkdir -p $out
    mkdir -p $bin/bin
    cp go/bin/* $bin/bin
  '';

  goDeps = null;
}
