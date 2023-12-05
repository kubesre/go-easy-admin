package dept

import "go-easy-admin/models"

type InterfaceDept interface {
	AddDept(dept *models.Dept) error
	DeptList() ([]models.Dept, error)
	DeptInfo(did string) ([]models.Dept, error)
	DelDept(did int) error
}

type deptInfo struct{}

func NewDeptInterface() InterfaceDept {
	return &deptInfo{}
}
