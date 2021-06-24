package models

import (
	"context"
	"testing"

	"github.com/moeabdol/birdpedia-golang/utils"
	"github.com/stretchr/testify/require"
)

func createTestBird(t *testing.T) Bird {
	arg := CreateBirdParams{
		Species:     utils.RandomString(6),
		Description: utils.RandomString(20),
	}

	bird, err := testStore.CreateBird(context.Background(), arg)
	if err != nil {
		t.Error(err)
	}

	return bird
}

func deleteTestBird(t *testing.T, bird Bird) {
	err := testStore.DeleteBird(context.Background(), bird.ID)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateBird(t *testing.T) {
	arg := CreateBirdParams{
		Species:     "H. leucocephalus",
		Description: "The bald eagle is a bird of prey found in North America.",
	}

	bird, err := testStore.CreateBird(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bird)
	require.Equal(t, bird.Species, arg.Species)
	require.Equal(t, bird.Description, arg.Description)
	require.NotZero(t, bird.CreatedAt)
	require.NotZero(t, bird.UpdatedAt)

	deleteTestBird(t, bird)
}

func TestDeleteBird(t *testing.T) {
	bird := createTestBird(t)

	err := testStore.DeleteBird(context.Background(), bird.ID)
	require.NoError(t, err)
}
