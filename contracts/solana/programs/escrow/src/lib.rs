use anchor_lang::prelude::*;
use anchor_spl::token::{Mint, Token, TokenAccount};

declare_id!("Fg6PaFpoGXkYsidMpWTK6W2BeZ7FEfcYkg476zPFsLnS");

const ADMIN_PUBKEY: Pubkey = pubkey!("GB9XNqUC32ZibLza8d7qMKBEv1hPZ142hzZ3sju7hG7b");

#[program]
pub mod escrow {
    use super::*;
    pub fn initialize(ctx: Context<Initialize>, x_amount: u64) -> Result<()> {
        let escrow = &mut ctx.accounts.escrow;
        escrow.bump = ctx.bumps.escrow;
        escrow.authority = ctx.accounts.user.key();
        escrow.escrowed_leea_tokens = ctx.accounts.escrowed_leea_tokens.key();


        // Transfer users's leea_token in program owned escrow token account
        anchor_spl::token::transfer(
            CpiContext::new(
                ctx.accounts.token_program.to_account_info(),
                anchor_spl::token::Transfer {
                    from: ctx.accounts.user_leea_token.to_account_info(),
                    to: ctx.accounts.escrowed_leea_tokens.to_account_info(),
                    authority: ctx.accounts.user.to_account_info(),
                },
            ),
            x_amount,
        )?;
        Ok(())
    }

    pub fn pay(ctx: Context<Pay>, y_amount: u64) -> Result<()> {
        // transfer escrowed_leea_token to agent
        anchor_spl::token::transfer(
            CpiContext::new_with_signer(
                ctx.accounts.token_program.to_account_info(),
                anchor_spl::token::Transfer {
                    from: ctx.accounts.escrowed_leea_tokens.to_account_info(),
                    to: ctx.accounts.agent_leea_tokens.to_account_info(),
                    authority: ctx.accounts.escrow.to_account_info(),
                },
                &[&["escrow".as_bytes(), ctx.accounts.escrow.authority.as_ref(), &[ctx.accounts.escrow.bump]]],
            ),
            y_amount,
        )?;
        Ok(())
    }
}


#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(mut)]
    user: Signer<'info>,
    leea_mint: Account<'info, Mint>,
    #[account(mut, constraint = user_leea_token.mint == leea_mint.key() && user_leea_token.owner == user.key())] 
    user_leea_token: Account<'info, TokenAccount>,
    #[account(
        init, 
        payer = user,  
        space = 8 + std::mem::size_of::<Escrow>(),
        seeds = ["escrow".as_bytes(), user.key().as_ref()],
        bump,
    )]
    pub escrow: Account<'info, Escrow>,
    #[account(
        init,
        payer = user,
        token::mint = leea_mint,
        token::authority = escrow,
    )]
    escrowed_leea_tokens: Account<'info, TokenAccount>,
    token_program: Program<'info, Token>,
    rent: Sysvar<'info, Rent>,
    system_program: Program<'info, System>,
}


#[account]
pub struct Escrow {
    authority: Pubkey,
    bump: u8,
    escrowed_leea_tokens: Pubkey,
}

#[derive(Accounts)]
pub struct Pay<'info> {
    pub admin: Signer<'info>,
    #[account(
        mut,
        seeds = ["escrow".as_bytes(), escrow.authority.as_ref()],
        bump = escrow.bump,
        address = ADMIN_PUBKEY
    )]
    pub escrow: Account<'info, Escrow>,
    #[account(mut, constraint = escrowed_leea_tokens.key() == escrow.escrowed_leea_tokens)]
    pub escrowed_leea_tokens: Account<'info, TokenAccount>,
    #[account(mut, constraint = agent_leea_tokens.mint == escrow.escrowed_leea_tokens)]
    pub agent_leea_tokens: Account<'info, TokenAccount>,
    pub token_program: Program<'info, Token>,
}