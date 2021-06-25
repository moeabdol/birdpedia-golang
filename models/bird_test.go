package models

import (
	"context"
	"database/sql"
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

	bird, err = testStore.GetBird(context.Background(), bird.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, bird)
}

func TestGetBird(t *testing.T) {
	bird1 := createTestBird(t)

	bird2, err := testStore.GetBird(context.Background(), bird1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, bird2)
	require.Equal(t, bird1.ID, bird2.ID)
	require.Equal(t, bird1.Species, bird2.Species)
	require.Equal(t, bird1.Description, bird2.Description)
	require.Equal(t, bird1.CreatedAt, bird2.CreatedAt)
	require.Equal(t, bird1.UpdatedAt, bird2.UpdatedAt)

	deleteTestBird(t, bird1)
}

func TestListBirds(t *testing.T) {
	const n = 10
	for i := 0; i < n; i++ {
		createTestBird(t)
	}

	arg := ListBirdsParams{
		Limit:  100,
		Offset: 0,
	}

	birds, err := testStore.ListBirds(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, birds)
	require.Len(t, birds, n)

	for _, bird := range birds {
		require.NotEmpty(t, bird)
		require.NotZero(t, bird.ID)
		require.NotEmpty(t, bird.Species)
		require.NotEmpty(t, bird.Description)
		require.NotEmpty(t, bird.CreatedAt)
		require.NotEmpty(t, bird.UpdatedAt)

		deleteTestBird(t, bird)
	}
}

func TestUpdateBird(t *testing.T) {
	bird1 := createTestBird(t)

	arg := UpdateBirdParams{
		ID:          bird1.ID,
		Description: utils.RandomString(20),
	}

	bird2, err := testStore.UpdateBird(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, bird2)
	require.Equal(t, bird1.ID, bird2.ID)
	require.Equal(t, bird1.Species, bird2.Species)
	require.NotEqual(t, bird1.Description, bird2.Description)
	require.Equal(t, bird1.CreatedAt, bird2.CreatedAt)
	require.NotEqual(t, bird1.UpdatedAt, bird2.UpdatedAt)

	deleteTestBird(t, bird2)
}
