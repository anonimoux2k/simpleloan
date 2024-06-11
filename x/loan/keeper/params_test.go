package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	keepertest "simpleloan/testutil/keeper"
	"simpleloan/x/loan/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := keepertest.LoanKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
