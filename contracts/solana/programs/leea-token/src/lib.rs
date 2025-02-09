use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::AssociatedToken,
    metadata::{create_metadata_accounts_v3, CreateMetadataAccountsV3, Metadata},
    token::{burn, mint_to, Burn, Mint, MintTo, Token, TokenAccount},
};
use mpl_token_metadata::types::DataV2;

use solana_program::system_instruction;
use solana_program::{pubkey, pubkey::Pubkey};

declare_id!("CpnJzw1VnfaKYrbHFcLgSEYKJXDhorgQAHKmueLNZxNq");

const ADMIN_PUBKEY: Pubkey = pubkey!("3M6Tp9boFAGWqM9FFeu6QTsgnsobe7UX79aryVtcAWpd");

pub const AGENT_SEED: &[u8] = b"leea_agent";
pub const AICO_SEED: &[u8] = b"aiCO_reward";
pub const UNLOCK_TIME: i64 = 1738153798;

pub const PREFIX: &str = "metadata";

fn find_metadata_account(mint: &Pubkey) -> (Pubkey, u8) {
    Pubkey::find_program_address(
        &[
            PREFIX.as_bytes(),
            mpl_token_metadata::ID.as_ref(),
            mint.as_ref(),
        ],
        &mpl_token_metadata::ID,
    )
}

#[program]
pub mod leea_token_aico {
    use super::*;

    // Create new token mint with PDA as mint authority
    pub fn create_mint(
        ctx: Context<CreateMint>,
        uri: String,
        name: String,
        symbol: String,
    ) -> Result<()> {
        // PDA seeds and bump to "sign" for CPI
        let seeds = AICO_SEED;
        let bump = ctx.bumps.reward_token_mint;
        let signer: &[&[&[u8]]] = &[&[seeds, &[bump]]];

        // On-chain token metadata for the mint
        let data_v2 = DataV2 {
            name: name,
            symbol: symbol,
            uri: uri,
            seller_fee_basis_points: 0,
            creators: None,
            collection: None,
            uses: None,
        };

        // CPI Context
        let cpi_ctx = CpiContext::new_with_signer(
            ctx.accounts.token_metadata_program.to_account_info(),
            CreateMetadataAccountsV3 {
                // the metadata account being created
                metadata: ctx.accounts.metadata_account.to_account_info(),
                // the mint account of the metadata account
                mint: ctx.accounts.reward_token_mint.to_account_info(),
                // the mint authority of the mint account
                mint_authority: ctx.accounts.reward_token_mint.to_account_info(),
                // the update authority of the metadata account
                update_authority: ctx.accounts.reward_token_mint.to_account_info(),
                // the payer for creating the metadata account
                payer: ctx.accounts.admin.to_account_info(),
                // the system program account
                system_program: ctx.accounts.system_program.to_account_info(),
                // the rent sysvar account
                rent: ctx.accounts.rent.to_account_info(),
            },
            signer,
        );

        create_metadata_accounts_v3(
            cpi_ctx, // cpi context
            data_v2, // token metadata
            true,    // is_mutable
            true,    // update_authority_is_signer
            None,    // collection details
        )?;

        Ok(())
    }

    pub fn initialize_agent(
        ctx: Context<InitializeAgent>,
        agent_name: String,
        amount: u64,
    ) -> Result<()> {
        if amount == 0 {
            return Err(error!(ErrorCode::AmountTooSmall));
        };
        let system_program = &ctx.accounts.system_program;
        let holder = &ctx.accounts.holder;
        let agent_account = &mut ctx.accounts.agent_account;
        // Create the transfer instruction
        let transfer_instruction =
            system_instruction::transfer(holder.key, ctx.accounts.admin.signer_key().as_ref().unwrap(), amount);
        // Invoke the transfer instruction
        anchor_lang::solana_program::program::invoke_signed(
            &transfer_instruction,
            &[
                holder.to_account_info(),
                ctx.accounts.admin.to_account_info(),
                ctx.accounts.system_program.to_account_info(),
            ],
            &[],
        )?;
        agent_account.holder = *holder.key;
        agent_account.balance += amount;
        agent_account.agent_name = agent_name;
        agent_account.created_at = Clock::get().unwrap().unix_timestamp;
        Ok(())
    }

    pub fn deposit(ctx: Context<UpdateBalance>, amount: u64) -> Result<()> {
        if amount == 0 {
            return Err(error!(ErrorCode::AmountTooSmall));
        };

        let agent_account = &mut ctx.accounts.agent_account;
        agent_account.balance += amount;
        Ok(())
    }

    pub fn aico_to_agent(ctx: Context<InitializeAgent>) -> Result<()> {
        // Check if agent deposited
        if ctx.accounts.agent_account.balance == 0 {
            return err!(ErrorCode::NotEnoughDeposit);
        }
        // Subtract SOL from agent // TODO
        ctx.accounts.agent_account.balance =
            ctx.accounts.agent_account.balance.checked_sub(1).unwrap();

        // PDA seeds and bump to "sign" for CPI
        let seeds = AICO_SEED;
        let bump = ctx.bumps.reward_token_mint;
        let signer: &[&[&[u8]]] = &[&[seeds, &[bump]]];

        // CPI Context
        let cpi_ctx = CpiContext::new_with_signer(
            ctx.accounts.token_program.to_account_info(),
            MintTo {
                mint: ctx.accounts.reward_token_mint.to_account_info(),
                to: ctx.accounts.agent_token_account.to_account_info(),
                authority: ctx.accounts.reward_token_mint.to_account_info(),
            },
            signer,
        );

        // Mint 1 token, accounting for decimals of mint
        let amount = (1u64)
            .checked_mul(10u64.pow(ctx.accounts.reward_token_mint.decimals as u32))
            .unwrap();

        mint_to(cpi_ctx, amount)?;
        Ok(())
    }
}

#[derive(Accounts)]
pub struct UpdateBalance<'info> {
    #[account(mut)]
    pub holder: Signer<'info>,

    #[account(mut, seeds = [AGENT_SEED, holder.key.as_ref()], bump)]
    pub agent_account: Account<'info, AgentAccount>,

    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct CreateMint<'info> {
    #[account(
        mut,
        address = ADMIN_PUBKEY
    )]
    pub admin: Signer<'info>,

    // The PDA is both the address of the mint account and the mint authority
    #[account(
        init,
        seeds = [AICO_SEED],
        bump,
        payer = admin,
        mint::decimals = 9,
        mint::authority = reward_token_mint,

    )]
    pub reward_token_mint: Account<'info, Mint>,

    ///CHECK: Using "address" constraint to validate metadata account address
    #[account(
        mut,
        address=find_metadata_account(&reward_token_mint.key()).0
    )]
    pub metadata_account: UncheckedAccount<'info>,

    pub token_program: Program<'info, Token>,
    pub token_metadata_program: Program<'info, Metadata>,
    pub system_program: Program<'info, System>,
    pub rent: Sysvar<'info, Rent>,
}

#[derive(Accounts)]
pub struct InitializeAgent<'info> {
    #[account(mut,address = ADMIN_PUBKEY)]
    pub admin: AccountInfo<'info>,

    #[account(mut)]
    pub holder: Signer<'info>,

    #[account(
        init,
        payer = holder,
        seeds = [AGENT_SEED, holder.key.as_ref()],
        bump,
        space = 8 + std::mem::size_of::<AgentAccount>(),
    )]
    pub agent_account: Account<'info, AgentAccount>,

    // Initialize agent token account if it doesn't exist
    #[account(
        init_if_needed,
        payer = holder,
        associated_token::mint = reward_token_mint,
        associated_token::authority = holder
    )]
    pub agent_token_account: Account<'info, TokenAccount>,

    #[account(
        mut,
        seeds = [AICO_SEED],
        bump,
    )]
    pub reward_token_mint: Account<'info, Mint>,

    pub token_program: Program<'info, Token>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    pub system_program: Program<'info, System>,
}

#[account]
#[derive(Default)]
pub struct AgentAccount {
    pub holder: Pubkey,
    pub agent_name: String,
    pub balance: u64,
    pub created_at: i64,
    pub updated_at: i64,
}

#[error_code]
pub enum ErrorCode {
    #[msg("Not enough deposit")]
    NotEnoughDeposit,
    #[msg("Amount must be greater than zero")]
    AmountTooSmall,
}
