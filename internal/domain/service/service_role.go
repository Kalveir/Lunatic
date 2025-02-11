// role service
package service

import (
	"api/internal/database"
	"api/internal/domain/models"
)

type (
	Role = models.Role
)

func GetRole() []Role {
	var role []Role
	database.DB.Find(&role)
	return role
}
func CreateRole(name string, permissionID []uint) Role {
	var permissions []Permission
	database.DB.Where("id IN ?", permissionID).Find(&permissions)
	role := Role{
		Name:        name,
		Permissions: permissions,
	}
	database.DB.Create(&role)
	roles := FindRole(role.ID)
	return roles

}
func FindRole(id uint) Role {
	var role Role
	database.DB.Preload("Permissions").Take(&role, id)
	return role
}
func UpdateRole(id uint, name string, permissionID []uint) Role {
	var role Role
	database.DB.Take(&role, id)

	var permission []Permission
	database.DB.Where("id IN ?", permissionID).Find(&permission)

	role.Name = name
	database.DB.Save(&role)

	database.DB.Model(&role).Association("Permissions").Replace(&permission)
	roles := FindRole(role.ID)
	return roles
}

func DeleteRole(id uint) error {
	var role models.Role
	if err := database.DB.Take(&role, id).Error; err != nil {
		return err
	}

	// Hapus relasi dari pivot table
	database.DB.Model(&role).Association("Permissions").Clear()

	// Hapus role dari database
	database.DB.Delete(&role)

	return nil
}
