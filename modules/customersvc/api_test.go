package customersvc

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/pkg/util"
	"github.com/TangSengDaoDao/TangSengDaoDaoServerLib/testutil"
	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	s, ctx := testutil.NewTestServer()
	m := New(ctx)
	m.Route(s.GetRoute())

	req, _ := http.NewRequest("POST", "/v1/reports", bytes.NewReader([]byte(util.ToJson(map[string]interface{}{
		"channel_id":   "1223",
		"channel_type": 1,
		"category_no":  "1223333",
		"imgs":         []string{"http://xdsdsd.com/wewe/dsd.png", "http://xdsdsd.com/wewe/1223.png"},
		"remark":       "this is remark",
	}))))
	w := httptest.NewRecorder()
	req.Header.Set("token", testutil.Token)
	s.GetRoute().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
