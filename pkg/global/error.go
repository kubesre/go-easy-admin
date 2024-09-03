/*
@auth: AnRuo
@source: 云原生运维圈
@website: https://www.kubesre.com/
@time: 2024/8/5
*/

package global

import "errors"

// 定义错误信息

func CreateErr(filed string, err error) error {
	e := errors.New(filed + "创建失败")
	GeaLogger.Error(e, err)
	return e
}

func UpdateErr(filed string, err error) error {
	e := errors.New(filed + "更新失败")
	GeaLogger.Error(e, err)
	return e
}

func DeleteErr(filed string, err error) error {
	e := errors.New(filed + "删除失败")
	GeaLogger.Error(e, err)
	return e
}
func NotFoundErr(filed string, err error) error {
	e := errors.New(filed + "不存在")
	GeaLogger.Error(e, err)
	return e
}

func GetErr(filed string, err error) error {
	e := errors.New(filed + "获取失败")
	GeaLogger.Error(e, err)
	return e
}

func OtherErr(err error, s ...string) error {
	GeaLogger.Error(err, s)
	return err
}
