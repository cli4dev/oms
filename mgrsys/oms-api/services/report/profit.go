package report

import (
	"bytes"
	"fmt"
	"os"

	common "gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/common"

	"github.com/micro-plat/hydra/component"
	"github.com/micro-plat/hydra/context"
	"github.com/micro-plat/lib4go/utility"
	"gitlab.100bm.cn/micro-plat/oms/mgrsys/oms-api/modules/report"
)

//ReportProfitHandler 利润报表
type ReportProfitHandler struct {
	container       component.IContainer
	ReportProfitLib report.IDbReportProfit
	commonLib       common.ICommonLib
}

//NewReportProfitHandler 创建利润报表对象
func NewReportProfitHandler(container component.IContainer) (u *ReportProfitHandler) {
	return &ReportProfitHandler{
		container:       container,
		ReportProfitLib: report.NewDbReportProfit(container),
		commonLib:       &common.CommonLib{},
	}
}

//QueryHandle  获取利润报表数据列表
func (u *ReportProfitHandler) QueryHandle(ctx *context.Context) (r interface{}) {

	ctx.Log.Info("--------获取利润报表数据列表--------")
	ctx.Log.Info("1.参数校验")

	var input report.QueryReportProfit
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.执行操作")
	data, count, total, err := u.ReportProfitLib.Query(&input)
	if err != nil {
		return context.NewError(context.ERR_NOT_IMPLEMENTED, err)
	}

	ctx.Log.Info("3.返回结果")
	return map[string]interface{}{
		"data":  data,
		"count": count,
		"total": total,
	}
}

func (u *ReportProfitHandler) ExportHandle(ctx *context.Context) interface{} {

	ctx.Log.Info("---------------导出利润统计---------------")
	ctx.Log.Info("1.参数校验")
	var input report.ExportReportProfit
	if err := ctx.Request.Bind(&input); err != nil {
		return context.NewError(context.ERR_NOT_ACCEPTABLE, err)
	}

	ctx.Log.Info("2.获取数据")
	data, header, err := u.ReportProfitLib.Export(&input)
	if err != nil {
		return err
	}

	ctx.Log.Info("3.构建数据")
	var buffer bytes.Buffer
	fileName := fmt.Sprintf("%s.xlsx", utility.GetGUID())
	u.commonLib.BuildXlsxData(data, header, fileName)
	f, err := os.Open(fileName)
	if err != nil {
		ctx.Log.Errorf("打开文件失败:err:%+v", err)
		return err
	}
	defer func() {
		f.Close()
		ctx.Log.Infof("删除文件:%s", fileName)
		os.Remove(fileName)
	}()
	buffer.ReadFrom(f)

	ctx.Log.Info("4.设置响应头")
	ctx.Response.SetHeader("Accept-Ranges", "bytes")
	ctx.Response.SetHeader("Content-Type", "application/vnd.ms-excel.numberformat:@")
	ctx.Response.SetHeader("Content-Disposition", "attachment;filename=利润统计报表.xlsx")
	ctx.Response.SetHeader("Charset", "utf-8")

	ctx.Log.Info("5.返回数据")
	return buffer.String()
}
