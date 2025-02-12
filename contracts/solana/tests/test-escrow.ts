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
  createTransferInstruction
} from "@solana/spl-token";
import { randomBytes } from "crypto";
import * as web3 from "@solana/web3.js";
import NodeWallet from "@coral-xyz/anchor/dist/cjs/nodewallet";
import path from 'path'

describe("escrow", () => {
  const provider = anchor.AnchorProvider.env();
  const connection = provider.connection;
  const program = anchor.workspace.Escrow as anchor.Program<Escrow>;

  const fullPath = path.resolve(process.cwd(), './tests/id.json')
  const secret = require(fullPath)
  if (!secret) {
    throw new Error(`No secret found at ${fullPath}`)
  }
  const testKey = Keypair.fromSecretKey(new Uint8Array(secret))
  const testWallet = new NodeWallet(testKey);
  const ADMIN_PUB_KEY = new web3.PublicKey(
    "GB9XNqUC32ZibLza8d7qMKBEv1hPZ142hzZ3sju7hG7b"
  );
  // Determine dummy account addresses
  const [initializer, taker] = Array.from({ length: 2 }, () => Keypair.generate());
  // Leea Token program
  const LEEA_TOKEN_PROGRAM = new web3.PublicKey(
    "6ZfJgJYt9pXZNRd9S5busKXpBYmKaQJBWAC3R1GVD5se"
  );
  // token mint PDA
  const [leeaTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
    [Buffer.from("aiCO_reward")],
    LEEA_TOKEN_PROGRAM
  );
  const initializerAtaA = getAssociatedTokenAddressSync(leeaTokenMintPDA, initializer.publicKey)
  const takerAtaA = getAssociatedTokenAddressSync(leeaTokenMintPDA, taker.publicKey)
  // Determined Escrow and Vault addresses
  const seed = new anchor.BN(randomBytes(8));
  const escrow = PublicKey.findProgramAddressSync(
    [Buffer.from("state"), seed.toArrayLike(Buffer, "le", 8)],
    program.programId
  )[0];
  const vault = getAssociatedTokenAddressSync(leeaTokenMintPDA, escrow, true);
  // Account Wrapper
  const accounts = {
    admin: ADMIN_PUB_KEY,
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
      ...[testKey, initializer].map((k) =>
        SystemProgram.transfer({
          fromPubkey: provider.publicKey,
          toPubkey: k.publicKey,
          lamports: 10 * LAMPORTS_PER_SOL,
        })
      )];

    await provider.sendAndConfirm(tx).then(log);
  })

  it("Check Leea Mint", async () => {
    const leeaMint = await getMint(connection, leeaTokenMintPDA);
    console.log("Token Mint: ", leeaMint.mintAuthority.toString());

    const userTokenAccount = getAssociatedTokenAddressSync(
      leeaTokenMintPDA,
      testKey.publicKey
    );

    const userBalance = await connection.getTokenAccountBalance(
      userTokenAccount
    );
    console.log("Test Key Leea Token Balance: ", userBalance.value.amount.toString());
  });

  it("Transfer Leea Token", async () => {
    const fromTokenAccount = await getOrCreateAssociatedTokenAccount(
      connection,
      testKey,
      leeaTokenMintPDA,
      testKey.publicKey
    );

    // Get the derived address of the destination wallet which will hold the custom token
    const associatedDestinationTokenAddr = await getOrCreateAssociatedTokenAccount(
      connection,
      testKey,
      leeaTokenMintPDA,
      initializer.publicKey
    );

    const receiverAccount = await connection.getAccountInfo(associatedDestinationTokenAddr.address);
    const instructions: web3.TransactionInstruction[] = [];
    instructions.push(
      createTransferInstruction(
        fromTokenAccount.address,
        associatedDestinationTokenAddr.address,
        testKey.publicKey,
        100000,
        [],
        TOKEN_PROGRAM_ID
      )
    );
    const transaction = new web3.Transaction().add(...instructions);

    await provider.sendAndConfirm(transaction, [testKey]).then(log);

    const initializerTokenAccount = getAssociatedTokenAddressSync(
      leeaTokenMintPDA,
      initializer.publicKey
    );

    const userBalance = await connection.getTokenAccountBalance(
      initializerTokenAccount
    );
    console.log("Initializer Leea Token Balance: ", userBalance.value.amount.toString());
  })

  it("Initialize", async () => {
    let initializerBalance = await provider.connection.getTokenAccountBalance(
      initializerAtaA
    );
    console.log("Initializer Leea Balance Before Escrow: ", initializerBalance.value.amount.toString());
    const initializerAmount = 100000;
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
    console.log("Initializer Leea Balance After Escrow: ", initializerBalance.value.amount.toString());
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
      .accounts({ ...accounts })
      // .signers([])
      .rpc()
      .then(confirm)
      .then(log);

    const takerBalance = await provider.connection.getTokenAccountBalance(
      takerAtaA
    );
    console.log("Agent Leea Balance After Payment: ", takerBalance.value.amount.toString());
    const escrowBalance = await provider.connection.getTokenAccountBalance(
      vault
    );
    console.log("Escrow Leea Balance Left: ", escrowBalance.value.amount.toString());
    // Degugging

    // const latestBlockhash = await anchor
    //   .getProvider()
    //   .connection.getLatestBlockhash();

    // const ix = await program.methods
    //   .exchange()
    //   .accounts({ ...accounts })
    //   .signers([taker])
    //   .instruction()

    // const msg = new TransactionMessage({
    //   payerKey: provider.publicKey,
    //   recentBlockhash: latestBlockhash.blockhash,
    //   instructions: [ix],
    // }).compileToV0Message();

    // const tx = new VersionedTransaction(msg);
    // tx.sign([taker]);

    // console.log(Buffer.from(tx.serialize()).toString("base64"));
    // await provider.sendAndConfirm(tx).then(log);
  });
});