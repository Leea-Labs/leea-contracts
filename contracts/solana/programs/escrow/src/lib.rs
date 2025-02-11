use anchor_lang::prelude::*;
mod contexts;
use contexts::*;
mod states;

declare_id!("4JGo2qhxDy4gXi7PXVAiAeaTUWKwK8cnfTNJwSpsQQ8Q");
#[program]
pub mod escrow {
    use super::*;

    pub fn initialize(
        ctx: Context<Initialize>,
        seed: u64,
        initializer_amount: u64,
    ) -> Result<()> {
        ctx.accounts
            .initialize_escrow(seed, &ctx.bumps, initializer_amount)?;
        ctx.accounts.deposit(initializer_amount)
    }

    pub fn cancel(ctx: Context<Cancel>) -> Result<()> {
        ctx.accounts.refund_and_close_vault()
    }

    pub fn pay_to_agent(ctx: Context<Pay>, amount: u64) -> Result<()> {
        ctx.accounts.pay_to_agent(amount)
    }
}