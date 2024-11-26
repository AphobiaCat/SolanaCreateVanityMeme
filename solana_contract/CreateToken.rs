use anchor_lang::prelude::*;

use anchor_spl::{
    associated_token::AssociatedToken,
    metadata::{
        create_metadata_accounts_v3, mpl_token_metadata::types::DataV2, CreateMetadataAccountsV3,
        Metadata as Metaplex,
    },
    token::{mint_to, Mint, MintTo, Token, TokenAccount},
};

// This is your program's public key and it will update
// automatically when you build the project.
declare_id!("");

#[program]
mod creator {
    use super::*;

    const TOTAL_SUPPLY_MEME: f64 = 10000000000.0;
    const DECIMALS: u64 = 1000000000;

    pub fn create_meme(
        ctx: Context<CreateMeme>,
        metadata: InitTokenParams,
    ) -> Result<()> {
        let random_num1_  = metadata.random_num1;
        let random_num2_  = metadata.random_num2;

        let rs_s_   = metadata.random_str;
        let rn1_s_  = random_num1_.to_le_bytes();
        let rn2_s_  = random_num2_.to_le_bytes();

        let seeds = &[
            rs_s_.as_bytes(),
            rn1_s_.as_ref(),
            rn2_s_.as_ref(),
            &[ctx.bumps.mint],
        ];
        let signer  = [&seeds[..]];

        let name_   = metadata.name.clone();
        let symbol_ = metadata.symbol.clone();
        let uri_    = metadata.uri.clone();

        let token_data: DataV2 = DataV2 {
            name: name_,
            symbol: symbol_,
            uri: uri_,
            seller_fee_basis_points: 0,
            creators: None,
            collection: None,
            uses: None,
        };

        let metadata_ctx = CpiContext::new_with_signer(
            ctx.accounts.token_metadata_program.to_account_info(),
            CreateMetadataAccountsV3 {
                payer: ctx.accounts.payer.to_account_info(),
                update_authority: ctx.accounts.mint.to_account_info(),
                mint: ctx.accounts.mint.to_account_info(),
                metadata: ctx.accounts.metadata.to_account_info(),
                mint_authority: ctx.accounts.mint.to_account_info(),
                system_program: ctx.accounts.system_program.to_account_info(),
                rent: ctx.accounts.rent.to_account_info(),
            },
            &signer,
        );

        create_metadata_accounts_v3(metadata_ctx, token_data, false, true, None)?;
        
        
        mint_to(
            CpiContext::new_with_signer(
                ctx.accounts.token_program.to_account_info(),
                MintTo {
                    authority: ctx.accounts.mint.to_account_info(),
                    to: ctx.accounts.payer_token_account.to_account_info(),
                    mint: ctx.accounts.mint.to_account_info(),
                },
                &signer,
            ),
            TOTAL_SUPPLY_MEME as u64 * DECIMALS,
        )?;

        anchor_spl::token::set_authority(
            CpiContext::new_with_signer(
                ctx.accounts.token_program.to_account_info(),
                anchor_spl::token::SetAuthority {
                    account_or_mint: ctx.accounts.mint.to_account_info(),
                    current_authority: ctx.accounts.mint.to_account_info(),
                },
                &signer,
            ),
            spl_token::instruction::AuthorityType::MintTokens,
            None,
        )?;

        Ok(())
    }
}

#[derive(Accounts)]
#[instruction(
    params: InitTokenParams
)]
pub struct CreateMeme<'info> {
    #[account(
        init,
        seeds = [params.random_str.as_bytes(), params.random_num1.to_le_bytes().as_ref(), params.random_num2.to_le_bytes().as_ref()],
        bump,
        payer = payer,
        mint::decimals = 9,
        mint::authority = mint,
    )]
    pub mint: Account<'info, Mint>,
    #[account(mut)]
    pub metadata: UncheckedAccount<'info>,
    #[account(
        init_if_needed,
        payer = payer,
        associated_token::mint = mint,
        associated_token::authority = payer,
    )]
    pub payer_token_account: Account<'info, TokenAccount>,

    #[account(mut)]
    pub payer: Signer<'info>,
    pub rent: Sysvar<'info, Rent>,
    pub system_program: Program<'info, System>,
    pub token_program: Program<'info, Token>,
    pub token_metadata_program: Program<'info, Metaplex>,
    pub associated_token_program: Program<'info, AssociatedToken>,
}

#[derive(AnchorSerialize, AnchorDeserialize, Debug, Clone)]
pub struct InitTokenParams {
    pub name        : String,
    pub symbol      : String,
    pub uri         : String,
    pub random_str  : String,
    pub random_num1 : u64,
    pub random_num2 : u64,
}