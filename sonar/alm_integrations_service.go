// Manage DevOps Platform Integrations
package sonar

import "net/http"

type AlmIntegrationsService struct {
	client *Client
}

type AlmIntegrationsGetGithubClientIdObject struct {
	ClientID string `json:"clientId,omitempty"`
}

type AlmIntegrationsListAzureProjectsObject struct {
	Projects []AlmIntegrationsListAzureProjectsObject_sub1 `json:"projects,omitempty"`
}

type AlmIntegrationsListAzureProjectsObject_sub1 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
}

type AlmIntegrationsListBitbucketserverProjectsObject struct {
	Projects []AlmIntegrationsListBitbucketserverProjectsObject_sub1 `json:"projects,omitempty"`
}

type AlmIntegrationsListBitbucketserverProjectsObject_sub1 struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

type AlmIntegrationsListGithubOrganizationsObject struct {
	Organizations []AlmIntegrationsListGithubOrganizationsObject_sub1 `json:"organizations,omitempty"`
	Paging        AlmIntegrationsListGithubOrganizationsObject_sub2   `json:"paging,omitempty"`
}

type AlmIntegrationsListGithubOrganizationsObject_sub1 struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

type AlmIntegrationsListGithubOrganizationsObject_sub2 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type AlmIntegrationsListGithubRepositoriesObject struct {
	Paging       AlmIntegrationsListGithubRepositoriesObject_sub1   `json:"paging,omitempty"`
	Repositories []AlmIntegrationsListGithubRepositoriesObject_sub2 `json:"repositories,omitempty"`
}

type AlmIntegrationsListGithubRepositoriesObject_sub2 struct {
	ID   int64  `json:"id,omitempty"`
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty"`
}

type AlmIntegrationsListGithubRepositoriesObject_sub1 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type AlmIntegrationsSearchAzureReposObject struct {
	Repositories []AlmIntegrationsSearchAzureReposObject_sub1 `json:"repositories,omitempty"`
}

type AlmIntegrationsSearchAzureReposObject_sub1 struct {
	Name        string `json:"name,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
}

type AlmIntegrationsSearchBitbucketcloudReposObject struct {
	IsLastPage   bool                                                  `json:"isLastPage,omitempty"`
	Paging       AlmIntegrationsSearchBitbucketcloudReposObject_sub1   `json:"paging,omitempty"`
	Repositories []AlmIntegrationsSearchBitbucketcloudReposObject_sub2 `json:"repositories,omitempty"`
}

type AlmIntegrationsSearchBitbucketcloudReposObject_sub2 struct {
	Name         string `json:"name,omitempty"`
	ProjectKey   string `json:"projectKey,omitempty"`
	Slug         string `json:"slug,omitempty"`
	SqProjectKey string `json:"sqProjectKey,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Workspace    string `json:"workspace,omitempty"`
}

type AlmIntegrationsSearchBitbucketcloudReposObject_sub1 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
}

type AlmIntegrationsSearchBitbucketserverReposObject struct {
	IsLastPage   bool                                                   `json:"isLastPage,omitempty"`
	Repositories []AlmIntegrationsSearchBitbucketserverReposObject_sub1 `json:"repositories,omitempty"`
}

type AlmIntegrationsSearchBitbucketserverReposObject_sub1 struct {
	Name         string `json:"name,omitempty"`
	ProjectKey   string `json:"projectKey,omitempty"`
	Slug         string `json:"slug,omitempty"`
	SqProjectKey string `json:"sqProjectKey,omitempty"`
	UUID         string `json:"uuid,omitempty"`
	Workspace    string `json:"workspace,omitempty"`
}

type AlmIntegrationsSearchGitlabReposObject struct {
	Paging       AlmIntegrationsSearchGitlabReposObject_sub1   `json:"paging,omitempty"`
	Repositories []AlmIntegrationsSearchGitlabReposObject_sub2 `json:"repositories,omitempty"`
}

type AlmIntegrationsSearchGitlabReposObject_sub2 struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	PathName string `json:"pathName,omitempty"`
	PathSlug string `json:"pathSlug,omitempty"`
	Slug     string `json:"slug,omitempty"`
	URL      string `json:"url,omitempty"`
}

type AlmIntegrationsSearchGitlabReposObject_sub1 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type AlmIntegrationsCheckPatOption struct {
	AlmSetting string `url:"almSetting,omitempty"` // Description:"DevOps Platform setting key",ExampleValue:""
}

// CheckPat Check validity of a Personal Access Token for the given DevOps Platform setting<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) CheckPat(opt *AlmIntegrationsCheckPatOption) (v *string, resp *http.Response, err error) {
	err = s.ValidateCheckPatOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/check_pat", opt)
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

type AlmIntegrationsGetGithubClientIdOption struct {
	AlmSetting string `url:"almSetting,omitempty"` // Description:"DevOps Platform setting key",ExampleValue:""
}

// GetGithubClientId Get the client id of a Github Integration.
func (s *AlmIntegrationsService) GetGithubClientId(opt *AlmIntegrationsGetGithubClientIdOption) (v *AlmIntegrationsGetGithubClientIdObject, resp *http.Response, err error) {
	err = s.ValidateGetGithubClientIdOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/get_github_client_id", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsGetGithubClientIdObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsImportAzureProjectOption struct {
	AlmSetting     string `url:"almSetting,omitempty"`     // Description:"DevOps Platform setting key",ExampleValue:""
	ProjectName    string `url:"projectName,omitempty"`    // Description:"Azure project name",ExampleValue:""
	RepositoryName string `url:"repositoryName,omitempty"` // Description:"Azure repository name",ExampleValue:""
}

// ImportAzureProject Create a SonarQube project with the information from the provided Azure DevOps project.<br/>Autoconfigure pull request decoration mechanism.<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ImportAzureProject(opt *AlmIntegrationsImportAzureProjectOption) (resp *http.Response, err error) {
	err = s.ValidateImportAzureProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/import_azure_project", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type AlmIntegrationsImportBitbucketcloudRepoOption struct {
	AlmSetting     string `url:"almSetting,omitempty"`     // Description:"DevOps Platform setting key",ExampleValue:""
	RepositorySlug string `url:"repositorySlug,omitempty"` // Description:"Bitbucket Cloud repository slug",ExampleValue:""
}

// ImportBitbucketcloudRepo Create a SonarQube project with the information from the provided Bitbucket Cloud repository.<br/>Autoconfigure pull request decoration mechanism.<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ImportBitbucketcloudRepo(opt *AlmIntegrationsImportBitbucketcloudRepoOption) (resp *http.Response, err error) {
	err = s.ValidateImportBitbucketcloudRepoOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/import_bitbucketcloud_repo", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type AlmIntegrationsImportBitbucketserverProjectOption struct {
	AlmSetting     string `url:"almSetting,omitempty"`     // Description:"DevOps Platform setting key",ExampleValue:""
	ProjectKey     string `url:"projectKey,omitempty"`     // Description:"BitbucketServer project key",ExampleValue:""
	RepositorySlug string `url:"repositorySlug,omitempty"` // Description:"BitbucketServer repository slug",ExampleValue:""
}

// ImportBitbucketserverProject Create a SonarQube project with the information from the provided BitbucketServer project.<br/>Autoconfigure pull request decoration mechanism.<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ImportBitbucketserverProject(opt *AlmIntegrationsImportBitbucketserverProjectOption) (resp *http.Response, err error) {
	err = s.ValidateImportBitbucketserverProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/import_bitbucketserver_project", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type AlmIntegrationsImportGithubProjectOption struct {
	AlmSetting    string `url:"almSetting,omitempty"`    // Description:"DevOps Platform setting key",ExampleValue:""
	Organization  string `url:"organization,omitempty"`  // Description:"GitHub organization",ExampleValue:""
	RepositoryKey string `url:"repositoryKey,omitempty"` // Description:"GitHub repository key",ExampleValue:""
}

// ImportGithubProject Create a SonarQube project with the information from the provided GitHub repository.<br/>Autoconfigure pull request decoration mechanism.<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ImportGithubProject(opt *AlmIntegrationsImportGithubProjectOption) (resp *http.Response, err error) {
	err = s.ValidateImportGithubProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/import_github_project", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type AlmIntegrationsImportGitlabProjectOption struct {
	AlmSetting      string `url:"almSetting,omitempty"`      // Description:"DevOps Platform setting key",ExampleValue:""
	GitlabProjectId string `url:"gitlabProjectId,omitempty"` // Description:"GitLab project ID",ExampleValue:""
}

// ImportGitlabProject Import a GitLab project to SonarQube, creating a new project and configuring MR decoration<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ImportGitlabProject(opt *AlmIntegrationsImportGitlabProjectOption) (resp *http.Response, err error) {
	err = s.ValidateImportGitlabProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/import_gitlab_project", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type AlmIntegrationsListAzureProjectsOption struct {
	AlmSetting string `url:"almSetting,omitempty"` // Description:"DevOps Platform setting key",ExampleValue:""
}

// ListAzureProjects List Azure projects<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ListAzureProjects(opt *AlmIntegrationsListAzureProjectsOption) (v *AlmIntegrationsListAzureProjectsObject, resp *http.Response, err error) {
	err = s.ValidateListAzureProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/list_azure_projects", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsListAzureProjectsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsListBitbucketserverProjectsOption struct {
	AlmSetting string `url:"almSetting,omitempty"` // Description:"DevOps Platform setting key",ExampleValue:""
}

// ListBitbucketserverProjects List the Bitbucket Server projects<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ListBitbucketserverProjects(opt *AlmIntegrationsListBitbucketserverProjectsOption) (v *AlmIntegrationsListBitbucketserverProjectsObject, resp *http.Response, err error) {
	err = s.ValidateListBitbucketserverProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/list_bitbucketserver_projects", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsListBitbucketserverProjectsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsListGithubOrganizationsOption struct {
	AlmSetting string `url:"almSetting,omitempty"` // Description:"DevOps Platform setting key",ExampleValue:""
	P          string `url:"p,omitempty"`          // Description:"Index of the page to display",ExampleValue:""
	Ps         string `url:"ps,omitempty"`         // Description:"Size for the paging to apply",ExampleValue:""
	Token      string `url:"token,omitempty"`      // Description:"Github authorization code",ExampleValue:""
}

// ListGithubOrganizations List GitHub organizations<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ListGithubOrganizations(opt *AlmIntegrationsListGithubOrganizationsOption) (v *AlmIntegrationsListGithubOrganizationsObject, resp *http.Response, err error) {
	err = s.ValidateListGithubOrganizationsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/list_github_organizations", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsListGithubOrganizationsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsListGithubRepositoriesOption struct {
	AlmSetting   string `url:"almSetting,omitempty"`   // Description:"DevOps Platform setting key",ExampleValue:""
	Organization string `url:"organization,omitempty"` // Description:"Github organization",ExampleValue:""
	P            string `url:"p,omitempty"`            // Description:"Index of the page to display",ExampleValue:""
	Ps           string `url:"ps,omitempty"`           // Description:"Size for the paging to apply",ExampleValue:""
	Q            string `url:"q,omitempty"`            // Description:"Limit search to repositories that contain the supplied string",ExampleValue:"Apache"
}

// ListGithubRepositories List the GitHub repositories for an organization<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) ListGithubRepositories(opt *AlmIntegrationsListGithubRepositoriesOption) (v *AlmIntegrationsListGithubRepositoriesObject, resp *http.Response, err error) {
	err = s.ValidateListGithubRepositoriesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/list_github_repositories", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsListGithubRepositoriesObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsSearchAzureReposOption struct {
	AlmSetting  string `url:"almSetting,omitempty"`  // Description:"DevOps Platform setting key",ExampleValue:""
	ProjectName string `url:"projectName,omitempty"` // Description:"Project name filter",ExampleValue:""
	SearchQuery string `url:"searchQuery,omitempty"` // Description:"Search query filter",ExampleValue:""
}

// SearchAzureRepos Search the Azure repositories<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) SearchAzureRepos(opt *AlmIntegrationsSearchAzureReposOption) (v *AlmIntegrationsSearchAzureReposObject, resp *http.Response, err error) {
	err = s.ValidateSearchAzureReposOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/search_azure_repos", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsSearchAzureReposObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsSearchBitbucketcloudReposOption struct {
	AlmSetting     string `url:"almSetting,omitempty"`     // Description:"DevOps Platform setting key",ExampleValue:""
	P              string `url:"p,omitempty"`              // Description:"1-based page number",ExampleValue:"42"
	Ps             string `url:"ps,omitempty"`             // Description:"Page size. Must be greater than 0 and less or equal than 100",ExampleValue:"20"
	RepositoryName string `url:"repositoryName,omitempty"` // Description:"Repository name filter",ExampleValue:""
}

// SearchBitbucketcloudRepos Search the Bitbucket Cloud repositories<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) SearchBitbucketcloudRepos(opt *AlmIntegrationsSearchBitbucketcloudReposOption) (v *AlmIntegrationsSearchBitbucketcloudReposObject, resp *http.Response, err error) {
	err = s.ValidateSearchBitbucketcloudReposOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/search_bitbucketcloud_repos", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsSearchBitbucketcloudReposObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsSearchBitbucketserverReposOption struct {
	AlmSetting     string `url:"almSetting,omitempty"`     // Description:"DevOps Platform setting key",ExampleValue:""
	ProjectName    string `url:"projectName,omitempty"`    // Description:"Project name filter",ExampleValue:""
	RepositoryName string `url:"repositoryName,omitempty"` // Description:"Repository name filter",ExampleValue:""
}

// SearchBitbucketserverRepos Search the Bitbucket Server repositories with REPO_ADMIN access<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) SearchBitbucketserverRepos(opt *AlmIntegrationsSearchBitbucketserverReposOption) (v *AlmIntegrationsSearchBitbucketserverReposObject, resp *http.Response, err error) {
	err = s.ValidateSearchBitbucketserverReposOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/search_bitbucketserver_repos", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsSearchBitbucketserverReposObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsSearchGitlabReposOption struct {
	AlmSetting  string `url:"almSetting,omitempty"`  // Description:"DevOps Platform setting key",ExampleValue:""
	P           string `url:"p,omitempty"`           // Description:"1-based page number",ExampleValue:"42"
	ProjectName string `url:"projectName,omitempty"` // Description:"Project name filter",ExampleValue:""
	Ps          string `url:"ps,omitempty"`          // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
}

// SearchGitlabRepos Search the GitLab projects.<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) SearchGitlabRepos(opt *AlmIntegrationsSearchGitlabReposOption) (v *AlmIntegrationsSearchGitlabReposObject, resp *http.Response, err error) {
	err = s.ValidateSearchGitlabReposOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "alm_integrations/search_gitlab_repos", opt)
	if err != nil {
		return
	}
	v = new(AlmIntegrationsSearchGitlabReposObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type AlmIntegrationsSetPatOption struct {
	AlmSetting string `url:"almSetting,omitempty"` // Description:"DevOps Platform setting key",ExampleValue:""
	Pat        string `url:"pat,omitempty"`        // Description:"Personal Access Token",ExampleValue:""
	Username   string `url:"username,omitempty"`   // Description:"Username",ExampleValue:""
}

// SetPat Set a Personal Access Token for the given DevOps Platform setting<br/>Only valid for Azure DevOps, Bitbucket Server, GitLab and Bitbucket Cloud Setting<br/>Requires the 'Create Projects' permission
func (s *AlmIntegrationsService) SetPat(opt *AlmIntegrationsSetPatOption) (resp *http.Response, err error) {
	err = s.ValidateSetPatOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "alm_integrations/set_pat", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}
