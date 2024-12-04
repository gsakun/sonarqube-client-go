// Manage quality profiles.
package sonar

import "net/http"

type QualityprofilesService struct {
	client *Client
}

type QualityprofilesChangelogObject struct {
	Events []QualityprofilesChangelogObject_sub2 `json:"events,omitempty"`
	Paging QualityprofilesChangelogObject_sub3   `json:"paging,omitempty"`
}

type QualityprofilesChangelogObject_sub2 struct {
	Action      string                              `json:"action,omitempty"`
	AuthorLogin string                              `json:"authorLogin,omitempty"`
	AuthorName  string                              `json:"authorName,omitempty"`
	Date        string                              `json:"date,omitempty"`
	Params      QualityprofilesChangelogObject_sub1 `json:"params,omitempty"`
	RuleKey     string                              `json:"ruleKey,omitempty"`
	RuleName    string                              `json:"ruleName,omitempty"`
}

type QualityprofilesChangelogObject_sub1 struct {
	Format   string `json:"format,omitempty"`
	Severity string `json:"severity,omitempty"`
}

type QualityprofilesChangelogObject_sub3 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type QualityprofilesCompareObject struct {
	InLeft   []QualityprofilesCompareObject_sub1 `json:"inLeft,omitempty"`
	InRight  []QualityprofilesCompareObject_sub1 `json:"inRight,omitempty"`
	Left     QualityprofilesCompareObject_sub2   `json:"left,omitempty"`
	Modified []QualityprofilesCompareObject_sub5 `json:"modified,omitempty"`
	Right    QualityprofilesCompareObject_sub2   `json:"right,omitempty"`
	Same     []QualityprofilesCompareObject_sub1 `json:"same,omitempty"`
}

type QualityprofilesCompareObject_sub5 struct {
	Key          string                            `json:"key,omitempty"`
	LanguageKey  string                            `json:"languageKey,omitempty"`
	LanguageName string                            `json:"languageName,omitempty"`
	Left         QualityprofilesCompareObject_sub4 `json:"left,omitempty"`
	Name         string                            `json:"name,omitempty"`
	PluginKey    string                            `json:"pluginKey,omitempty"`
	PluginName   string                            `json:"pluginName,omitempty"`
	Right        QualityprofilesCompareObject_sub4 `json:"right,omitempty"`
}

type QualityprofilesCompareObject_sub1 struct {
	Key          string `json:"key,omitempty"`
	LanguageKey  string `json:"languageKey,omitempty"`
	LanguageName string `json:"languageName,omitempty"`
	Name         string `json:"name,omitempty"`
	PluginKey    string `json:"pluginKey,omitempty"`
	PluginName   string `json:"pluginName,omitempty"`
	Severity     string `json:"severity,omitempty"`
}

type QualityprofilesCompareObject_sub2 struct {
	Key  string `json:"key,omitempty"`
	Name string `json:"name,omitempty"`
}

type QualityprofilesCompareObject_sub3 struct {
	MaximumFunctionParameters string `json:"maximumFunctionParameters,omitempty"`
}

type QualityprofilesCompareObject_sub4 struct {
	Params   QualityprofilesCompareObject_sub3 `json:"params,omitempty"`
	Severity string                            `json:"severity,omitempty"`
}

type QualityprofilesCopyObject struct {
	IsDefault   bool   `json:"isDefault,omitempty"`
	IsInherited bool   `json:"isInherited,omitempty"`
	Key         string `json:"key,omitempty"`
	Language    string `json:"language,omitempty"`
	Name        string `json:"name,omitempty"`
	ParentKey   string `json:"parentKey,omitempty"`
}

type QualityprofilesCreateObject struct {
	Profile  QualityprofilesCreateObject_sub1 `json:"profile,omitempty"`
	Warnings []string                         `json:"warnings,omitempty"`
}

type QualityprofilesCreateObject_sub1 struct {
	IsDefault    bool   `json:"isDefault,omitempty"`
	IsInherited  bool   `json:"isInherited,omitempty"`
	Key          string `json:"key,omitempty"`
	Language     string `json:"language,omitempty"`
	LanguageName string `json:"languageName,omitempty"`
	Name         string `json:"name,omitempty"`
}

type QualityprofilesExportersObject struct {
	Exporters []QualityprofilesExportersObject_sub1 `json:"exporters,omitempty"`
}

type QualityprofilesExportersObject_sub1 struct {
	Key       string   `json:"key,omitempty"`
	Languages []string `json:"languages,omitempty"`
	Name      string   `json:"name,omitempty"`
}

type QualityprofilesImportersObject struct {
	Importers []QualityprofilesImportersObject_sub1 `json:"importers,omitempty"`
}

type QualityprofilesImportersObject_sub1 struct {
	Key       string   `json:"key,omitempty"`
	Languages []string `json:"languages,omitempty"`
	Name      string   `json:"name,omitempty"`
}

type QualityprofilesInheritanceObject struct {
	Ancestors []QualityprofilesInheritanceObject_sub1 `json:"ancestors,omitempty"`
	Children  []QualityprofilesInheritanceObject_sub2 `json:"children,omitempty"`
	Profile   QualityprofilesInheritanceObject_sub3   `json:"profile,omitempty"`
}

type QualityprofilesInheritanceObject_sub3 struct {
	ActiveRuleCount     int64  `json:"activeRuleCount,omitempty"`
	IsBuiltIn           bool   `json:"isBuiltIn,omitempty"`
	Key                 string `json:"key,omitempty"`
	Name                string `json:"name,omitempty"`
	OverridingRuleCount int64  `json:"overridingRuleCount,omitempty"`
	Parent              string `json:"parent,omitempty"`
}

type QualityprofilesInheritanceObject_sub2 struct {
	ActiveRuleCount     int64  `json:"activeRuleCount,omitempty"`
	IsBuiltIn           bool   `json:"isBuiltIn,omitempty"`
	Key                 string `json:"key,omitempty"`
	Name                string `json:"name,omitempty"`
	OverridingRuleCount int64  `json:"overridingRuleCount,omitempty"`
}

type QualityprofilesInheritanceObject_sub1 struct {
	ActiveRuleCount int64  `json:"activeRuleCount,omitempty"`
	IsBuiltIn       bool   `json:"isBuiltIn,omitempty"`
	Key             string `json:"key,omitempty"`
	Name            string `json:"name,omitempty"`
	Parent          string `json:"parent,omitempty"`
}

type QualityprofilesProjectsObject struct {
	Paging  QualityprofilesProjectsObject_sub1   `json:"paging,omitempty"`
	Results []QualityprofilesProjectsObject_sub2 `json:"results,omitempty"`
}

type QualityprofilesProjectsObject_sub2 struct {
	ID       string `json:"id,omitempty"`
	Key      string `json:"key,omitempty"`
	Name     string `json:"name,omitempty"`
	Selected bool   `json:"selected,omitempty"`
}

type QualityprofilesProjectsObject_sub1 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type QualityprofilesSearchObject struct {
	Actions  QualityprofilesSearchObject_sub1   `json:"actions,omitempty"`
	Profiles []QualityprofilesSearchObject_sub3 `json:"profiles,omitempty"`
}

type QualityprofilesSearchObject_sub3 struct {
	Actions                   QualityprofilesSearchObject_sub2 `json:"actions,omitempty"`
	ActiveDeprecatedRuleCount int64                            `json:"activeDeprecatedRuleCount,omitempty"`
	ActiveRuleCount           int64                            `json:"activeRuleCount,omitempty"`
	IsBuiltIn                 bool                             `json:"isBuiltIn,omitempty"`
	IsDefault                 bool                             `json:"isDefault,omitempty"`
	IsInherited               bool                             `json:"isInherited,omitempty"`
	Key                       string                           `json:"key,omitempty"`
	Language                  string                           `json:"language,omitempty"`
	LanguageName              string                           `json:"languageName,omitempty"`
	LastUsed                  string                           `json:"lastUsed,omitempty"`
	Name                      string                           `json:"name,omitempty"`
	ParentKey                 string                           `json:"parentKey,omitempty"`
	ParentName                string                           `json:"parentName,omitempty"`
	ProjectCount              int64                            `json:"projectCount,omitempty"`
	RuleUpdatedAt             string                           `json:"ruleUpdatedAt,omitempty"`
	UserUpdatedAt             string                           `json:"userUpdatedAt,omitempty"`
}

type QualityprofilesSearchObject_sub2 struct {
	AssociateProjects bool `json:"associateProjects,omitempty"`
	Copy              bool `json:"copy,omitempty"`
	Delete            bool `json:"delete,omitempty"`
	Edit              bool `json:"edit,omitempty"`
	SetAsDefault      bool `json:"setAsDefault,omitempty"`
}

type QualityprofilesSearchObject_sub1 struct {
	Create bool `json:"create,omitempty"`
}

type QualityprofilesSearchGroupsObject struct {
	Groups []QualityprofilesSearchGroupsObject_sub1 `json:"groups,omitempty"`
	Paging QualityprofilesSearchGroupsObject_sub2   `json:"paging,omitempty"`
}

type QualityprofilesSearchGroupsObject_sub1 struct {
	Description string `json:"description,omitempty"`
	Name        string `json:"name,omitempty"`
	Selected    bool   `json:"selected,omitempty"`
}

type QualityprofilesSearchGroupsObject_sub2 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type QualityprofilesSearchUsersObject struct {
	Paging QualityprofilesSearchUsersObject_sub1   `json:"paging,omitempty"`
	Users  []QualityprofilesSearchUsersObject_sub2 `json:"users,omitempty"`
}

type QualityprofilesSearchUsersObject_sub2 struct {
	Avatar   string `json:"avatar,omitempty"`
	Login    string `json:"login,omitempty"`
	Name     string `json:"name,omitempty"`
	Selected bool   `json:"selected,omitempty"`
}

type QualityprofilesSearchUsersObject_sub1 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type QualityprofilesShowObject struct {
	Profile QualityprofilesShowObject_sub1 `json:"profile,omitempty"`
}

type QualityprofilesShowObject_sub1 struct {
	ActiveDeprecatedRuleCount int64  `json:"activeDeprecatedRuleCount,omitempty"`
	ActiveRuleCount           int64  `json:"activeRuleCount,omitempty"`
	IsBuiltIn                 bool   `json:"isBuiltIn,omitempty"`
	IsDefault                 bool   `json:"isDefault,omitempty"`
	IsInherited               bool   `json:"isInherited,omitempty"`
	Key                       string `json:"key,omitempty"`
	Language                  string `json:"language,omitempty"`
	LanguageName              string `json:"languageName,omitempty"`
	LastUsed                  string `json:"lastUsed,omitempty"`
	Name                      string `json:"name,omitempty"`
	ProjectCount              int64  `json:"projectCount,omitempty"`
	RulesUpdatedAt            string `json:"rulesUpdatedAt,omitempty"`
}

type QualityprofilesActivateRuleOption struct {
	Key      string `url:"key,omitempty"`      // Description:"Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Params   string `url:"params,omitempty"`   // Description:"Parameters as semi-colon list of <code>key=value</code>. Ignored if parameter reset is true.",ExampleValue:"params=key1=v1;key2=v2"
	Reset    string `url:"reset,omitempty"`    // Description:"Reset severity and parameters of activated rule. Set the values defined on parent profile or from rule default values.",ExampleValue:""
	Rule     string `url:"rule,omitempty"`     // Description:"Rule key",ExampleValue:"java:AvoidCycles"
	Severity string `url:"severity,omitempty"` // Description:"Severity. Ignored if parameter reset is true.",ExampleValue:""
}

// ActivateRule Activate a rule on a Quality Profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) ActivateRule(opt *QualityprofilesActivateRuleOption) (resp *http.Response, err error) {
	err = s.ValidateActivateRuleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/activate_rule", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesActivateRulesOption struct {
	Activation          string `url:"activation,omitempty"`          // Description:"Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.",ExampleValue:""
	ActiveSeverities    string `url:"active_severities,omitempty"`   // Description:"Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	Asc                 string `url:"asc,omitempty"`                 // Description:"Ascending sort",ExampleValue:""
	AvailableSince      string `url:"available_since,omitempty"`     // Description:"Filters rules added since date. Format is yyyy-MM-dd",ExampleValue:"2014-06-22"
	CompareToProfile    string `url:"compareToProfile,omitempty"`    // Description:"Quality profile key to filter rules that are activated. Meant to compare easily to profile set in 'qprofile'",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
	Cwe                 string `url:"cwe,omitempty"`                 // Description:"Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE.",ExampleValue:"12,125,unknown"
	Inheritance         string `url:"inheritance,omitempty"`         // Description:"Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.",ExampleValue:"INHERITED,OVERRIDES"
	IsTemplate          string `url:"is_template,omitempty"`         // Description:"Filter template rules",ExampleValue:""
	Languages           string `url:"languages,omitempty"`           // Description:"Comma-separated list of languages",ExampleValue:"java,js"
	OwaspTop10          string `url:"owaspTop10,omitempty"`          // Description:"Comma-separated list of OWASP Top 10 2017 lowercase categories.",ExampleValue:""
	OwaspTop102021      string `url:"owaspTop10-2021,omitempty"`     // Description:"Comma-separated list of OWASP Top 10 2021 lowercase categories.",ExampleValue:""
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
	TargetKey           string `url:"targetKey,omitempty"`           // Description:"Quality Profile key on which the rule activation is done. To retrieve a quality profile key please see <code>api/qualityprofiles/search</code>",ExampleValue:"AU-TpxcA-iU5OvuD2FL0"
	TargetSeverity      string `url:"targetSeverity,omitempty"`      // Description:"Severity to set on the activated rules",ExampleValue:""
	TemplateKey         string `url:"template_key,omitempty"`        // Description:"Key of the template rule to filter on. Used to search for the custom rules based on this template.",ExampleValue:"java:S001"
	Types               string `url:"types,omitempty"`               // Description:"Comma-separated list of types. Returned rules match any of the tags (OR operator)",ExampleValue:"BUG"
}

// ActivateRules Bulk-activate rules on one quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) ActivateRules(opt *QualityprofilesActivateRulesOption) (resp *http.Response, err error) {
	err = s.ValidateActivateRulesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/activate_rules", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesAddGroupOption struct {
	Group          string `url:"group,omitempty"`          // Description:"Group name",ExampleValue:"sonar-administrators"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality Profile name",ExampleValue:"Recommended quality profile"
}

// AddGroup Allow a group to edit a Quality Profile.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) AddGroup(opt *QualityprofilesAddGroupOption) (resp *http.Response, err error) {
	err = s.ValidateAddGroupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/add_group", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesAddProjectOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language.",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name.",ExampleValue:"Sonar way"
}

// AddProject Associate a project with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func (s *QualityprofilesService) AddProject(opt *QualityprofilesAddProjectOption) (resp *http.Response, err error) {
	err = s.ValidateAddProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/add_project", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesAddUserOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	Login          string `url:"login,omitempty"`          // Description:"User login",ExampleValue:"john.doe"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality Profile name",ExampleValue:"Recommended quality profile"
}

// AddUser Allow a user to edit a Quality Profile.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) AddUser(opt *QualityprofilesAddUserOption) (resp *http.Response, err error) {
	err = s.ValidateAddUserOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/add_user", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesBackupOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language.",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name.",ExampleValue:"Sonar way"
}

// Backup Backup a quality profile in XML form. The exported profile can be restored through api/qualityprofiles/restore.
func (s *QualityprofilesService) Backup(opt *QualityprofilesBackupOption) (v *string, resp *http.Response, err error) {
	err = s.ValidateBackupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/backup", opt)
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

type QualityprofilesChangeParentOption struct {
	Language             string `url:"language,omitempty"`             // Description:"Quality profile language.",ExampleValue:""
	ParentQualityProfile string `url:"parentQualityProfile,omitempty"` // Description:"New parent profile name. <br> If no profile is provided, the inheritance link with current parent profile (if any) is broken, which deactivates all rules which come from the parent and are not overridden.",ExampleValue:"Sonar way"
	QualityProfile       string `url:"qualityProfile,omitempty"`       // Description:"Quality profile name.",ExampleValue:"Sonar way"
}

// ChangeParent Change a quality profile's parent.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) ChangeParent(opt *QualityprofilesChangeParentOption) (resp *http.Response, err error) {
	err = s.ValidateChangeParentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/change_parent", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesChangelogOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language.",ExampleValue:""
	P              string `url:"p,omitempty"`              // Description:"1-based page number",ExampleValue:"42"
	Ps             string `url:"ps,omitempty"`             // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name.",ExampleValue:"Sonar way"
	Since          string `url:"since,omitempty"`          // Description:"Start date for the changelog (inclusive). <br>Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	To             string `url:"to,omitempty"`             // Description:"End date for the changelog (exclusive, strictly before). <br>Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
}

// Changelog Get the history of changes on a quality profile: rule activation/deactivation, change in parameters/severity. Events are ordered by date in descending order (most recent first).
func (s *QualityprofilesService) Changelog(opt *QualityprofilesChangelogOption) (v *QualityprofilesChangelogObject, resp *http.Response, err error) {
	err = s.ValidateChangelogOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/changelog", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesChangelogObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesCompareOption struct {
	LeftKey  string `url:"leftKey,omitempty"`  // Description:"Profile key.",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	RightKey string `url:"rightKey,omitempty"` // Description:"Another profile key.",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
}

// Compare Compare two quality profiles.
func (s *QualityprofilesService) Compare(opt *QualityprofilesCompareOption) (v *QualityprofilesCompareObject, resp *http.Response, err error) {
	err = s.ValidateCompareOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/compare", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesCompareObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesCopyOption struct {
	FromKey string `url:"fromKey,omitempty"` // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ToName  string `url:"toName,omitempty"`  // Description:"Name for the new quality profile.",ExampleValue:"My Sonar way"
}

// Copy Copy a quality profile.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Copy(opt *QualityprofilesCopyOption) (v *QualityprofilesCopyObject, resp *http.Response, err error) {
	err = s.ValidateCopyOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/copy", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesCopyObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesCreateOption struct {
	Language string `url:"language,omitempty"` // Description:"Quality profile language",ExampleValue:"js"
	Name     string `url:"name,omitempty"`     // Description:"Quality profile name",ExampleValue:"My Sonar way"
}

// Create Create a quality profile.<br>Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Create(opt *QualityprofilesCreateOption) (v *QualityprofilesCreateObject, resp *http.Response, err error) {
	err = s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/create", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesCreateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesDeactivateRuleOption struct {
	Key  string `url:"key,omitempty"`  // Description:"Quality Profile key. Can be obtained through <code>api/qualityprofiles/search</code>",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Rule string `url:"rule,omitempty"` // Description:"Rule key",ExampleValue:"java:S1144"
}

// DeactivateRule Deactivate a rule on a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) DeactivateRule(opt *QualityprofilesDeactivateRuleOption) (resp *http.Response, err error) {
	err = s.ValidateDeactivateRuleOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/deactivate_rule", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesDeactivateRulesOption struct {
	Activation          string `url:"activation,omitempty"`          // Description:"Filter rules that are activated or deactivated on the selected Quality profile. Ignored if the parameter 'qprofile' is not set.",ExampleValue:""
	ActiveSeverities    string `url:"active_severities,omitempty"`   // Description:"Comma-separated list of activation severities, i.e the severity of rules in Quality profiles.",ExampleValue:"CRITICAL,BLOCKER"
	Asc                 string `url:"asc,omitempty"`                 // Description:"Ascending sort",ExampleValue:""
	AvailableSince      string `url:"available_since,omitempty"`     // Description:"Filters rules added since date. Format is yyyy-MM-dd",ExampleValue:"2014-06-22"
	CompareToProfile    string `url:"compareToProfile,omitempty"`    // Description:"Quality profile key to filter rules that are activated. Meant to compare easily to profile set in 'qprofile'",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
	Cwe                 string `url:"cwe,omitempty"`                 // Description:"Comma-separated list of CWE identifiers. Use 'unknown' to select rules not associated to any CWE.",ExampleValue:"12,125,unknown"
	Inheritance         string `url:"inheritance,omitempty"`         // Description:"Comma-separated list of values of inheritance for a rule within a quality profile. Used only if the parameter 'activation' is set.",ExampleValue:"INHERITED,OVERRIDES"
	IsTemplate          string `url:"is_template,omitempty"`         // Description:"Filter template rules",ExampleValue:""
	Languages           string `url:"languages,omitempty"`           // Description:"Comma-separated list of languages",ExampleValue:"java,js"
	OwaspTop10          string `url:"owaspTop10,omitempty"`          // Description:"Comma-separated list of OWASP Top 10 2017 lowercase categories.",ExampleValue:""
	OwaspTop102021      string `url:"owaspTop10-2021,omitempty"`     // Description:"Comma-separated list of OWASP Top 10 2021 lowercase categories.",ExampleValue:""
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
	TargetKey           string `url:"targetKey,omitempty"`           // Description:"Quality Profile key on which the rule deactivation is done. To retrieve a profile key please see <code>api/qualityprofiles/search</code>",ExampleValue:"AU-TpxcA-iU5OvuD2FL1"
	TemplateKey         string `url:"template_key,omitempty"`        // Description:"Key of the template rule to filter on. Used to search for the custom rules based on this template.",ExampleValue:"java:S001"
	Types               string `url:"types,omitempty"`               // Description:"Comma-separated list of types. Returned rules match any of the tags (OR operator)",ExampleValue:"BUG"
}

// DeactivateRules Bulk deactivate rules on Quality profiles.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) DeactivateRules(opt *QualityprofilesDeactivateRulesOption) (resp *http.Response, err error) {
	err = s.ValidateDeactivateRulesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/deactivate_rules", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesDeleteOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language.",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name.",ExampleValue:"Sonar way"
}

// Delete Delete a quality profile and all its descendants. The default quality profile cannot be deleted.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Delete(opt *QualityprofilesDeleteOption) (resp *http.Response, err error) {
	err = s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/delete", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesExportOption struct {
	ExporterKey    string `url:"exporterKey,omitempty"`    // Description:"Output format. If left empty, the same format as api/qualityprofiles/backup is used. Possible values are described by api/qualityprofiles/exporters.",ExampleValue:""
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:"py"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name to export. If left empty, the default profile for the language is exported.",ExampleValue:"My Sonar way"
}

// Export Export a quality profile.
func (s *QualityprofilesService) Export(opt *QualityprofilesExportOption) (v *string, resp *http.Response, err error) {
	err = s.ValidateExportOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/export", opt)
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

// Exporters Lists available profile export formats.
func (s *QualityprofilesService) Exporters() (v *QualityprofilesExportersObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "qualityprofiles/exporters", nil)
	if err != nil {
		return
	}
	v = new(QualityprofilesExportersObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Importers List supported importers.
func (s *QualityprofilesService) Importers() (v *QualityprofilesImportersObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "qualityprofiles/importers", nil)
	if err != nil {
		return
	}
	v = new(QualityprofilesImportersObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesInheritanceOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language.",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name.",ExampleValue:"Sonar way"
}

// Inheritance Show a quality profile's ancestors and children.
func (s *QualityprofilesService) Inheritance(opt *QualityprofilesInheritanceOption) (v *QualityprofilesInheritanceObject, resp *http.Response, err error) {
	err = s.ValidateInheritanceOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/inheritance", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesInheritanceObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesProjectsOption struct {
	Key      string `url:"key,omitempty"`      // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q        string `url:"q,omitempty"`        // Description:"Limit search to projects that contain the supplied string.",ExampleValue:"sonar"
	Selected string `url:"selected,omitempty"` // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Projects List projects with their association status regarding a quality profile <br/>See api/qualityprofiles/search in order to get the Quality Profile Key
func (s *QualityprofilesService) Projects(opt *QualityprofilesProjectsOption) (v *QualityprofilesProjectsObject, resp *http.Response, err error) {
	err = s.ValidateProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/projects", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesProjectsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesRemoveGroupOption struct {
	Group          string `url:"group,omitempty"`          // Description:"Group name",ExampleValue:"sonar-administrators"
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality Profile name",ExampleValue:"Recommended quality profile"
}

// RemoveGroup Remove the ability from a group to edit a Quality Profile.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) RemoveGroup(opt *QualityprofilesRemoveGroupOption) (resp *http.Response, err error) {
	err = s.ValidateRemoveGroupOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/remove_group", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesRemoveProjectOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language.",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name.",ExampleValue:"Sonar way"
}

// RemoveProject Remove a project's association with a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li>  <li>Administer right on the specified project</li></ul>
func (s *QualityprofilesService) RemoveProject(opt *QualityprofilesRemoveProjectOption) (resp *http.Response, err error) {
	err = s.ValidateRemoveProjectOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/remove_project", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesRemoveUserOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	Login          string `url:"login,omitempty"`          // Description:"User login",ExampleValue:"john.doe"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality Profile name",ExampleValue:"Recommended quality profile"
}

// RemoveUser Remove the ability from a user to edit a Quality Profile.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) RemoveUser(opt *QualityprofilesRemoveUserOption) (resp *http.Response, err error) {
	err = s.ValidateRemoveUserOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/remove_user", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesRenameOption struct {
	Key  string `url:"key,omitempty"`  // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Name string `url:"name,omitempty"` // Description:"New quality profile name",ExampleValue:"My Sonar way"
}

// Rename Rename a quality profile.<br> Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) Rename(opt *QualityprofilesRenameOption) (resp *http.Response, err error) {
	err = s.ValidateRenameOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/rename", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesRestoreOption struct {
	Backup string `url:"backup,omitempty"` // Description:"A profile backup file in XML format, as generated by api/qualityprofiles/backup or the former api/profiles/backup.",ExampleValue:""
}

// Restore Restore a quality profile using an XML file. The restored profile name is taken from the backup file, so if a profile with the same name and language already exists, it will be overwritten.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) Restore(opt *QualityprofilesRestoreOption) (resp *http.Response, err error) {
	err = s.ValidateRestoreOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/restore", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesSearchOption struct {
	Defaults       string `url:"defaults,omitempty"`       // Description:"If set to true, return only the quality profiles marked as default for each language",ExampleValue:""
	Language       string `url:"language,omitempty"`       // Description:"Language key. If provided, only profiles for the given language are returned.",ExampleValue:""
	Project        string `url:"project,omitempty"`        // Description:"Project key",ExampleValue:"my_project"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name",ExampleValue:"SonarQube Way"
}

// Search Search quality profiles
func (s *QualityprofilesService) Search(opt *QualityprofilesSearchOption) (v *QualityprofilesSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/search", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesSearchGroupsOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	P              string `url:"p,omitempty"`              // Description:"1-based page number",ExampleValue:"42"
	Ps             string `url:"ps,omitempty"`             // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q              string `url:"q,omitempty"`              // Description:"Limit search to group names that contain the supplied string.",ExampleValue:"sonar"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality Profile name",ExampleValue:"Recommended quality profile"
	Selected       string `url:"selected,omitempty"`       // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// SearchGroups List the groups that are allowed to edit a Quality Profile.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) SearchGroups(opt *QualityprofilesSearchGroupsOption) (v *QualityprofilesSearchGroupsObject, resp *http.Response, err error) {
	err = s.ValidateSearchGroupsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/search_groups", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesSearchGroupsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesSearchUsersOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language",ExampleValue:""
	P              string `url:"p,omitempty"`              // Description:"1-based page number",ExampleValue:"42"
	Ps             string `url:"ps,omitempty"`             // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q              string `url:"q,omitempty"`              // Description:"Limit search to names or logins that contain the supplied string.",ExampleValue:"freddy"
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality Profile name",ExampleValue:"Recommended quality profile"
	Selected       string `url:"selected,omitempty"`       // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// SearchUsers List the users that are allowed to edit a Quality Profile.<br>Requires one of the following permissions:<ul>  <li>'Administer Quality Profiles'</li>  <li>Edit right on the specified quality profile</li></ul>
func (s *QualityprofilesService) SearchUsers(opt *QualityprofilesSearchUsersOption) (v *QualityprofilesSearchUsersObject, resp *http.Response, err error) {
	err = s.ValidateSearchUsersOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/search_users", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesSearchUsersObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type QualityprofilesSetDefaultOption struct {
	Language       string `url:"language,omitempty"`       // Description:"Quality profile language.",ExampleValue:""
	QualityProfile string `url:"qualityProfile,omitempty"` // Description:"Quality profile name.",ExampleValue:"Sonar way"
}

// SetDefault Select the default profile for a given language.<br> Requires to be logged in and the 'Administer Quality Profiles' permission.
func (s *QualityprofilesService) SetDefault(opt *QualityprofilesSetDefaultOption) (resp *http.Response, err error) {
	err = s.ValidateSetDefaultOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "qualityprofiles/set_default", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type QualityprofilesShowOption struct {
	CompareToSonarWay string `url:"compareToSonarWay,omitempty"` // Description:"Add the number of missing rules from the related Sonar way profile in the response",ExampleValue:""
	Key               string `url:"key,omitempty"`               // Description:"Quality profile key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// Show Show a quality profile
func (s *QualityprofilesService) Show(opt *QualityprofilesShowOption) (v *QualityprofilesShowObject, resp *http.Response, err error) {
	err = s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "qualityprofiles/show", opt)
	if err != nil {
		return
	}
	v = new(QualityprofilesShowObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
