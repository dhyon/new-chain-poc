package types_test

import (
	"testing"

	"github.com/codchen/new-chain-poc/x/dex/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{

				LongBookList: []types.LongBook{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				LongBookCount: 2,
				ShortBookList: []types.ShortBook{
					{
						Id: 0,
					},
					{
						Id: 1,
					},
				},
				ShortBookCount: 2,
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated longBook",
			genState: &types.GenesisState{
				LongBookList: []types.LongBook{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid longBook count",
			genState: &types.GenesisState{
				LongBookList: []types.LongBook{
					{
						Id: 1,
					},
				},
				LongBookCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated shortBook",
			genState: &types.GenesisState{
				ShortBookList: []types.ShortBook{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid shortBook count",
			genState: &types.GenesisState{
				ShortBookList: []types.ShortBook{
					{
						Id: 1,
					},
				},
				ShortBookCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
