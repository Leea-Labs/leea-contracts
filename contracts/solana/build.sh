#!bin/bash

cargo build --release --target x86_64-apple-darwin

cp target/x86_64-apple-darwin/release/escrow-client client-rs/escrow/bin/x86_64-apple-darwin/escrow-client
cp target/x86_64-apple-darwin/release/mint client-rs/mint/bin/x86_64-apple-darwin/mint
cp target/x86_64-apple-darwin/release/registry-client client-rs/registry/bin/x86_64-apple-darwin/registry-client

cargo build --release --target aarch64-apple-darwin

cp target/aarch64-apple-darwin/release/escrow-client client-rs/escrow/bin/aarch64-apple-darwin/escrow-client
cp target/aarch64-apple-darwin/release/mint client-rs/mint/bin/aarch64-apple-darwin/mint
cp target/aarch64-apple-darwin/release/registry-client client-rs/registry/bin/aarch64-apple-darwin/registry-client

cp target/idl/* client-rs/escrow/idls
cp target/idl/* client-rs/mint/idls
cp target/idl/* client-rs/registry/idls