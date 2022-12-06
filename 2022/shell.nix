{ pkgs ? import <nixpkgs> {} }:
  pkgs.mkShell {
    nativeBuildInputs = [ 
      pkgs.cargo 
      pkgs.rustc 
      pkgs.rustfmt
      pkgs.rust.packages.stable.rustPlatform.rustcSrc
      pkgs.rust.packages.stable.rustPlatform.rustLibSrc
      pkgs.python3 
      pkgs.entr
    ];
    RUST_SRC_PATH = "${pkgs.rust.packages.stable.rustPlatform.rustLibSrc}";
}