use anchor_client::{
    solana_client::rpc_client::RpcClient,
    solana_sdk::{
        commitment_config::CommitmentConfig, pubkey::Pubkey, signature::read_keypair_file, signer::Signer, system_program,
    },
    Client, Cluster,
};
use anchor_lang::prelude::*;
use std::rc::Rc;
use std::{env, str::FromStr};

use anchor_spl::associated_token::get_associated_token_address;

declare_program!(escrow);

declare_program!(aico);

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let args: Vec<String> = env::args().collect();

    let user_key_path = &args[1];
    let amount: u64 = args[2].parse::<u64>().unwrap();
    let rpc_url = &args[3];

    let user = read_keypair_file(&*shellexpand::tilde(user_key_path))
        .expect("Escrow requires user keypair file");

    let cluster = Cluster::from_str(rpc_url)?;

    // Create program client
    let provider =
        Client::new_with_options(cluster, Rc::new(&user), CommitmentConfig::confirmed());
    let escrow_program = provider.program(escrow::ID)?;
    let aico_program = provider.program(aico::ID)?;

    let mint_seeds: &[&[u8]] = &[b"aiCO_reward"];
    let leea_token_mint_pda = Pubkey::find_program_address(mint_seeds, &aico_program.id()).0;
    let user_token_account = get_associated_token_address(&user.pubkey(), &leea_token_mint_pda);
    
    let escrow_seeds: &[&[u8]; 2] = &[b"state", &user.pubkey().to_bytes()];
    let escrow_pub = Pubkey::find_program_address(escrow_seeds, &escrow_program.id()).0;
    let vault = get_associated_token_address(&escrow_pub, &leea_token_mint_pda);

    //  Build and send instructions
    let init_ix = escrow_program
        .request()
        .accounts(escrow::client::accounts::Initialize {
            initializer: user.pubkey(),
            mint_a: leea_token_mint_pda,
            initializer_ata_a: user_token_account,
            escrow: escrow_pub,
            system_program: system_program::ID,
            associated_token_program: anchor_spl::associated_token::ID,
            token_program: anchor_spl::token::ID,
        })
        .args(escrow::client::args::Initialize {})
        .instructions()?
        .remove(0);

    let deposit_ix = escrow_program
        .request()
        .accounts(escrow::client::accounts::Deposit {
            initializer: user.pubkey(),
            mint_a: leea_token_mint_pda,
            initializer_ata_a: user_token_account,
            escrow: escrow_pub,
            vault: vault,
            system_program: system_program::ID,
            associated_token_program: anchor_spl::associated_token::ID,
            token_program: anchor_spl::token::ID,
        })
        .args(escrow::client::args::Deposit {
            amount: amount,
        })
        .instructions()?
        .remove(0);

    let signature = escrow_program
        .request()
        .instruction(init_ix)
        .instruction(deposit_ix)
        .signer(&user)
        .send()
        .await?;
    println!("   Transaction confirmed: {}", signature);

    let connection = RpcClient::new_with_commitment(
        rpc_url, // Local validator URL
        CommitmentConfig::confirmed(),
    );

    let vault_balance = connection.get_token_account_balance(&vault)?;
    println!("   Vault balance: {}", vault_balance.ui_amount_string.as_str());
    Ok(())
}
