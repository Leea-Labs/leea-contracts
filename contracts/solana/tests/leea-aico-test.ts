import BN from "bn.js";
import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import * as web3 from "@solana/web3.js";
import { Metaplex } from "@metaplex-foundation/js";
import { getMint, getAssociatedTokenAddressSync } from "@solana/spl-token";
import type { LeeaTokenAico } from "../target/types/leea_token_aico";
import { print_address, print_thread, print_tx, stream_program_logs } from "./utils";
import assert from "assert";
import { Keypair, SystemProgram, Transaction, LAMPORTS_PER_SOL } from "@solana/web3.js";
import path from 'path'

describe("leea-aico", async () => {
  // Configure the client to use the local cluster.
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);
  const program = anchor.workspace.LeeaTokenAico as Program<LeeaTokenAico>;
  const connection = provider.connection;
  print_address("🔗 Leea aiCO program", program.programId.toString());

  // Create key pairs
  // 1. Admin key
  let fullPath = path.resolve(process.cwd(), './tests/admin.json')
  let secret = require(path.resolve(process.cwd(), './tests/admin.json'))
  if (!secret) {
    throw new Error(`No secret found at ${fullPath}`)
  }
  const adminKey = Keypair.fromSecretKey(new Uint8Array(secret))
  console.log(`Admin key: ${adminKey.publicKey.toString()}'`); // GB9XNqUC32ZibLza8d7qMKBEv1hPZ142hzZ3sju7hG7b
  // 2. Funds key
  fullPath = path.resolve(process.cwd(), './tests/fund.json')
  secret = require(fullPath)
  if (!secret) {
    throw new Error(`No secret found at ${fullPath}`)
  }
  const fundsKey = Keypair.fromSecretKey(new Uint8Array(secret))
  console.log(`Fund key: ${fundsKey.publicKey.toString()}'`); // EFqAYmaZPwHdsBc8PCZYo5DywDVMrCHU79K6Fq1MrbQU
  // 3. Agent key pairs
  fullPath = path.resolve(process.cwd(), './tests/agent.json')
  secret = require(fullPath)
  if (!secret) {
    throw new Error(`No secret found at ${fullPath}`)
  }
  const agent1 = Keypair.fromSecretKey(new Uint8Array(secret))
  console.log(`Agent key: ${agent1.publicKey.toString()}'`); // 55rm9zK2YZumXeXH6XSXXjgtkSzL4MDeh9K8XFmWrs28
 
  // const [agent1, agent2] = Array.from({ length: 2 }, () => Keypair.generate());

  // Get required PDAs
  // 1. Reward token mint PDA
  const [leeaTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
    [Buffer.from("aiCO_reward")],
    program.programId
  );
  // 2. Agent data account PDA
  const [agentPDA] = anchor.web3.PublicKey.findProgramAddressSync(
    [Buffer.from("leea_agent"), agent1.publicKey.toBuffer()],
    program.programId
  );
  // 3. Agent token account address
  const agentTokenAccount = getAssociatedTokenAddressSync(
    leeaTokenMintPDA,
    agent1.publicKey
  );

  // Metaplex data for token
  const TOKEN_METADATA_PROGRAM_ID = new web3.PublicKey(
    "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"
  );
  // 1. Metaplex setup
  const metaplex = Metaplex.make(provider.connection);
  // 2. Token metadata
  const metadata = {
    uri: "https://raw.githubusercontent.com/solana-developers/program-examples/new-examples/tokens/tokens/.assets/spl-token.json",
    name: "Leea ai",
    symbol: "LEEA",
  };
  // 3. Reward token mint metadata account address
  const rewardTokenMintMetadataPDA = await metaplex
    .nfts()
    .pdas()
    .metadata({ mint: leeaTokenMintPDA });

  // Utility functions
  async function logTransaction(txHash) {
    const { blockhash, lastValidBlockHeight } =
      await provider.connection.getLatestBlockhash();

    await provider.connection.confirmTransaction({
      blockhash,
      lastValidBlockHeight,
      signature: txHash,
    });

    console.log(`Use 'solana confirm -v ${txHash}' to see the logs`);
  }

  let txHash;

  const log = async (signature: string): Promise<string> => {
    console.log(
      `Your transaction signature: https://explorer.solana.com/transaction/${signature}?cluster=custom&customUrl=${connection.rpcEndpoint}`
    );
    return signature;
  };

  it("Add lamports to keypairs", async () => {
    let tx = new Transaction();
    tx.instructions = [
      ...[adminKey, agent1].map((k) =>
        SystemProgram.transfer({
          fromPubkey: provider.publicKey,
          toPubkey: k.publicKey,
          lamports: 10 * LAMPORTS_PER_SOL,
        })
      )];
    await provider.sendAndConfirm(tx).then(log);
  })

  it("Mint Leea SPL Token", async () => {
    try {
      const mintData = await getMint(provider.connection, leeaTokenMintPDA);
      console.log("Token Already Minted: ", mintData.address.toString());
    } catch (e) {
      txHash = await program.methods
        .createMint(metadata.uri, metadata.name, metadata.symbol)
        .accounts({
          leeaTokenMint: leeaTokenMintPDA,
          metadataAccount: rewardTokenMintMetadataPDA,
          tokenMetadataProgram: TOKEN_METADATA_PROGRAM_ID,
        })
        .signers([adminKey])
        .rpc();
      await logTransaction(txHash);
      console.log("Token Minted: ", leeaTokenMintPDA.toString());
      const mintData = await getMint(provider.connection, leeaTokenMintPDA);
      // Assertions
      assert.equal(mintData.mintAuthority.toString(), leeaTokenMintPDA.toString());
    }
  });

  it("Register Agent1 at aiCO program", async () => {
    try {
      const agentData = await program.account.agentAccount.fetch(agentPDA);
      console.log("Agent Already Exists");
      console.log("Agent: ", agentData.holder.toString());
      assert.equal(agentData.holder.toString(), agent1.publicKey.toString())
    } catch (e) {
      // probably not exist try to create
      txHash = await program.methods
        .initializeAgent("GPT4")
        .accounts({
          holder: agent1.publicKey,
          agentAccount: agentPDA,
        })
        .signers([agent1])
        .rpc();
      await logTransaction(txHash);
      console.log("Agent Account Created");
      const agentData = await program.account.agentAccount.fetch(agentPDA);
      console.log("Agent: ", agentData.holder.toString());
      assert.equal(agentData.agentName, "GPT4")
      assert.equal(agentData.holder.toString(), agent1.publicKey.toString())
    }
  });

  it("Deposit to agent account for aiCO", async () => {
    txHash = await program.methods
      .deposit(new BN(10000000))
      .accounts({
        recipient: fundsKey.publicKey,
        agentAccount: agentPDA,
        holder: agent1.publicKey
      })
      .signers([agent1])
      .rpc();
    await logTransaction(txHash);
    const agentData = await program.account.agentAccount.fetch(agentPDA);
    console.log("Agent Balance: ", agentData.balance.toString());
    assert.equal(agentData.agentName, "GPT4")
    assert.equal(agentData.holder.toString(), agent1.publicKey.toString())
  });

  it("Claim Leea tokens at aiCO time", async () => {
    txHash = await program.methods
      .aicoToAgent()
      .accounts({
        holder: agent1.publicKey,
        leeaTokenMint: leeaTokenMintPDA,
        agentTokenAccount: agentTokenAccount,
        agentAccount: agentPDA,
      })
      .signers([agent1])
      .rpc();
    await logTransaction(txHash);
    const [agentBalance, agentData] = await Promise.all([
      provider.connection.getTokenAccountBalance(agentTokenAccount),
      program.account.agentAccount.fetch(agentPDA),
    ]);
    console.log("Agent Token Balance: ", agentBalance.value.uiAmount);
    console.log("Agent Name: ", agentData.agentName);
    assert.equal(agentData.agentName, "GPT4");
    assert.equal(agentData.holder.toString(), agent1.publicKey.toString())
    assert.notEqual(agentBalance.value.uiAmount, 0);
    assert.equal(agentData.balance, 0);
    const fundBalance = await provider.connection.getBalance(fundsKey.publicKey);
    console.log(`Leea fund wallet balance ${fundBalance}`)
  });
});