{
  description = "Go development environment";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs }:
  let
    system = "aarch64-linux";
    pkgs = nixpkgs.legacyPackages.${system};
  in {
    devShells.${system}.default = pkgs.mkShell {
      buildInputs = with pkgs; [
        go
        gopls
        gotools
        delve
      ];

      shellHook = ''
        export GOPATH=$(pwd)/.gopath
        export PATH=$GOPATH/bin:$PATH
        echo "Go development environment for aarch64-darwin ready!"
      '';
    };
  };
}
