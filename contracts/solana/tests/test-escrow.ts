import * as anchor from "@coral-xyz/anchor";
import { Escrow } from "../target/types/escrow";
import { Keypair, LAMPORTS_PER_SOL, PublicKey, SystemProgram, Transaction, TransactionMessage, VersionedTransaction } from "@solana/web3.js";
import {
  ASSOCIATED_TOKEN_PROGRAM_ID,
  MINT_SIZE,
  TOKEN_PROGRAM_ID,
  createAssociatedTokenAccountIdempotentInstruction,
  createInitializeMint2Instruction,
  createMintToInstruction,
  getAssociatedTokenAddressSync,
  getMinimumBalanceForRentExemptMint
} from "@solana/spl-token";
import { randomBytes } from "crypto";
import * as web3 from "@solana/web3.js";

describe("escrow", () => {
  anchor.setProvider(anchor.AnchorProvider.env());
  const provider = anchor.getProvider();
  const connection = provider.connection;
  const program = anchor.workspace.Escrow as anchor.Program<Escrow>;
  const ADMIN_PUB_KEY = new web3.PublicKey(
    "GB9XNqUC32ZibLza8d7qMKBEv1hPZ142hzZ3sju7hG7b"
  );
  // Determine dummy token mints and token account addresses
  const [initializer, taker, mintA, mintB] = Array.from({ length: 4 }, () => Keypair.generate());
  const [initializerAtaA, initializerAtaB, takerAtaA, takerAtaB] = [initializer, taker]
    .map((a) => [mintA, mintB].map((m) => getAssociatedTokenAddressSync(m.publicKey, a.publicKey)))
    .flat();

  // Determined Escrow and Vault addresses
  const seed = new anchor.BN(randomBytes(8));
  const escrow = PublicKey.findProgramAddressSync(
    [Buffer.from("state"), seed.toArrayLike(Buffer, "le", 8)],
    program.programId
  )[0];
  const vault = getAssociatedTokenAddressSync(mintA.publicKey, escrow, true);

  // Account Wrapper
  const accounts = {
    admin: ADMIN_PUB_KEY,
    initializer: initializer.publicKey,
    taker: taker.publicKey,
    mintA: mintA.publicKey,
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

  it("Airdrop and create mints", async () => {
    let lamports = await getMinimumBalanceForRentExemptMint(connection);
    let tx = new Transaction();
    tx.instructions = [
      ...[initializer, taker].map((k) =>
        SystemProgram.transfer({
          fromPubkey: provider.publicKey,
          toPubkey: k.publicKey,
          lamports: 0.01 * LAMPORTS_PER_SOL,
        })
      ),
      ...[mintA, mintB].map((m) =>
        SystemProgram.createAccount({
          fromPubkey: provider.publicKey,
          newAccountPubkey: m.publicKey,
          lamports,
          space: MINT_SIZE,
          programId: TOKEN_PROGRAM_ID,
        })
      ),
      ...[
        [mintA.publicKey, initializer.publicKey, initializerAtaA],
        [mintB.publicKey, taker.publicKey, takerAtaB],
      ].flatMap((x) => [
        createInitializeMint2Instruction(x[0], 6, x[1], null),
        createAssociatedTokenAccountIdempotentInstruction(provider.publicKey, x[2], x[1], x[0]),
        createMintToInstruction(x[0], x[2], x[1], 1e9),
      ]),
    ];

    await provider.sendAndConfirm(tx, [mintA, mintB, initializer, taker]).then(log);
  });

  it("Initialize", async () => {
    let initializerBalance = await provider.connection.getTokenAccountBalance(
      initializerAtaA
    );
    console.log("Initializer TokenA Balance Before Escrow: ", initializerBalance.value.amount.toString());
    const initializerAmount = 1e6;
    const takerAmount = 1e6;
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
    console.log("Initializer TokenA Balance After Escrow: ", initializerBalance.value.amount.toString());
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
    console.log("Agent TokenA Balance After Payment: ", takerBalance.value.amount.toString());

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