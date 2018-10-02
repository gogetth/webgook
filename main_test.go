package main

import (
	"encoding/json"
	"net/http/httptest"
	"testing"

	"github.com/gogetth/webgook/api"
	"github.com/gogetth/webgook/helper"
	"github.com/labstack/echo"
)

func TestWebhook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/webhook", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("token")
	c.SetParamValues("test")

	api := &api.API{
		ScriptRunner: &helper.ScriptRunnerMockup{},
	}

	err := api.Webhook(c)
	if err != nil {
		t.Errorf("error : %q", err)
	}

	var resp map[string]interface{}
	json.Unmarshal(rec.Body.Bytes(), &resp)

	if resp["message"] != "ok" {
		t.Errorf("Want the word `ok` but found : %s", resp["message"])
	}
}
