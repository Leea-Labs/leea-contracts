import * as anchor from "@coral-xyz/anchor";
import { Escrow } from "../target/types/escrow";
import {
  Keypair,
  LAMPORTS_PER_SOL,
  PublicKey,
  SystemProgram,
} from "@solana/web3.js";
import {
  ASSOCIATED_TOKEN_PROGRAM_ID,
  TOKEN_PROGRAM_ID,
  getAssociatedTokenAddressSync,
} from "@solana/spl-token";
import { randomBytes } from "crypto";
import * as web3 from "@solana/web3.js";
import path from "path";
import assert from "assert";
import { log, confirm, print_address } from "../tests/utils";
import NodeWallet from "@coral-xyz/anchor/dist/cjs/nodewallet";
import type { Aico } from "../target/types/aico";
import bs58 from "bs58";

// Connect to solana
// 1. Admin key
let fullPath = path.resolve(process.cwd(), "./admin.json");
let secret = require(fullPath);
if (!secret) {
  throw new Error(`No secret found at ${fullPath}`);
}
const adminKey = Keypair.fromSecretKey(new Uint8Array(secret));
console.log(`Admin key: ${adminKey.publicKey.toString()}'`);

const wallet = new NodeWallet(Keypair.fromSecretKey(new Uint8Array(secret)));
const solanaConnection = new web3.Connection(
  "https://api.devnet.solana.com",
  "confirmed"
);
const provider = new anchor.AnchorProvider(solanaConnection, wallet, {
  commitment: "processed",
});
anchor.setProvider(provider);
const program = anchor.workspace.Escrow as anchor.Program<Escrow>;
const connection = provider.connection;
const leeaAiCOprogram = anchor.workspace.Aico as anchor.Program<Aico>;
print_address("🔗 Leea Escrow program", program.programId.toString());

const [leeaTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
  [Buffer.from("aiCO_reward")],
  leeaAiCOprogram.programId
);

// Helper to get all PDAs and ATAs for a user
function getEscrowAddressesForUser(userPubkey: PublicKey, mint: PublicKey) {
  const escrow = PublicKey.findProgramAddressSync(
    [Buffer.from("state"), userPubkey.toBuffer()],
    program.programId
  )[0];
  const vault = getAssociatedTokenAddressSync(mint, escrow, true);
  const userAta = getAssociatedTokenAddressSync(mint, userPubkey);
  return { escrow, vault, userAta };
}

// Test: Initialize escrow for a user (admin as signer)
export async function testInitializeEscrowForUser(userPubkey: PublicKey) {
  const { escrow } = getEscrowAddressesForUser(userPubkey, leeaTokenMintPDA);
  const accounts = {
    admin: adminKey.publicKey,
    initializer: userPubkey,
    mintA: leeaTokenMintPDA,
    escrow,
    associatedTokenProgram: ASSOCIATED_TOKEN_PROGRAM_ID,
    tokenProgram: TOKEN_PROGRAM_ID,
    systemProgram: SystemProgram.programId,
  };
  await program.methods
    .initialize()
    .accounts(accounts)
    .signers([adminKey])
    .rpc()
    .then((t) => confirm(t, connection))
    .then((t) => log(t, connection));
  console.log(`Initialized escrow for user: ${userPubkey.toBase58()}`);
}

// Test: Deposit tokens into escrow for a user (admin or user as signer)
export async function testDepositForUser(
  userPubkey: PublicKey,
  authority: Keypair, // can be admin or user
  amount: number
) {
  const { escrow, vault } = getEscrowAddressesForUser(
    userPubkey,
    leeaTokenMintPDA
  );
  const authorityAta = getAssociatedTokenAddressSync(
    leeaTokenMintPDA,
    authority.publicKey
  );
  const accounts = {
    authority: authority.publicKey,
    initializer: userPubkey,
    mintA: leeaTokenMintPDA,
    fromAta: authorityAta,
    escrow,
    vault,
    associatedTokenProgram: ASSOCIATED_TOKEN_PROGRAM_ID,
    tokenProgram: TOKEN_PROGRAM_ID,
    systemProgram: SystemProgram.programId,
  };
  // Check authority's token balance before
  let authorityBalance =
    await provider.connection.getTokenAccountBalance(authorityAta);
  console.log(
    `Authority (${authority.publicKey.toBase58()}) token balance before deposit: `,
    authorityBalance.value.amount.toString()
  );
  await program.methods
    .deposit(new anchor.BN(amount))
    .accounts(accounts)
    .signers([authority])
    .rpc()
    .then((t) => confirm(t, connection))
    .then((t) => log(t, connection));
  // Check balances after
  authorityBalance =
    await provider.connection.getTokenAccountBalance(authorityAta);
  let vaultBalance = await provider.connection.getTokenAccountBalance(vault);
  console.log(
    `Authority (${authority.publicKey.toBase58()}) token balance after deposit: `,
    authorityBalance.value.amount.toString()
  );
  console.log("Vault token balance: ", vaultBalance.value.amount.toString());
}

// Example usage for testing
(async () => {
  // Replace with your user public key
  const userPubkey = new web3.PublicKey("");
  // Replace with your user secret if you want to test user deposit
  // const userSecret = ...
  // const userKeypair = Keypair.fromSecretKey(new Uint8Array(userSecret));

  // Admin initializes escrow for user
  await testInitializeEscrowForUser(userPubkey);
  // Admin deposits for user
  await testDepositForUser(userPubkey, adminKey, 1 * LAMPORTS_PER_SOL);
  // Uncomment below to test user deposit (user must have tokens)
  // await testDepositForUser(userPubkey, userKeypair, 1 * LAMPORTS_PER_SOL);
})();
