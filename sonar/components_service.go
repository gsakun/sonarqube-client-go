// Get information about a component (file, directory, project, ...) and its ancestors or descendants. Update a project or module key.
package sonar

import "net/http"

type ComponentsService struct {
	client *Client
}

type ComponentsAppObject struct {
	CanCreateManualIssue bool                     `json:"canCreateManualIssue,omitempty"`
	CanMarkAsFavorite    bool                     `json:"canMarkAsFavorite,omitempty"`
	Fav                  bool                     `json:"fav,omitempty"`
	Key                  string                   `json:"key,omitempty"`
	LongName             string                   `json:"longName,omitempty"`
	Measures             ComponentsAppObject_sub1 `json:"measures,omitempty"`
	Name                 string                   `json:"name,omitempty"`
	Project              string                   `json:"project,omitempty"`
	ProjectName          string                   `json:"projectName,omitempty"`
	Q                    string                   `json:"q,omitempty"`
	UUID                 string                   `json:"uuid,omitempty"`
}

type ComponentsAppObject_sub1 struct {
	Debt               string `json:"debt,omitempty"`
	DebtRatio          string `json:"debtRatio,omitempty"`
	DuplicationDensity string `json:"duplicationDensity,omitempty"`
	Issues             string `json:"issues,omitempty"`
	Lines              string `json:"lines,omitempty"`
	SqaleRating        string `json:"sqaleRating,omitempty"`
}

type ComponentsSearchObject struct {
	Components []ComponentsSearchObject_sub1 `json:"components,omitempty"`
	Paging     ComponentsSearchObject_sub2   `json:"paging,omitempty"`
}

type ComponentsSearchObject_sub1 struct {
	Key       string `json:"key,omitempty"`
	Name      string `json:"name,omitempty"`
	Project   string `json:"project,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type ComponentsSearchObject_sub2 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type ComponentsSearchProjectsObject struct {
	Components []ComponentsSearchProjectsObject_sub1 `json:"components,omitempty"`
	Facets     []ComponentsSearchProjectsObject_sub3 `json:"facets,omitempty"`
	Paging     ComponentsSearchProjectsObject_sub4   `json:"paging,omitempty"`
}

type ComponentsSearchProjectsObject_sub2 struct {
	Count int64  `json:"count,omitempty"`
	Val   string `json:"val,omitempty"`
}

type ComponentsSearchProjectsObject_sub1 struct {
	IsFavorite    bool     `json:"isFavorite,omitempty"`
	Key           string   `json:"key,omitempty"`
	Name          string   `json:"name,omitempty"`
	NeedIssueSync bool     `json:"needIssueSync,omitempty"`
	Qualifier     string   `json:"qualifier,omitempty"`
	Tags          []string `json:"tags,omitempty"`
	Visibility    string   `json:"visibility,omitempty"`
}

type ComponentsSearchProjectsObject_sub4 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type ComponentsSearchProjectsObject_sub3 struct {
	Property string                                `json:"property,omitempty"`
	Values   []ComponentsSearchProjectsObject_sub2 `json:"values,omitempty"`
}

type ComponentsShowObject struct {
	Ancestors []ComponentsShowObject_sub1 `json:"ancestors,omitempty"`
	Component ComponentsShowObject_sub2   `json:"component,omitempty"`
}

type ComponentsShowObject_sub1 struct {
	AnalysisDate string   `json:"analysisDate,omitempty"`
	Description  string   `json:"description,omitempty"`
	Key          string   `json:"key,omitempty"`
	Name         string   `json:"name,omitempty"`
	Path         string   `json:"path,omitempty"`
	Qualifier    string   `json:"qualifier,omitempty"`
	Tags         []string `json:"tags,omitempty"`
	Version      string   `json:"version,omitempty"`
	Visibility   string   `json:"visibility,omitempty"`
}

type ComponentsShowObject_sub2 struct {
	AnalysisDate   string `json:"analysisDate,omitempty"`
	Key            string `json:"key,omitempty"`
	Language       string `json:"language,omitempty"`
	LeakPeriodDate string `json:"leakPeriodDate,omitempty"`
	Name           string `json:"name,omitempty"`
	Path           string `json:"path,omitempty"`
	Qualifier      string `json:"qualifier,omitempty"`
	Version        string `json:"version,omitempty"`
}

type ComponentsSuggestionsObject struct {
	Projects []interface{}                      `json:"projects,omitempty"`
	Results  []ComponentsSuggestionsObject_sub2 `json:"results,omitempty"`
}

type ComponentsSuggestionsObject_sub1 struct {
	IsFavorite        bool   `json:"isFavorite,omitempty"`
	IsRecentlyBrowsed bool   `json:"isRecentlyBrowsed,omitempty"`
	Key               string `json:"key,omitempty"`
	Match             string `json:"match,omitempty"`
	Name              string `json:"name,omitempty"`
	Project           string `json:"project,omitempty"`
}

type ComponentsSuggestionsObject_sub2 struct {
	Items []ComponentsSuggestionsObject_sub1 `json:"items,omitempty"`
	More  int64                              `json:"more,omitempty"`
	Q     string                             `json:"q,omitempty"`
}

type ComponentsTreeObject struct {
	BaseComponent ComponentsTreeObject_sub1   `json:"baseComponent,omitempty"`
	Components    []ComponentsTreeObject_sub2 `json:"components,omitempty"`
	Paging        ComponentsTreeObject_sub3   `json:"paging,omitempty"`
}

type ComponentsTreeObject_sub1 struct {
	Description string   `json:"description,omitempty"`
	Key         string   `json:"key,omitempty"`
	Qualifier   string   `json:"qualifier,omitempty"`
	Tags        []string `json:"tags,omitempty"`
	Visibility  string   `json:"visibility,omitempty"`
}

type ComponentsTreeObject_sub2 struct {
	Key       string `json:"key,omitempty"`
	Language  string `json:"language,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type ComponentsTreeObject_sub3 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type ComponentsAppOption struct {
	Branch      string `url:"branch,omitempty"`      // Description:"Branch key. Not available in the community edition.",ExampleValue:"feature/my_branch"
	Component   string `url:"component,omitempty"`   // Description:"Component key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id. Not available in the community edition.",ExampleValue:"5461"
}

// App Coverage data required for rendering the component viewer.<br>Either branch or pull request can be provided, not both<br>Requires the following permission: 'Browse'.
func (s *ComponentsService) App(opt *ComponentsAppOption) (v *ComponentsAppObject, resp *http.Response, err error) {
	err = s.ValidateAppOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/app", opt)
	if err != nil {
		return
	}
	v = new(ComponentsAppObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ComponentsSearchOption struct {
	P          string `url:"p,omitempty"`          // Description:"1-based page number",ExampleValue:"42"
	Ps         string `url:"ps,omitempty"`         // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q          string `url:"q,omitempty"`          // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"sonar"
	Qualifiers string `url:"qualifiers,omitempty"` // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>TRK - Projects</li></ul>",ExampleValue:""
}

// Search Search for components
func (s *ComponentsService) Search(opt *ComponentsSearchOption) (v *ComponentsSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/search", opt)
	if err != nil {
		return
	}
	v = new(ComponentsSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ComponentsSearchProjectsOption struct {
	Asc    string `url:"asc,omitempty"`    // Description:"Ascending sort",ExampleValue:""
	F      string `url:"f,omitempty"`      // Description:"Comma-separated list of the fields to be returned in response",ExampleValue:""
	Facets string `url:"facets,omitempty"` // Description:"Comma-separated list of the facets to be computed. No facet is computed by default.",ExampleValue:""
	Filter string `url:"filter,omitempty"` // Description:"Filter of projects on name, key, measure value, quality gate, language, tag or whether a project is a favorite or not.<br>The filter must be encoded to form a valid URL (for example '=' must be replaced by '%3D').<br>Examples of use:<ul> <li>to filter my favorite projects with a failed quality gate and a coverage greater than or equals to 60% and a coverage strictly lower than 80%:<br>   <code>filter="alert_status = ERROR and isFavorite and coverage >= 60 and coverage < 80"</code></li> <li>to filter projects with a reliability, security and maintainability rating equals or worse than B:<br>   <code>filter="reliability_rating>=2 and security_rating>=2 and sqale_rating>=2"</code></li> <li>to filter projects without duplication data:<br>   <code>filter="duplicated_lines_density = NO_DATA"</code></li></ul>To filter on project name or key, use the 'query' keyword, for instance : <code>filter='query = "Sonar"'</code>.<br><br>To filter on a numeric metric, provide the metric key.<br>These are the supported metric keys:<br><ul><li>alert_status</li><li>coverage</li><li>duplicated_lines_density</li><li>lines</li><li>ncloc</li><li>ncloc_language_distribution</li><li>new_coverage</li><li>new_duplicated_lines_density</li><li>new_lines</li><li>new_maintainability_rating</li><li>new_reliability_rating</li><li>new_security_hotspots_reviewed</li><li>new_security_rating</li><li>new_security_review_rating</li><li>reliability_rating</li><li>security_hotspots_reviewed</li><li>security_rating</li><li>security_review_rating</li><li>sqale_rating</li></ul><br>To filter on a rating, provide the corresponding metric key (ex: reliability_rating for reliability rating).<br>The possible values are:<ul> <li>'1' for rating A</li> <li>'2' for rating B</li> <li>'3' for rating C</li> <li>'4' for rating D</li> <li>'5' for rating E</li></ul>To filter on a Quality Gate status use the metric key 'alert_status'. Only the '=' operator can be used.<br>The possible values are:<ul> <li>'OK' for Passed</li> <li>'WARN' for Warning</li> <li>'ERROR' for Failed</li></ul>To filter on language keys use the language key: <ul> <li>to filter on a single language you can use 'language = java'</li> <li>to filter on several languages you must use 'language IN (java, js)'</li></ul>Use the WS api/languages/list to find the key of a language.<br> To filter on tags use the 'tags' keyword:<ul> <li>to filter on one tag you can use <code>tags = finance</code></li> <li>to filter on several tags you must use <code>tags in (offshore, java)</code></li></ul>To filter on a qualifier use key 'qualifier'. Only the '=' operator can be used.<br>The possible values are:<ul> <li>TRK - for projects</li> <li>APP - for applications</li></ul>",ExampleValue:""
	P      string `url:"p,omitempty"`      // Description:"1-based page number",ExampleValue:"42"
	Ps     string `url:"ps,omitempty"`     // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	S      string `url:"s,omitempty"`      // Description:"Sort projects by numeric metric key, quality gate status (using 'alert_status'), last analysis date (using 'analysisDate'), or by project name.",ExampleValue:""
}

// SearchProjects Search for projects
func (s *ComponentsService) SearchProjects(opt *ComponentsSearchProjectsOption) (v *ComponentsSearchProjectsObject, resp *http.Response, err error) {
	err = s.ValidateSearchProjectsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/search_projects", opt)
	if err != nil {
		return
	}
	v = new(ComponentsSearchProjectsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ComponentsShowOption struct {
	Branch      string `url:"branch,omitempty"`      // Description:"Branch key. Not available in the community edition.",ExampleValue:"feature/my_branch"
	Component   string `url:"component,omitempty"`   // Description:"Component key",ExampleValue:"my_project"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id. Not available in the community edition.",ExampleValue:"5461"
}

// Show Returns a component (file, directory, project, portfolio…) and its ancestors. The ancestors are ordered from the parent to the root project. Requires the following permission: 'Browse' on the project of the specified component.
func (s *ComponentsService) Show(opt *ComponentsShowOption) (v *ComponentsShowObject, resp *http.Response, err error) {
	err = s.ValidateShowOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/show", opt)
	if err != nil {
		return
	}
	v = new(ComponentsShowObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ComponentsSuggestionsOption struct {
	More            string `url:"more,omitempty"`            // Description:"Category, for which to display the next 20 results (skipping the first 6 results)",ExampleValue:""
	RecentlyBrowsed string `url:"recentlyBrowsed,omitempty"` // Description:"Comma separated list of component keys, that have recently been browsed by the user. Only the first 50 items will be used. Order is not taken into account.",ExampleValue:"org.sonarsource:sonarqube,some.other:project"
	S               string `url:"s,omitempty"`               // Description:"Search query: can contain several search tokens separated by spaces.",ExampleValue:"sonar"
}

// Suggestions Internal WS for the top-right search engine. The result will contain component search results, grouped by their qualifiers.<p>Each result contains:<ul><li>the component key</li><li>the component's name (unescaped)</li><li>optionally a display name, which puts emphasis to matching characters (this text contains html tags and parts of the html-escaped name)</li></ul>
func (s *ComponentsService) Suggestions(opt *ComponentsSuggestionsOption) (v *ComponentsSuggestionsObject, resp *http.Response, err error) {
	err = s.ValidateSuggestionsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/suggestions", opt)
	if err != nil {
		return
	}
	v = new(ComponentsSuggestionsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type ComponentsTreeOption struct {
	Asc         string `url:"asc,omitempty"`         // Description:"Ascending sort",ExampleValue:""
	Branch      string `url:"branch,omitempty"`      // Description:"Branch key. Not available in the community edition.",ExampleValue:"feature/my_branch"
	Component   string `url:"component,omitempty"`   // Description:"Base component key. The search is based on this component.",ExampleValue:"my_project"
	P           string `url:"p,omitempty"`           // Description:"1-based page number",ExampleValue:"42"
	Ps          string `url:"ps,omitempty"`          // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	PullRequest string `url:"pullRequest,omitempty"` // Description:"Pull request id. Not available in the community edition.",ExampleValue:"5461"
	Q           string `url:"q,omitempty"`           // Description:"Limit search to: <ul><li>component names that contain the supplied string</li><li>component keys that are exactly the same as the supplied string</li></ul>",ExampleValue:"FILE_NAM"
	Qualifiers  string `url:"qualifiers,omitempty"`  // Description:"Comma-separated list of component qualifiers. Filter the results with the specified qualifiers. Possible values are:<ul><li>BRC - no description available</li><li>UTS - Test Files</li><li>FIL - Files</li><li>DIR - Directories</li><li>TRK - Projects</li></ul>",ExampleValue:""
	S           string `url:"s,omitempty"`           // Description:"Comma-separated list of sort fields",ExampleValue:"name, path"
	Strategy    string `url:"strategy,omitempty"`    // Description:"Strategy to search for base component descendants:<ul><li>children: return the children components of the base component. Grandchildren components are not returned</li><li>all: return all the descendants components of the base component. Grandchildren are returned.</li><li>leaves: return all the descendant components (files, in general) which don't have other children. They are the leaves of the component tree.</li></ul>",ExampleValue:""
}

// Tree Navigate through components based on the chosen strategy.<br>Requires the following permission: 'Browse' on the specified project.<br>When limiting search with the q parameter, directories are not returned.
func (s *ComponentsService) Tree(opt *ComponentsTreeOption) (v *ComponentsTreeObject, resp *http.Response, err error) {
	err = s.ValidateTreeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "components/tree", opt)
	if err != nil {
		return
	}
	v = new(ComponentsTreeObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
