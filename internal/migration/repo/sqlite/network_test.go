package sqlite_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	dbschema "github.com/FuturFusion/migration-manager/internal/db"
	dbdriver "github.com/FuturFusion/migration-manager/internal/db/sqlite"
	"github.com/FuturFusion/migration-manager/internal/migration"
	"github.com/FuturFusion/migration-manager/internal/migration/repo/sqlite"
	"github.com/FuturFusion/migration-manager/internal/migration/repo/sqlite/entities"
	"github.com/FuturFusion/migration-manager/internal/transaction"
	"github.com/FuturFusion/migration-manager/shared/api"
)

func TestNetworkDatabaseActions(t *testing.T) {
	networkA := migration.Network{Identifier: "networkA", Type: api.NETWORKTYPE_VMWARE_STANDARD, Location: "/path/to/networkA", Source: testSource.Name, Properties: []byte("{}")}
	networkB := migration.Network{Identifier: "networkB", Type: api.NETWORKTYPE_VMWARE_STANDARD, Location: "/path/to/networkA", Source: testSource.Name, Overrides: api.NetworkOverride{Name: "foo"}, Properties: []byte("{}")}
	networkC := migration.Network{Identifier: "networkC", Type: api.NETWORKTYPE_VMWARE_STANDARD, Location: "/path/to/networkC", Source: testSource.Name, Overrides: api.NetworkOverride{Name: "bar"}, Properties: []byte("{}")}

	ctx := context.Background()

	tmpDir := t.TempDir()

	db, err := dbdriver.Open(tmpDir)
	require.NoError(t, err)

	t.Cleanup(func() {
		err = db.Close()
		require.NoError(t, err)
	})

	_, err = dbschema.EnsureSchema(db, tmpDir)
	require.NoError(t, err)

	tx := transaction.Enable(db)
	entities.PreparedStmts, err = entities.PrepareStmts(tx, false)
	require.NoError(t, err)

	sourceSvc := migration.NewSourceService(sqlite.NewSource(tx))
	_, err = sourceSvc.Create(ctx, testSource)
	require.NoError(t, err)

	network := sqlite.NewNetwork(tx)

	// Add networkA.
	networkA.ID, err = network.Create(ctx, networkA)
	require.NoError(t, err)

	// Add networkB.
	networkB.ID, err = network.Create(ctx, networkB)
	require.NoError(t, err)

	// Add networkC.
	_, err = network.Create(ctx, networkC)
	require.NoError(t, err)

	// Ensure we have three entries
	networks, err := network.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, networks, 3)

	// Should get back networkA unchanged.
	dbNetworkA, err := network.GetByNameAndSource(ctx, networkA.Identifier, networkA.Source)
	require.NoError(t, err)
	require.Equal(t, networkA, *dbNetworkA)

	dbNetworkA, err = network.GetByNameAndSource(ctx, networkA.Identifier, networkA.Source)
	require.NoError(t, err)
	require.Equal(t, networkA, *dbNetworkA)

	// Test updating a network.
	networkB.Overrides.Name = "baz"
	err = network.Update(ctx, networkB)
	require.NoError(t, err)
	dbNetworkB, err := network.GetByNameAndSource(ctx, networkB.Identifier, networkB.Source)
	require.NoError(t, err)
	require.Equal(t, networkB, *dbNetworkB)

	// Delete a network.
	err = network.DeleteByNameAndSource(ctx, networkA.Identifier, networkA.Source)
	require.NoError(t, err)
	_, err = network.GetByNameAndSource(ctx, networkA.Identifier, networkA.Source)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Should have two networks remaining.
	networks, err = network.GetAll(ctx)
	require.NoError(t, err)
	require.Len(t, networks, 2)

	// Can't delete a network that doesn't exist.
	err = network.DeleteByNameAndSource(ctx, "BazBiz", "something")
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Can't update a network that doesn't exist.
	err = network.Update(ctx, networkA)
	require.ErrorIs(t, err, migration.ErrNotFound)

	// Can't add a duplicate network.
	networkB.ID, err = network.Create(ctx, networkB)
	require.ErrorIs(t, err, migration.ErrConstraintViolation)
}
