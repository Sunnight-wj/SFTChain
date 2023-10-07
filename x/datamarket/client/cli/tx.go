package cli

import (
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
	"github.com/cosmos/cosmos-sdk/client"
)

// NewTxCmd returns a root CLI command handler for certain modules/DataMarket
// transaction commands.
func NewTxCmd() *cobra.Command {
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
