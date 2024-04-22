package migration

import (
	"context"
	"github.com/dhiemaz/AccountService/constant"
	"github.com/dhiemaz/AccountService/infrastructures/logger"
	accessModel "github.com/dhiemaz/AccountService/internal/domain/access/model"
	officeModel "github.com/dhiemaz/AccountService/internal/domain/office/model"
	roleModel "github.com/dhiemaz/AccountService/internal/domain/role/model"
	"go.mongodb.org/mongo-driver/mongo"
	"os"
	"strings"
)

func MigrateData(direction string, db *mongo.Database) {
	var collectionName string

	switch strings.ToUpper(direction) {
	case "UP":
		{
			for i := 1; i <= 4; i++ {
				var datas []interface{}

				switch i {
				case 1:
					{
						collectionName = constant.AccessCollection
						keys := []accessModel.Access{
							{
								AccessId:   1,
								AccessName: "CREATE",
							},
							{
								AccessId:   2,
								AccessName: "UPDATE",
							},
							{
								AccessId:   3,
								AccessName: "DELETE",
							},
							{
								AccessId:   4,
								AccessName: "READ",
							},
						}

						for _, val := range keys {
							datas = append(datas, val)
						}
						break
					}
				case 2:
					{
						collectionName = constant.OfficeCollection
						keys := []officeModel.Office{
							{
								BranchId:   1,
								BranchName: "Head Office (HO)",
							},
							{
								BranchId:   2,
								BranchName: "Branch Office 1",
							},
							{
								BranchId:   3,
								BranchName: "Branch Office 2",
							},
							{
								BranchId:   4,
								BranchName: "Branch Office 3",
							},
							{
								BranchId:   5,
								BranchName: "Branch Office 4",
							},
							{
								BranchId:   6,
								BranchName: "Branch Office 5",
							},
						}

						for _, val := range keys {
							datas = append(datas, val)
						}
						break
					}
				case 3:
					{
						collectionName = constant.RoleCollection
						keys := []roleModel.Role{
							{
								RoleId:   1,
								RoleName: "Admin",
							},
							{
								RoleId:   2,
								RoleName: "Finance Staff",
							},
							{
								RoleId:   3,
								RoleName: "Operation Admin",
							},
							{
								RoleId:   5,
								RoleName: "Helpdesk",
							},
							{
								RoleId:   6,
								RoleName: "Warehouse Admin",
							},
							{
								RoleId:   7,
								RoleName: "Other Staff",
							},
						}

						for _, val := range keys {
							datas = append(datas, val)
						}
						break
					}
				default:
					{
						os.Exit(0)
					}
				}

				logger.WithFields(logger.Fields{"component": "migration-up", "action": "MigrateData"}).
					Infof("running migration up for data name : %s", collectionName)

				if _, err := db.Collection(collectionName).InsertMany(context.TODO(), datas); err != nil {
					logger.WithFields(logger.Fields{"component": "migration-up", "action": "MigrateData"}).
						Errorf("failed running migration up for %s, error : %v", collectionName, err)
				}
			}
		}
	case "DOWN":
		{
			for i := 1; i <= 4; i++ {
				switch i {
				case 1:
					{
						collectionName = constant.AccessCollection
						break
					}
				case 2:
					{
						collectionName = constant.OfficeCollection

						break
					}
				case 3:
					{
						collectionName = constant.RoleCollection
						break
					}
				default:
					{
						os.Exit(0)
					}
				}

				logger.WithFields(logger.Fields{"component": "migration-down", "action": "MigrateData"}).
					Infof("running migration down for data name : %s", collectionName)

				if err := db.Collection(collectionName).Drop(context.TODO()); err != nil {
					logger.WithFields(logger.Fields{"component": "migration-down", "action": "MigrateData"}).
						Errorf("failed running migration down for %s, error : %v", collectionName, err)
				}
			}
		}

	}
}
