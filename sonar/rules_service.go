// Get and update some details of automatic rules, and manage custom rules.
package sonar

import "net/http"

type RulesService struct {
	client *Client
}

type RulesAppObject struct {
	CanWrite        bool                  `json:"canWrite,omitempty"`
	Characteristics []RulesAppObject_sub1 `json:"characteristics,omitempty"`
	Languages       RulesAppObject_sub2   `json:"languages,omitempty"`
	Repositories    []RulesAppObject_sub3 `json:"repositories,omitempty"`
	Statuses        RulesAppObject_sub4   `json:"statuses,omitempty"`
}

type RulesAppObject_sub4 struct {
	Beta       string `json:"BETA,omitempty"`
	Deprecated string `json:"DEPRECATED,omitempty"`
	Ready      string `json:"READY,omitempty"`
}

type RulesAppObject_sub2 struct {
	Java string `json:"java,omitempty"`
	Js   string `json:"js,omitempty"`
	Php  string `json:"php,omitempty"`
}

type RulesAppObject_sub3 struct {
	Key      string `json:"key,omitempty"`
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`
}

type RulesAppObject_sub1 struct {
	Key    string `json:"key,omitempty"`
	Name   string `json:"name,omitempty"`
	Parent string `json:"parent,omitempty"`
}

type RulesCreateObject struct {
	Rule RulesCreateObject_sub2 `json:"rule,omitempty"`
}

type RulesCreateObject_sub2 struct {
	CreatedAt   string                   `json:"createdAt,omitempty"`
	HTMLDesc    string                   `json:"htmlDesc,omitempty"`
	IsExternal  bool                     `json:"isExternal,omitempty"`
	IsTemplate  bool                     `json:"isTemplate,omitempty"`
	Key         string                   `json:"key,omitempty"`
	Lang        string                   `json:"lang,omitempty"`
	LangName    string                   `json:"langName,omitempty"`
	MdDesc      string                   `json:"mdDesc,omitempty"`
	Name        string                   `json:"name,omitempty"`
	Params      []RulesCreateObject_sub1 `json:"params,omitempty"`
	Repo        string                   `json:"repo,omitempty"`
	Scope       string                   `json:"scope,omitempty"`
	Severity    string                   `json:"severity,omitempty"`
	Status      string                   `json:"status,omitempty"`
	SysTags     []interface{}            `json:"sysTags,omitempty"`
	TemplateKey string                   `json:"templateKey,omitempty"`
	Type        string                   `json:"type,omitempty"`
}

type RulesCreateObject_sub1 struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	HTMLDesc     string `json:"htmlDesc,omitempty"`
	Key          string `json:"key,omitempty"`
	Type         string `json:"type,omitempty"`
}

type RulesRepositoriesObject struct {
	Repositories []RulesRepositoriesObject_sub1 `json:"repositories,omitempty"`
}

type RulesRepositoriesObject_sub1 struct {
	Key      string `json:"key,omitempty"`
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`
}

type RulesSearchObject struct {
	Actives RulesSearchObject_sub3    `json:"actives,omitempty"`
	Facets  []RulesSearchObject_sub5  `json:"facets,omitempty"`
	Paging  RulesSearchObject_sub6    `json:"paging,omitempty"`
	Rules   []RulesSearchObject_sub10 `json:"rules,omitempty"`
}

type RulesSearchObject_sub8 struct {
	Content string                 `json:"content,omitempty"`
	Context RulesSearchObject_sub7 `json:"context,omitempty"`
	Key     string                 `json:"key,omitempty"`
}

type RulesSearchObject_sub4 struct {
	Count int64  `json:"count,omitempty"`
	Val   string `json:"val,omitempty"`
}

type RulesSearchObject_sub10 struct {
	CreatedAt           string                   `json:"createdAt,omitempty"`
	DescriptionSections []RulesSearchObject_sub8 `json:"descriptionSections,omitempty"`
	HTMLDesc            string                   `json:"htmlDesc,omitempty"`
	HTMLNote            string                   `json:"htmlNote,omitempty"`
	InternalKey         string                   `json:"internalKey,omitempty"`
	IsExternal          bool                     `json:"isExternal,omitempty"`
	IsTemplate          bool                     `json:"isTemplate,omitempty"`
	Key                 string                   `json:"key,omitempty"`
	Lang                string                   `json:"lang,omitempty"`
	LangName            string                   `json:"langName,omitempty"`
	MdNote              string                   `json:"mdNote,omitempty"`
	Name                string                   `json:"name,omitempty"`
	NoteLogin           string                   `json:"noteLogin,omitempty"`
	Params              []RulesSearchObject_sub9 `json:"params,omitempty"`
	Repo                string                   `json:"repo,omitempty"`
	Scope               string                   `json:"scope,omitempty"`
	Severity            string                   `json:"severity,omitempty"`
	Status              string                   `json:"status,omitempty"`
	SysTags             []string                 `json:"sysTags,omitempty"`
	Tags                []interface{}            `json:"tags,omitempty"`
	TemplateKey         string                   `json:"templateKey,omitempty"`
	Type                string                   `json:"type,omitempty"`
	UpdatedAt           string                   `json:"updatedAt,omitempty"`
}

type RulesSearchObject_sub9 struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	Desc         string `json:"desc,omitempty"`
	Key          string `json:"key,omitempty"`
}

type RulesSearchObject_sub7 struct {
	DisplayName string `json:"displayName,omitempty"`
	Key         string `json:"key,omitempty"`
}

type RulesSearchObject_sub2 struct {
	Inherit  string                   `json:"inherit,omitempty"`
	Params   []RulesSearchObject_sub1 `json:"params,omitempty"`
	QProfile string                   `json:"qProfile,omitempty"`
	Severity string                   `json:"severity,omitempty"`
}

type RulesSearchObject_sub1 struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type RulesSearchObject_sub5 struct {
	Name   string                   `json:"name,omitempty"`
	Values []RulesSearchObject_sub4 `json:"values,omitempty"`
}

type RulesSearchObject_sub6 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type RulesSearchObject_sub3 struct {
	Squid_ClassCyclomaticComplexity  []RulesSearchObject_sub2 `json:"squid:ClassCyclomaticComplexity,omitempty"`
	Squid_MethodCyclomaticComplexity []RulesSearchObject_sub2 `json:"squid:MethodCyclomaticComplexity,omitempty"`
	Squid_S1067                      []RulesSearchObject_sub2 `json:"squid:S1067,omitempty"`
}

type RulesShowObject struct {
	Actives []RulesShowObject_sub2 `json:"actives,omitempty"`
	Rule    RulesShowObject_sub6   `json:"rule,omitempty"`
}

type RulesShowObject_sub4 struct {
	Content string               `json:"content,omitempty"`
	Context RulesShowObject_sub3 `json:"context,omitempty"`
	Key     string               `json:"key,omitempty"`
}

type RulesShowObject_sub6 struct {
	DefaultRemFnBaseEffort    string                 `json:"defaultRemFnBaseEffort,omitempty"`
	DefaultRemFnGapMultiplier string                 `json:"defaultRemFnGapMultiplier,omitempty"`
	DefaultRemFnType          string                 `json:"defaultRemFnType,omitempty"`
	DescriptionSections       []RulesShowObject_sub4 `json:"descriptionSections,omitempty"`
	GapDescription            string                 `json:"gapDescription,omitempty"`
	HTMLDesc                  string                 `json:"htmlDesc,omitempty"`
	InternalKey               string                 `json:"internalKey,omitempty"`
	IsExternal                bool                   `json:"isExternal,omitempty"`
	Key                       string                 `json:"key,omitempty"`
	Lang                      string                 `json:"lang,omitempty"`
	LangName                  string                 `json:"langName,omitempty"`
	Name                      string                 `json:"name,omitempty"`
	Params                    []RulesShowObject_sub5 `json:"params,omitempty"`
	RemFnBaseEffort           string                 `json:"remFnBaseEffort,omitempty"`
	RemFnGapMultiplier        string                 `json:"remFnGapMultiplier,omitempty"`
	RemFnOverloaded           bool                   `json:"remFnOverloaded,omitempty"`
	RemFnType                 string                 `json:"remFnType,omitempty"`
	Repo                      string                 `json:"repo,omitempty"`
	Scope                     string                 `json:"scope,omitempty"`
	Severity                  string                 `json:"severity,omitempty"`
	Status                    string                 `json:"status,omitempty"`
	SysTags                   []string               `json:"sysTags,omitempty"`
	Tags                      []interface{}          `json:"tags,omitempty"`
	Template                  bool                   `json:"template,omitempty"`
	Type                      string                 `json:"type,omitempty"`
}

type RulesShowObject_sub5 struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	Desc         string `json:"desc,omitempty"`
	Key          string `json:"key,omitempty"`
}

type RulesShowObject_sub3 struct {
	DisplayName string `json:"displayName,omitempty"`
	Key         string `json:"key,omitempty"`
}

type RulesShowObject_sub2 struct {
	Inherit  string                 `json:"inherit,omitempty"`
	Params   []RulesShowObject_sub1 `json:"params,omitempty"`
	QProfile string                 `json:"qProfile,omitempty"`
	Severity string                 `json:"severity,omitempty"`
}

type RulesShowObject_sub1 struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
}

type RulesTagsObject struct {
	Tags []string `json:"tags,omitempty"`
}

type RulesUpdateObject struct {
	Rule RulesUpdateObject_sub2 `json:"rule,omitempty"`
}

type RulesUpdateObject_sub2 struct {
	CreatedAt       string                   `json:"createdAt,omitempty"`
	DebtOverloaded  bool                     `json:"debtOverloaded,omitempty"`
	HTMLDesc        string                   `json:"htmlDesc,omitempty"`
	IsExternal      bool                     `json:"isExternal,omitempty"`
	IsTemplate      bool                     `json:"isTemplate,omitempty"`
	Key             string                   `json:"key,omitempty"`
	Lang            string                   `json:"lang,omitempty"`
	LangName        string                   `json:"langName,omitempty"`
	MdDesc          string                   `json:"mdDesc,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Params          []RulesUpdateObject_sub1 `json:"params,omitempty"`
	RemFnOverloaded bool                     `json:"remFnOverloaded,omitempty"`
	Repo            string                   `json:"repo,omitempty"`
	Scope           string                   `json:"scope,omitempty"`
	Severity        string                   `json:"severity,omitempty"`
	Status          string                   `json:"status,omitempty"`
	SysTags         []interface{}            `json:"sysTags,omitempty"`
	Tags            []interface{}            `json:"tags,omitempty"`
	TemplateKey     string                   `json:"templateKey,omitempty"`
	Type            string                   `json:"type,omitempty"`
}

type RulesUpdateObject_sub1 struct {
	DefaultValue string `json:"defaultValue,omitempty"`
	HTMLDesc     string `json:"htmlDesc,omitempty"`
	Key          string `json:"key,omitempty"`
	Type         string `json:"type,omitempty"`
}

// App Get data required for rendering the page 'Coding Rules'.
func (s *RulesService) App() (v *RulesAppObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "rules/app", nil)
	if err != nil {
		return
	}
	v = new(RulesAppObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RulesCreateOption struct {
	CustomKey           string `url:"customKey,omitempty"`           // Description:"Key of the custom rule",ExampleValue:"Todo_should_not_be_used"
	MarkdownDescription string `url:"markdownDescription,omitempty"` // Description:"Rule description in <a href='/formatting/help'>markdown format</a>",ExampleValue:"Description of my custom rule"
	Name                string `url:"name,omitempty"`                // Description:"Rule name",ExampleValue:"My custom rule"
	Params              string `url:"params,omitempty"`              // Description:"Parameters as semi-colon list of <key>=<value>, for example 'params=key1=v1;key2=v2' (Only for custom rule)",ExampleValue:""
	PreventReactivation string `url:"preventReactivation,omitempty"` // Description:"If set to true and if the rule has been deactivated (status 'REMOVED'), a status 409 will be returned",ExampleValue:""
	Severity            string `url:"severity,omitempty"`            // Description:"Rule severity",ExampleValue:""
	Status              string `url:"status,omitempty"`              // Description:"Rule status",ExampleValue:""
	TemplateKey         string `url:"templateKey,omitempty"`         // Description:"Key of the template rule in order to create a custom rule (mandatory for custom rule)",ExampleValue:"java:XPath"
	Type                string `url:"type,omitempty"`                // Description:"Rule type",ExampleValue:""
}

// Create Create a custom rule.<br>Requires the 'Administer Quality Profiles' permission
func (s *RulesService) Create(opt *RulesCreateOption) (v *RulesCreateObject, resp *http.Response, err error) {
	err = s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "rules/create", opt)
	if err != nil {
		return
	}
	v = new(RulesCreateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RulesDeleteOption struct {
	Key string `url:"key,omitempty"` // Description:"Rule key",ExampleValue:"java:S1144"
}

// Delete Delete custom rule.<br/>Requires the 'Administer Quality Profiles' permission
func (s *RulesService) Delete(opt *RulesDeleteOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "rules/delete", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

// List List rules, excluding the manual rules and the rules with status REMOVED. JSON format is not supported for response.
func (s *RulesService) List() (v *string, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "rules/list", nil)
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

type RulesRepositoriesOption struct {
	Language string `url:"language,omitempty"` // Description:"A language key; if provided, only repositories for the given language will be returned",ExampleValue:"java"
	Q        string `url:"q,omitempty"`        // Description:"A pattern to match repository keys/names against",ExampleValue:"java"
}

// Repositories List available rule repositories
func (s *RulesService) Repositories(opt *RulesRepositoriesOption) (v *RulesRepositoriesObject, resp *http.Response, err error) {
	err = s.ValidateRepositoriesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "rules/repositories", opt)
	if err != nil {
		return
	}
	v = new(RulesRepositoriesObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RulesSearchOption struct {
	Activation          string `url:"activation,omitempty"`          // Description:"Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.",ExampleValue:""
	ActiveSeverities    string `url:"active_severities,omitempty"`   // Description:"Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	Asc                 string `url:"asc,omitempty"`                 // Description:"Ascending sort",ExampleValue:""
	AvailableSince      string `url:"available_since,omitempty"`     // Description:"Filters rules added since date. Format is yyyy-MM-dd",ExampleValue:"2014-06-22"
	CompareToProfile    string `url:"compareToProfile,omitempty"`    // Description:"Quality profile key to filter rules that are activated. Meant to compare easily to profile set in 'qprofile'",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
	Cwe                 string `url:"cwe,omitempty"`                 // Description:"Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE.",ExampleValue:"12,125,unknown"
	F                   string `url:"f,omitempty"`                   // Description:"Comma-separated list of additional fields to be returned in the response. All the fields are returned by default, except actives.",ExampleValue:"effortToFixDescription,defaultDebtRemFn"
	Facets              string `url:"facets,omitempty"`              // Description:"Comma-separated list of the facets to be computed. No facet is computed by default.",ExampleValue:"languages,repositories"
	IncludeExternal     string `url:"include_external,omitempty"`    // Description:"Include external engine rules in the results",ExampleValue:""
	Inheritance         string `url:"inheritance,omitempty"`         // Description:"Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.",ExampleValue:"INHERITED,OVERRIDES"
	IsTemplate          string `url:"is_template,omitempty"`         // Description:"Filter template rules",ExampleValue:""
	Languages           string `url:"languages,omitempty"`           // Description:"Comma-separated list of languages",ExampleValue:"java,js"
	OwaspTop10          string `url:"owaspTop10,omitempty"`          // Description:"Comma-separated list of OWASP Top 10 2017 lowercase categories.",ExampleValue:""
	OwaspTop102021      string `url:"owaspTop10-2021,omitempty"`     // Description:"Comma-separated list of OWASP Top 10 2021 lowercase categories.",ExampleValue:""
	P                   string `url:"p,omitempty"`                   // Description:"1-based page number",ExampleValue:"42"
	Ps                  string `url:"ps,omitempty"`                  // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q                   string `url:"q,omitempty"`                   // Description:"UTF-8 search query",ExampleValue:"xpath"
	Qprofile            string `url:"qprofile,omitempty"`            // Description:"Quality profile key to filter on. Used only if the parameter 'activation' is set.",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Repositories        string `url:"repositories,omitempty"`        // Description:"Comma-separated list of repositories",ExampleValue:"java,html"
	RuleKey             string `url:"rule_key,omitempty"`            // Description:"Key of rule to search for",ExampleValue:"java:S1144"
	S                   string `url:"s,omitempty"`                   // Description:"Sort field",ExampleValue:"name"
	SansTop25           string `url:"sansTop25,omitempty"`           // Description:"Comma-separated list of SANS Top 25 categories.",ExampleValue:""
	Severities          string `url:"severities,omitempty"`          // Description:"Comma-separated list of default severities. Not the same than severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	SonarsourceSecurity string `url:"sonarsourceSecurity,omitempty"` // Description:"Comma-separated list of SonarSource security categories. Use 'others' to select rules not associated with any category",ExampleValue:"sql-injection,command-injection,others"
	Statuses            string `url:"statuses,omitempty"`            // Description:"Comma-separated list of status codes",ExampleValue:"READY"
	Tags                string `url:"tags,omitempty"`                // Description:"Comma-separated list of tags. Returned rules match any of the tags (OR operator)",ExampleValue:"security,java8"
	TemplateKey         string `url:"template_key,omitempty"`        // Description:"Key of the template rule to filter on. Used to search for the custom rules based on this template.",ExampleValue:"java:S001"
	Types               string `url:"types,omitempty"`               // Description:"Comma-separated list of types. Returned rules match any of the tags (OR operator)",ExampleValue:"BUG"
}

// Search Search for a collection of relevant rules matching a specified query.<br/>
func (s *RulesService) Search(opt *RulesSearchOption) (v *RulesSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "rules/search", opt)
	if err != nil {
		return
	}
	v = new(RulesSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RulesShowOption struct {
	Actives string `url:"actives,omitempty"` // Description:"Show rule's activations for all profiles ("active rules")",ExampleValue:""
	Key     string `url:"key,omitempty"`     // Description:"Rule key",ExampleValue:"javascript:EmptyBlock"
}

// Show Get detailed information about a rule<br>
func (s *RulesService) Show(opt *RulesShowOption) (v *RulesShowObject, resp *http.Response, err error) {
	err = s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "rules/show", opt)
	if err != nil {
		return
	}
	v = new(RulesShowObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RulesTagsOption struct {
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Limit search to tags that contain the supplied string.",ExampleValue:"misra"
}

// Tags List rule tags
func (s *RulesService) Tags(opt *RulesTagsOption) (v *RulesTagsObject, resp *http.Response, err error) {
	err = s.ValidateTagsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "rules/tags", opt)
	if err != nil {
		return
	}
	v = new(RulesTagsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type RulesUpdateOption struct {
	Key                        string `url:"key,omitempty"`                           // Description:"Key of the rule to update",ExampleValue:"javascript:NullCheck"
	MarkdownDescription        string `url:"markdown_description,omitempty"`          // Description:"Rule description (mandatory for custom rule and manual rule) in <a href='/formatting/help'>markdown format</a>",ExampleValue:"Description of my custom rule"
	MarkdownNote               string `url:"markdown_note,omitempty"`                 // Description:"Optional note in <a href='/formatting/help'>markdown format</a>. Use empty value to remove current note. Note is not changed if the parameter is not set.",ExampleValue:"my *note*"
	Name                       string `url:"name,omitempty"`                          // Description:"Rule name (mandatory for custom rule)",ExampleValue:"My custom rule"
	Params                     string `url:"params,omitempty"`                        // Description:"Parameters as semi-colon list of <key>=<value>, for example 'params=key1=v1;key2=v2' (Only when updating a custom rule)",ExampleValue:""
	RemediationFnBaseEffort    string `url:"remediation_fn_base_effort,omitempty"`    // Description:"Base effort of the remediation function of the rule",ExampleValue:"1d"
	RemediationFnType          string `url:"remediation_fn_type,omitempty"`           // Description:"Type of the remediation function of the rule",ExampleValue:""
	RemediationFyGapMultiplier string `url:"remediation_fy_gap_multiplier,omitempty"` // Description:"Gap multiplier of the remediation function of the rule",ExampleValue:"3min"
	Severity                   string `url:"severity,omitempty"`                      // Description:"Rule severity (Only when updating a custom rule)",ExampleValue:""
	Status                     string `url:"status,omitempty"`                        // Description:"Rule status (Only when updating a custom rule)",ExampleValue:""
	Tags                       string `url:"tags,omitempty"`                          // Description:"Optional comma-separated list of tags to set. Use blank value to remove current tags. Tags are not changed if the parameter is not set.",ExampleValue:"java8,security"
}

// Update Update an existing rule.<br>Requires the 'Administer Quality Profiles' permission
func (s *RulesService) Update(opt *RulesUpdateOption) (v *RulesUpdateObject, resp *http.Response, err error) {
	err = s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "rules/update", opt)
	if err != nil {
		return
	}
	v = new(RulesUpdateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
