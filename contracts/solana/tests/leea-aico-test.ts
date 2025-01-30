import BN from "bn.js";
import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import * as web3 from "@solana/web3.js";
import { Metaplex } from "@metaplex-foundation/js";
import { getMint, getAssociatedTokenAddressSync } from "@solana/spl-token";
import type { LeeaTokenAico } from "../target/types/leea_token_aico";
import { print_address, print_thread, print_tx, stream_program_logs } from "./utils";

describe("leea-aico", async () => {
  // Configure the client to use the local cluster.
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);
  const program = anchor.workspace.LeeaTokenAico as Program<LeeaTokenAico>;

  print_address("🔗 Leea aiCO program", program.programId.toString());

  // metaplex token metadata program ID
  const TOKEN_METADATA_PROGRAM_ID = new web3.PublicKey(
    "metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s"
  );

  const ADMIN_PUB_KEY = new web3.PublicKey(
    "3M6Tp9boFAGWqM9FFeu6QTsgnsobe7UX79aryVtcAWpd"
  );
  // metaplex setup
  const metaplex = Metaplex.make(provider.connection);

  // token metadata
  const metadata = {
    uri: "https://raw.githubusercontent.com/solana-developers/program-examples/new-examples/tokens/tokens/.assets/spl-token.json",
    name: "Leea ai",
    symbol: "LEEA",
  };

  // reward token mint PDA
  const [rewardTokenMintPDA] = anchor.web3.PublicKey.findProgramAddressSync(
    [Buffer.from("aiCO_reward")],
    program.programId
  );

  // agent data account PDA
  const [agentPDA] = anchor.web3.PublicKey.findProgramAddressSync(
    [Buffer.from("leea_agent"), provider.wallet.publicKey.toBuffer()],
    program.programId
  );

  // reward token mint metadata account address
  const rewardTokenMintMetadataPDA = await metaplex
    .nfts()
    .pdas()
    .metadata({ mint: rewardTokenMintPDA });

  // agent token account address
  const agentTokenAccount = getAssociatedTokenAddressSync(
    rewardTokenMintPDA,
    provider.wallet.publicKey
  );

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

  async function fetchAccountData() {
    const [agentBalance, agentData] = await Promise.all([
      provider.connection.getTokenAccountBalance(agentTokenAccount),
      program.account.agentAccount.fetch(agentPDA),
    ]);

    console.log("Agent Token Balance: ", agentBalance.value.uiAmount);
    console.log("Agent Name: ", agentData.agentName);
  }

  let txHash;

  it("Mint Leea Token", async () => {
    try {
      const mintData = await getMint(provider.connection, rewardTokenMintPDA);
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
  });

  it("Register Agent", async () => {
    try {
      const agentData = await program.account.agentAccount.fetch(agentPDA);
      console.log("Agent Already Exists");
      console.log("Agent: ", agentData.holder.toString());
    } catch (e) {
      console.log("error", e);
      // probably not exist try to create
      try {
        txHash = await program.methods
          .initializeAgent("GPT4")
          .accounts({
            agentAccount: agentPDA,
          })
          .rpc();
        await logTransaction(txHash);
        console.log("Agent Account Created");
      } catch (e) {
        console.log("error", e);
      }
    }
  });

  it("Deposit to agent account for aiCO", async () => {
    try {
      txHash = await program.methods
        .deposit(new BN(10000000))
        .accounts({
          recipient: ADMIN_PUB_KEY,
          agentAccount: agentPDA,
        })
        .rpc();
      await logTransaction(txHash);
      const agentData = await program.account.agentAccount.fetch(agentPDA);
      console.log("Agent Balance: ", agentData.balance.toString());
    } catch (e) {
      console.log("error", e);
    }
  });

  it("Mint Leea tokens at aiCO time", async () => {
    try {
      txHash = await program.methods
        .aicoToAgent()
        .accounts({
          rewardTokenMint: rewardTokenMintPDA,
          agentTokenAccount: agentTokenAccount,
          agentAccount: agentPDA,
        })
        .rpc();
      await logTransaction(txHash);
      await fetchAccountData();
    } catch (e) {
      console.log("error", e);
    }
  });
});

