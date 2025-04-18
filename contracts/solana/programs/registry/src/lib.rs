use anchor_lang::prelude::*;
use solana_program::{pubkey, pubkey::Pubkey};

declare_id!("4JzjcmJSqWi2gpk8tZT93JHtyMimn3bubQMYs8FLAopM");

const ADMIN_PUBKEY: Pubkey = pubkey!("NAnMTmTdCYoszmR7cWsd1DodzfUCadnjXoHoHqwkPak");
const AGENT_SEED: &[u8] = b"leea_agent";

#[program]
pub mod registry {
    use super::*;
    pub fn register_agent(
        ctx: Context<RegisterAgent>,
        agent_name: String,
        description: String,
        fee: u64,
    ) -> Result<()> {
        let holder = &ctx.accounts.holder;
        let agent_account = &mut ctx.accounts.agent_account;
        agent_account.holder = *holder.key;
        agent_account.agent_name = agent_name;
        agent_account.description = description;
        agent_account.fee = fee;
        agent_account.created_at = Clock::get().unwrap().unix_timestamp;
        Ok(())
    }

    pub fn update_agent_score(ctx: Context<UpdateAgentScore>, score: u64) -> Result<()> {
        let agent_account = &mut ctx.accounts.agent_account;
        agent_account.score = score;
        agent_account.updated_at = Clock::get().unwrap().unix_timestamp;
        Ok(())
    }

    pub fn register_agent_by_admin(
        ctx: Context<RegisterAgentByAdmin>,
        agent_name: String,
        description: String,
        fee: u64,
    ) -> Result<()> {
        let admin = &ctx.accounts.admin;
        let holder = &ctx.accounts.holder;
        let agent_account = &mut ctx.accounts.agent_account;
        
        agent_account.holder = *holder.key;
        agent_account.agent_name = agent_name;
        agent_account.description = description;
        agent_account.fee = fee;
        agent_account.created_at = Clock::get().unwrap().unix_timestamp;
        Ok(())
    }
}

#[derive(Accounts)]
pub struct RegisterAgent<'info> {
    #[account(mut)]
    pub holder: Signer<'info>,
    #[account(
        init,
        payer = holder,
        seeds = [AGENT_SEED, holder.key().as_ref()],
        bump,
        space = 8 + std::mem::size_of::<AgentAccount>(),
    )]
    pub agent_account: Account<'info, AgentAccount>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct UpdateAgentScore<'info> {
    #[account(mut,address = ADMIN_PUBKEY)]
    pub admin: Signer<'info>,

    #[account(mut)]
    pub holder: AccountInfo<'info>,

    #[account(
        mut,
        seeds = [AGENT_SEED, holder.key().as_ref()],
        bump,
    )]
    pub agent_account: Account<'info, AgentAccount>,
    pub system_program: Program<'info, System>,
}

#[derive(Accounts)]
pub struct RegisterAgentByAdmin<'info> {
    #[account(mut, address = ADMIN_PUBKEY)]
    pub admin: Signer<'info>,

    /// CHECK: The holder account that will own the agent
    pub holder: AccountInfo<'info>,

    #[account(
        init,
        payer = admin,
        seeds = [AGENT_SEED, holder.key().as_ref()],
        bump,
        space = 8 + std::mem::size_of::<AgentAccount>(),
    )]
    pub agent_account: Account<'info, AgentAccount>,
    pub system_program: Program<'info, System>,
}

#[account]
#[derive(Default)]
pub struct AgentAccount {
    pub holder: Pubkey,
    pub agent_name: String,
    pub description: String,
    pub fee: u64,
    pub score: u64,
    pub created_at: i64,
    pub updated_at: i64,
}
