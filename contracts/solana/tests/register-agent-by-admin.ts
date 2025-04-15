import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import { Registry } from "../target/types/registry";
import { assert } from "chai";

describe("registry", () => {
  const provider = anchor.AnchorProvider.env();
  anchor.setProvider(provider);

  const program = anchor.workspace.Registry as Program<Registry>;
  const adminKeypair = anchor.web3.Keypair.fromSecretKey(
    Buffer.from(JSON.parse(process.env.ADMIN_PRIVATE_KEY || "[]"))
  );

  it("Registers an agent by admin", async () => {
    // Generate a new keypair for the agent holder
    const holderKeypair = anchor.web3.Keypair.generate();

    // Get the PDA for the agent account
    const [agentPda] = anchor.web3.PublicKey.findProgramAddressSync(
      [Buffer.from("leea_agent"), holderKeypair.publicKey.toBuffer()],
      program.programId
    );

    // Register the agent using admin
    const agentName = "Test Agent";
    const description = "Test Description";
    const fee = new anchor.BN(1000);

    await program.methods
      .registerAgentByAdmin(agentName, description, fee)
      .accounts({
        //@ts-ignore
        admin: adminKeypair.publicKey,
        holder: holderKeypair.publicKey,
        agentAccount: agentPda,
        systemProgram: anchor.web3.SystemProgram.programId,
      })
      .signers([adminKeypair])
      .rpc();

    // Fetch the created account
    const agentAccount = await program.account.agentAccount.fetch(agentPda);

    // Verify the account data
    assert.equal(
      agentAccount.holder.toString(),
      holderKeypair.publicKey.toString()
    );
    assert.equal(agentAccount.agentName, agentName);
    assert.equal(agentAccount.description, description);
    assert.equal(agentAccount.fee.toString(), fee.toString());
    assert.isAbove(agentAccount.createdAt.toNumber(), 0);
  });
});
