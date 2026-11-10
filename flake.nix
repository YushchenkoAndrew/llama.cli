
{
  description = "llama.cli - CLI interface for llama.cpp";
  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { self, nixpkgs, ... }: {
    packages.x86_64-linux.default =
      let
        pkgs = nixpkgs.legacyPackages.x86_64-linux;
      in
      pkgs.buildGoModule {
        pname = "llc";
        version = "master";
        src = self;
        vendorHash = null;
        postInstall = ''
          mv $out/bin/llama-cli $out/bin/llc
        '';
      };
  };
}

