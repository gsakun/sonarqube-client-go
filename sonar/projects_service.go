// Manage project existence.
package sonar

import "net/http"

const (
	ProjectVisibilityPublic  = "public"
	ProjectVisibilityPrivate = "private"
)

type ProjectsService struct {
	client *Client
}

type ProjectsCreateObject struct {
	Project ProjectsCreateObject_sub1 `json:"project,omitempty"`
}

type ProjectsCreateObject_sub1 struct {
	Key       string `json:"key,omitempty"`
	Name      string `json:"name,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type ProjectsSearchObject struct {
	Components []ProjectsSearchObject_sub1 `json:"components,omitempty"`
	Paging     ProjectsSearchObject_sub2   `json:"paging,omitempty"`
}

type ProjectsSearchObject_sub1 struct {
	Key              string `json:"key,omitempty"`
	LastAnalysisDate string `json:"lastAnalysisDate,omitempty"`
	Name             string `json:"name,omitempty"`
	Qualifier        string `json:"qualifier,omitempty"`
	Revision         string `json:"revision,omitempty"`
	Visibility       string `json:"visibility,omitempty"`
}

type ProjectsSearchObject_sub2 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type ProjectsSearchMyProjectsObject struct {
	Paging   ProjectsSearchMyProjectsObject_sub1   `json:"paging,omitempty"`
	Projects []ProjectsSearchMyProjectsObject_sub3 `json:"projects,omitempty"`
}

type ProjectsSearchMyProjectsObject_sub3 struct {
	Description      string                                `json:"description,omitempty"`
	Key              string                                `json:"key,omitempty"`
	LastAnalysisDate string                                `json:"lastAnalysisDate,omitempty"`
	Links            []ProjectsSearchMyProjectsObject_sub2 `json:"links,omitempty"`
	Name             string                                `json:"name,omitempty"`
	QualityGate      string                                `json:"qualityGate,omitempty"`
}

type ProjectsSearchMyProjectsObject_sub2 struct {
	Href string `json:"href,omitempty"`
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
}

type ProjectsSearchMyProjectsObject_sub1 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type ProjectsSearchMyScannableProjectsObject struct {
	Projects []ProjectsSearchMyScannableProjectsObject_sub1 `json:"projects,omitempty"`
}

type ProjectsSearchMyScannableProjectsObject_sub1 struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

type ProjectsBulkDeleteOption struct {
	AnalyzedBefore    string `url:"analyzedBefore,omitempty"`    // Description:"Filter the projects for which last analysis of any branch is older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	OnProvisionedOnly string `url:"onProvisionedOnly,omitempty"` // Description:"Filter the projects that are provisioned",ExampleValue:""
	Projects          string `url:"projects,omitempty"`          // Description:"Comma-separated list of project keys",ExampleValue:"my_project,another_project"
	Q                 string `url:"q,omitempty"`                 // Description:"Limit to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>",ExampleValue:"sonar"
	Qualifiers        string `url:"qualifiers,omitempty"`        // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers",ExampleValue:""
	Visibility        string `url:"visibility,omitempty"`        // Description:"Filter the projects that should be visible to everyone (public), or only specific user/groups (private).<br/>If no visibility is specified, the default project visibility will be used.",ExampleValue:""
}

// BulkDelete Delete one or several projects.<br />Only the 1'000 first items in project filters are taken into account.<br />Requires 'Administer System' permission.<br />At least one parameter is required among analyzedBefore, projects and q
func (s *ProjectsService) BulkDelete(opt *ProjectsBulkDeleteOption) (resp *http.Response, err error) {
	err = s.ValidateBulkDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/bulk_delete", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectsCreateOption struct {
	MainBranch string `url:"mainBranch,omitempty"` // Description:"Key of the main branch of the project. If not provided, the default main branch key will be used.",ExampleValue:"develop"
	Name       string `url:"name,omitempty"`       // Description:"Name of the project. If name is longer than 500, it is abbreviated.",ExampleValue:"SonarQube"
	Project    string `url:"project,omitempty"`    // Description:"Key of the project",ExampleValue:"my_project"
	Visibility string `url:"visibility,omitempty"` // Description:"Whether the created project should be visible to everyone, or only specific user/groups.<br/>If no visibility is specified, the default project visibility will be used.",ExampleValue:""
}

// Create Create a project.<br/>Requires 'Create Projects' permission
func (s *ProjectsService) Create(opt *ProjectsCreateOption) (v *ProjectsCreateObject, resp *http.Response, err error) {
	err = s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/create", opt)
	if err != nil {
		return
	}
	v = new(ProjectsCreateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectsDeleteOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Delete Delete a project.<br> Requires 'Administer System' permission or 'Administer' permission on the project.
func (s *ProjectsService) Delete(opt *ProjectsDeleteOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/delete", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectsSearchOption struct {
	AnalyzedBefore    string `url:"analyzedBefore,omitempty"`    // Description:"Filter the projects for which the last analysis of all branches are older than the given date (exclusive).<br> Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	OnProvisionedOnly string `url:"onProvisionedOnly,omitempty"` // Description:"Filter the projects that are provisioned",ExampleValue:""
	P                 string `url:"p,omitempty"`                 // Description:"1-based page number",ExampleValue:"42"
	Projects          string `url:"projects,omitempty"`          // Description:"Comma-separated list of project keys",ExampleValue:"my_project,another_project"
	Ps                string `url:"ps,omitempty"`                // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q                 string `url:"q,omitempty"`                 // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that contain the supplied string</li></ul>",ExampleValue:"sonar"
	Qualifiers        string `url:"qualifiers,omitempty"`        // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers",ExampleValue:""
	Visibility        string `url:"visibility,omitempty"`        // Description:"Filter the projects that should be visible to everyone (public), or only specific user/groups (private).<br/>If no visibility is specified, the default project visibility will be used.",ExampleValue:""
}

// Search Search for projects or views to administrate them.<br>Requires 'Administer System' permission
func (s *ProjectsService) Search(opt *ProjectsSearchOption) (v *ProjectsSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "projects/search", opt)
	if err != nil {
		return
	}
	v = new(ProjectsSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectsSearchMyProjectsOption struct {
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// SearchMyProjects Return list of projects for which the current user has 'Administer' permission. Maximum 1'000 projects are returned.
func (s *ProjectsService) SearchMyProjects(opt *ProjectsSearchMyProjectsOption) (v *ProjectsSearchMyProjectsObject, resp *http.Response, err error) {
	err = s.ValidateSearchMyProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "projects/search_my_projects", opt)
	if err != nil {
		return
	}
	v = new(ProjectsSearchMyProjectsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectsSearchMyScannableProjectsOption struct {
	Q string `url:"q,omitempty"` // Description:"Limit search to project names that contain the supplied string.",ExampleValue:"project"
}

// SearchMyScannableProjects List projects that a user can scan.
func (s *ProjectsService) SearchMyScannableProjects(opt *ProjectsSearchMyScannableProjectsOption) (v *ProjectsSearchMyScannableProjectsObject, resp *http.Response, err error) {
	err = s.ValidateSearchMyScannableProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "projects/search_my_scannable_projects", opt)
	if err != nil {
		return
	}
	v = new(ProjectsSearchMyScannableProjectsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectsUpdateDefaultVisibilityOption struct {
	ProjectVisibility string `url:"projectVisibility,omitempty"` // Description:"Default visibility for projects",ExampleValue:""
}

// UpdateDefaultVisibility Update the default visibility for new projects.<br/>Requires System Administrator privileges
func (s *ProjectsService) UpdateDefaultVisibility(opt *ProjectsUpdateDefaultVisibilityOption) (resp *http.Response, err error) {
	err = s.ValidateUpdateDefaultVisibilityOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/update_default_visibility", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectsUpdateKeyOption struct {
	From string `url:"from,omitempty"` // Description:"Project key",ExampleValue:"my_old_project"
	To   string `url:"to,omitempty"`   // Description:"New project key",ExampleValue:"my_new_project"
}

// UpdateKey Update a project all its sub-components keys.<br>Requires one of the following permissions: <ul><li>'Administer System'</li><li>'Administer' rights on the specified project</li></ul>
func (s *ProjectsService) UpdateKey(opt *ProjectsUpdateKeyOption) (resp *http.Response, err error) {
	err = s.ValidateUpdateKeyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/update_key", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectsUpdateVisibilityOption struct {
	Project    string `url:"project,omitempty"`    // Description:"Project key",ExampleValue:"my_project"
	Visibility string `url:"visibility,omitempty"` // Description:"New visibility",ExampleValue:""
}

// UpdateVisibility Updates visibility of a project or view.<br>Requires 'Project administer' permission on the specified project or view
func (s *ProjectsService) UpdateVisibility(opt *ProjectsUpdateVisibilityOption) (resp *http.Response, err error) {
	err = s.ValidateUpdateVisibilityOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "projects/update_visibility", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}
