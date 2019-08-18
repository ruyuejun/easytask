package test

import (
	"dcs-gocron/router"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserRouter_userLogin(t *testing.T) {
	r := router.NewRouters()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/login", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"code":10001,"data":null,"msg":"登录成功"}`, w.Body.String())
}
