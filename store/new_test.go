package store_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/sonofbytes/stscreds/store"
	"github.com/spf13/afero"
)

func TestNew(t *testing.T) {

	appFS := afero.NewMemMapFs()

	s := store.New(&store.StoreInput{
		StoreFS: appFS,
	})
	assert.IsType(t, &store.Store{}, s)
}

func TestStore_StoreDir(t *testing.T) {
	// test default .aws
	appFS := afero.NewMemMapFs()

	s := store.New(&store.StoreInput{
		StoreFS: appFS,
	})
	assert.Equal(t, ".aws", s.StoreDir())

	// test alternate location
	s = store.New(&store.StoreInput{
		StoreFS: appFS,
		StoreDir: ".somewhere",
	})
	assert.Equal(t, ".somewhere", s.StoreDir())
}