package cli

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/spf13/cobra"

	"github.com/CosmosContracts/juno/v17/x/datamarket/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd() *cobra.Command {
	// Group tokenfactory queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		GetCmdQueryParams(),
		GetCmdQueryData(),
		GetCmdQueryShare(),
		GetCmdQueryDataSetFromBuyer(),
	)

	return cmd
}

// GetCmdQueryParams returns the params for the module
func GetCmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params [flags]",
		Short: "Get the params for the x/datamarket module",
		Args:  cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(cmd.Context(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryData returns the all the urls of an dataSet
func GetCmdQueryData() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data [class] [flags]",
		Short: "Get all the data urls of an dataSet according the class",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Data(cmd.Context(), &types.QueryDataRequest{
				Class: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryShare returns the uploader share in an dataSet
func GetCmdQueryShare() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "share [class] [uploader_address] [flags]",
		Short: "Get the uploader share in an dataSet",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Share(cmd.Context(), &types.QueryShareRequest{
				Uploader: args[0],
				Class:    args[1],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

// GetCmdQueryDataSetFromBuyer returns all the dataSet the buyer own
func GetCmdQueryDataSetFromBuyer() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "data-set-from-buyer [buyer address] [flags]",
		Short: "Get all the dataSet the buyer own",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.DataSetFromBuyer(cmd.Context(), &types.QueryDataSetFromBuyerRequest{
				Buyer: args[0],
			})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
