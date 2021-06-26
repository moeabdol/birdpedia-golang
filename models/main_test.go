package models

import (
	"os"
	"testing"

	"github.com/moeabdol/birdpedia-golang/utils"
)

func TestMain(m *testing.M) {
	utils.LoadConfig()
	ConnectDatabase()

	os.Exit(m.Run())
}
