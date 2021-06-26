package controllers

import (
	"os"
	"testing"

	"github.com/moeabdol/birdpedia-golang/models"
	"github.com/moeabdol/birdpedia-golang/utils"
)

func TestMain(m *testing.M) {
	utils.LoadConfig()
	models.ConnectDatabase()

	os.Exit(m.Run())
}
