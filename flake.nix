{
  description = "Go development shell";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixpkgs-unstable";  # use the unstable packages for more up-to-date packages
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
  flake-utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
    in {
      devShells.default = pkgs.mkShell {
        buildInputs = with pkgs; [
          delve
          go
          gopls
          gofumpt
          gotools
          gotests
        ];
        shellHook = ''
          export GOPATH=$(pwd)/.gopath
          export PATH=$GOPATH/bin:$PATH

          if [ -f go.mod ]; then
            echo "Found go.mod, running tidy"
            go mod tidy
          fi

          echo "Go development shell"
          echo "$(go version)"
          echo "GOPATH=$GOPATH"
        '';
      };
    }
  );
}
