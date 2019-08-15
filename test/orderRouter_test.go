package test

import (
	"Demo1/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestOrderRouter_orderList(t *testing.T) {
	r := router.InitRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/order/list", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"code":10002,"data":{},"msg":"获取信息列表成功"}`, w.Body.String())
}
