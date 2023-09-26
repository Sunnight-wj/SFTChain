package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
)

type BankKeeper interface {
}

type AccountKeeper interface {
	SetModuleAccount(ctx sdk.Context, macc authtypes.ModuleAccountI)
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
}

