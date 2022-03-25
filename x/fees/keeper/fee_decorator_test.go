package keeper_test

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/desmos-labs/desmos/v3/x/fees/keeper"
	"github.com/desmos-labs/desmos/v3/x/fees/types"
)

func (suite *KeeperTestSuite) TestFeeDecorator() {
	suite.SetupTest()
	params := types.NewParams(
		[]types.MinFee{
			types.NewMinFee(
				"desmos.profiles.v2.MsgSaveProfile",
				sdk.NewCoins(sdk.NewCoin("stake", sdk.NewInt(10000)))),
		},
	)
	suite.keeper.SetParams(suite.ctx, params)

	feeDecorator := keeper.NewMinFeeDecorator(suite.keeper)

	testsCases := []struct {
		name      string
		txFee     sdk.Coins
		shouldErr bool
	}{
		{
			name:      "Minimum fees doesn't match and returns error",
			txFee:     sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 150)),
			shouldErr: true,
		},
		{
			name:      "Minimum fees match and returns no error",
			txFee:     sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 13000)),
			shouldErr: false,
		},
	}

	for _, tc := range testsCases {
		anteHandler := sdk.ChainAnteDecorators(feeDecorator)
		tx := legacytx.NewStdTx([]sdk.Msg{}, legacytx.NewStdFee(
			10,
			tc.txFee,
		), []legacytx.StdSignature{}, "")

		_, err := anteHandler(suite.ctx, tx, false)
		if tc.shouldErr {
			suite.Require().Error(err)
		} else {
			suite.Require().NoError(err)
		}
	}

}
