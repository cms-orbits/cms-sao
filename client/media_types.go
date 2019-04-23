// Code generated by goagen v1.4.1, DO NOT EDIT.
//
// API "SAO": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.4.1

package client

import (
	"github.com/goadesign/goa"
	"net/http"
)

// Any source code or input to be compiled and executed against the user test case (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft+json; view=default
type ComJossemargtSaoDraft struct {
	// Contest unique and human readable string identifier
	ContestSlug string `form:"contestSlug" json:"contestSlug" yaml:"contestSlug" xml:"contestSlug"`
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Task unique and human readable string identifier
	TaskSlug string `form:"taskSlug" json:"taskSlug" yaml:"taskSlug" xml:"taskSlug"`
}

// Validate validates the ComJossemargtSaoDraft media type instance.
func (mt *ComJossemargtSaoDraft) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.ContestSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.contestSlug`, mt.ContestSlug, `[_a-zA-Z0-9\-]+`))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.TaskSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.taskSlug`, mt.TaskSlug, `[_a-zA-Z0-9\-]+`))
	}
	return
}

// Any source code or input to be compiled and executed against the user test case (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft+json; view=full
type ComJossemargtSaoDraftFull struct {
	// Contest ID where this Entry has been submitted
	ContestID int `form:"contestID" json:"contestID" yaml:"contestID" xml:"contestID"`
	// Contest unique and human readable string identifier
	ContestSlug string `form:"contestSlug" json:"contestSlug" yaml:"contestSlug" xml:"contestSlug"`
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Links to related resources
	Links *ComJossemargtSaoDraftLinks `form:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty" xml:"links,omitempty"`
	// Task ID where this Entry has been submitted
	TaskID int `form:"taskID" json:"taskID" yaml:"taskID" xml:"taskID"`
	// Task unique and human readable string identifier
	TaskSlug string `form:"taskSlug" json:"taskSlug" yaml:"taskSlug" xml:"taskSlug"`
	// User ID of the Entry's owner
	UserID int `form:"userID" json:"userID" yaml:"userID" xml:"userID"`
}

// Validate validates the ComJossemargtSaoDraftFull media type instance.
func (mt *ComJossemargtSaoDraftFull) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.ContestSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.contestSlug`, mt.ContestSlug, `[_a-zA-Z0-9\-]+`))
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.TaskSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.taskSlug`, mt.TaskSlug, `[_a-zA-Z0-9\-]+`))
	}
	return
}

// Any source code or input to be compiled and executed against the user test case (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft+json; view=link
type ComJossemargtSaoDraftLink struct {
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the ComJossemargtSaoDraftLink media type instance.
func (mt *ComJossemargtSaoDraftLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// ComJossemargtSaoDraftLinks contains links to related resources of ComJossemargtSaoDraft.
type ComJossemargtSaoDraftLinks struct {
	Result *ComJossemargtSaoDraftResultLink `form:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty" xml:"result,omitempty"`
}

// Validate validates the ComJossemargtSaoDraftLinks type instance.
func (ut *ComJossemargtSaoDraftLinks) Validate() (err error) {
	if ut.Result != nil {
		if err2 := ut.Result.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeComJossemargtSaoDraft decodes the ComJossemargtSaoDraft instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraft(resp *http.Response) (*ComJossemargtSaoDraft, error) {
	var decoded ComJossemargtSaoDraft
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoDraftFull decodes the ComJossemargtSaoDraftFull instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftFull(resp *http.Response) (*ComJossemargtSaoDraftFull, error) {
	var decoded ComJossemargtSaoDraftFull
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoDraftLink decodes the ComJossemargtSaoDraftLink instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftLink(resp *http.Response) (*ComJossemargtSaoDraftLink, error) {
	var decoded ComJossemargtSaoDraftLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// The representation of the result of an entry draft compile, execution and evaluation process (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft-result+json; view=default
type ComJossemargtSaoDraftResult struct {
	// Entry evaluation result
	Evaluation *EvaluationResult `form:"evaluation,omitempty" json:"evaluation,omitempty" yaml:"evaluation,omitempty" xml:"evaluation,omitempty"`
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Compound Result ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the ComJossemargtSaoDraftResult media type instance.
func (mt *ComJossemargtSaoDraftResult) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Evaluation != nil {
		if err2 := mt.Evaluation.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// The representation of the result of an entry draft compile, execution and evaluation process (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft-result+json; view=full
type ComJossemargtSaoDraftResultFull struct {
	// Entry compilation result
	Compilation *CompilationResult `form:"compilation,omitempty" json:"compilation,omitempty" yaml:"compilation,omitempty" xml:"compilation,omitempty"`
	// Entry evaluation result
	Evaluation *EvaluationResult `form:"evaluation,omitempty" json:"evaluation,omitempty" yaml:"evaluation,omitempty" xml:"evaluation,omitempty"`
	// Entry execution result
	Execution *ExecutionResult `form:"execution" json:"execution" yaml:"execution" xml:"execution"`
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Compound Result ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Links to related resources
	Links *ComJossemargtSaoDraftResultLinks `form:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty" xml:"links,omitempty"`
}

// Validate validates the ComJossemargtSaoDraftResultFull media type instance.
func (mt *ComJossemargtSaoDraftResultFull) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Execution == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "execution"))
	}
	if mt.Compilation != nil {
		if err2 := mt.Compilation.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Evaluation != nil {
		if err2 := mt.Evaluation.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// The representation of the result of an entry draft compile, execution and evaluation process (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft-result+json; view=link
type ComJossemargtSaoDraftResultLink struct {
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Compound Result ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the ComJossemargtSaoDraftResultLink media type instance.
func (mt *ComJossemargtSaoDraftResultLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// ComJossemargtSaoDraft-ResultLinks contains links to related resources of ComJossemargtSaoDraft-Result.
type ComJossemargtSaoDraftResultLinks struct {
}

// DecodeComJossemargtSaoDraftResult decodes the ComJossemargtSaoDraftResult instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftResult(resp *http.Response) (*ComJossemargtSaoDraftResult, error) {
	var decoded ComJossemargtSaoDraftResult
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoDraftResultFull decodes the ComJossemargtSaoDraftResultFull instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftResultFull(resp *http.Response) (*ComJossemargtSaoDraftResultFull, error) {
	var decoded ComJossemargtSaoDraftResultFull
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoDraftResultLink decodes the ComJossemargtSaoDraftResultLink instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftResultLink(resp *http.Response) (*ComJossemargtSaoDraftResultLink, error) {
	var decoded ComJossemargtSaoDraftResultLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ComJossemargtSaoDraft-ResultCollection is the media type for an array of ComJossemargtSaoDraft-Result (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft-result+json; type=collection; view=default
type ComJossemargtSaoDraftResultCollection []*ComJossemargtSaoDraftResult

// Validate validates the ComJossemargtSaoDraftResultCollection media type instance.
func (mt ComJossemargtSaoDraftResultCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoDraft-ResultCollection is the media type for an array of ComJossemargtSaoDraft-Result (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft-result+json; type=collection; view=full
type ComJossemargtSaoDraftResultFullCollection []*ComJossemargtSaoDraftResultFull

// Validate validates the ComJossemargtSaoDraftResultFullCollection media type instance.
func (mt ComJossemargtSaoDraftResultFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoDraft-ResultCollection is the media type for an array of ComJossemargtSaoDraft-Result (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft-result+json; type=collection; view=link
type ComJossemargtSaoDraftResultLinkCollection []*ComJossemargtSaoDraftResultLink

// Validate validates the ComJossemargtSaoDraftResultLinkCollection media type instance.
func (mt ComJossemargtSaoDraftResultLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoDraft-ResultLinksArray contains links to related resources of ComJossemargtSaoDraft-ResultCollection.
type ComJossemargtSaoDraftResultLinksArray []*ComJossemargtSaoDraftResultLinks

// DecodeComJossemargtSaoDraftResultCollection decodes the ComJossemargtSaoDraftResultCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftResultCollection(resp *http.Response) (ComJossemargtSaoDraftResultCollection, error) {
	var decoded ComJossemargtSaoDraftResultCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoDraftResultFullCollection decodes the ComJossemargtSaoDraftResultFullCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftResultFullCollection(resp *http.Response) (ComJossemargtSaoDraftResultFullCollection, error) {
	var decoded ComJossemargtSaoDraftResultFullCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoDraftResultLinkCollection decodes the ComJossemargtSaoDraftResultLinkCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftResultLinkCollection(resp *http.Response) (ComJossemargtSaoDraftResultLinkCollection, error) {
	var decoded ComJossemargtSaoDraftResultLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// ComJossemargtSaoDraftCollection is the media type for an array of ComJossemargtSaoDraft (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft+json; type=collection; view=default
type ComJossemargtSaoDraftCollection []*ComJossemargtSaoDraft

// Validate validates the ComJossemargtSaoDraftCollection media type instance.
func (mt ComJossemargtSaoDraftCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoDraftCollection is the media type for an array of ComJossemargtSaoDraft (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft+json; type=collection; view=full
type ComJossemargtSaoDraftFullCollection []*ComJossemargtSaoDraftFull

// Validate validates the ComJossemargtSaoDraftFullCollection media type instance.
func (mt ComJossemargtSaoDraftFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoDraftCollection is the media type for an array of ComJossemargtSaoDraft (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.draft+json; type=collection; view=link
type ComJossemargtSaoDraftLinkCollection []*ComJossemargtSaoDraftLink

// Validate validates the ComJossemargtSaoDraftLinkCollection media type instance.
func (mt ComJossemargtSaoDraftLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoDraftLinksArray contains links to related resources of ComJossemargtSaoDraftCollection.
type ComJossemargtSaoDraftLinksArray []*ComJossemargtSaoDraftLinks

// Validate validates the ComJossemargtSaoDraftLinksArray type instance.
func (ut ComJossemargtSaoDraftLinksArray) Validate() (err error) {
	for _, e := range ut {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeComJossemargtSaoDraftCollection decodes the ComJossemargtSaoDraftCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftCollection(resp *http.Response) (ComJossemargtSaoDraftCollection, error) {
	var decoded ComJossemargtSaoDraftCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoDraftFullCollection decodes the ComJossemargtSaoDraftFullCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftFullCollection(resp *http.Response) (ComJossemargtSaoDraftFullCollection, error) {
	var decoded ComJossemargtSaoDraftFullCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoDraftLinkCollection decodes the ComJossemargtSaoDraftLinkCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoDraftLinkCollection(resp *http.Response) (ComJossemargtSaoDraftLinkCollection, error) {
	var decoded ComJossemargtSaoDraftLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// Any source code or input to be compiled, executed and evaluated (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; view=default
type ComJossemargtSaoEntry struct {
	// Contest unique and human readable string identifier
	ContestSlug string `form:"contestSlug" json:"contestSlug" yaml:"contestSlug" xml:"contestSlug"`
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Task unique and human readable string identifier
	TaskSlug string `form:"taskSlug" json:"taskSlug" yaml:"taskSlug" xml:"taskSlug"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Token bool `form:"token" json:"token" yaml:"token" xml:"token"`
}

// Validate validates the ComJossemargtSaoEntry media type instance.
func (mt *ComJossemargtSaoEntry) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.ContestSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.contestSlug`, mt.ContestSlug, `[_a-zA-Z0-9\-]+`))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.TaskSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.taskSlug`, mt.TaskSlug, `[_a-zA-Z0-9\-]+`))
	}
	return
}

// Any source code or input to be compiled, executed and evaluated (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; view=full
type ComJossemargtSaoEntryFull struct {
	// Contest ID where this Entry has been submitted
	ContestID int `form:"contestID" json:"contestID" yaml:"contestID" xml:"contestID"`
	// Contest unique and human readable string identifier
	ContestSlug string `form:"contestSlug" json:"contestSlug" yaml:"contestSlug" xml:"contestSlug"`
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
	// Links to related resources
	Links *ComJossemargtSaoEntryLinks `form:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty" xml:"links,omitempty"`
	// Task ID where this Entry has been submitted
	TaskID int `form:"taskID" json:"taskID" yaml:"taskID" xml:"taskID"`
	// Task unique and human readable string identifier
	TaskSlug string `form:"taskSlug" json:"taskSlug" yaml:"taskSlug" xml:"taskSlug"`
	// Identifies when an Entry has been processed using a CMS Entry Token. The default value is true, in other words
	// 		any submitted Entry will use a CMS Token
	Token bool `form:"token" json:"token" yaml:"token" xml:"token"`
	// User ID of the Entry's owner
	UserID int `form:"userID" json:"userID" yaml:"userID" xml:"userID"`
}

// Validate validates the ComJossemargtSaoEntryFull media type instance.
func (mt *ComJossemargtSaoEntryFull) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.ContestSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.contestSlug`, mt.ContestSlug, `[_a-zA-Z0-9\-]+`))
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, mt.TaskSlug); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.taskSlug`, mt.TaskSlug, `[_a-zA-Z0-9\-]+`))
	}
	return
}

// Any source code or input to be compiled, executed and evaluated (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; view=link
type ComJossemargtSaoEntryLink struct {
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID int `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the ComJossemargtSaoEntryLink media type instance.
func (mt *ComJossemargtSaoEntryLink) Validate() (err error) {

	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// ComJossemargtSaoEntryLinks contains links to related resources of ComJossemargtSaoEntry.
type ComJossemargtSaoEntryLinks struct {
	Result *ComJossemargtSaoResultLink `form:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty" xml:"result,omitempty"`
}

// Validate validates the ComJossemargtSaoEntryLinks type instance.
func (ut *ComJossemargtSaoEntryLinks) Validate() (err error) {
	if ut.Result != nil {
		if err2 := ut.Result.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// DecodeComJossemargtSaoEntry decodes the ComJossemargtSaoEntry instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoEntry(resp *http.Response) (*ComJossemargtSaoEntry, error) {
	var decoded ComJossemargtSaoEntry
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoEntryFull decodes the ComJossemargtSaoEntryFull instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoEntryFull(resp *http.Response) (*ComJossemargtSaoEntryFull, error) {
	var decoded ComJossemargtSaoEntryFull
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoEntryLink decodes the ComJossemargtSaoEntryLink instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoEntryLink(resp *http.Response) (*ComJossemargtSaoEntryLink, error) {
	var decoded ComJossemargtSaoEntryLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ComJossemargtSaoEntryCollection is the media type for an array of ComJossemargtSaoEntry (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; type=collection; view=default
type ComJossemargtSaoEntryCollection []*ComJossemargtSaoEntry

// Validate validates the ComJossemargtSaoEntryCollection media type instance.
func (mt ComJossemargtSaoEntryCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoEntryCollection is the media type for an array of ComJossemargtSaoEntry (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; type=collection; view=full
type ComJossemargtSaoEntryFullCollection []*ComJossemargtSaoEntryFull

// Validate validates the ComJossemargtSaoEntryFullCollection media type instance.
func (mt ComJossemargtSaoEntryFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoEntryCollection is the media type for an array of ComJossemargtSaoEntry (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; type=collection; view=link
type ComJossemargtSaoEntryLinkCollection []*ComJossemargtSaoEntryLink

// Validate validates the ComJossemargtSaoEntryLinkCollection media type instance.
func (mt ComJossemargtSaoEntryLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoEntryLinksArray contains links to related resources of ComJossemargtSaoEntryCollection.
type ComJossemargtSaoEntryLinksArray []*ComJossemargtSaoEntryLinks

// Validate validates the ComJossemargtSaoEntryLinksArray type instance.
func (ut ComJossemargtSaoEntryLinksArray) Validate() (err error) {
	for _, e := range ut {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeComJossemargtSaoEntryCollection decodes the ComJossemargtSaoEntryCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoEntryCollection(resp *http.Response) (ComJossemargtSaoEntryCollection, error) {
	var decoded ComJossemargtSaoEntryCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoEntryFullCollection decodes the ComJossemargtSaoEntryFullCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoEntryFullCollection(resp *http.Response) (ComJossemargtSaoEntryFullCollection, error) {
	var decoded ComJossemargtSaoEntryFullCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoEntryLinkCollection decodes the ComJossemargtSaoEntryLinkCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoEntryLinkCollection(resp *http.Response) (ComJossemargtSaoEntryLinkCollection, error) {
	var decoded ComJossemargtSaoEntryLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// The representation of the result of an entry compile, evaluation and grading process (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; view=default
type ComJossemargtSaoResult struct {
	// Entry evaluation result
	Evaluation *EvaluationResult `form:"evaluation" json:"evaluation" yaml:"evaluation" xml:"evaluation"`
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Compound Result ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the ComJossemargtSaoResult media type instance.
func (mt *ComJossemargtSaoResult) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Evaluation == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "evaluation"))
	}
	if mt.Evaluation != nil {
		if err2 := mt.Evaluation.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// The representation of the result of an entry compile, evaluation and grading process (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; view=full
type ComJossemargtSaoResultFull struct {
	// Entry compilation result
	Compilation *CompilationResult `form:"compilation,omitempty" json:"compilation,omitempty" yaml:"compilation,omitempty" xml:"compilation,omitempty"`
	// Entry evaluation result
	Evaluation *EvaluationResult `form:"evaluation" json:"evaluation" yaml:"evaluation" xml:"evaluation"`
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Compound Result ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Links to related resources
	Links *ComJossemargtSaoResultLinks `form:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty" xml:"links,omitempty"`
	// Entry graded score
	Score *ScoreResult `form:"score,omitempty" json:"score,omitempty" yaml:"score,omitempty" xml:"score,omitempty"`
}

// Validate validates the ComJossemargtSaoResultFull media type instance.
func (mt *ComJossemargtSaoResultFull) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.Evaluation == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "evaluation"))
	}
	if mt.Compilation != nil {
		if err2 := mt.Compilation.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Evaluation != nil {
		if err2 := mt.Evaluation.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// The representation of the result of an entry compile, evaluation and grading process (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; view=link
type ComJossemargtSaoResultLink struct {
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Compound Result ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the ComJossemargtSaoResultLink media type instance.
func (mt *ComJossemargtSaoResultLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// ComJossemargtSaoResultLinks contains links to related resources of ComJossemargtSaoResult.
type ComJossemargtSaoResultLinks struct {
}

// DecodeComJossemargtSaoResult decodes the ComJossemargtSaoResult instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoResult(resp *http.Response) (*ComJossemargtSaoResult, error) {
	var decoded ComJossemargtSaoResult
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoResultFull decodes the ComJossemargtSaoResultFull instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoResultFull(resp *http.Response) (*ComJossemargtSaoResultFull, error) {
	var decoded ComJossemargtSaoResultFull
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeComJossemargtSaoResultLink decodes the ComJossemargtSaoResultLink instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoResultLink(resp *http.Response) (*ComJossemargtSaoResultLink, error) {
	var decoded ComJossemargtSaoResultLink
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ComJossemargtSaoResultCollection is the media type for an array of ComJossemargtSaoResult (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; type=collection; view=default
type ComJossemargtSaoResultCollection []*ComJossemargtSaoResult

// Validate validates the ComJossemargtSaoResultCollection media type instance.
func (mt ComJossemargtSaoResultCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoResultCollection is the media type for an array of ComJossemargtSaoResult (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; type=collection; view=full
type ComJossemargtSaoResultFullCollection []*ComJossemargtSaoResultFull

// Validate validates the ComJossemargtSaoResultFullCollection media type instance.
func (mt ComJossemargtSaoResultFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoResultCollection is the media type for an array of ComJossemargtSaoResult (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; type=collection; view=link
type ComJossemargtSaoResultLinkCollection []*ComJossemargtSaoResultLink

// Validate validates the ComJossemargtSaoResultLinkCollection media type instance.
func (mt ComJossemargtSaoResultLinkCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// ComJossemargtSaoResultLinksArray contains links to related resources of ComJossemargtSaoResultCollection.
type ComJossemargtSaoResultLinksArray []*ComJossemargtSaoResultLinks

// DecodeComJossemargtSaoResultCollection decodes the ComJossemargtSaoResultCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoResultCollection(resp *http.Response) (ComJossemargtSaoResultCollection, error) {
	var decoded ComJossemargtSaoResultCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoResultFullCollection decodes the ComJossemargtSaoResultFullCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoResultFullCollection(resp *http.Response) (ComJossemargtSaoResultFullCollection, error) {
	var decoded ComJossemargtSaoResultFullCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeComJossemargtSaoResultLinkCollection decodes the ComJossemargtSaoResultLinkCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoResultLinkCollection(resp *http.Response) (ComJossemargtSaoResultLinkCollection, error) {
	var decoded ComJossemargtSaoResultLinkCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// The representation of a summarized entry's score (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.score-sum+json; view=default
type ComJossemargtSaoScoreSum struct {
	// Contest Identifier associated with this score
	ContestID *int `form:"contestID,omitempty" json:"contestID,omitempty" yaml:"contestID,omitempty" xml:"contestID,omitempty"`
	// The graded value relative to the contest
	ContestValue float64 `form:"contestValue" json:"contestValue" yaml:"contestValue" xml:"contestValue"`
	// Contest Identifier associated with this score
	TaskID *int `form:"taskID,omitempty" json:"taskID,omitempty" yaml:"taskID,omitempty" xml:"taskID,omitempty"`
	// The graded value relative to the Task total score
	TaskValue float64 `form:"taskValue" json:"taskValue" yaml:"taskValue" xml:"taskValue"`
	// Contest Identifier associated with this score
	UserID *int `form:"userID,omitempty" json:"userID,omitempty" yaml:"userID,omitempty" xml:"userID,omitempty"`
}

// Validate validates the ComJossemargtSaoScoreSum media type instance.
func (mt *ComJossemargtSaoScoreSum) Validate() (err error) {

	if mt.ContestID != nil {
		if *mt.ContestID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.contestID`, *mt.ContestID, 1, true))
		}
	}
	if mt.TaskID != nil {
		if *mt.TaskID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.taskID`, *mt.TaskID, 1, true))
		}
	}
	if mt.UserID != nil {
		if *mt.UserID < 1 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`response.userID`, *mt.UserID, 1, true))
		}
	}
	return
}

// DecodeComJossemargtSaoScoreSum decodes the ComJossemargtSaoScoreSum instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoScoreSum(resp *http.Response) (*ComJossemargtSaoScoreSum, error) {
	var decoded ComJossemargtSaoScoreSum
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ComJossemargtSaoScore-SumCollection is the media type for an array of ComJossemargtSaoScore-Sum (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.score-sum+json; type=collection; view=default
type ComJossemargtSaoScoreSumCollection []*ComJossemargtSaoScoreSum

// Validate validates the ComJossemargtSaoScoreSumCollection media type instance.
func (mt ComJossemargtSaoScoreSumCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// DecodeComJossemargtSaoScoreSumCollection decodes the ComJossemargtSaoScoreSumCollection instance encoded in resp body.
func (c *Client) DecodeComJossemargtSaoScoreSumCollection(resp *http.Response) (ComJossemargtSaoScoreSumCollection, error) {
	var decoded ComJossemargtSaoScoreSumCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeErrorResponse decodes the ErrorResponse instance encoded in resp body.
func (c *Client) DecodeErrorResponse(resp *http.Response) (*goa.ErrorResponse, error) {
	var decoded goa.ErrorResponse
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
