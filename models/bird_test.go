package models

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBird(t *testing.T) {
	arg := CreateBirdParams{
		Species:     "hello",
		Description: "world",
	}

	bird, err := testStore.CreateBird(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bird)
}
