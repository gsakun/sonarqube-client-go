// Manage branch
package sonar

import "net/http"

type ProjectBranchesService struct {
	client *Client
}

type ProjectBranchesListObject struct {
	Branches []ProjectBranchesListObject_sub2 `json:"branches,omitempty"`
}

type ProjectBranchesListObject_sub2 struct {
	AnalysisDate      string                         `json:"analysisDate,omitempty"`
	ExcludedFromPurge bool                           `json:"excludedFromPurge,omitempty"`
	IsMain            bool                           `json:"isMain,omitempty"`
	Name              string                         `json:"name,omitempty"`
	Status            ProjectBranchesListObject_sub1 `json:"status,omitempty"`
	Type              string                         `json:"type,omitempty"`
}

type ProjectBranchesListObject_sub1 struct {
	QualityGateStatus string `json:"qualityGateStatus,omitempty"`
}

type ProjectBranchesDeleteOption struct {
	Branch  string `url:"branch,omitempty"`  // Description:"Branch key",ExampleValue:"feature/my_branch"
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Delete Delete a non-main branch of a project or application.<br/>Requires 'Administer' rights on the specified project or application.
func (s *ProjectBranchesService) Delete(opt *ProjectBranchesDeleteOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_branches/delete", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectBranchesListOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// List List the branches of a project or application.<br/>Requires 'Browse' or 'Execute analysis' rights on the specified project or application.
func (s *ProjectBranchesService) List(opt *ProjectBranchesListOption) (v *ProjectBranchesListObject, resp *http.Response, err error) {
	err = s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_branches/list", opt)
	if err != nil {
		return
	}
	v = new(ProjectBranchesListObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectBranchesRenameOption struct {
	Name    string `url:"name,omitempty"`    // Description:"New name of the main branch",ExampleValue:"branch1"
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Rename Rename the main branch of a project or application.<br/>Requires 'Administer' permission on the specified project or application.
func (s *ProjectBranchesService) Rename(opt *ProjectBranchesRenameOption) (resp *http.Response, err error) {
	err = s.ValidateRenameOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_branches/rename", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectBranchesSetAutomaticDeletionProtectionOption struct {
	Branch  string `url:"branch,omitempty"`  // Description:"Branch key",ExampleValue:"feature/my_branch"
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Value   string `url:"value,omitempty"`   // Description:"Sets whether the branch should be protected from automatic deletion when it hasn't been analyzed for a set period of time.",ExampleValue:"true"
}

// SetAutomaticDeletionProtection Protect a specific branch from automatic deletion. Protection can't be disabled for the main branch.<br/>Requires 'Administer' permission on the specified project or application.
func (s *ProjectBranchesService) SetAutomaticDeletionProtection(opt *ProjectBranchesSetAutomaticDeletionProtectionOption) (resp *http.Response, err error) {
	err = s.ValidateSetAutomaticDeletionProtectionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_branches/set_automatic_deletion_protection", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}
