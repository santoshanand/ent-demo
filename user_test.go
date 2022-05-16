package entdemo

import (
	"context"
	"log"
	"testing"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
	"github.com/santoshanand/ent-demo/ent"
	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Output:

	user1, err := client.User.Create().SetName("test user").Save(ctx)
	if err != nil {
		log.Fatalf("failed cr	eating a todo: %v", err)
	}
	assert.Nil(t, err)
	assert.NotNil(t, user1)
}
