package customersvc

import (
	"net/http"

	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/config"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/wkhttp"
)

type CustomersvcGroup struct {
	service IService
}

func New(ctx *config.Context) *CustomersvcGroup {
	return &CustomersvcGroup{
		service: NewService(ctx),
	}
}

func (c *CustomersvcGroup) Route(r *wkhttp.WKHttp) {
	r.POST("/awakenTheGroup", c.AwakenTheGroup)
}

func (c *CustomersvcGroup) AwakenTheGroup(r *wkhttp.Context) {

	externalNo := r.MustGet("externalNo").(string)

	resp, err := c.service.AwakenTheGroup(externalNo)
	if err != nil {
		r.ResponseError(err)
		return
	}

	r.JSON(http.StatusOK, &CustomersvcGroupResp{
		GroupId: resp,
	})
}

type CustomersvcGroupResp struct {
	GroupId string `json:"group_id,omitempty"`
}
