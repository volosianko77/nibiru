package pricefeed_test

import (
	"testing"

	keepertest "github.com/MatrixDao/matrix/testutil/keeper"
	"github.com/MatrixDao/matrix/testutil/nullify"
	"github.com/MatrixDao/matrix/x/pricefeed"
	"github.com/MatrixDao/matrix/x/pricefeed/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.PricefeedKeeper(t)
	pricefeed.InitGenesis(ctx, k, genesisState)
	got := pricefeed.ExportGenesis(ctx, k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
