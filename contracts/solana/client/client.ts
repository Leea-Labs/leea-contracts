import * as web3 from "@solana/web3.js";
import * as anchor from "@coral-xyz/anchor";
import { Metaplex } from "@metaplex-foundation/js";
import { getMint, getAssociatedTokenAddressSync } from "@solana/spl-token";
import type { AnchorToken } from "../target/types/anchor_token";

// Configure the client to use the local cluster
anchor.setProvider(anchor.AnchorProvider.env());

const program = anchor.workspace.AnchorToken as anchor.Program<AnchorToken>;


// metaplex token metadata program ID
const TOKEN_METADATA_PROGRAM_ID = new web3.PublicKey(
  "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"
);

// metaplex setup
const metaplex = Metaplex.make(program.provider.connection);

// token metadata
const metadata = {
  uri: "https://raw.githubusercontent.com/solana-developers/program-examples/new-examples/tokens/tokens/.assets/spl-token.json",
  name: "Solana Gold",
  symbol: "GOLDSOL",
};

// reward token mint PDA
const [rewardTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
  [Buffer.from("reward")],
  program.programId
);

// player data account PDA
const [playerPDA] = anchor.web3.PublicKey.findProgramAddressSync(
  [Buffer.from("player"), program.provider.publicKey.toBuffer()],
  program.programId
);

// reward token mint metadata account address
const rewardTokenMintMetadataPDA = await metaplex
  .nfts()
  .pdas()
  .metadata({ mint: rewardTokenMintPDA });

// player token account address
const playerTokenAccount = getAssociatedTokenAddressSync(
  rewardTokenMintPDA,
  program.provider.publicKey
);

async function logTransaction(txHash) {
  const { blockhash, lastValidBlockHeight } =
    await program.provider.connection.getLatestBlockhash();

  await program.provider.connection.confirmTransaction({
    blockhash,
    lastValidBlockHeight,
    signature: txHash,
  });

  console.log(`Use 'solana confirm -v ${txHash}' to see the logs`);
}

async function fetchAccountData() {
  const [playerBalance, playerData] = await Promise.all([
    program.provider.connection.getTokenAccountBalance(playerTokenAccount),
    program.account.playerData.fetch(playerPDA),
  ]);

  console.log("Player Token Balance: ", playerBalance.value.uiAmount);
  console.log("Player Health: ", playerData.health);
}

let txHash;

try {
  const mintData = await getMint(program.provider.connection, rewardTokenMintPDA);
  console.log("Mint Already Exists");
} catch {
  txHash = await program.methods
    .createMint(metadata.uri, metadata.name, metadata.symbol)
    .accounts({
      rewardTokenMint: rewardTokenMintPDA,
      metadataAccount: rewardTokenMintMetadataPDA,
      tokenMetadataProgram: TOKEN_METADATA_PROGRAM_ID,
    })
    .rpc();
  await logTransaction(txHash);
}
console.log("Token Mint: ", rewardTokenMintPDA.toString());

try {
  const playerData = await program.account.playerData.fetch(playerPDA);
  console.log("Player Already Exists");
  console.log("Player Health: ", playerData.health);
} catch {
  txHash = await program.methods
    .initPlayer()
    .accounts({
      playerData: playerPDA,
      player: program.provider.publicKey,
    })
    .rpc();
  await logTransaction(txHash);
  console.log("Player Account Created");
}

txHash = await program.methods
  .killEnemy()
  .accounts({
    playerData: playerPDA,
    playerTokenAccount: playerTokenAccount,
    rewardTokenMint: rewardTokenMintPDA,
  })
  .rpc();
await logTransaction(txHash);
console.log("Enemy Defeated");
await fetchAccountData();

txHash = await program.methods
  .heal()
  .accounts({
    playerData: playerPDA,
    playerTokenAccount: playerTokenAccount,
    rewardTokenMint: rewardTokenMintPDA,
  })
  .rpc();
await logTransaction(txHash);
console.log("Player Healed");
await fetchAccountData();