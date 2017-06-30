{ pkgs ? (import <nixpkgs> {}) }:
with pkgs;
buildGoPackage rec {
  name = "kubernix-dev";
  goPackagePath = "github.com/moretea/kubernix";

  subPackages = ["cmd/kubernix"];
  src = ./.;

  installPhase = ''
    mkdir -p $out
    mkdir -p $bin/bin
    cp go/bin/* $bin/bin
  '';

  goDeps = null;
}
