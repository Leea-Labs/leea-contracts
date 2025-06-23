use anchor_lang::prelude::*;
use anchor_spl::{
    associated_token::AssociatedToken,
    token::{transfer_checked, Mint, Token, TokenAccount, TransferChecked},
};

use crate::states::Escrow;
use crate::global::ADMIN_PUBKEY;

#[derive(Accounts)]
pub struct Initialize<'info> {
    #[account(mut, address = ADMIN_PUBKEY)]
    pub admin: Signer<'info>,
    /// The user for whom the escrow is being created
    #[account(mut)]
    pub initializer: UncheckedAccount<'info>,
    pub mint_a: Account<'info, Mint>,
    #[account(
        init_if_needed,
        payer = admin,
        space = Escrow::INIT_SPACE,
        seeds = [b"state".as_ref(), initializer.key().as_ref()],
        bump
    )]
    pub escrow: Account<'info, Escrow>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    pub token_program: Program<'info, Token>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct Deposit<'info> {
    #[account(mut)]
    /// The authority funding the escrow (can be admin or user)
    pub authority: Signer<'info>,
    /// The user for whom the escrow is being funded
    #[account(mut)]
    pub initializer: UncheckedAccount<'info>,
    pub mint_a: Account<'info, Mint>,
    /// Source token account (must belong to authority)
    #[account(
        mut,
        associated_token::mint = mint_a,
        associated_token::authority = authority
    )]
    pub from_ata: Account<'info, TokenAccount>,
    #[account(
        mut,
        has_one = mint_a,
        seeds=[b"state", initializer.key().as_ref()],
        bump = escrow.bump
    )]
    pub escrow: Box<Account<'info, Escrow>>,
    #[account(
        init_if_needed,
        payer = authority,
        associated_token::mint = mint_a,
        associated_token::authority = escrow
    )]
    pub vault: Account<'info, TokenAccount>,
    pub associated_token_program: Program<'info, AssociatedToken>,
    pub token_program: Program<'info, Token>,
    pub system_program: Program<'info, System>,
}

impl<'info> Initialize<'info> {
    pub fn initialize_escrow(&mut self, bumps: &InitializeBumps) -> Result<()> {
        self.escrow.set_inner(Escrow {
            bump: bumps.escrow,
            initializer: self.initializer.key(),
            mint_a: self.mint_a.key(),
        });
        Ok(())
    }
}

impl<'info> Deposit<'info> {
    pub fn deposit(&mut self, amount: u64) -> Result<()> {
        transfer_checked(self.into_deposit_context(), amount, self.mint_a.decimals)
    }

    fn into_deposit_context(&self) -> CpiContext<'_, '_, '_, 'info, TransferChecked<'info>> {
        let cpi_accounts = TransferChecked {
            from: self.from_ata.to_account_info(),
            mint: self.mint_a.to_account_info(),
            to: self.vault.to_account_info(),
            authority: self.authority.to_account_info(),
        };
        CpiContext::new(self.token_program.to_account_info(), cpi_accounts)
    }
}
