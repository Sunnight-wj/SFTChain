package cli

import (
	"strconv"

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
		NewMintTo(),
		NewUpdateVipDiscount(),
	)
	return txCmd
}

func NewUploadData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "upload [data_url] [tag] [flags]",
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
		Use:   "buy [tag] [flags]",
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
		Use:   "update-params [data_price] [fee_percentage] [flags]",
		Short: "update dataPrice and feePercentage.",
		Long:  "update dataPrice and feePercentage. only authority can update",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			authority := cliCtx.GetFromAddress()
			dataPrice, err := sdk.ParseCoinNormalized(args[0])
			if err != nil {
				return err
			}
			feePercentage, err := strconv.ParseFloat(args[1], 64)
			if err != nil {
				return err
			}
			msg := &types.MsgUpdateParams{
				Authority: authority.String(),
				Params: types.Params{
					DataPrice:     dataPrice,
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

func NewMintTo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-to [to_address] [amount] [flags]",
		Short: "Mint a denom to an address.",
		Long:  "Mint a denom to an address.",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			sender := cliCtx.GetFromAddress()

			toAddr, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}
			amount, err := sdk.ParseCoinNormalized(args[1])
			if err != nil {
				return err
			}
			msg := &types.MsgMintTo{
				Sender:        sender.String(),
				Amount:        amount,
				MintToAddress: toAddr.String(),
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

func NewUpdateVipDiscount() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-vip-discount [discount] [flags]",
		Short: "Update the VIP discount for buy data.",
		Long:  "Update the VIP discount for buy data.",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			//sender := cliCtx.GetFromAddress()

			discount, err := strconv.ParseFloat(args[0], 64)
			if err != nil {
				return err
			}
			msg := &types.MsgUpdateVipDiscount{
				Discount: discount,
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
