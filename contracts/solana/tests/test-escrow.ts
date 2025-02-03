import * as anchor from "@coral-xyz/anchor";
import * as splToken from "@solana/spl-token";
import { Program } from "@coral-xyz/anchor";
import { Escrow } from "../target/types/escrow";
import type { LeeaTokenAico } from "../target/types/leea_token_aico";
import { LAMPORTS_PER_SOL, SYSVAR_RENT_PUBKEY } from "@solana/web3.js";
import NodeWallet from "@coral-xyz/anchor/dist/cjs/nodewallet";
import * as web3 from "@solana/web3.js";

describe("Escrow", () => {
    const provider = anchor.AnchorProvider.env();
    anchor.setProvider(provider);
    const program = anchor.workspace.Escrow as Program<Escrow>;

    const user = provider.wallet.publicKey; // anchor.web3.Keypair.generate();
    const payer = (provider.wallet as NodeWallet).payer;
    const agent = anchor.web3.Keypair.generate();
    const escrowedXTokens = anchor.web3.Keypair.generate();
    let leea_mint;
    let user_leea_token;
    let agent_leea_token;
    let escrow: anchor.web3.PublicKey;
      const ADMIN_PUB_KEY = new web3.PublicKey(
        "GB9XNqUC32ZibLza8d7qMKBEv1hPZ142hzZ3sju7hG7b"
      );

    before(async () => {
        await provider.connection.requestAirdrop(agent.publicKey, 1 * LAMPORTS_PER_SOL);
        // Derive escrow address
        [escrow] = await anchor.web3.PublicKey.findProgramAddress([
            anchor.utils.bytes.utf8.encode("escrow"),
            user.toBuffer()
        ],
            program.programId)
        leea_mint = await splToken.Token.createMint(
            provider.connection,
            payer,
            provider.wallet.publicKey,
            provider.wallet.publicKey,
            6,
            splToken.TOKEN_PROGRAM_ID
        );
        user_leea_token = await leea_mint.createAccount(user);
        await leea_mint.mintTo(user_leea_token, payer, [], 10_000_000_000);
        agent_leea_token = await leea_mint.createAccount(agent.publicKey);
    })


    it("Initialize escrow", async () => {
        const x_amount = new anchor.BN(40);
        const tx = await program.methods.initialize(x_amount)
            .accounts({
                user: user,
                leeaMint: leea_mint.publicKey,
                userLeeaToken: user_leea_token,
                escrow: escrow,
                escrowedXTokens: escrowedXTokens.publicKey,
                tokenProgram: splToken.TOKEN_PROGRAM_ID,
                rent: SYSVAR_RENT_PUBKEY,
                systemProgram: anchor.web3.SystemProgram.programId
            })
            .signers([escrowedXTokens])
            .rpc()
    });


    it("Pay to agent", async () => {
        const x_amount = new anchor.BN(40);
        const tx = await program.methods.pay(x_amount)
            .accounts({
                agentLeeaTokens: user_leea_token,
                escrow: escrow,
                escrowedXTokens: escrowedXTokens.publicKey,
                tokenProgram: splToken.TOKEN_PROGRAM_ID
            })
            .signers([ADMIN_PUB_KEY])
            .rpc()
    });
});
