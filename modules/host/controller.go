package host

import (
	"awesome/controller"
	"awesome/models"
	"awesome/utils"
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateTenant(ctx *fiber.Ctx) error {
	return controller.Create[models.Tenant](
		models.TENANT_MODEL_NAME,
		controller.CreateOptions[models.Tenant]{
			GetDB: utils.GetHostDB,
			PreCreate: func(ctx *fiber.Ctx, db *gorm.DB, tenant *models.Tenant) error {
				tenant.TenantOwnerID = ctx.Locals("userId").(uint)
				// TODO: change this to the actual url (bit of a hack for now)
				tenant.TenantUrl = "http://localhost:300" + strconv.FormatUint(uint64(ctx.Locals("userId").(uint)), 10)
				return nil
			},
			PostCreate: func(ctx *fiber.Ctx, db *gorm.DB, tenant *models.Tenant) error {
				dbConnectionString, err := utils.CreateDatabase(tenant.Name, db)
				if err != nil {
					if strings.Contains(err.Error(), "database already exists") {
						return errors.New("database already exists")
					} else {
						return err
					}
				}

				err = db.Table(models.TENANT_MODEL_NAME).Where("id = ?", tenant.ID).Updates(map[string]interface{}{
					"tenantDBConnectionString": dbConnectionString,
				}).Error
				if err != nil {
					return err
				}

				tenantDB, err := utils.GetDbConnection(dbConnectionString)
				if err != nil {
					return err
				}

				err = utils.GormMigrate(tenantDB)
				if err != nil {
					return err
				}

				hostDB := utils.GetHostDB()
				var user models.TenantOwner
				err = hostDB.First(&user, tenant.TenantOwnerID).Error
				if err != nil {
					return err
				}

				if user.ID == 0 {
					return fiber.ErrNotFound
				}

				tenantUser := models.User{
					Name:     user.Name,
					Password: user.Password,
					Email:    user.Email,
				}

				err = tenantDB.Create(&tenantUser).Error
				if err != nil {
					return err
				}

				newResourceIndex := models.Resource{
					Name:         tenantUser.Name,
					ResourceType: "users",
					ResourceID:   tenantUser.ID,
				}

				err = tenantDB.Create(&newResourceIndex).Error
				return err
			},
		},
	)(ctx)
}
