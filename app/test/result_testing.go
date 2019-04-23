// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": result TestHelpers
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

package test

import (
	"bytes"
	"context"
	"fmt"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/goatest"
	"github.com/jossemargt/cms-sao/app"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
)

// GetResultBadRequest runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetResultBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, resultID string) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	getCtx, _err := app.NewGetResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// GetResultNotFound runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetResultNotFound(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, resultID string) http.ResponseWriter {
	// Setup service
	var (
		logBuf bytes.Buffer

		respSetter goatest.ResponseSetterFunc = func(r interface{}) {}
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	getCtx, _err := app.NewGetResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 404 {
		t.Errorf("invalid response status code: got %+v, expected 404", rw.Code)
	}

	// Return results
	return rw
}

// GetResultOK runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetResultOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, resultID string) (http.ResponseWriter, *app.ComJossemargtSaoResult) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	getCtx, _err := app.NewGetResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoResult
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoResult)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoResult", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetResultOKFull runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetResultOKFull(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, resultID string) (http.ResponseWriter, *app.ComJossemargtSaoResultFull) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	getCtx, _err := app.NewGetResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoResultFull
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoResultFull)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoResultFull", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// GetResultOKLink runs the method Get of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func GetResultOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, resultID string) (http.ResponseWriter, *app.ComJossemargtSaoResultLink) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	u := &url.URL{
		Path: fmt.Sprintf("/sao/v1/results/%v", resultID),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	prms["resultID"] = []string{fmt.Sprintf("%v", resultID)}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	getCtx, _err := app.NewGetResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Get(getCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt *app.ComJossemargtSaoResultLink
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(*app.ComJossemargtSaoResultLink)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoResultLink", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowResultBadRequest runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowResultBadRequest(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, contest *int, entry *int, page int, pageSize int, sort string, task *int, user *int) (http.ResponseWriter, error) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		query["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		query["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		prms["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		prms["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	showCtx, _err := app.NewShowResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		return nil, e
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 400 {
		t.Errorf("invalid response status code: got %+v, expected 400", rw.Code)
	}
	var mt error
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(error)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of error", resp, resp)
		}
	}

	// Return results
	return rw, mt
}

// ShowResultOK runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowResultOK(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, contest *int, entry *int, page int, pageSize int, sort string, task *int, user *int) (http.ResponseWriter, app.ComJossemargtSaoResultCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		query["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		query["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		prms["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		prms["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	showCtx, _err := app.NewShowResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoResultCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoResultCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoResultCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowResultOKFull runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowResultOKFull(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, contest *int, entry *int, page int, pageSize int, sort string, task *int, user *int) (http.ResponseWriter, app.ComJossemargtSaoResultFullCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		query["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		query["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		prms["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		prms["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	showCtx, _err := app.NewShowResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoResultFullCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoResultFullCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoResultFullCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}

// ShowResultOKLink runs the method Show of the given controller with the given parameters.
// It returns the response writer so it's possible to inspect the response headers and the media type struct written to the response.
// If ctx is nil then context.Background() is used.
// If service is nil then a default service is created.
func ShowResultOKLink(t goatest.TInterface, ctx context.Context, service *goa.Service, ctrl app.ResultController, contest *int, entry *int, page int, pageSize int, sort string, task *int, user *int) (http.ResponseWriter, app.ComJossemargtSaoResultLinkCollection) {
	// Setup service
	var (
		logBuf bytes.Buffer
		resp   interface{}

		respSetter goatest.ResponseSetterFunc = func(r interface{}) { resp = r }
	)
	if service == nil {
		service = goatest.Service(&logBuf, respSetter)
	} else {
		logger := log.New(&logBuf, "", log.Ltime)
		service.WithLogger(goa.NewLogger(logger))
		newEncoder := func(io.Writer) goa.Encoder { return respSetter }
		service.Encoder = goa.NewHTTPEncoder() // Make sure the code ends up using this decoder
		service.Encoder.Register(newEncoder, "*/*")
	}

	// Setup request context
	rw := httptest.NewRecorder()
	query := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		query["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		query["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		query["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		query["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		query["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		query["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		query["user"] = sliceVal
	}
	u := &url.URL{
		Path:     fmt.Sprintf("/sao/v1/results/"),
		RawQuery: query.Encode(),
	}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		panic("invalid test " + err.Error()) // bug
	}
	prms := url.Values{}
	if contest != nil {
		sliceVal := []string{strconv.Itoa(*contest)}
		prms["contest"] = sliceVal
	}
	if entry != nil {
		sliceVal := []string{strconv.Itoa(*entry)}
		prms["entry"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(page)}
		prms["page"] = sliceVal
	}
	{
		sliceVal := []string{strconv.Itoa(pageSize)}
		prms["page_size"] = sliceVal
	}
	{
		sliceVal := []string{sort}
		prms["sort"] = sliceVal
	}
	if task != nil {
		sliceVal := []string{strconv.Itoa(*task)}
		prms["task"] = sliceVal
	}
	if user != nil {
		sliceVal := []string{strconv.Itoa(*user)}
		prms["user"] = sliceVal
	}
	if ctx == nil {
		ctx = context.Background()
	}
	goaCtx := goa.NewContext(goa.WithAction(ctx, "ResultTest"), rw, req, prms)
	showCtx, _err := app.NewShowResultContext(goaCtx, req, service)
	if _err != nil {
		e, ok := _err.(goa.ServiceError)
		if !ok {
			panic("invalid test data " + _err.Error()) // bug
		}
		t.Errorf("unexpected parameter validation error: %+v", e)
		return nil, nil
	}

	// Perform action
	_err = ctrl.Show(showCtx)

	// Validate response
	if _err != nil {
		t.Fatalf("controller returned %+v, logs:\n%s", _err, logBuf.String())
	}
	if rw.Code != 200 {
		t.Errorf("invalid response status code: got %+v, expected 200", rw.Code)
	}
	var mt app.ComJossemargtSaoResultLinkCollection
	if resp != nil {
		var _ok bool
		mt, _ok = resp.(app.ComJossemargtSaoResultLinkCollection)
		if !_ok {
			t.Fatalf("invalid response media: got variable of type %T, value %+v, expected instance of app.ComJossemargtSaoResultLinkCollection", resp, resp)
		}
		_err = mt.Validate()
		if _err != nil {
			t.Errorf("invalid response media type: %s", _err)
		}
	}

	// Return results
	return rw, mt
}
