// Manage users.
package sonar

import "net/http"

type UsersService struct {
	client *Client
}

type UsersCreateObject struct {
	User UsersCreateObject_sub1 `json:"user,omitempty"`
}

type UsersCreateObject_sub1 struct {
	Active      bool     `json:"active,omitempty"`
	Email       string   `json:"email,omitempty"`
	Local       bool     `json:"local,omitempty"`
	Login       string   `json:"login,omitempty"`
	Name        string   `json:"name,omitempty"`
	ScmAccounts []string `json:"scmAccounts,omitempty"`
}

type UsersCurrentObject struct {
	Avatar                      string                  `json:"avatar,omitempty"`
	DismissedNotices            UsersCurrentObject_sub1 `json:"dismissedNotices,omitempty"`
	Email                       string                  `json:"email,omitempty"`
	ExternalIdentity            string                  `json:"externalIdentity,omitempty"`
	ExternalProvider            string                  `json:"externalProvider,omitempty"`
	Groups                      []string                `json:"groups,omitempty"`
	Homepage                    UsersCurrentObject_sub2 `json:"homepage,omitempty"`
	IsLoggedIn                  bool                    `json:"isLoggedIn,omitempty"`
	Local                       bool                    `json:"local,omitempty"`
	Login                       string                  `json:"login,omitempty"`
	Name                        string                  `json:"name,omitempty"`
	Permissions                 UsersCurrentObject_sub3 `json:"permissions,omitempty"`
	ScmAccounts                 []string                `json:"scmAccounts,omitempty"`
	UsingSonarLintConnectedMode bool                    `json:"usingSonarLintConnectedMode,omitempty"`
}

type UsersCurrentObject_sub2 struct {
	Component string `json:"component,omitempty"`
	Type      string `json:"type,omitempty"`
}

type UsersCurrentObject_sub1 struct {
	EducationPrinciples bool `json:"educationPrinciples,omitempty"`
	SonarlintAd         bool `json:"sonarlintAd,omitempty"`
}

type UsersCurrentObject_sub3 struct {
	Global []string `json:"global,omitempty"`
}

type UsersDeactivateObject struct {
	User UsersDeactivateObject_sub1 `json:"user,omitempty"`
}

type UsersDeactivateObject_sub1 struct {
	Active      bool          `json:"active,omitempty"`
	Groups      []interface{} `json:"groups,omitempty"`
	Local       bool          `json:"local,omitempty"`
	Login       string        `json:"login,omitempty"`
	Name        string        `json:"name,omitempty"`
	ScmAccounts []interface{} `json:"scmAccounts,omitempty"`
}

type UsersGroupsObject struct {
	Groups []UsersGroupsObject_sub1 `json:"groups,omitempty"`
	Paging UsersGroupsObject_sub2   `json:"paging,omitempty"`
}

type UsersGroupsObject_sub1 struct {
	Default     bool   `json:"default,omitempty"`
	Description string `json:"description,omitempty"`
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Selected    bool   `json:"selected,omitempty"`
}

type UsersGroupsObject_sub2 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type UsersIdentityProvidersObject struct {
	IdentityProviders []UsersIdentityProvidersObject_sub1 `json:"identityProviders,omitempty"`
}

type UsersIdentityProvidersObject_sub1 struct {
	BackgroundColor string `json:"backgroundColor,omitempty"`
	HelpMessage     string `json:"helpMessage,omitempty"`
	IconPath        string `json:"iconPath,omitempty"`
	Key             string `json:"key,omitempty"`
	Name            string `json:"name,omitempty"`
}

type UsersSearchObject struct {
	Paging UsersSearchObject_sub1   `json:"paging,omitempty"`
	Users  []UsersSearchObject_sub2 `json:"users,omitempty"`
}

type UsersSearchObject_sub2 struct {
	Active           bool     `json:"active,omitempty"`
	Avatar           string   `json:"avatar,omitempty"`
	Email            string   `json:"email,omitempty"`
	ExternalIdentity string   `json:"externalIdentity,omitempty"`
	ExternalProvider string   `json:"externalProvider,omitempty"`
	Groups           []string `json:"groups,omitempty"`
	Local            bool     `json:"local,omitempty"`
	Login            string   `json:"login,omitempty"`
	Name             string   `json:"name,omitempty"`
	ScmAccounts      []string `json:"scmAccounts,omitempty"`
	TokensCount      int64    `json:"tokensCount,omitempty"`
}

type UsersSearchObject_sub1 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type UsersUpdateObject struct {
	User UsersUpdateObject_sub1 `json:"user,omitempty"`
}

type UsersUpdateObject_sub1 struct {
	Active      bool     `json:"active,omitempty"`
	Email       string   `json:"email,omitempty"`
	Local       bool     `json:"local,omitempty"`
	Login       string   `json:"login,omitempty"`
	Name        string   `json:"name,omitempty"`
	ScmAccounts []string `json:"scmAccounts,omitempty"`
}

type UsersAnonymizeOption struct {
	Login string `url:"login,omitempty"` // Description:"User login",ExampleValue:"myuser"
}

// Anonymize Anonymize a deactivated user. Requires Administer System permission
func (s *UsersService) Anonymize(opt *UsersAnonymizeOption) (resp *http.Response, err error) {
	err = s.ValidateAnonymizeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/anonymize", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type UsersChangePasswordOption struct {
	Login            string `url:"login,omitempty"`            // Description:"User login",ExampleValue:"myuser"
	Password         string `url:"password,omitempty"`         // Description:"New password",ExampleValue:"mypassword"
	PreviousPassword string `url:"previousPassword,omitempty"` // Description:"Previous password. Required when changing one's own password.",ExampleValue:"oldpassword"
}

// ChangePassword Update a user's password. Authenticated users can change their own password, provided that the account is not linked to an external authentication system. Administer System permission is required to change another user's password.
func (s *UsersService) ChangePassword(opt *UsersChangePasswordOption) (resp *http.Response, err error) {
	err = s.ValidateChangePasswordOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/change_password", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type UsersCreateOption struct {
	Email      string `url:"email,omitempty"`      // Description:"User email",ExampleValue:"myname@email.com"
	Local      string `url:"local,omitempty"`      // Description:"Specify if the user should be authenticated from SonarQube server or from an external authentication system. Password should not be set when local is set to false.",ExampleValue:""
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"myuser"
	Name       string `url:"name,omitempty"`       // Description:"User name",ExampleValue:"My Name"
	Password   string `url:"password,omitempty"`   // Description:"User password. Only mandatory when creating local user, otherwise it should not be set",ExampleValue:"mypassword"
	ScmAccount string `url:"scmAccount,omitempty"` // Description:"List of SCM accounts. To set several values, the parameter must be called once for each value.",ExampleValue:"scmAccount=firstValue&scmAccount=secondValue&scmAccount=thirdValue"
}

// Create Create a user.<br/>If a deactivated user account exists with the given login, it will be reactivated.<br/>Requires Administer System permission
func (s *UsersService) Create(opt *UsersCreateOption) (v *UsersCreateObject, resp *http.Response, err error) {
	err = s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/create", opt)
	if err != nil {
		return
	}
	v = new(UsersCreateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// Current Get the details of the current authenticated user.
func (s *UsersService) Current() (v *UsersCurrentObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "users/current", nil)
	if err != nil {
		return
	}
	v = new(UsersCurrentObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UsersDeactivateOption struct {
	Anonymize string `url:"anonymize,omitempty"` // Description:"Anonymize user in addition to deactivating it",ExampleValue:""
	Login     string `url:"login,omitempty"`     // Description:"User login",ExampleValue:"myuser"
}

// Deactivate Deactivate a user. Requires Administer System permission
func (s *UsersService) Deactivate(opt *UsersDeactivateOption) (v *UsersDeactivateObject, resp *http.Response, err error) {
	err = s.ValidateDeactivateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/deactivate", opt)
	if err != nil {
		return
	}
	v = new(UsersDeactivateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UsersDismissNoticeOption struct {
	Notice string `url:"notice,omitempty"` // Description:"notice key to dismiss",ExampleValue:"educationPrinciples"
}

// DismissNotice Dismiss a notice for the current user. Silently ignore if the notice is already dismissed.
func (s *UsersService) DismissNotice(opt *UsersDismissNoticeOption) (resp *http.Response, err error) {
	err = s.ValidateDismissNoticeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/dismiss_notice", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

// DismissSonarlintAd Dismiss SonarLint advertisement. Deprecated since 9.6, replaced api/users/dismiss_notice
func (s *UsersService) DismissSonarlintAd() (resp *http.Response, err error) {
	req, err := s.client.NewRequest("POST", "users/dismiss_sonarlint_ad", nil)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type UsersGroupsOption struct {
	Login    string `url:"login,omitempty"`    // Description:"A user login",ExampleValue:"admin"
	P        string `url:"p,omitempty"`        // Description:"1-based page number",ExampleValue:"42"
	Ps       string `url:"ps,omitempty"`       // Description:"Page size. Must be greater than 0.",ExampleValue:"20"
	Q        string `url:"q,omitempty"`        // Description:"Limit search to group names that contain the supplied string.",ExampleValue:"users"
	Selected string `url:"selected,omitempty"` // Description:"Depending on the value, show only selected items (selected=selected), deselected items (selected=deselected), or all items with their selection status (selected=all).",ExampleValue:""
}

// Groups Lists the groups a user belongs to. <br/>Requires Administer System permission.
func (s *UsersService) Groups(opt *UsersGroupsOption) (v *UsersGroupsObject, resp *http.Response, err error) {
	err = s.ValidateGroupsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "users/groups", opt)
	if err != nil {
		return
	}
	v = new(UsersGroupsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

// IdentityProviders List the external identity providers
func (s *UsersService) IdentityProviders() (v *UsersIdentityProvidersObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "users/identity_providers", nil)
	if err != nil {
		return
	}
	v = new(UsersIdentityProvidersObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UsersSearchOption struct {
	Deactivated      string `url:"deactivated,omitempty"`      // Description:"Return deactivated users instead of active users",ExampleValue:""
	ExternalIdentity string `url:"externalIdentity,omitempty"` /*
	Description:"Find a user by its external identity (ie. its login in the Identity Provider).
	This is case sensitive and only available with Administer System permission.
	",ExampleValue:""
	*/
	P  string `url:"p,omitempty"`  // Description:"1-based page number",ExampleValue:"42"
	Ps string `url:"ps,omitempty"` // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q  string `url:"q,omitempty"`  // Description:"Filter on login, name and email.<br />This parameter can either be case sensitive and perform an exact match, or case insensitive and perform a partial match (contains), depending on the scenario:<br /><ul>  <li>    If the search query is <em>less or equal to 15 characters</em>, then the query is <em>case insensitive</em>, and will match any login, name, or email, that     <em>contains</em> the search query.  </li>  <li>    If the search query is <em>greater than 15 characters</em>, then the query becomes <em>case sensitive</em>, and will match any login, name, or email, that     <em>exactly matches</em> the search query.  </li></ul>",ExampleValue:""
}

// Search Get a list of users. By default, only active users are returned.<br/>The following fields are only returned when user has Administer System permission or for logged-in in user :<ul>   <li>'email'</li>   <li>'externalIdentity'</li>   <li>'externalProvider'</li>   <li>'groups'</li>   <li>'lastConnectionDate'</li>   <li>'tokensCount'</li></ul>Field 'lastConnectionDate' is only updated every hour, so it may not be accurate, for instance when a user authenticates many times in less than one hour.
func (s *UsersService) Search(opt *UsersSearchOption) (v *UsersSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "users/search", opt)
	if err != nil {
		return
	}
	v = new(UsersSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UsersSetHomepageOption struct {
	Branch    string `url:"branch,omitempty"`    // Description:"Branch key. It can only be used when parameter 'type' is set to 'PROJECT'",ExampleValue:"feature/my_branch"
	Component string `url:"component,omitempty"` // Description:"Project key. It should only be used when parameter 'type' is set to 'PROJECT'",ExampleValue:"my_project"
	Type      string `url:"type,omitempty"`      // Description:"Type of the requested page",ExampleValue:""
}

// SetHomepage Set homepage of current user.<br> Requires authentication.
func (s *UsersService) SetHomepage(opt *UsersSetHomepageOption) (resp *http.Response, err error) {
	err = s.ValidateSetHomepageOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/set_homepage", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type UsersUpdateOption struct {
	Email      string `url:"email,omitempty"`      // Description:"User email",ExampleValue:"myname@email.com"
	Login      string `url:"login,omitempty"`      // Description:"User login",ExampleValue:"myuser"
	Name       string `url:"name,omitempty"`       // Description:"User name",ExampleValue:"My Name"
	ScmAccount string `url:"scmAccount,omitempty"` // Description:"SCM accounts. To set several values, the parameter must be called once for each value.",ExampleValue:"scmAccount=firstValue&scmAccount=secondValue&scmAccount=thirdValue"
}

// Update Update a user.<br/>Requires Administer System permission
func (s *UsersService) Update(opt *UsersUpdateOption) (v *UsersUpdateObject, resp *http.Response, err error) {
	err = s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/update", opt)
	if err != nil {
		return
	}
	v = new(UsersUpdateObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type UsersUpdateIdentityProviderOption struct {
	Login               string `url:"login,omitempty"`               // Description:"User login",ExampleValue:""
	NewExternalIdentity string `url:"newExternalIdentity,omitempty"` // Description:"New external identity, usually the login used in the authentication system. If not provided previous identity will be used.",ExampleValue:""
	NewExternalProvider string `url:"newExternalProvider,omitempty"` // Description:"New external provider. Only authentication system installed are available. Use 'LDAP' identity provider for single server LDAP setup.User 'LDAP_{serverKey}' identity provider for multiple LDAP server setup.",ExampleValue:""
}

// UpdateIdentityProvider Update identity provider information. <br/>It's only possible to migrate to an installed identity provider. Be careful that as soon as this information has been updated for a user, the user will only be able to authenticate on the new identity provider. It is not possible to migrate external user to local one.<br/>Requires Administer System permission.
func (s *UsersService) UpdateIdentityProvider(opt *UsersUpdateIdentityProviderOption) (resp *http.Response, err error) {
	err = s.ValidateUpdateIdentityProviderOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/update_identity_provider", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type UsersUpdateLoginOption struct {
	Login    string `url:"login,omitempty"`    // Description:"The current login (case-sensitive)",ExampleValue:"mylogin"
	NewLogin string `url:"newLogin,omitempty"` // Description:"The new login. It must not already exist.",ExampleValue:"mynewlogin"
}

// UpdateLogin Update a user login. A login can be updated many times.<br/>Requires Administer System permission
func (s *UsersService) UpdateLogin(opt *UsersUpdateLoginOption) (resp *http.Response, err error) {
	err = s.ValidateUpdateLoginOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "users/update_login", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}
