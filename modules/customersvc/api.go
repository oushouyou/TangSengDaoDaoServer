package customersvc

import (
	"errors"
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
	r.POST("/v1/awakenTheGroup", c.AwakenTheGroup)
}

type custormerReq struct {
	Username string `json:"username"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
}

func (c *CustomersvcGroup) AwakenTheGroup(r *wkhttp.Context) {
	var creq custormerReq
	if err := r.BindJSON(&creq); err != nil {
		r.ResponseError(errors.New("请求数据格式有误！"))
		return
	}
	customer := Customer{
		Username: creq.Username,
		Phone:    creq.Phone,
		Email:    creq.Email,
	}
	resp, err := c.service.AwakenTheGroup(customer)
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
