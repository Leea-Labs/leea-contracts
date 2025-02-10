import * as anchor from "@coral-xyz/anchor";
import { Escrow } from "../target/types/escrow";
import { Keypair, LAMPORTS_PER_SOL, PublicKey, SystemProgram, Transaction } from "@solana/web3.js";
import {
  ASSOCIATED_TOKEN_PROGRAM_ID,
  MINT_SIZE,
  TOKEN_PROGRAM_ID,
  createAssociatedTokenAccountIdempotentInstruction,
  createInitializeMint2Instruction,
  createMintToInstruction,
  getAssociatedTokenAddressSync,
  getMinimumBalanceForRentExemptMint,
  getMint,
  getOrCreateAssociatedTokenAccount,
  createTransferInstruction,
  mintToChecked
} from "@solana/spl-token";
import { randomBytes } from "crypto";
import * as web3 from "@solana/web3.js";
import path from 'path'
import assert from "assert";
import type { LeeaTokenAico } from "../target/types/leea_token_aico";

describe("escrow", () => {
  const provider = anchor.AnchorProvider.env();
  const connection = provider.connection;
  const program = anchor.workspace.Escrow as anchor.Program<Escrow>;
  const leeaAiCOprogram = anchor.workspace.LeeaTokenAico as anchor.Program<LeeaTokenAico>;

  // Create key pairs #######################################
  // 1. Admin key
  let fullPath = path.resolve(process.cwd(), './tests/admin.json')
  let secret = require(path.resolve(process.cwd(), './tests/admin.json'))
  if (!secret) {
    throw new Error(`No secret found at ${fullPath}`)
  }
  const adminKey = Keypair.fromSecretKey(new Uint8Array(secret))
  console.log(`Admin key: ${adminKey.publicKey.toString()}'`); // GB9XNqUC32ZibLza8d7qMKBEv1hPZ142hzZ3sju7hG7b
  // 2. Agent key pairs
  fullPath = path.resolve(process.cwd(), './tests/agent.json')
  secret = require(fullPath)
  if (!secret) {
    throw new Error(`No secret found at ${fullPath}`)
  }
  const agent1 = Keypair.fromSecretKey(new Uint8Array(secret))
  console.log(`Agent key: ${agent1.publicKey.toString()}'`); // 55rm9zK2YZumXeXH6XSXXjgtkSzL4MDeh9K8XFmWrs28

  // 2. Escrow initializer (user) and taker (agent)
  const [initializer, taker] = Array.from({ length: 2 }, () => Keypair.generate());

  // Get required PDAs ######################################
  // 1. Token mint PDA
  const [leeaTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
    [Buffer.from("aiCO_reward")],
    leeaAiCOprogram.programId
  );
  // 2. Initializer Token account
  const initializerAtaA = getAssociatedTokenAddressSync(leeaTokenMintPDA, initializer.publicKey)
  // 3. Taker(agent) Token account
  const takerAtaA = getAssociatedTokenAddressSync(leeaTokenMintPDA, taker.publicKey)
  // 4. Admin Token account
  const adminTokenAccount = getAssociatedTokenAddressSync(
    leeaTokenMintPDA,
    adminKey.publicKey
  );
  // 5. Determined Escrow and Vault addresses
  const seed = new anchor.BN(randomBytes(8));
  const escrow = PublicKey.findProgramAddressSync(
    [Buffer.from("state"), seed.toArrayLike(Buffer, "le", 8)],
    program.programId
  )[0];
  const vault = getAssociatedTokenAddressSync(leeaTokenMintPDA, escrow, true);

  // Account Wrapper #######################################
  const accounts = {
    admin: adminKey.publicKey,
    initializer: initializer.publicKey,
    taker: taker.publicKey,
    mintA: leeaTokenMintPDA,
    initializerAtaA: initializerAtaA,
    takerAtaA,
    escrow,
    vault,
    associatedTokenprogram: ASSOCIATED_TOKEN_PROGRAM_ID,
    tokenProgram: TOKEN_PROGRAM_ID,
    systemProgram: SystemProgram.programId,
  };

  // Utility functions #########################################
  const confirm = async (signature: string): Promise<string> => {
    const block = await connection.getLatestBlockhash();
    await connection.confirmTransaction({
      signature,
      ...block,
    });
    return signature;
  };

  const log = async (signature: string): Promise<string> => {
    console.log(
      `Your transaction signature: https://explorer.solana.com/transaction/${signature}?cluster=custom&customUrl=${connection.rpcEndpoint}`
    );
    return signature;
  };
  it("Top up test wallet", async () => {
    let tx = new Transaction();
    tx.instructions = [
      ...[adminKey, initializer, taker].map((k) =>
        SystemProgram.transfer({
          fromPubkey: provider.publicKey,
          toPubkey: k.publicKey,
          lamports: 10 * LAMPORTS_PER_SOL,
        })
      )];
    await provider.sendAndConfirm(tx).then(log);
  })
  //########################################################

  it("Check Leea Mint", async () => {
    const leeaMint = await getMint(connection, leeaTokenMintPDA);
    assert.equal(leeaMint.address, leeaTokenMintPDA)
    console.log("Token Mint: ", leeaMint.mintAuthority.toString());
  });

  // it("Mint some Leea tokens to admin", async () => {
  //   let txhash = await mintToChecked(
  //     connection, // connection
  //     adminKey, // fee payer
  //     leeaTokenMintPDA, // mint
  //     initializerAtaA, // receiver (should be a token account)
  //     leeaTokenMintPDA, // mint authority
  //     1e8, // amount. if your decimals are 8, you mint 10^8 for 1 token.
  //     8, // decimals
  //   );
  //   console.log(`txhash: ${txhash}`);
  //   const adminTokenBalance = await connection.getTokenAccountBalance(
  //     initializerAtaA
  //   );
  //   console.log("Initializer Leea Token Balance: ", adminTokenBalance.value.amount.toString());
  // })

  it("Transfer some Leea tokens to initializer (user)", async () => {
    const agentTokenAccount = getAssociatedTokenAddressSync(
      leeaTokenMintPDA,
      agent1.publicKey
    );
    const agentTokenBalance = await connection.getTokenAccountBalance(
      agentTokenAccount
    );
    console.log("Agent Leea token balance: ", agentTokenBalance.value.amount.toString());
    const associatedDestinationTokenAddr = await getOrCreateAssociatedTokenAccount(
      connection,
      initializer,
      leeaTokenMintPDA,
      initializer.publicKey
    );
    const instructions: web3.TransactionInstruction[] = [];
    instructions.push(
      createTransferInstruction(
        agentTokenAccount,
        associatedDestinationTokenAddr.address,
        agent1.publicKey,
        100000,
        [],
        TOKEN_PROGRAM_ID
      )
    );
    const transaction = new web3.Transaction().add(...instructions);
    await provider.sendAndConfirm(transaction, [agent1]).then(log);
    const initializerBalance = await connection.getTokenAccountBalance(
      initializerAtaA
    );
    console.log("Initializer Leea token balance: ", initializerBalance.value.amount.toString());
    assert.equal(initializerBalance.value.amount, 100000)
  })

  it("Initialize", async () => {
    let initializerBalance = await provider.connection.getTokenAccountBalance(
      initializerAtaA
    );
    console.log("Initializer token balance before escrowing: ", initializerBalance.value.amount.toString());
    const initializerAmount = 1e5;
    await program.methods
      .initialize(seed, new anchor.BN(initializerAmount))
      .accounts({ ...accounts })
      .signers([initializer])
      .rpc()
      .then(confirm)
      .then(log);
    initializerBalance = await provider.connection.getTokenAccountBalance(
      initializerAtaA
    );
    console.log("Initializer token balance after escrowing: ", initializerBalance.value.amount.toString());
    let vaultBalance = await provider.connection.getTokenAccountBalance(
      vault
    );
    console.log("Vault token balance: ", vaultBalance.value.amount.toString());
    assert.equal(vaultBalance.value.amount, initializerAmount)
  });

  xit("Cancel", async () => {
    await program.methods
      .cancel()
      .accounts({ ...accounts })
      .signers([initializer])
      .rpc()
      .then(confirm)
      .then(log);
  });

  it("Pay to agent", async () => {
    await program.methods
      .payToAgent(new anchor.BN(1e3))
      .accounts({
        admin: adminKey.publicKey,
        initializer: initializer.publicKey,
        taker: taker.publicKey,
        mintA: leeaTokenMintPDA,
        takerAtaA,
        escrow,
        vault,
        associatedTokenprogram: ASSOCIATED_TOKEN_PROGRAM_ID,
        tokenProgram: TOKEN_PROGRAM_ID,
        systemProgram: SystemProgram.programId,
      })
      .signers([adminKey])
      .rpc()
      .then(confirm)
      .then(log);

    const takerBalance = await provider.connection.getTokenAccountBalance(
      takerAtaA
    );
    console.log("Agent token balance after work is done: ", takerBalance.value.amount.toString());
    const vaultBalance = await provider.connection.getTokenAccountBalance(
      vault
    );
    console.log("Escrow Leea Balance Left: ", vaultBalance.value.amount.toString());
    assert.equal(vaultBalance.value.amount, 1e5 - 1e3)
  });
});