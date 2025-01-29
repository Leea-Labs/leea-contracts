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
  name: "Leea ai",
  symbol: "LEEA",
};

// reward token mint PDA
const [rewardTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
  [Buffer.from("aiCO_reward")],
  pg.PROGRAM_ID
);

// agent data account PDA
const [agentPDA] = anchor.web3.PublicKey.findProgramAddressSync(
  [Buffer.from("leea_agent"), pg.wallet.publicKey.toBuffer()],
  pg.PROGRAM_ID
);

// reward token mint metadata account address
const rewardTokenMintMetadataPDA = await metaplex
  .nfts()
  .pdas()
  .metadata({ mint: rewardTokenMintPDA });

// agent token account address
const agentTokenAccount = getAssociatedTokenAddressSync(
  rewardTokenMintPDA,
  pg.wallet.publicKey
);

async function logTransaction(txHash) {
  const { blockhash, lastValidBlockHeight } =
    await pg.connection.getLatestBlockhash();

  await pg.connection.confirmTransaction({
    blockhash,
    lastValidBlockHeight,
    signature: txHash,
  });

  console.log(`Use 'solana confirm -v ${txHash}' to see the logs`);
}

async function fetchAccountData() {
  const [agentBalance, agentData] = await Promise.all([
    pg.connection.getTokenAccountBalance(agentTokenAccount),
    pg.program.account.agentAccount.fetch(agentPDA),
  ]);

  console.log("Agent Token Balance: ", agentBalance.value.uiAmount);
  console.log("Agent Name: ", agentData.agentName);
}

let txHash;

try {
  const mintData = await getMint(pg.connection, rewardTokenMintPDA);
  console.log("Mint Already Exists");
} catch {
  txHash = await pg.program.methods
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
  const agentData = await pg.program.account.agentAccount.fetch(agentPDA);
  console.log("Agent Already Exists");
  console.log("Agent Name: ", agentData.agentName);
} catch {
  txHash = await pg.program.methods
    .initializeAgent("GPT4", new BN(1))
    .accounts({
      agentAccount: agentPDA,
      admin: pg.wallet.publicKey,
    })
    .rpc();
  await logTransaction(txHash);
  console.log("Agent Account Created");
}
await fetchAccountData();
