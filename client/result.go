// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "SAO v1": result Resource Client
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --force=true
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.3.1

package client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// GetResultPath computes a request path to the get action of result.
func GetResultPath(resultID string) string {
	param0 := resultID

	return fmt.Sprintf("/sao/v1/results/%s", param0)
}

// Returns an specific result with the given entry and testcase ID
func (c *Client) GetResult(ctx context.Context, path string) (*http.Response, error) {
	req, err := c.NewGetResultRequest(ctx, path)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewGetResultRequest create the request corresponding to the get action endpoint of the result resource.
func (c *Client) NewGetResultRequest(ctx context.Context, path string) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// ShowResultPath computes a request path to the show action of result.
func ShowResultPath() string {

	return fmt.Sprintf("/sao/v1/results/")
}

// List all the results delimited by the query params
func (c *Client) ShowResult(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, ranked *string, sort *string, task *int, user *int) (*http.Response, error) {
	req, err := c.NewShowResultRequest(ctx, path, contest, entry, page, pageSize, ranked, sort, task, user)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowResultRequest create the request corresponding to the show action endpoint of the result resource.
func (c *Client) NewShowResultRequest(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, ranked *string, sort *string, task *int, user *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if contest != nil {
		tmp3 := strconv.Itoa(*contest)
		values.Set("contest", tmp3)
	}
	if entry != nil {
		tmp4 := strconv.Itoa(*entry)
		values.Set("entry", tmp4)
	}
	if page != nil {
		tmp5 := strconv.Itoa(*page)
		values.Set("page", tmp5)
	}
	if pageSize != nil {
		tmp6 := strconv.Itoa(*pageSize)
		values.Set("page_size", tmp6)
	}
	if ranked != nil {
		values.Set("ranked", *ranked)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	if task != nil {
		tmp7 := strconv.Itoa(*task)
		values.Set("task", tmp7)
	}
	if user != nil {
		tmp8 := strconv.Itoa(*user)
		values.Set("user", tmp8)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
