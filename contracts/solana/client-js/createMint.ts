import BN from "bn.js";
import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import * as web3 from "@solana/web3.js";
import type { Aico } from "../target/types/aico";
import { print_address } from "../tests/utils";
import { Keypair, SYSVAR_RENT_PUBKEY } from "@solana/web3.js";
import path from "path";
import { TOKEN_PROGRAM_ID } from "@solana/spl-token";
import { Connection } from "@solana/web3.js";
import NodeWallet from "@coral-xyz/anchor/dist/cjs/nodewallet";

const METAPLEX_METADATA_PROGRAM_ID = new web3.PublicKey(
  "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"
);

// Connect to solana
// 1. Admin key
let fullPath = path.resolve(__dirname, "../admin.json");
let secret = require(fullPath);
if (!secret) {
  throw new Error(`No secret found at ${fullPath}`);
}
const adminKey = Keypair.fromSecretKey(new Uint8Array(secret));
console.log(`Admin key: ${adminKey.publicKey.toString()}'`);

const wallet = new NodeWallet(Keypair.fromSecretKey(new Uint8Array(secret)));
const solanaConnection = new Connection(
  "https://api.devnet.solana.com",
  "confirmed"
);
const provider = new anchor.AnchorProvider(solanaConnection, wallet, {
  commitment: "processed",
});
anchor.setProvider(provider);
const program = anchor.workspace.Aico as Program<Aico>;
const connection = provider.connection;
print_address("🔗 Leea aiCO program", program.programId.toString());

const [leeaTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
  [Buffer.from("aiCO_reward")],
  program.programId
);

async function main() {
  await program.methods
    .createMint("Leea Test Token", "LEEA", "")
    .accounts({
      //@ts-ignore
      admin: adminKey.publicKey,
      leeaTokenMint: leeaTokenMintPDA,
      metadataAccount: web3.PublicKey.findProgramAddressSync(
        [
          Buffer.from("metadata"),
          METAPLEX_METADATA_PROGRAM_ID.toBuffer(),
          leeaTokenMintPDA.toBuffer(),
        ],
        METAPLEX_METADATA_PROGRAM_ID
      )[0],
      tokenProgram: TOKEN_PROGRAM_ID,
      tokenMetadataProgram: METAPLEX_METADATA_PROGRAM_ID,
      systemProgram: web3.SystemProgram.programId,
      rent: SYSVAR_RENT_PUBKEY,
    })
    .signers([adminKey])
    .rpc();
}

main();
