package keeper

import (
	// this line is used by starport scaffolding # 1
	"github.com/sapiens-cosmos/ibb/x/ibb/types"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	abci "github.com/tendermint/tendermint/abci/types"
)

func NewQuerier(k Keeper, legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		var (
			res []byte
			err error
		)

		switch path[0] {
		// this line is used by starport scaffolding # 2
		case types.QueryGetRepay:
			return getRepay(ctx, path[1], k, legacyQuerierCdc)

		case types.QueryListRepay:
			return listRepay(ctx, k, legacyQuerierCdc)

		case types.QueryGetWithdraw:
			return getWithdraw(ctx, path[1], k, legacyQuerierCdc)

		case types.QueryListWithdraw:
			return listWithdraw(ctx, k, legacyQuerierCdc)

		//query Pool
		case types.QueryListPool:
			return listPool(ctx, k, legacyQuerierCdc)

		case types.QueryGetPool:
			return getPool(ctx, path[1], k, legacyQuerierCdc)

		case types.QueryLoadPool:
			return loadPool(ctx, k, legacyQuerierCdc)

		//query User
		case types.QueryLoadUser:
			return loadUser(ctx, path[1], k, legacyQuerierCdc)

		case types.QueryListCollection:
			return listCollection(ctx, k, legacyQuerierCdc)
		default:
			err = sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unknown %s query endpoint: %s", types.ModuleName, path[0])
		}

		return res, err
	}
}
