package cli

import (
	"github.com/CosmosContracts/juno/v17/x/hello/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"strconv"
)

func GetTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Hello subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewUpdateParams(),
	)
	return txCmd
}

func NewUpdateParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-params [name] [age] [flags]",
		Short: "update name and age.",
		Long:  "update name and age. only authority can update",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			authority := cliCtx.GetFromAddress()
			name := args[0]
			age, err := strconv.ParseUint(args[1], 10, 64)
			if err != nil {
				return err
			}
			msg := &types.MsgUpdateParams{
				Authority: authority.String(),
				Params: types.Params{
					Name: name,
					Age:  age,
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
