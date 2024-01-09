/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/1/9
*/

package operationLogs

import (
	"errors"
	"go-easy-admin/common/global"
	"go-easy-admin/dao"
	"go-easy-admin/models"
)

type OperationLogService interface {
	GetOperationLogList(limit, page int) (*models.OperationLogList, error)
}
type operationLogService struct{}

func NewOperationLogService() OperationLogService {
	return &operationLogService{}
}

func (s *operationLogService) GetOperationLogList(limit, page int) (*models.OperationLogList, error) {
	data, err := dao.NewOperationLogService().GetOperationLogList(limit, page)
	if err != nil {
		global.TPLogger.Error("获取操作日志列表失败：", err)
		return nil, errors.New("获取操作日志列表")
	}
	return data, nil
}
