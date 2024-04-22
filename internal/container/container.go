package container

import (
	"github.com/dhiemaz/AccountService/infrastructures/database"
	accessAdapter "github.com/dhiemaz/AccountService/internal/domain/access/adapter/mongo"
	accessPresenter "github.com/dhiemaz/AccountService/internal/domain/access/presenter/grpc"
	accessUC "github.com/dhiemaz/AccountService/internal/domain/access/usecase"
	accountAdapter "github.com/dhiemaz/AccountService/internal/domain/account/adapter/mongo"
	accountPresenter "github.com/dhiemaz/AccountService/internal/domain/account/presenter/grpc"
	accountUC "github.com/dhiemaz/AccountService/internal/domain/account/usecase"
	officeAdapter "github.com/dhiemaz/AccountService/internal/domain/office/adapter/mongo"
	officePresenter "github.com/dhiemaz/AccountService/internal/domain/office/presenter/grpc"
	officeUC "github.com/dhiemaz/AccountService/internal/domain/office/usecase"
	roleAdapter "github.com/dhiemaz/AccountService/internal/domain/role/adapter/mongo"
	rolePresenter "github.com/dhiemaz/AccountService/internal/domain/role/presenter/grpc"
	roleUC "github.com/dhiemaz/AccountService/internal/domain/role/usecase"
	"github.com/sarulabs/di"
)

// Container data
type Container struct {
	ctn di.Container
}

// NewContainer create a new container object
func NewContainer() *Container {
	builder, _ := di.NewBuilder()

	_ = builder.Add([]di.Def{
		{Name: "accessSvc", Build: accessUsecase},
		{Name: "accountSvc", Build: accountUsecase},
		{Name: "officeSvc", Build: officeUsecase},
		{Name: "roleSvc", Build: roleUsecase},
	}...)

	return &Container{
		ctn: builder.Build(),
	}
}

// Resolve container by name
func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}

func accessUsecase(_ di.Container) (interface{}, error) {
	accessRepository := accessAdapter.NewAccessMongoAdapter(database.GetMongoConnection())
	presenter := &accessPresenter.AccessPresenter{}
	return accessUC.NewAccessUsecase(accessRepository, presenter), nil
}

func accountUsecase(_ di.Container) (interface{}, error) {
	accountRepository := accountAdapter.NewAccountMongoAdapter(database.GetMongoConnection())
	presenter := &accountPresenter.AccountPresenter{}
	return accountUC.NewAccountUsecase(accountRepository, presenter), nil
}

func officeUsecase(_ di.Container) (interface{}, error) {
	officeRepository := officeAdapter.NewOfficeMongoAdapter(database.GetMongoConnection())
	presenter := &officePresenter.OfficePresenter{}
	return officeUC.NewOfficeUsecase(officeRepository, presenter), nil
}

func roleUsecase(_ di.Container) (interface{}, error) {
	roleRepository := roleAdapter.NewRoleMongoAdapter(database.GetMongoConnection())
	presenter := &rolePresenter.RolePresenter{}
	return roleUC.NewRoleUsecase(roleRepository, presenter), nil
}
