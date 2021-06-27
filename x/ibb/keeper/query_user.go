package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/sapiens-cosmos/ibb/x/ibb/types"
)

func loadUser(ctx sdk.Context, walletAddress string, keeper Keeper, legacyQuerierCdc *codec.LegacyAmino) ([]byte, error) {
	//getUser
	userList := keeper.GetAllUser(ctx)
	var queryUser types.User
	for _, user := range userList {
		if user.Creator == walletAddress {
			queryUser = user
		}
	}

	//TODO: add append user logic when the user has not been registered before
	//the creator var for user should be set as walletAddress variable we received from parameter
	// if queryUser.Creator == "" {
	// 	queryUser.Creator = walletAddress
	// 	queryUser.Collateral = []bool{false, false, false, false, false, false, false, false}
	// 	keeper.AppendUser(ctx, queryUser)
	// }

	// var userAssetList []types.LoadUserResponse
	// var userAsset types.LoadUserResponse
	// assetBalances := keeper.bankKeeper.GetAllBalances(ctx, sdk.AccAddress(walletAddress))
	// for i := 0; i < 8; i++ {
	// 	userAsset.AssetBalance = assetBalances[0].Amount
	// 	userAssetList = append(userAssetList, userAsset)
	// }

	bz, err := codec.MarshalJSONIndent(legacyQuerierCdc, queryUser)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return bz, nil
}