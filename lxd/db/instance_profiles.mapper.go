//go:build linux && cgo && !agent
// +build linux,cgo,!agent

package db

// The code below was generated by lxd-generate - DO NOT EDIT!

import (
	"fmt"

	"github.com/lxc/lxd/lxd/db/cluster"
	"github.com/lxc/lxd/lxd/db/query"
	"github.com/lxc/lxd/shared/api"
)

var _ = api.ServerEnvironment{}

var instanceProfileObjects = cluster.RegisterStmt(`
SELECT instances_profiles.instance_id, instances_profiles.profile_id, instances_profiles.apply_order
  FROM instances_profiles
  ORDER BY instances_profiles.instance_id
`)

var instanceProfileCreate = cluster.RegisterStmt(`
INSERT INTO instances_profiles (instance_id, profile_id, apply_order)
  VALUES (?, ?, ?)
`)

var instanceProfileDeleteByInstanceID = cluster.RegisterStmt(`
DELETE FROM instances_profiles WHERE instance_id = ?
`)

// GetProfileInstances returns all available instance_profiles.
// generator: instance_profile GetMany
func (c *ClusterTx) GetProfileInstances() (map[int][]int, error) {
	var err error

	// Result slice.
	objects := make([]InstanceProfile, 0)

	stmt := c.stmt(instanceProfileObjects)
	args := []interface{}{}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, InstanceProfile{})
		return []interface{}{
			&objects[i].InstanceID,
			&objects[i].ProfileID,
			&objects[i].ApplyOrder,
		}
	}

	// Select.
	err = query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	resultMap := map[int][]int{}
	for _, object := range objects {
		resultMap[object.ProfileID] = append(resultMap[object.ProfileID], object.InstanceID)
	}

	return resultMap, nil
}

// GetInstanceProfiles returns all available instance_profiles.
// generator: instance_profile GetMany
func (c *ClusterTx) GetInstanceProfiles() (map[int][]int, error) {
	var err error

	// Result slice.
	objects := make([]InstanceProfile, 0)

	stmt := c.stmt(instanceProfileObjects)
	args := []interface{}{}

	// Dest function for scanning a row.
	dest := func(i int) []interface{} {
		objects = append(objects, InstanceProfile{})
		return []interface{}{
			&objects[i].InstanceID,
			&objects[i].ProfileID,
			&objects[i].ApplyOrder,
		}
	}

	// Select.
	err = query.SelectObjects(stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"instances_profiles\" table: %w", err)
	}

	resultMap := map[int][]int{}
	for _, object := range objects {
		resultMap[object.InstanceID] = append(resultMap[object.InstanceID], object.ProfileID)
	}

	return resultMap, nil
}

// CreateInstanceProfile adds a new instance_profile to the database.
// generator: instance_profile Create
func (c *ClusterTx) CreateInstanceProfile(object InstanceProfile) (int64, error) {
	args := make([]interface{}, 3)

	// Populate the statement arguments.
	args[0] = object.InstanceID
	args[1] = object.ProfileID
	args[2] = object.ApplyOrder

	// Prepared statement to use.
	stmt := c.stmt(instanceProfileCreate)

	// Execute the statement.
	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, fmt.Errorf("Failed to create \"instances_profiles\" entry: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Failed to fetch \"instances_profiles\" entry ID: %w", err)
	}

	return id, nil
}

// DeleteInstanceProfiles deletes the instance_profile matching the given key parameters.
// generator: instance_profile DeleteMany
func (c *ClusterTx) DeleteInstanceProfiles(object Instance) error {
	stmt := c.stmt(instanceProfileDeleteByInstanceID)
	result, err := stmt.Exec(int(object.ID))
	if err != nil {
		return fmt.Errorf("Delete \"instances_profiles\" entry failed: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	return nil
}
