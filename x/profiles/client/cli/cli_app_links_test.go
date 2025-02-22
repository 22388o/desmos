package cli_test

import (
	"fmt"
	"time"

	"github.com/cosmos/cosmos-sdk/client/flags"
	clitestutil "github.com/cosmos/cosmos-sdk/testutil/cli"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/golang/protobuf/proto"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	"github.com/desmos-labs/desmos/v3/x/profiles/client/cli"
	"github.com/desmos-labs/desmos/v3/x/profiles/types"
)

func (s *IntegrationTestSuite) TestCmdQueryApplicationsLinks() {
	val := s.network.Validators[0]
	testCases := []struct {
		name           string
		args           []string
		shouldErr      bool
		expectedOutput types.QueryApplicationLinksResponse
	}{
		{
			name: "existing links are returned properly",
			args: []string{
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryApplicationLinksResponse{
				Links: []types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
			},
		},
		{
			name: "existing links of the given user address are not found",
			args: []string{
				"cosmos122u6u9gpdr2rp552fkkvlgyecjlmtqhkascl5a",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryApplicationLinksResponse{
				Links: []types.ApplicationLink{},
			},
		},
		{
			name: "existing links of the given user are returned properly",
			args: []string{
				"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
				fmt.Sprintf("--%s=json", tmcli.OutputFlag),
			},
			shouldErr: false,
			expectedOutput: types.QueryApplicationLinksResponse{
				Links: []types.ApplicationLink{
					types.NewApplicationLink(
						"cosmos1ftkjv8njvkekk00ehwdfl5sst8zgdpenjfm4hs",
						types.NewData("reddit", "reddit-user"),
						types.ApplicationLinkStateInitialized,
						types.NewOracleRequest(
							0,
							1,
							types.NewOracleRequestCallData("twitter", "call_data"),
							"client_id",
						),
						nil,
						time.Date(2020, 1, 1, 00, 00, 00, 000, time.UTC),
					),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdQueryApplicationsLinks()
			clientCtx := val.ClientCtx
			out, err := clitestutil.ExecTestCLICmd(clientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)

				var response types.QueryApplicationLinksResponse
				s.Require().NoError(clientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), &response), out.String())
				s.Require().Equal(tc.expectedOutput.Links, response.Links)
			}
		})
	}
}

func (s *IntegrationTestSuite) TestCmdUnlinkApplication() {
	val := s.network.Validators[0]
	testCases := []struct {
		name      string
		args      []string
		shouldErr bool
		respType  proto.Message
	}{
		{
			name:      "empty app name returns error",
			args:      []string{"", "twitter"},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name:      "empty username returns error",
			args:      []string{"twitter", ""},
			shouldErr: true,
			respType:  &sdk.TxResponse{},
		},
		{
			name: "valid request works properly",
			args: []string{
				"reddit",
				"reddit-user",
				fmt.Sprintf("--%s=%s", flags.FlagFrom, val.Address.String()),
				fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
				fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
			},
			shouldErr: false,
			respType:  &sdk.TxResponse{},
		},
	}

	for _, tc := range testCases {
		tc := tc

		s.Run(tc.name, func() {
			cmd := cli.GetCmdUnlinkApplication()
			out, err := clitestutil.ExecTestCLICmd(val.ClientCtx, cmd, tc.args)

			if tc.shouldErr {
				s.Require().Error(err)
			} else {
				s.Require().NoError(err)
				s.Require().NoError(val.ClientCtx.JSONCodec.UnmarshalJSON(out.Bytes(), tc.respType), out.String())
			}
		})
	}
}
