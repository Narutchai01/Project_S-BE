package repositories

import "github.com/Narutchai01/Project_S-BE/entities"

type AdminRepository interface {
	CreateAdmin(admin entities.Admin) (entities.Admin, error)
	GetAdmins() ([]entities.Admin, error)
	GetAdmin(id int) (entities.Admin, error)
	UpdateAdmin(id int, admin entities.Admin) (entities.Admin, error)
	DeleteAdmin(id int) (entities.Admin, error)
	GetAdminByEmail(email string) (entities.Admin, error)
}
