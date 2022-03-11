package cli

import (
	"strconv"

	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
)

var _ = strconv.Itoa(0)

func CmdLimitSell() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "limit-sell [price] [quantity]",
		Short: "Broadcast message limitSell",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			argPrice, err := cast.ToInt32E(args[0])
			if err != nil {
				return err
			}
			argQuantity, err := cast.ToInt32E(args[1])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgLimitSell(
				clientCtx.GetFromAddress().String(),
				argPrice,
				argQuantity,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
