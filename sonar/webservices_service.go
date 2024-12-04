// Get information on the web api supported on this instance.
package sonar

import "net/http"

type WebservicesService struct {
	client *Client
}

type WebservicesListObject struct {
	WebServices []WebservicesListObject_sub4 `json:"webServices,omitempty"`
}

type WebservicesListObject_sub4 struct {
	Actions     []WebservicesListObject_sub3 `json:"actions,omitempty"`
	Description string                       `json:"description,omitempty"`
	Path        string                       `json:"path,omitempty"`
	Since       string                       `json:"since,omitempty"`
}

type WebservicesListObject_sub3 struct {
	Changelog          []WebservicesListObject_sub1 `json:"changelog,omitempty"`
	DeprecatedSince    string                       `json:"deprecatedSince,omitempty"`
	Description        string                       `json:"description,omitempty"`
	HasResponseExample bool                         `json:"hasResponseExample,omitempty"`
	Internal           bool                         `json:"internal,omitempty"`
	Key                string                       `json:"key,omitempty"`
	Params             []WebservicesListObject_sub2 `json:"params,omitempty"`
	Post               bool                         `json:"post,omitempty"`
	Since              string                       `json:"since,omitempty"`
}

type WebservicesListObject_sub2 struct {
	DefaultValue       string   `json:"defaultValue,omitempty"`
	DeprecatedKey      string   `json:"deprecatedKey,omitempty"`
	DeprecatedKeySince string   `json:"deprecatedKeySince,omitempty"`
	DeprecatedSince    string   `json:"deprecatedSince,omitempty"`
	Description        string   `json:"description,omitempty"`
	ExampleValue       string   `json:"exampleValue,omitempty"`
	Internal           bool     `json:"internal,omitempty"`
	Key                string   `json:"key,omitempty"`
	MaximumLength      int64    `json:"maximumLength,omitempty"`
	MaximumValue       int64    `json:"maximumValue,omitempty"`
	MinimumLength      int64    `json:"minimumLength,omitempty"`
	PossibleValues     []string `json:"possibleValues,omitempty"`
	Required           bool     `json:"required,omitempty"`
	Since              string   `json:"since,omitempty"`
}

type WebservicesListObject_sub1 struct {
	Description string `json:"description,omitempty"`
	Version     string `json:"version,omitempty"`
}

type WebservicesResponseExampleObject struct {
	Example WebservicesResponseExampleObject_sub9 `json:"example,omitempty"`
	Format  string                                `json:"format,omitempty"`
}

type WebservicesResponseExampleObject_sub5 struct {
	Actions      []string                                `json:"actions,omitempty"`
	Attr         WebservicesResponseExampleObject_sub2   `json:"attr,omitempty"`
	Author       string                                  `json:"author,omitempty"`
	Comments     []WebservicesResponseExampleObject_sub3 `json:"comments,omitempty"`
	Component    string                                  `json:"component,omitempty"`
	CreationDate string                                  `json:"creationDate,omitempty"`
	Debt         string                                  `json:"debt,omitempty"`
	Key          string                                  `json:"key,omitempty"`
	Line         int64                                   `json:"line,omitempty"`
	Message      string                                  `json:"message,omitempty"`
	Project      string                                  `json:"project,omitempty"`
	Resolution   string                                  `json:"resolution,omitempty"`
	Rule         string                                  `json:"rule,omitempty"`
	Severity     string                                  `json:"severity,omitempty"`
	Status       string                                  `json:"status,omitempty"`
	Tags         []string                                `json:"tags,omitempty"`
	TextRange    WebservicesResponseExampleObject_sub4   `json:"textRange,omitempty"`
	Transitions  []string                                `json:"transitions,omitempty"`
	UpdateDate   string                                  `json:"updateDate,omitempty"`
}

type WebservicesResponseExampleObject_sub8 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type WebservicesResponseExampleObject_sub9 struct {
	Components []WebservicesResponseExampleObject_sub1 `json:"components,omitempty"`
	Issues     []WebservicesResponseExampleObject_sub5 `json:"issues,omitempty"`
	Paging     WebservicesResponseExampleObject_sub6   `json:"paging,omitempty"`
	Rules      []WebservicesResponseExampleObject_sub7 `json:"rules,omitempty"`
	Users      []WebservicesResponseExampleObject_sub8 `json:"users,omitempty"`
}

type WebservicesResponseExampleObject_sub3 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type WebservicesResponseExampleObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type WebservicesResponseExampleObject_sub4 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type WebservicesResponseExampleObject_sub2 struct {
	Jira_issue_key string `json:"jira-issue-key,omitempty"`
}

type WebservicesResponseExampleObject_sub7 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type WebservicesResponseExampleObject_sub6 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type WebservicesListOption struct {
	IncludeInternals string `url:"include_internals,omitempty"` // Description:"Include web services that are implemented for internal use only. Their forward-compatibility is not assured",ExampleValue:""
}

// List List web services
func (s *WebservicesService) List(opt *WebservicesListOption) (v *WebservicesListObject, resp *http.Response, err error) {
	err = s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/list", opt)
	if err != nil {
		return
	}
	v = new(WebservicesListObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type WebservicesResponseExampleOption struct {
	Action     string `url:"action,omitempty"`     // Description:"Action of the web service",ExampleValue:"search"
	Controller string `url:"controller,omitempty"` // Description:"Controller of the web service",ExampleValue:"api/issues"
}

// ResponseExample Display web service response example
func (s *WebservicesService) ResponseExample(opt *WebservicesResponseExampleOption) (v *WebservicesResponseExampleObject, resp *http.Response, err error) {
	err = s.ValidateResponseExampleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webservices/response_example", opt)
	if err != nil {
		return
	}
	v = new(WebservicesResponseExampleObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
