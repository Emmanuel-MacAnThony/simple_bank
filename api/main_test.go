package api

import (
	"os"
	"testing"
	"time"

	db "github.com/Emmanuel-MacAnThony/simple_bank/db/sqlc"
	"github.com/Emmanuel-MacAnThony/simple_bank/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
	// _ "github.com/lib/pq"
)

func newTestServer(t *testing.T, store db.Store) *Server {
	config := utils.Config{
		TokenSymetricKey:    utils.RandomString(32),
		AccessTokenDuration: time.Minute,
	}

	server, err := NewServer(config, store)
	require.NoError(t, err)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())

}
