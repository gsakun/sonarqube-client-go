// Manage project analyses.
package sonar

import "net/http"

type ProjectAnalysesService struct {
	client *Client
}

type ProjectAnalysesCreateEventObject struct {
	Event ProjectAnalysesCreateEventObject_sub1 `json:"event,omitempty"`
}

type ProjectAnalysesCreateEventObject_sub1 struct {
	Analysis string `json:"analysis,omitempty"`
	Category string `json:"category,omitempty"`
	Key      string `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ProjectAnalysesSearchObject struct {
	Analyses []ProjectAnalysesSearchObject_sub4 `json:"analyses,omitempty"`
	Paging   ProjectAnalysesSearchObject_sub5   `json:"paging,omitempty"`
}

type ProjectAnalysesSearchObject_sub1 struct {
	Branch string `json:"branch,omitempty"`
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
}

type ProjectAnalysesSearchObject_sub4 struct {
	BuildString                 string                             `json:"buildString,omitempty"`
	Date                        string                             `json:"date,omitempty"`
	DetectedCI                  string                             `json:"detectedCI,omitempty"`
	Events                      []ProjectAnalysesSearchObject_sub3 `json:"events,omitempty"`
	Key                         string                             `json:"key,omitempty"`
	ManualNewCodePeriodBaseline bool                               `json:"manualNewCodePeriodBaseline,omitempty"`
	ProjectVersion              string                             `json:"projectVersion,omitempty"`
	Revision                    string                             `json:"revision,omitempty"`
}

type ProjectAnalysesSearchObject_sub3 struct {
	Category    string                           `json:"category,omitempty"`
	Description string                           `json:"description,omitempty"`
	Key         string                           `json:"key,omitempty"`
	Name        string                           `json:"name,omitempty"`
	QualityGate ProjectAnalysesSearchObject_sub2 `json:"qualityGate,omitempty"`
}

type ProjectAnalysesSearchObject_sub2 struct {
	Failing      []ProjectAnalysesSearchObject_sub1 `json:"failing,omitempty"`
	Status       string                             `json:"status,omitempty"`
	StillFailing bool                               `json:"stillFailing,omitempty"`
}

type ProjectAnalysesSearchObject_sub5 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type ProjectAnalysesUpdateEventObject struct {
	Event ProjectAnalysesUpdateEventObject_sub1 `json:"event,omitempty"`
}

type ProjectAnalysesUpdateEventObject_sub1 struct {
	Analysis string `json:"analysis,omitempty"`
	Category string `json:"category,omitempty"`
	Key      string `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
}

type ProjectAnalysesCreateEventOption struct {
	Analysis string `url:"analysis,omitempty"` // Description:"Analysis key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Category string `url:"category,omitempty"` // Description:"Category",ExampleValue:""
	Name     string `url:"name,omitempty"`     // Description:"Name",ExampleValue:"5.6"
}

// CreateEvent Create a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be created.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) CreateEvent(opt *ProjectAnalysesCreateEventOption) (v *ProjectAnalysesCreateEventObject, resp *http.Response, err error) {
	err = s.ValidateCreateEventOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/create_event", opt)
	if err != nil {
		return
	}
	v = new(ProjectAnalysesCreateEventObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectAnalysesDeleteOption struct {
	Analysis string `url:"analysis,omitempty"` // Description:"Analysis key",ExampleValue:"AU-TpxcA-iU5OvuD2FL1"
}

// Delete Delete a project analysis.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the project of the specified analysis</li></ul>
func (s *ProjectAnalysesService) Delete(opt *ProjectAnalysesDeleteOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/delete", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesDeleteEventOption struct {
	Event string `url:"event,omitempty"` // Description:"Event key",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
}

// DeleteEvent Delete a project analysis event.<br>Only event of category 'VERSION' and 'OTHER' can be deleted.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) DeleteEvent(opt *ProjectAnalysesDeleteEventOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteEventOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/delete_event", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesSearchOption struct {
	Branch   string `url:"branch,omitempty"`   // Description:"Key of a branch",ExampleValue:"feature/my_branch"
	Category string `url:"category,omitempty"` // Description:"Event category. Filter analyses that have at least one event of the category specified.",ExampleValue:"OTHER"
	From     string `url:"from,omitempty"`     // Description:"Filter analyses created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided",ExampleValue:"2013-05-01"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Project  string `url:"project,omitempty"`  // Description:"Project key",ExampleValue:"my_project"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Statuses string `url:"statuses,omitempty"` // Description:"List of statuses of desired analysis.",ExampleValue:"P,U,L"
	To       string `url:"to,omitempty"`       // Description:"Filter analyses created before the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
}

// Search Search a project analyses and attached events.<br>Requires the following permission: 'Browse' on the specified project. <br>For applications, it also requires 'Browse' permission on its child projects.
func (s *ProjectAnalysesService) Search(opt *ProjectAnalysesSearchOption) (v *ProjectAnalysesSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_analyses/search", opt)
	if err != nil {
		return
	}
	v = new(ProjectAnalysesSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectAnalysesSetBaselineOption struct {
	Analysis string `url:"analysis,omitempty"` // Description:"Analysis key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Branch   string `url:"branch,omitempty"`   // Description:"Branch key",ExampleValue:""
	Project  string `url:"project,omitempty"`  // Description:"Project key",ExampleValue:""
}

// SetBaseline Set an analysis as the baseline of the New Code Period on a project or a branch.<br/>This manually set baseline.<br/>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) SetBaseline(opt *ProjectAnalysesSetBaselineOption) (resp *http.Response, err error) {
	err = s.ValidateSetBaselineOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/set_baseline", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesUnsetBaselineOption struct {
	Branch  string `url:"branch,omitempty"`  // Description:"Branch key",ExampleValue:""
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:""
}

// UnsetBaseline Unset any manually-set New Code Period baseline on a project or a branch.<br/>Unsetting a manual baseline restores the use of the default new code period setting.<br/>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) UnsetBaseline(opt *ProjectAnalysesUnsetBaselineOption) (resp *http.Response, err error) {
	err = s.ValidateUnsetBaselineOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/unset_baseline", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectAnalysesUpdateEventOption struct {
	Event string `url:"event,omitempty"` // Description:"Event key",ExampleValue:"AU-TpxcA-iU5OvuD2FL5"
	Name  string `url:"name,omitempty"`  // Description:"New name",ExampleValue:"5.6"
}

// UpdateEvent Update a project analysis event.<br>Only events of category 'VERSION' and 'OTHER' can be updated.<br>Requires one of the following permissions:<ul>  <li>'Administer System'</li>  <li>'Administer' rights on the specified project</li></ul>
func (s *ProjectAnalysesService) UpdateEvent(opt *ProjectAnalysesUpdateEventOption) (v *ProjectAnalysesUpdateEventObject, resp *http.Response, err error) {
	err = s.ValidateUpdateEventOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_analyses/update_event", opt)
	if err != nil {
		return
	}
	v = new(ProjectAnalysesUpdateEventObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
