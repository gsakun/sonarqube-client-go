// Read and update Security Hotspots.
package sonar

import "net/http"

type HotspotsService struct {
	client *Client
}

type HotspotsEditCommentObject struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type HotspotsSearchObject struct {
	Components []HotspotsSearchObject_sub1 `json:"components,omitempty"`
	Hotspots   []HotspotsSearchObject_sub2 `json:"hotspots,omitempty"`
	Paging     HotspotsSearchObject_sub3   `json:"paging,omitempty"`
}

type HotspotsSearchObject_sub2 struct {
	Assignee                 string        `json:"assignee,omitempty"`
	Author                   string        `json:"author,omitempty"`
	Component                string        `json:"component,omitempty"`
	CreationDate             string        `json:"creationDate,omitempty"`
	Flows                    []interface{} `json:"flows,omitempty"`
	Key                      string        `json:"key,omitempty"`
	Line                     int64         `json:"line,omitempty"`
	Message                  string        `json:"message,omitempty"`
	MessageFormattings       []interface{} `json:"messageFormattings,omitempty"`
	Project                  string        `json:"project,omitempty"`
	RuleKey                  string        `json:"ruleKey,omitempty"`
	SecurityCategory         string        `json:"securityCategory,omitempty"`
	Status                   string        `json:"status,omitempty"`
	UpdateDate               string        `json:"updateDate,omitempty"`
	VulnerabilityProbability string        `json:"vulnerabilityProbability,omitempty"`
}

type HotspotsSearchObject_sub1 struct {
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type HotspotsSearchObject_sub3 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type HotspotsShowObject struct {
	Assignee           string                    `json:"assignee,omitempty"`
	Author             string                    `json:"author,omitempty"`
	CanChangeStatus    bool                      `json:"canChangeStatus,omitempty"`
	Changelog          []HotspotsShowObject_sub2 `json:"changelog,omitempty"`
	Comment            []HotspotsShowObject_sub3 `json:"comment,omitempty"`
	Component          HotspotsShowObject_sub4   `json:"component,omitempty"`
	CreationDate       string                    `json:"creationDate,omitempty"`
	Hash               string                    `json:"hash,omitempty"`
	Key                string                    `json:"key,omitempty"`
	Line               int64                     `json:"line,omitempty"`
	Message            string                    `json:"message,omitempty"`
	MessageFormattings []HotspotsShowObject_sub5 `json:"messageFormattings,omitempty"`
	Project            HotspotsShowObject_sub6   `json:"project,omitempty"`
	Rule               HotspotsShowObject_sub7   `json:"rule,omitempty"`
	Status             string                    `json:"status,omitempty"`
	UpdateDate         string                    `json:"updateDate,omitempty"`
	Users              []HotspotsShowObject_sub8 `json:"users,omitempty"`
}

type HotspotsShowObject_sub8 struct {
	Active bool   `json:"active,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type HotspotsShowObject_sub2 struct {
	Avatar       string                    `json:"avatar,omitempty"`
	CreationDate string                    `json:"creationDate,omitempty"`
	Diffs        []HotspotsShowObject_sub1 `json:"diffs,omitempty"`
	IsUserActive bool                      `json:"isUserActive,omitempty"`
	User         string                    `json:"user,omitempty"`
	UserName     string                    `json:"userName,omitempty"`
}

type HotspotsShowObject_sub3 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
}

type HotspotsShowObject_sub5 struct {
	End   int64  `json:"end,omitempty"`
	Start int64  `json:"start,omitempty"`
	Type  string `json:"type,omitempty"`
}

type HotspotsShowObject_sub4 struct {
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type HotspotsShowObject_sub6 struct {
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type HotspotsShowObject_sub7 struct {
	Key                      string `json:"key,omitempty"`
	Name                     string `json:"name,omitempty"`
	SecurityCategory         string `json:"securityCategory,omitempty"`
	VulnerabilityProbability string `json:"vulnerabilityProbability,omitempty"`
}

type HotspotsShowObject_sub1 struct {
	Key      string `json:"key,omitempty"`
	NewValue string `json:"newValue,omitempty"`
	OldValue string `json:"oldValue,omitempty"`
}

type HotspotsAddCommentOption struct {
	Comment string `url:"comment,omitempty"` // Description:"Comment text.",ExampleValue:"This is safe because user input is validated by the calling code"
	Hotspot string `url:"hotspot,omitempty"` // Description:"Key of the Security Hotspot",ExampleValue:"AU-TpxcA-iU5OvuD2FL0"
}

// AddComment Add a comment to a Security Hotpot.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified Security Hotspot.
func (s *HotspotsService) AddComment(opt *HotspotsAddCommentOption) (resp *http.Response, err error) {
	err = s.ValidateAddCommentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "hotspots/add_comment", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type HotspotsAssignOption struct {
	Assignee string `url:"assignee,omitempty"` // Description:"Login of the assignee with 'Browse' project permission",ExampleValue:"admin"
	Comment  string `url:"comment,omitempty"`  // Description:"A comment provided with assign action",ExampleValue:"Hey Bob! Could you please have a look and confirm my assertion that we are safe here, please"
	Hotspot  string `url:"hotspot,omitempty"`  // Description:"Hotspot key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// Assign Assign a hotspot to an active user. Requires authentication and Browse permission on project
func (s *HotspotsService) Assign(opt *HotspotsAssignOption) (resp *http.Response, err error) {
	err = s.ValidateAssignOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "hotspots/assign", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type HotspotsChangeStatusOption struct {
	Comment    string `url:"comment,omitempty"`    // Description:"Comment text.",ExampleValue:"This is safe because user input is validated by the calling code"
	Hotspot    string `url:"hotspot,omitempty"`    // Description:"Key of the Security Hotspot",ExampleValue:"AU-TpxcA-iU5OvuD2FL0"
	Resolution string `url:"resolution,omitempty"` // Description:"Resolution of the Security Hotspot when new status is REVIEWED, otherwise must not be set.",ExampleValue:""
	Status     string `url:"status,omitempty"`     // Description:"New status of the Security Hotspot.",ExampleValue:""
}

// ChangeStatus Change the status of a Security Hotpot.<br/>Requires the 'Administer Security Hotspot' permission.
func (s *HotspotsService) ChangeStatus(opt *HotspotsChangeStatusOption) (resp *http.Response, err error) {
	err = s.ValidateChangeStatusOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "hotspots/change_status", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type HotspotsDeleteCommentOption struct {
	Comment string `url:"comment,omitempty"` // Description:"Comment key.",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// DeleteComment Delete comment from Security Hotpot.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified Security Hotspot.
func (s *HotspotsService) DeleteComment(opt *HotspotsDeleteCommentOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteCommentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "hotspots/delete_comment", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type HotspotsEditCommentOption struct {
	Comment string `url:"comment,omitempty"` // Description:"Comment key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Text    string `url:"text,omitempty"`    // Description:"Comment text",ExampleValue:"Safe because it doesn't apply to the context"
}

// EditComment Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified hotspot.
func (s *HotspotsService) EditComment(opt *HotspotsEditCommentOption) (v *HotspotsEditCommentObject, resp *http.Response, err error) {
	err = s.ValidateEditCommentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "hotspots/edit_comment", opt)
	if err != nil {
		return
	}
	v = new(HotspotsEditCommentObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type HotspotsSearchOption struct {
	Branch              string `url:"branch,omitempty"`              // Description:"Branch key. Not available in the community edition.",ExampleValue:"feature/my_branch"
	Cwe                 string `url:"cwe,omitempty"`                 // Description:"Comma-separated list of CWE numbers",ExampleValue:"89,434,352"
	Files               string `url:"files,omitempty"`               // Description:"Comma-separated list of files. Returns only hotspots found in those files",ExampleValue:"src/main/java/org/sonar/server/Test.java"
	Hotspots            string `url:"hotspots,omitempty"`            // Description:"Comma-separated list of Security Hotspot keys. This parameter is required unless projectKey is provided.",ExampleValue:"AWhXpLoInp4On-Y3xc8x"
	InNewCodePeriod     string `url:"inNewCodePeriod,omitempty"`     // Description:"If 'inNewCodePeriod' is provided, only Security Hotspots created in the new code period are returned.",ExampleValue:""
	OnlyMine            string `url:"onlyMine,omitempty"`            // Description:"If 'projectKey' is provided, returns only Security Hotspots assigned to the current user",ExampleValue:""
	OwaspAsvs40         string `url:"owaspAsvs-4.0,omitempty"`       // Description:"Comma-separated list of OWASP ASVS v4.0 categories or rules.",ExampleValue:"6,6.1.2"
	OwaspAsvsLevel      string `url:"owaspAsvsLevel,omitempty"`      // Description:"Filters hotspots with lower or equal OWASP ASVS level to the parameter value. Should be used in combination with the 'owaspAsvs-4.0' parameter.",ExampleValue:"2"
	OwaspTop10          string `url:"owaspTop10,omitempty"`          // Description:"Comma-separated list of OWASP 2017 Top 10 lowercase categories.",ExampleValue:""
	OwaspTop102021      string `url:"owaspTop10-2021,omitempty"`     // Description:"Comma-separated list of OWASP 2021 Top 10 lowercase categories.",ExampleValue:""
	P                   string `url:"p,omitempty"`                   // Description:"1-based page number",ExampleValue:"42"
	PciDss32            string `url:"pciDss-3.2,omitempty"`          // Description:"Comma-separated list of PCI DSS v3.2 categories.",ExampleValue:"4,6.5.8,10.1"
	PciDss40            string `url:"pciDss-4.0,omitempty"`          // Description:"Comma-separated list of PCI DSS v4.0 categories.",ExampleValue:"4,6.5.8,10.1"
	ProjectKey          string `url:"projectKey,omitempty"`          // Description:"Key of the project or application. This parameter is required unless hotspots is provided.",ExampleValue:"my_project"
	Ps                  string `url:"ps,omitempty"`                  // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	PullRequest         string `url:"pullRequest,omitempty"`         // Description:"Pull request id. Not available in the community edition.",ExampleValue:"5461"
	Resolution          string `url:"resolution,omitempty"`          // Description:"If 'projectKey' is provided and if status is 'REVIEWED', only Security Hotspots with the specified resolution are returned.",ExampleValue:""
	SansTop25           string `url:"sansTop25,omitempty"`           // Description:"Comma-separated list of SANS Top 25 categories.",ExampleValue:""
	SonarsourceSecurity string `url:"sonarsourceSecurity,omitempty"` // Description:"Comma-separated list of SonarSource security categories. Use 'others' to select issues not associated with any category",ExampleValue:""
	Status              string `url:"status,omitempty"`              // Description:"If 'projectKey' is provided, only Security Hotspots with the specified status are returned.",ExampleValue:""
}

// Search Search for Security Hotpots. <br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects. <br>When issue indexation is in progress returns 503 service unavailable HTTP code.
func (s *HotspotsService) Search(opt *HotspotsSearchOption) (v *HotspotsSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "hotspots/search", opt)
	if err != nil {
		return
	}
	v = new(HotspotsSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type HotspotsShowOption struct {
	Hotspot string `url:"hotspot,omitempty"` // Description:"Key of the Security Hotspot",ExampleValue:"AU-TpxcA-iU5OvuD2FL0"
}

// Show Provides the details of a Security Hotspot.
func (s *HotspotsService) Show(opt *HotspotsShowOption) (v *HotspotsShowObject, resp *http.Response, err error) {
	err = s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "hotspots/show", opt)
	if err != nil {
		return
	}
	v = new(HotspotsShowObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
