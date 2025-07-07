package authz

import (
	"errors"
	"spoke7-go/pkg/logger"
	"sync"

	"github.com/casbin/casbin/v2"
	fileadapter "github.com/casbin/casbin/v2/persist/file-adapter"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type Authz struct {
	enforcer    *casbin.Enforcer
	config      AuthzConfig
	logger      logger.Logger
	db          *gorm.DB
	controllers []AuthResolver
	mu          sync.Mutex
}

func NewAuthz(cfg AuthzConfig, db *gorm.DB, logger logger.Logger) *Authz {
	return &Authz{
		config: cfg,
		db:     db,
		logger: logger,
	}
}

func (a *Authz) RegisterController(controllers ...AuthResolver) {
	a.mu.Lock()
	defer a.mu.Unlock()
	for _, controller := range controllers {
		a.controllers = append(a.controllers, controller)
	}

}

func (a *Authz) InitCasbin() error {

	if a.config.InitRulePath != "" {
		// 1. Prepare the enforcer with File first
		fileAdapter := fileadapter.NewAdapter(a.config.InitRulePath)
		fileEnforcer, err := casbin.NewEnforcer(a.config.ModelPath, fileAdapter)
		if err != nil {
			return err
		}
		fileEnforcer.AddFunction("contains", Contains)
		fileEnforcer.AddFunction("eq", Eq)

		if err = fileEnforcer.LoadPolicy(); err != nil {
			return err
		}
		filePolicies, err := fileEnforcer.GetPolicy()
		if err != nil {

		}
		fileGrouping, err := fileEnforcer.GetGroupingPolicy()
		if err != nil {
			return nil
		}

		// 2. Prepare the Database enforcer
		dbAdapter, err := gormadapter.NewAdapterByDB(a.db)
		if err != nil {
			return err
		}
		dbEnforcer, err := casbin.NewEnforcer(a.config.ModelPath, dbAdapter)
		if err != nil {
			return err
		}
		dbEnforcer.AddFunction("contains", Contains)
		dbEnforcer.AddFunction("eq", Eq)

		if err = dbEnforcer.LoadPolicy(); err != nil {
			return err
		}

		// 3. Determine which policies to add

		for _, p := range filePolicies {
			ok, err := dbEnforcer.HasPolicy(p[0], p[1], p[2])
			if err != nil {
				return err
			}
			if !ok {
				dbEnforcer.AddPolicy(p[0], p[1], p[2])
			}
		}

		for _, g := range fileGrouping {
			if ok, _ := dbEnforcer.HasGroupingPolicy(g[0], g[1]); !ok {
				dbEnforcer.AddGroupingPolicy(g[0], g[1])
			}
		}

		// 4. Save newly added policies to the DB
		if err := dbEnforcer.SavePolicy(); err != nil {
			return err
		}
		a.enforcer = dbEnforcer
	} else {
		dbAdapter, err := gormadapter.NewAdapterByDB(a.db)
		if err != nil {
			return err
		}
		dbEnforcer, err := casbin.NewEnforcer(a.config.ModelPath, dbAdapter)
		if err != nil {
			return err
		}

		dbEnforcer.AddFunction("contains", Contains)
		dbEnforcer.AddFunction("eq", Eq)

		err = dbEnforcer.LoadPolicy()
		if err != nil {
			return err
		}
		a.enforcer = dbEnforcer
	}

	//a.seedAdmin()
	return nil
}

func (a *Authz) GetEnforcer() *casbin.Enforcer {
	return a.enforcer
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

func Eq(args ...interface{}) (interface{}, error) {
	if len(args) != 2 {
		return nil, errors.New("KeyMatchFunc expects exactly two arguments")
	}

	value1, ok1 := args[0].(string)
	value2, ok2 := args[1].(string)
	if !ok1 || !ok2 {
		return nil, errors.New("arguments must be ([]string, string)")
	}

	return value1 == value2, nil

}
