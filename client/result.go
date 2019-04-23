// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": result Resource Client
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

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

// Get complete Entry Result data for the given entry and testcase ID.
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

// List the Results delimited and grouped by contest, task, entry or user identifier
func (c *Client) ShowResult(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Response, error) {
	req, err := c.NewShowResultRequest(ctx, path, contest, entry, page, pageSize, sort, task, user)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowResultRequest create the request corresponding to the show action endpoint of the result resource.
func (c *Client) NewShowResultRequest(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, sort *string, task *int, user *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if contest != nil {
		tmp22 := strconv.Itoa(*contest)
		values.Set("contest", tmp22)
	}
	if entry != nil {
		tmp23 := strconv.Itoa(*entry)
		values.Set("entry", tmp23)
	}
	if page != nil {
		tmp24 := strconv.Itoa(*page)
		values.Set("page", tmp24)
	}
	if pageSize != nil {
		tmp25 := strconv.Itoa(*pageSize)
		values.Set("page_size", tmp25)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	if task != nil {
		tmp26 := strconv.Itoa(*task)
		values.Set("task", tmp26)
	}
	if user != nil {
		tmp27 := strconv.Itoa(*user)
		values.Set("user", tmp27)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
