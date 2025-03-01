use anchor_client::{
    solana_sdk::{
        commitment_config::CommitmentConfig, pubkey::Pubkey, signature::read_keypair_file,
        signature::Keypair, signer::Signer, system_program
    },
    Client, Cluster,
};
use anchor_lang::prelude::*;
use std::fs::File;
use std::io::prelude::*;
use std::path::Path;
use std::rc::Rc;
use std::{env, str::FromStr};

use anchor_spl::{
    associated_token::{AssociatedToken, get_associated_token_address},
    token::{
        close_account, transfer_checked, CloseAccount, Mint, Token, TokenAccount, TransferChecked,
    },
};

declare_program!(aico);
use aico::{accounts::AgentAccount, client::accounts, client::args};

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let args: Vec<String> = env::args().collect();

    let admin_key_path = &args[1];
    let receiver_pub_key = Pubkey::from_str(&args[2]).unwrap();
    let amount: u64 = args[3].parse::<u64>().unwrap();
    let rpc_url = &args[4];

    let admin = read_keypair_file(&*shellexpand::tilde(admin_key_path))
        .expect("Mint requires an admin keypair file");

    let cluster = Cluster::from_str(rpc_url)?;

    // Create program client
    let provider =
        Client::new_with_options(cluster, Rc::new(&admin), CommitmentConfig::confirmed());
    let program = provider.program(aico::ID)?;

    let seeds: &[&[u8]]  = &[b"aiCO_reward"];

    let leeaTokenMintPDA = Pubkey::find_program_address(seeds,  &program.id()).0;
    let receiver_token_account = get_associated_token_address(&receiver_pub_key, &leeaTokenMintPDA);
    //  Build and send instructions
    let mint_ix = program
        .request()
        .accounts(accounts::MintToReceiver {
            admin: admin.pubkey(),
            receiver: receiver_pub_key,
            receiver_token_account: receiver_token_account,
            leea_token_mint: leeaTokenMintPDA,
            system_program: system_program::ID,
            associated_token_program: anchor_spl::associated_token::ID,
            token_program: anchor_spl::token::ID
        })
        .args(args::MintToReceiver {
            amount: amount,
        })
        .instructions()?
        .remove(0);

    let signature = program
        .request()
        .instruction(mint_ix)
        .signer(&admin)
        .send()
        .await?;
    println!("   Transaction confirmed: {}", signature);

    Ok(())
}
