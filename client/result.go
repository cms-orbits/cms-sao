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

// Get complete result data for the given entry and testcase ID.
// The "re" and "ut" prefix delimits the entry type as "ranked entry" or "user test" respectively.
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

// List the results delimited and grouped by contest, task, entry or user identifier
func (c *Client) ShowResult(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, ranked *bool, sort *string, task *int, user *int) (*http.Response, error) {
	req, err := c.NewShowResultRequest(ctx, path, contest, entry, page, pageSize, ranked, sort, task, user)
	if err != nil {
		return nil, err
	}
	return c.Client.Do(ctx, req)
}

// NewShowResultRequest create the request corresponding to the show action endpoint of the result resource.
func (c *Client) NewShowResultRequest(ctx context.Context, path string, contest *int, entry *int, page *int, pageSize *int, ranked *bool, sort *string, task *int, user *int) (*http.Request, error) {
	scheme := c.Scheme
	if scheme == "" {
		scheme = "http"
	}
	u := url.URL{Host: c.Host, Scheme: scheme, Path: path}
	values := u.Query()
	if contest != nil {
		tmp8 := strconv.Itoa(*contest)
		values.Set("contest", tmp8)
	}
	if entry != nil {
		tmp9 := strconv.Itoa(*entry)
		values.Set("entry", tmp9)
	}
	if page != nil {
		tmp10 := strconv.Itoa(*page)
		values.Set("page", tmp10)
	}
	if pageSize != nil {
		tmp11 := strconv.Itoa(*pageSize)
		values.Set("page_size", tmp11)
	}
	if ranked != nil {
		tmp12 := strconv.FormatBool(*ranked)
		values.Set("ranked", tmp12)
	}
	if sort != nil {
		values.Set("sort", *sort)
	}
	if task != nil {
		tmp13 := strconv.Itoa(*task)
		values.Set("task", tmp13)
	}
	if user != nil {
		tmp14 := strconv.Itoa(*user)
		values.Set("user", tmp14)
	}
	u.RawQuery = values.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	return req, nil
}
