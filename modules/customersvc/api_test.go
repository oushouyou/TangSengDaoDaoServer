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

	req, _ := http.NewRequest("POST", "/v1/awakenTheGroup", bytes.NewReader([]byte(util.ToJson(map[string]interface{}{
		"external_no":  "1223",
		"channel_type": 1,
		"category_no":  "1223333",
	}))))
	w := httptest.NewRecorder()
	req.Header.Set("token", testutil.Token)
	s.GetRoute().ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}
