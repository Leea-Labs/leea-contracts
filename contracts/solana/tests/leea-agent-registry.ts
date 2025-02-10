import BN from "bn.js";
import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import * as web3 from "@solana/web3.js";
import type { LeeaAgentRegistry } from "../target/types/leea_agent_registry";
import { print_address, print_thread, print_tx, stream_program_logs } from "./utils";
import assert from "assert";

describe("leea-agent-registry", async () => {
    // Configure the client to use the local cluster.
    const provider = anchor.AnchorProvider.env();
    anchor.setProvider(provider);
    const program = anchor.workspace.LeeaTokenAico as Program<LeeaAgentRegistry>;

    print_address("🔗 Leea Agent Registry program", program.programId.toString());
});