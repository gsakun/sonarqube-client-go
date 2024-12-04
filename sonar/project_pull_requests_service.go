package sonar

import "net/http"

type ProjectPullRequestsService struct {
	client *Client
}

type ProjectPullRequestsDeleteOption struct {
	Project     string `url:"project,omitempty"`     // Description:"",ExampleValue:""
	PullRequest string `url:"pullRequest,omitempty"` // Description:"",ExampleValue:""
}

// Delete
func (s *ProjectPullRequestsService) Delete(opt *ProjectPullRequestsDeleteOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_pull_requests/delete", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type ProjectPullRequestsListOption struct {
	Project string `url:"project,omitempty"` // Description:"",ExampleValue:""
}

// List
func (s *ProjectPullRequestsService) List(opt *ProjectPullRequestsListOption) (v *string, resp *http.Response, err error) {
	err = s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_pull_requests/list", opt)
	if err != nil {
		return
	}
	v = new(string)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
