// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "SAO v1": Application Media Types
//
// Command:
// $ goagen
// --design=github.com/jossemargt/cms-sao/design
// --force=true
// --notool=true
// --out=$(GOPATH)/src/github.com/jossemargt/cms-sao
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
)

// Any source code or input to be compiled, executed and evaluated (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; view=default
type ComJossemargtSaoEntry struct {
	// Contest unique and human readble string identifier
	ContestSlug *string `form:"contestSlug,omitempty" json:"contestSlug,omitempty" yaml:"contestSlug,omitempty" xml:"contestSlug,omitempty"`
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Indenties if the entry should be ranked or taken as an user test
	Ranked bool `form:"ranked" json:"ranked" yaml:"ranked" xml:"ranked"`
	// Task unique and human readble string identifier
	TaskSlug *string `form:"taskSlug,omitempty" json:"taskSlug,omitempty" yaml:"taskSlug,omitempty" xml:"taskSlug,omitempty"`
}

// Validate validates the ComJossemargtSaoEntry media type instance.
func (mt *ComJossemargtSaoEntry) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.ContestSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *mt.ContestSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.contestSlug`, *mt.ContestSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	if mt.TaskSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *mt.TaskSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.taskSlug`, *mt.TaskSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	return
}

// Any source code or input to be compiled, executed and evaluated (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.entry+json; view=full
type ComJossemargtSaoEntryFull struct {
	// Contest unique and human readble string identifier
	ContestSlug *string `form:"contestSlug,omitempty" json:"contestSlug,omitempty" yaml:"contestSlug,omitempty" xml:"contestSlug,omitempty"`
	// API href for making requests on the entry
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique entry ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Links to related resources
	Links *ComJossemargtSaoEntryLinks `form:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty" xml:"links,omitempty"`
	// Indenties if the entry should be ranked or taken as an user test
	Ranked bool `form:"ranked" json:"ranked" yaml:"ranked" xml:"ranked"`
	// Task unique and human readble string identifier
	TaskSlug *string `form:"taskSlug,omitempty" json:"taskSlug,omitempty" yaml:"taskSlug,omitempty" xml:"taskSlug,omitempty"`
}

// Validate validates the ComJossemargtSaoEntryFull media type instance.
func (mt *ComJossemargtSaoEntryFull) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	if mt.ContestSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *mt.ContestSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.contestSlug`, *mt.ContestSlug, `[_a-zA-Z0-9\-]+`))
		}
	}
	if mt.Links != nil {
		if err2 := mt.Links.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.TaskSlug != nil {
		if ok := goa.ValidatePattern(`[_a-zA-Z0-9\-]+`, *mt.TaskSlug); !ok {
			err = goa.MergeErrors(err, goa.InvalidPatternError(`response.taskSlug`, *mt.TaskSlug, `[_a-zA-Z0-9\-]+`))
		}
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
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
}

// Validate validates the ComJossemargtSaoEntryLink media type instance.
func (mt *ComJossemargtSaoEntryLink) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Href == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "href"))
	}
	return
}

// ComJossemargtSaoEntryLinks contains links to related resources of ComJossemargtSaoEntry.
type ComJossemargtSaoEntryLinks struct {
	Result *ComJossemargtSaoResultLink `form:"result,omitempty" json:"result,omitempty" yaml:"result,omitempty" xml:"result,omitempty"`
	Score  *ComJossemargtSaoScoreLink  `form:"score,omitempty" json:"score,omitempty" yaml:"score,omitempty" xml:"score,omitempty"`
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

// The representation of the result of an entry compile, execute or evaluation process (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; view=default
type ComJossemargtSaoResult struct {
	// Entry evaluation result
	Evaluation *EvaluationResult `form:"evaluation" json:"evaluation" yaml:"evaluation" xml:"evaluation"`
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique result ID
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

// The representation of the result of an entry compile, execute or evaluation process (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; view=full
type ComJossemargtSaoResultFull struct {
	// Entry compilation result
	Compilation *CompilationResult `form:"compilation,omitempty" json:"compilation,omitempty" yaml:"compilation,omitempty" xml:"compilation,omitempty"`
	// Entry evaluation result
	Evaluation *EvaluationResult `form:"evaluation" json:"evaluation" yaml:"evaluation" xml:"evaluation"`
	// Entry execution result
	Execution []*ExecutionResult `form:"execution,omitempty" json:"execution,omitempty" yaml:"execution,omitempty" xml:"execution,omitempty"`
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique result ID
	ID string `form:"id" json:"id" yaml:"id" xml:"id"`
	// Links to related resources
	Links *ComJossemargtSaoResultLinks `form:"links,omitempty" json:"links,omitempty" yaml:"links,omitempty" xml:"links,omitempty"`
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

// The representation of the result of an entry compile, execute or evaluation process (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.result+json; view=link
type ComJossemargtSaoResultLink struct {
	// API href for making requests on the result
	Href string `form:"href" json:"href" yaml:"href" xml:"href"`
	// Unique result ID
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
	Score *ComJossemargtSaoScoreLink `form:"score,omitempty" json:"score,omitempty" yaml:"score,omitempty" xml:"score,omitempty"`
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

// The representation of the entry's scoring after being evaluated (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.score+json; view=default
type ComJossemargtSaoScore struct {
	// API href for making requests on the score
	Href *string `form:"href,omitempty" json:"href,omitempty" yaml:"href,omitempty" xml:"href,omitempty"`
	// Unique score ID
	ID *string `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// An official graded score with a token
	Value *float64 `form:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty" xml:"value,omitempty"`
}

// The representation of the entry's scoring after being evaluated (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.score+json; view=full
type ComJossemargtSaoScoreFull struct {
	// API href for making requests on the score
	Href *string `form:"href,omitempty" json:"href,omitempty" yaml:"href,omitempty" xml:"href,omitempty"`
	// Unique score ID
	ID *string `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
	// An un-official graded score
	UntokenedValue *float64 `form:"untokenedValue,omitempty" json:"untokenedValue,omitempty" yaml:"untokenedValue,omitempty" xml:"untokenedValue,omitempty"`
	// An official graded score with a token
	Value *float64 `form:"value,omitempty" json:"value,omitempty" yaml:"value,omitempty" xml:"value,omitempty"`
}

// The representation of the entry's scoring after being evaluated (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.score+json; view=link
type ComJossemargtSaoScoreLink struct {
	// API href for making requests on the score
	Href *string `form:"href,omitempty" json:"href,omitempty" yaml:"href,omitempty" xml:"href,omitempty"`
	// Unique score ID
	ID *string `form:"id,omitempty" json:"id,omitempty" yaml:"id,omitempty" xml:"id,omitempty"`
}

// ComJossemargtSaoScoreCollection is the media type for an array of ComJossemargtSaoScore (default view)
//
// Identifier: application/vnd.com.jossemargt.sao.score+json; type=collection; view=default
type ComJossemargtSaoScoreCollection []*ComJossemargtSaoScore

// ComJossemargtSaoScoreCollection is the media type for an array of ComJossemargtSaoScore (full view)
//
// Identifier: application/vnd.com.jossemargt.sao.score+json; type=collection; view=full
type ComJossemargtSaoScoreFullCollection []*ComJossemargtSaoScoreFull

// ComJossemargtSaoScoreCollection is the media type for an array of ComJossemargtSaoScore (link view)
//
// Identifier: application/vnd.com.jossemargt.sao.score+json; type=collection; view=link
type ComJossemargtSaoScoreLinkCollection []*ComJossemargtSaoScoreLink