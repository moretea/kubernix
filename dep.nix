# Some minimal tooling to get github.com/golang/dep to work.
{ pkgs ? import <nixpkgs> {} }:
with pkgs;
rec {
  dep = buildGoPackage rec {
    name = "dep-unstable-${version}";
    version = "0.3.0";

    src = fetchFromGitHub {
      owner = "golang";
      repo = "dep";
      rev = "v${version}";
      sha256 = "1mkyc2z2zidh5h4vwwwc71cbyyi48c0n8gh2imjxbyy0g5i1vbgm";
    };

    goPackagePath = "github.com/golang/dep";
    subPackages = [ "cmd/dep" ];
    goDeps = null;
  };

  dep2vendor = lib.overrideDerivation dep (attrs: {
    name = "dep2vendor-unstable-${attrs.version}";

    patches = [./dep2vendor.patch];
    buildInputs = [ makeWrapper ];

    postInstall = ''
      wrapProgram $bin/bin/dep \
        --prefix PATH : ${lib.makeBinPath [git go]}
    '';
  });

  buildGoDepPackage = { name, Gopkg_toml, Gopkg_lock, depsSha256, ... }@args:
    let
      vendor = stdenv.mkDerivation {
        name = name + "-vendor";
        NIX_SSL_CERT_FILE = "${cacert}/etc/ssl/certs/ca-bundle.crt";
        buildCommand = ''
          export GOPATH=`pwd`
          export DUMMY_PATH=$GOPATH/src/dummy/
          mkdir -p $DUMMY_PATH
          cp ${Gopkg_toml} $DUMMY_PATH/Gopkg.toml
          cp ${Gopkg_lock} $DUMMY_PATH/Gopkg.lock
          cd $DUMMY_PATH
          ${dep2vendor}/bin/dep ensure -v -vendor-only
          mv vendor $out
        '';

        outputHashMode = "recursive";
        outputHashAlgo = "sha256";
        outputHash = depsSha256;
      };

      src = stdenv.mkDerivation {
        name = name + "-src";
        buildCommand = ''
          mkdir $out
          cp -r ${args.src}/* $out/
          ln -s ${vendor} $out/vendor || (echo "You probably have a 'vendor/' directory left in your source dir!"; exit 1)
        '';
      };
    in buildGoPackage (args // { goDeps = null; inherit src; });
}
