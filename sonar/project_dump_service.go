// Project export/import
package sonar

import "net/http"

type ProjectDumpService struct {
	client *Client
}

type ProjectDumpExportObject struct {
	ProjectID   string `json:"projectId,omitempty"`
	ProjectKey  string `json:"projectKey,omitempty"`
	ProjectName string `json:"projectName,omitempty"`
	TaskID      string `json:"taskId,omitempty"`
}

type ProjectDumpStatusObject struct {
	CanBeExported bool   `json:"canBeExported,omitempty"`
	CanBeImported bool   `json:"canBeImported,omitempty"`
	DumpToImport  string `json:"dumpToImport,omitempty"`
	ExportedDump  string `json:"exportedDump,omitempty"`
}

type ProjectDumpExportOption struct {
	Key string `url:"key,omitempty"` // Description:"",ExampleValue:"my_project"
}

// Export Triggers project dump so that the project can be imported to another SonarQube server (see api/project_dump/import, available in Enterprise Edition). Requires the 'Administer' permission.
func (s *ProjectDumpService) Export(opt *ProjectDumpExportOption) (v *ProjectDumpExportObject, resp *http.Response, err error) {
	err = s.ValidateExportOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "project_dump/export", opt)
	if err != nil {
		return
	}
	v = new(ProjectDumpExportObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ProjectDumpStatusOption struct {
	Id  string `url:"id,omitempty"`  // Description:"Project id",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Key string `url:"key,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Status Provide the import and export status of a project. Permission 'Administer' is required. The project id or project key must be provided.
func (s *ProjectDumpService) Status(opt *ProjectDumpStatusOption) (v *ProjectDumpStatusObject, resp *http.Response, err error) {
	err = s.ValidateStatusOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "project_dump/status", opt)
	if err != nil {
		return
	}
	v = new(ProjectDumpStatusObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
