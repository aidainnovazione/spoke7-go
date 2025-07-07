package main

import (
	"errors"
	"testing"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	"github.com/stretchr/testify/assert"
)

type Subject struct {
	Name   string
	Groups []string
}

type Object struct {
	Type   string
	Owner  string
	Groups []string
}

func Contains(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, errors.New("KeyMatchFunc expects exactly two arguments")
	}

	slice, ok1 := args[0].([]string)
	value, ok2 := args[1].(string)
	if !ok1 || !ok2 {
		return nil, errors.New("arguments must be ([]string, string)")
	}

	for _, s := range slice {
		if s == value {
			return true, nil
		}
	}
	return false, nil
}

func TestAccessControl(t *testing.T) {
	m, err := model.NewModelFromFile("model.conf")
	if err != nil {
		t.Fatalf("Failed to load model: %v", err)
	}

	adapter := fileadapter.NewAdapter("init_rules.csv")

	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		t.Fatalf("Failed to create enforcer: %v", err)
	}

	e.AddFunction("contains", Contains)

	// Manually add g-function for Casbin
	// No extra setup needed if you're using the default one

	tests := []struct {
		name     string
		sub      Subject
		obj      Object
		act      string
		expected bool
	}{
		{
			name:     "alice creates her own invoice - allowed",
			sub:      Subject{Name: "alice", Groups: []string{}},
			obj:      Object{Type: "invoice", Owner: "alice", Groups: []string{}},
			act:      "create",
			expected: true,
		},
		{
			name:     "alice reads her own invoice - allowed (via group editor)",
			sub:      Subject{Name: "alice", Groups: []string{}},
			obj:      Object{Type: "invoice", Owner: "alice", Groups: []string{}},
			act:      "read",
			expected: true,
		},
		{
			name:     "bob reads alice's invoice being collaborator - allower",
			sub:      Subject{Name: "bob", Groups: []string{"collaborator"}},
			obj:      Object{Type: "invoice", Owner: "alice", Groups: []string{"collaborator"}},
			act:      "read",
			expected: true,
		},
		{
			name:     "bob creates alice's invoice - denied",
			sub:      Subject{Name: "bob", Groups: []string{}},
			obj:      Object{Type: "invoice", Owner: "alice", Groups: []string{}},
			act:      "create",
			expected: false,
		},
		{
			name:     "bob reads alice's invoice - denied",
			sub:      Subject{Name: "bob", Groups: []string{}},
			obj:      Object{Type: "invoice", Owner: "alice", Groups: []string{}},
			act:      "read",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, err := e.Enforce(tt.sub, tt.obj, tt.act)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, ok)
		})
	}
}

func TestGetFilteredGroupingPolicy(t *testing.T) {
	m, err := model.NewModelFromFile("model.conf")
	if err != nil {
		t.Fatalf("Failed to load model: %v", err)
	}

	adapter := fileadapter.NewAdapter("init_rules.csv")

	e, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		t.Fatalf("Failed to create enforcer: %v", err)
	}

	// Example: get all groups/roles for user "bob"
	policies, err := e.GetFilteredPolicy(0, "alice")

	expected := [][]string{
		{"bob", "collaborator"},
	}

	assert.ElementsMatch(t, expected, policies)

	//groupingPolicies, err := e.GetImplicitRolesForUser("alice")

	//groupingPolicies, err := e.GetImplicitPermissionsForUser("alice")

	groupingPolicies, err := e.GetImplicitRolesForUser("alice")

	assert.ElementsMatch(t, expected, groupingPolicies)
}
