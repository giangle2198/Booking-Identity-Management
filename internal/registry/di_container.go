package registry

import (
	"booking-identity-management/config"
	"booking-identity-management/internal/adapter"
	"booking-identity-management/internal/api"
	"booking-identity-management/internal/helper"
	"booking-identity-management/internal/repository"
	"booking-identity-management/internal/repository/model"
	"booking-identity-management/internal/usecase"
	libRegistry "booking-libs/registry"

	"github.com/sarulabs/di"
)

const (
	// config
	ConfigDIName string = "Config"
	// repository
	BaseRepositoryDIName          string = "baseRepository"
	RoleRepositoryDIName          string = "roleRepository"
	UserRepositoryDIName          string = "userRepository"
	UserAccountRepositoryDIName   string = "userAccountRepository"
	JWTManagementRepositoryDIName string = "jwtManagementRepository"
	// usecase
	BaseUsecaseDIName     string = "baseUsecase"
	RoleUsecaseDIName     string = "roleUsecase"
	UserUsecaseDIName     string = "userUsecase"
	UserRoleUsecaseDIName string = "userRoleUsecase"
	// helper
	sqlDBHelperDIName    string = "SQLDBHelper"
	pbConvertDIName      string = "PBConverter"
	modelConverterDIName string = "ModelConverter"
	// adapter
	PasetoAdapterDIName string = "PasetoAdapter"
	JWTAdapterDIName    string = "JwtAdapter"
	// api
	apiDIName string = "API"
)

// BuildDIContainer Wrapper for lib registry BuildDIContainer
func BuildDIContainer() {
	initBuilder()
	libRegistry.BuildDIContainer()
}

// GetDependency Wrapper for lib registry GetDependency
func GetDependency(name string) interface{} {
	return libRegistry.GetDependency(name)
}

func initBuilder() {
	// init for configs builder
	libRegistry.ConfigsBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  ConfigDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg, err := config.Load()
				return cfg, err
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for helpers builder
	libRegistry.HelpersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  sqlDBHelperDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return helper.NewDBHelper(cfg.MySQL.Host, int(cfg.MySQL.Port), cfg.MySQL.Username, cfg.MySQL.Password, cfg.MySQL.Database, cfg.MySQL.IsMigrate, model.Role{}, model.UserAccount{}, model.JWTManagment{}, model.User{}), nil
			},
			Close: func(obj interface{}) error {
				return nil
			}}, di.Def{
			Name:  pbConvertDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				return helper.NewPbConverter(), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  modelConverterDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				return helper.NewModelConverter(), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for apdeters builder
	libRegistry.AdaptersBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  JWTAdapterDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return adapter.NewJWTAdapter(cfg), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  PasetoAdapterDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				return adapter.NewPasetoAdapter(cfg), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for usecases builder
	libRegistry.UsecasesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  BaseUsecaseDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				baseRepository := di.Get(BaseRepositoryDIName).(repository.BaseRepository)
				roleRepository := di.Get(RoleRepositoryDIName).(repository.RoleRepository)
				userAccountRepository := di.Get(UserAccountRepositoryDIName).(repository.UserAccountRepository)
				userRepository := di.Get(UserRepositoryDIName).(repository.UserActionRepository)
				return usecase.NewBaseUsecase(cfg, baseRepository, roleRepository, userAccountRepository, userRepository), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  RoleUsecaseDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				roleRepository := di.Get(RoleRepositoryDIName).(repository.RoleRepository)
				return usecase.NewRoleUsecase(cfg, roleRepository), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  UserUsecaseDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				userAccountRepository := di.Get(UserAccountRepositoryDIName).(repository.UserAccountRepository)
				userRepository := di.Get(UserRepositoryDIName).(repository.UserActionRepository)
				jwtAdapter := di.Get(JWTAdapterDIName).(adapter.JWTAdapter)
				pasetoAdapter := di.Get(PasetoAdapterDIName).(adapter.PasetoAdapter)
				jwtManagementRepository := di.Get(JWTManagementRepositoryDIName).(repository.JWTManagementRepository)
				return usecase.NewUserActionUsecase(cfg, userRepository, userAccountRepository, jwtAdapter, jwtManagementRepository, pasetoAdapter), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  UserRoleUsecaseDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				baseUsecase := di.Get(BaseUsecaseDIName).(usecase.BaseUsecase)
				roleRepository := di.Get(RoleRepositoryDIName).(repository.RoleRepository)
				userAccountRepository := di.Get(UserAccountRepositoryDIName).(repository.UserAccountRepository)
				userRepository := di.Get(UserRepositoryDIName).(repository.UserActionRepository)
				return usecase.NewUserRoleUsecase(baseUsecase, cfg, roleRepository, userAccountRepository, userRepository), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})

		return arr
	}

	// init for repositories builder
	libRegistry.RepositoriesBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  UserRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				dbHelper := di.Get(sqlDBHelperDIName).(helper.DBHelper)
				return repository.NewUserActionRepository(cfg, dbHelper), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  UserAccountRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				dbHelper := di.Get(sqlDBHelperDIName).(helper.DBHelper)
				return repository.NewUserAccount(cfg, dbHelper), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  BaseRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				dbHelper := di.Get(sqlDBHelperDIName).(helper.DBHelper)
				return repository.NewBaseRepository(cfg, dbHelper), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  RoleRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				dbHelper := di.Get(sqlDBHelperDIName).(helper.DBHelper)
				return repository.NewRoleRepository(cfg, dbHelper), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		}, di.Def{
			Name:  JWTManagementRepositoryDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				dbHelper := di.Get(sqlDBHelperDIName).(helper.DBHelper)
				return repository.NewJWTManagement(cfg, dbHelper), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}

	// init for apis builder
	libRegistry.APIsBuilder = func() []di.Def {
		arr := []di.Def{}
		arr = append(arr, di.Def{
			Name:  apiDIName,
			Scope: di.App,
			Build: func(di di.Container) (interface{}, error) {
				cfg := di.Get(ConfigDIName).(*config.Config)
				baseUsecase := di.Get(BaseUsecaseDIName).(usecase.BaseUsecase)
				roleUsecase := di.Get(RoleUsecaseDIName).(usecase.RoleUsecase)
				userUsecase := di.Get(UserUsecaseDIName).(usecase.UserActionUsecase)
				userRoleUsecase := di.Get(UserRoleUsecaseDIName).(usecase.UserRoleUsecase)
				pbConvert := di.Get(pbConvertDIName).(helper.PbConverter)
				return api.NewAPI(cfg, pbConvert, baseUsecase, userUsecase, userRoleUsecase, roleUsecase), nil
			},
			Close: func(obj interface{}) error {
				return nil
			},
		})
		return arr
	}
}
