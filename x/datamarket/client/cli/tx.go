package cli

import (
	"strconv"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetTxCmd returns a root CLI command handler for certain modules/DataMarket
// transaction commands.
func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "DataMarket subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewUploadData(),
		NewBuyData(),
		NewUpdateParams(),
	)
	return txCmd
}

func NewUploadData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload [data_url] [tag]",
		Short: "Upload data url with a tag.",
		Long:  "Upload data url with a tag.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			uploader := cliCtx.GetFromAddress()

			dataUrl := args[0]
			class := args[1]

			msg := &types.MsgUploadData{
				Uploader: uploader.String(),
				Class:    class,
				Url:      dataUrl,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewBuyData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "buy [tag]",
		Short: "buy dataSet with a tag.",
		Long:  "buy dataSet with a tag.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			buyer := cliCtx.GetFromAddress()

			class := args[0]

			msg := &types.MsgBuyData{
				Buyer: buyer.String(),
				Class: class,
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func NewUpdateParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-params [data_price_denom] [data_price_amount] [fee_percentage]",
		Short: "update dataPrice and feePercentage.",
		Long:  "update dataPrice and feePercentage. only authority can update",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			authority := cliCtx.GetFromAddress()

			dataPriceDenom := args[0]

			dataPriceAmount, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}
			feePercentage, err := strconv.ParseFloat(args[2], 64)
			if err != nil {
				return err
			}
			msg := &types.MsgUpdateParams{
				Authority: authority.String(),
				Params: &types.Params{
					DataPrice: &sdk.Coin{
						Denom:  dataPriceDenom,
						Amount: math.NewInt(dataPriceAmount),
					},
					FeePercentage: feePercentage,
				},
			}

			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
