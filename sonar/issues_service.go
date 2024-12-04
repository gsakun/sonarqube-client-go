// Read and update issues.
package sonar

import "net/http"

type IssuesService struct {
	client *Client
}

type IssuesAddCommentObject struct {
	Components []IssuesAddCommentObject_sub1 `json:"components,omitempty"`
	Issue      IssuesAddCommentObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesAddCommentObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesAddCommentObject_sub6 `json:"users,omitempty"`
}

type IssuesAddCommentObject_sub4 struct {
	Actions                   []string                      `json:"actions,omitempty"`
	Assignee                  string                        `json:"assignee,omitempty"`
	Author                    string                        `json:"author,omitempty"`
	Comments                  []IssuesAddCommentObject_sub2 `json:"comments,omitempty"`
	Component                 string                        `json:"component,omitempty"`
	CreationDate              string                        `json:"creationDate,omitempty"`
	Debt                      string                        `json:"debt,omitempty"`
	Effort                    string                        `json:"effort,omitempty"`
	Flows                     []interface{}                 `json:"flows,omitempty"`
	Key                       string                        `json:"key,omitempty"`
	Line                      int64                         `json:"line,omitempty"`
	Message                   string                        `json:"message,omitempty"`
	Project                   string                        `json:"project,omitempty"`
	Rule                      string                        `json:"rule,omitempty"`
	RuleDescriptionContextKey string                        `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                        `json:"severity,omitempty"`
	Status                    string                        `json:"status,omitempty"`
	Tags                      []string                      `json:"tags,omitempty"`
	TextRange                 IssuesAddCommentObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                      `json:"transitions,omitempty"`
	Type                      string                        `json:"type,omitempty"`
	UpdateDate                string                        `json:"updateDate,omitempty"`
}

type IssuesAddCommentObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesAddCommentObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesAddCommentObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	ID        int64  `json:"id,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

type IssuesAddCommentObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesAddCommentObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesAssignObject struct {
	Components []IssuesAssignObject_sub1 `json:"components,omitempty"`
	Issue      IssuesAssignObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesAssignObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesAssignObject_sub6 `json:"users,omitempty"`
}

type IssuesAssignObject_sub4 struct {
	Actions                   []string                  `json:"actions,omitempty"`
	Assignee                  string                    `json:"assignee,omitempty"`
	Author                    string                    `json:"author,omitempty"`
	Comments                  []IssuesAssignObject_sub2 `json:"comments,omitempty"`
	Component                 string                    `json:"component,omitempty"`
	CreationDate              string                    `json:"creationDate,omitempty"`
	Debt                      string                    `json:"debt,omitempty"`
	Effort                    string                    `json:"effort,omitempty"`
	Flows                     []interface{}             `json:"flows,omitempty"`
	Key                       string                    `json:"key,omitempty"`
	Line                      int64                     `json:"line,omitempty"`
	Message                   string                    `json:"message,omitempty"`
	Project                   string                    `json:"project,omitempty"`
	Rule                      string                    `json:"rule,omitempty"`
	RuleDescriptionContextKey string                    `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                    `json:"severity,omitempty"`
	Status                    string                    `json:"status,omitempty"`
	Tags                      []string                  `json:"tags,omitempty"`
	TextRange                 IssuesAssignObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                  `json:"transitions,omitempty"`
	Type                      string                    `json:"type,omitempty"`
	UpdateDate                string                    `json:"updateDate,omitempty"`
}

type IssuesAssignObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesAssignObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesAssignObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

type IssuesAssignObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesAssignObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesAuthorsObject struct {
	Authors []string `json:"authors,omitempty"`
}

type IssuesBulkChangeObject struct {
	Failures int64 `json:"failures,omitempty"`
	Ignored  int64 `json:"ignored,omitempty"`
	Success  int64 `json:"success,omitempty"`
	Total    int64 `json:"total,omitempty"`
}

type IssuesChangelogObject struct {
	Changelog []IssuesChangelogObject_sub2 `json:"changelog,omitempty"`
}

type IssuesChangelogObject_sub2 struct {
	Avatar        string                       `json:"avatar,omitempty"`
	CreationDate  string                       `json:"creationDate,omitempty"`
	Diffs         []IssuesChangelogObject_sub1 `json:"diffs,omitempty"`
	ExternalUser  string                       `json:"externalUser,omitempty"`
	IsUserActive  bool                         `json:"isUserActive,omitempty"`
	User          string                       `json:"user,omitempty"`
	UserName      string                       `json:"userName,omitempty"`
	WebhookSource string                       `json:"webhookSource,omitempty"`
}

type IssuesChangelogObject_sub1 struct {
	Key      string `json:"key,omitempty"`
	NewValue string `json:"newValue,omitempty"`
	OldValue string `json:"oldValue,omitempty"`
}

type IssuesComponentTagsObject struct {
	Tags []IssuesComponentTagsObject_sub1 `json:"tags,omitempty"`
}

type IssuesComponentTagsObject_sub1 struct {
	Key   string `json:"key,omitempty"`
	Value int64  `json:"value,omitempty"`
}

type IssuesDeleteCommentObject struct {
	Components []IssuesDeleteCommentObject_sub1 `json:"components,omitempty"`
	Issue      IssuesDeleteCommentObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesDeleteCommentObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesDeleteCommentObject_sub6 `json:"users,omitempty"`
}

type IssuesDeleteCommentObject_sub4 struct {
	Actions                   []string                         `json:"actions,omitempty"`
	Assignee                  string                           `json:"assignee,omitempty"`
	Author                    string                           `json:"author,omitempty"`
	Comments                  []IssuesDeleteCommentObject_sub2 `json:"comments,omitempty"`
	Component                 string                           `json:"component,omitempty"`
	CreationDate              string                           `json:"creationDate,omitempty"`
	Debt                      string                           `json:"debt,omitempty"`
	Effort                    string                           `json:"effort,omitempty"`
	Flows                     []interface{}                    `json:"flows,omitempty"`
	Key                       string                           `json:"key,omitempty"`
	Line                      int64                            `json:"line,omitempty"`
	Message                   string                           `json:"message,omitempty"`
	Project                   string                           `json:"project,omitempty"`
	Rule                      string                           `json:"rule,omitempty"`
	RuleDescriptionContextKey string                           `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                           `json:"severity,omitempty"`
	Status                    string                           `json:"status,omitempty"`
	Tags                      []string                         `json:"tags,omitempty"`
	TextRange                 IssuesDeleteCommentObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                         `json:"transitions,omitempty"`
	Type                      string                           `json:"type,omitempty"`
	UpdateDate                string                           `json:"updateDate,omitempty"`
}

type IssuesDeleteCommentObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesDeleteCommentObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesDeleteCommentObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

type IssuesDeleteCommentObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesDeleteCommentObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesDoTransitionObject struct {
	Components []IssuesDoTransitionObject_sub1 `json:"components,omitempty"`
	Issue      IssuesDoTransitionObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesDoTransitionObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesDoTransitionObject_sub6 `json:"users,omitempty"`
}

type IssuesDoTransitionObject_sub4 struct {
	Actions                   []string                        `json:"actions,omitempty"`
	Assignee                  string                          `json:"assignee,omitempty"`
	Author                    string                          `json:"author,omitempty"`
	Comments                  []IssuesDoTransitionObject_sub2 `json:"comments,omitempty"`
	Component                 string                          `json:"component,omitempty"`
	CreationDate              string                          `json:"creationDate,omitempty"`
	Debt                      string                          `json:"debt,omitempty"`
	Effort                    string                          `json:"effort,omitempty"`
	Flows                     []interface{}                   `json:"flows,omitempty"`
	Key                       string                          `json:"key,omitempty"`
	Line                      int64                           `json:"line,omitempty"`
	Message                   string                          `json:"message,omitempty"`
	Project                   string                          `json:"project,omitempty"`
	Rule                      string                          `json:"rule,omitempty"`
	RuleDescriptionContextKey string                          `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                          `json:"severity,omitempty"`
	Status                    string                          `json:"status,omitempty"`
	Tags                      []string                        `json:"tags,omitempty"`
	TextRange                 IssuesDoTransitionObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                        `json:"transitions,omitempty"`
	Type                      string                          `json:"type,omitempty"`
	UpdateDate                string                          `json:"updateDate,omitempty"`
}

type IssuesDoTransitionObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesDoTransitionObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesDoTransitionObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

type IssuesDoTransitionObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesDoTransitionObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesEditCommentObject struct {
	Components []IssuesEditCommentObject_sub1 `json:"components,omitempty"`
	Issue      IssuesEditCommentObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesEditCommentObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesEditCommentObject_sub6 `json:"users,omitempty"`
}

type IssuesEditCommentObject_sub4 struct {
	Actions                   []string                       `json:"actions,omitempty"`
	Assignee                  string                         `json:"assignee,omitempty"`
	Author                    string                         `json:"author,omitempty"`
	Comments                  []IssuesEditCommentObject_sub2 `json:"comments,omitempty"`
	Component                 string                         `json:"component,omitempty"`
	CreationDate              string                         `json:"creationDate,omitempty"`
	Debt                      string                         `json:"debt,omitempty"`
	Effort                    string                         `json:"effort,omitempty"`
	Flows                     []interface{}                  `json:"flows,omitempty"`
	Key                       string                         `json:"key,omitempty"`
	Line                      int64                          `json:"line,omitempty"`
	Message                   string                         `json:"message,omitempty"`
	Project                   string                         `json:"project,omitempty"`
	Rule                      string                         `json:"rule,omitempty"`
	RuleDescriptionContextKey string                         `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                         `json:"severity,omitempty"`
	Status                    string                         `json:"status,omitempty"`
	Tags                      []string                       `json:"tags,omitempty"`
	TextRange                 IssuesEditCommentObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                       `json:"transitions,omitempty"`
	Type                      string                         `json:"type,omitempty"`
	UpdateDate                string                         `json:"updateDate,omitempty"`
}

type IssuesEditCommentObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesEditCommentObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesEditCommentObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

type IssuesEditCommentObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesEditCommentObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesSearchObject struct {
	Components []IssuesSearchObject_sub1  `json:"components,omitempty"`
	Issues     []IssuesSearchObject_sub8  `json:"issues,omitempty"`
	Paging     IssuesSearchObject_sub9    `json:"paging,omitempty"`
	Rules      []IssuesSearchObject_sub10 `json:"rules,omitempty"`
	Users      []IssuesSearchObject_sub11 `json:"users,omitempty"`
}

type IssuesSearchObject_sub8 struct {
	Actions                   []string                  `json:"actions,omitempty"`
	Attr                      IssuesSearchObject_sub2   `json:"attr,omitempty"`
	Author                    string                    `json:"author,omitempty"`
	Comments                  []IssuesSearchObject_sub3 `json:"comments,omitempty"`
	Component                 string                    `json:"component,omitempty"`
	CreationDate              string                    `json:"creationDate,omitempty"`
	Effort                    string                    `json:"effort,omitempty"`
	Flows                     []IssuesSearchObject_sub7 `json:"flows,omitempty"`
	Hash                      string                    `json:"hash,omitempty"`
	Key                       string                    `json:"key,omitempty"`
	Line                      int64                     `json:"line,omitempty"`
	Message                   string                    `json:"message,omitempty"`
	MessageFormattings        []IssuesSearchObject_sub4 `json:"messageFormattings,omitempty"`
	Project                   string                    `json:"project,omitempty"`
	QuickFixAvailable         bool                      `json:"quickFixAvailable,omitempty"`
	Resolution                string                    `json:"resolution,omitempty"`
	Rule                      string                    `json:"rule,omitempty"`
	RuleDescriptionContextKey string                    `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                    `json:"severity,omitempty"`
	Status                    string                    `json:"status,omitempty"`
	Tags                      []string                  `json:"tags,omitempty"`
	TextRange                 IssuesSearchObject_sub5   `json:"textRange,omitempty"`
	Transitions               []string                  `json:"transitions,omitempty"`
	Type                      string                    `json:"type,omitempty"`
	UpdateDate                string                    `json:"updateDate,omitempty"`
}

type IssuesSearchObject_sub11 struct {
	Active bool   `json:"active,omitempty"`
	Avatar string `json:"avatar,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesSearchObject_sub3 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesSearchObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
}

type IssuesSearchObject_sub4 struct {
	End   int64  `json:"end,omitempty"`
	Start int64  `json:"start,omitempty"`
	Type  string `json:"type,omitempty"`
}

type IssuesSearchObject_sub5 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesSearchObject_sub2 struct {
	Jira_issue_key string `json:"jira-issue-key,omitempty"`
}

type IssuesSearchObject_sub10 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesSearchObject_sub7 struct {
	Locations []IssuesSearchObject_sub6 `json:"locations,omitempty"`
}

type IssuesSearchObject_sub6 struct {
	Msg            string                    `json:"msg,omitempty"`
	MsgFormattings []IssuesSearchObject_sub4 `json:"msgFormattings,omitempty"`
	TextRange      IssuesSearchObject_sub5   `json:"textRange,omitempty"`
}

type IssuesSearchObject_sub9 struct {
	PageIndex int64 `json:"pageIndex,omitempty"`
	PageSize  int64 `json:"pageSize,omitempty"`
	Total     int64 `json:"total,omitempty"`
}

type IssuesSetSeverityObject struct {
	Components []IssuesSetSeverityObject_sub1 `json:"components,omitempty"`
	Issue      IssuesSetSeverityObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesSetSeverityObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesSetSeverityObject_sub6 `json:"users,omitempty"`
}

type IssuesSetSeverityObject_sub4 struct {
	Actions                   []string                       `json:"actions,omitempty"`
	Assignee                  string                         `json:"assignee,omitempty"`
	Author                    string                         `json:"author,omitempty"`
	Comments                  []IssuesSetSeverityObject_sub2 `json:"comments,omitempty"`
	Component                 string                         `json:"component,omitempty"`
	CreationDate              string                         `json:"creationDate,omitempty"`
	Debt                      string                         `json:"debt,omitempty"`
	Effort                    string                         `json:"effort,omitempty"`
	Flows                     []interface{}                  `json:"flows,omitempty"`
	Key                       string                         `json:"key,omitempty"`
	Line                      int64                          `json:"line,omitempty"`
	Message                   string                         `json:"message,omitempty"`
	Project                   string                         `json:"project,omitempty"`
	Rule                      string                         `json:"rule,omitempty"`
	RuleDescriptionContextKey string                         `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                         `json:"severity,omitempty"`
	Status                    string                         `json:"status,omitempty"`
	Tags                      []string                       `json:"tags,omitempty"`
	TextRange                 IssuesSetSeverityObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                       `json:"transitions,omitempty"`
	Type                      string                         `json:"type,omitempty"`
	UpdateDate                string                         `json:"updateDate,omitempty"`
}

type IssuesSetSeverityObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesSetSeverityObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesSetSeverityObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

type IssuesSetSeverityObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesSetSeverityObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesSetTagsObject struct {
	Components []IssuesSetTagsObject_sub1 `json:"components,omitempty"`
	Issue      IssuesSetTagsObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesSetTagsObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesSetTagsObject_sub6 `json:"users,omitempty"`
}

type IssuesSetTagsObject_sub4 struct {
	Actions                   []string                   `json:"actions,omitempty"`
	Assignee                  string                     `json:"assignee,omitempty"`
	Author                    string                     `json:"author,omitempty"`
	Comments                  []IssuesSetTagsObject_sub2 `json:"comments,omitempty"`
	Component                 string                     `json:"component,omitempty"`
	CreationDate              string                     `json:"creationDate,omitempty"`
	Debt                      string                     `json:"debt,omitempty"`
	Effort                    string                     `json:"effort,omitempty"`
	Flows                     []interface{}              `json:"flows,omitempty"`
	Key                       string                     `json:"key,omitempty"`
	Line                      int64                      `json:"line,omitempty"`
	Message                   string                     `json:"message,omitempty"`
	Project                   string                     `json:"project,omitempty"`
	Rule                      string                     `json:"rule,omitempty"`
	RuleDescriptionContextKey string                     `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                     `json:"severity,omitempty"`
	Status                    string                     `json:"status,omitempty"`
	Tags                      []string                   `json:"tags,omitempty"`
	TextRange                 IssuesSetTagsObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                   `json:"transitions,omitempty"`
	Type                      string                     `json:"type,omitempty"`
	UpdateDate                string                     `json:"updateDate,omitempty"`
}

type IssuesSetTagsObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesSetTagsObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesSetTagsObject_sub1 struct {
	Enabled      bool   `json:"enabled,omitempty"`
	Key          string `json:"key,omitempty"`
	LongName     string `json:"longName,omitempty"`
	Name         string `json:"name,omitempty"`
	Path         string `json:"path,omitempty"`
	ProjectID    int64  `json:"projectId,omitempty"`
	Qualifier    string `json:"qualifier,omitempty"`
	SubProjectID int64  `json:"subProjectId,omitempty"`
	UUID         string `json:"uuid,omitempty"`
}

type IssuesSetTagsObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesSetTagsObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesSetTypeObject struct {
	Components []IssuesSetTypeObject_sub1 `json:"components,omitempty"`
	Issue      IssuesSetTypeObject_sub4   `json:"issue,omitempty"`
	Rules      []IssuesSetTypeObject_sub5 `json:"rules,omitempty"`
	Users      []IssuesSetTypeObject_sub6 `json:"users,omitempty"`
}

type IssuesSetTypeObject_sub4 struct {
	Actions                   []string                   `json:"actions,omitempty"`
	Assignee                  string                     `json:"assignee,omitempty"`
	Author                    string                     `json:"author,omitempty"`
	Comments                  []IssuesSetTypeObject_sub2 `json:"comments,omitempty"`
	Component                 string                     `json:"component,omitempty"`
	CreationDate              string                     `json:"creationDate,omitempty"`
	Debt                      string                     `json:"debt,omitempty"`
	Effort                    string                     `json:"effort,omitempty"`
	Flows                     []interface{}              `json:"flows,omitempty"`
	Key                       string                     `json:"key,omitempty"`
	Line                      int64                      `json:"line,omitempty"`
	Message                   string                     `json:"message,omitempty"`
	Project                   string                     `json:"project,omitempty"`
	Rule                      string                     `json:"rule,omitempty"`
	RuleDescriptionContextKey string                     `json:"ruleDescriptionContextKey,omitempty"`
	Severity                  string                     `json:"severity,omitempty"`
	Status                    string                     `json:"status,omitempty"`
	Tags                      []string                   `json:"tags,omitempty"`
	TextRange                 IssuesSetTypeObject_sub3   `json:"textRange,omitempty"`
	Transitions               []string                   `json:"transitions,omitempty"`
	Type                      string                     `json:"type,omitempty"`
	UpdateDate                string                     `json:"updateDate,omitempty"`
}

type IssuesSetTypeObject_sub6 struct {
	Active bool   `json:"active,omitempty"`
	Email  string `json:"email,omitempty"`
	Login  string `json:"login,omitempty"`
	Name   string `json:"name,omitempty"`
}

type IssuesSetTypeObject_sub2 struct {
	CreatedAt string `json:"createdAt,omitempty"`
	HTMLText  string `json:"htmlText,omitempty"`
	Key       string `json:"key,omitempty"`
	Login     string `json:"login,omitempty"`
	Markdown  string `json:"markdown,omitempty"`
	Updatable bool   `json:"updatable,omitempty"`
}

type IssuesSetTypeObject_sub1 struct {
	Enabled   bool   `json:"enabled,omitempty"`
	Key       string `json:"key,omitempty"`
	LongName  string `json:"longName,omitempty"`
	Name      string `json:"name,omitempty"`
	Path      string `json:"path,omitempty"`
	Qualifier string `json:"qualifier,omitempty"`
	UUID      string `json:"uuid,omitempty"`
}

type IssuesSetTypeObject_sub3 struct {
	EndLine     int64 `json:"endLine,omitempty"`
	EndOffset   int64 `json:"endOffset,omitempty"`
	StartLine   int64 `json:"startLine,omitempty"`
	StartOffset int64 `json:"startOffset,omitempty"`
}

type IssuesSetTypeObject_sub5 struct {
	Key      string `json:"key,omitempty"`
	Lang     string `json:"lang,omitempty"`
	LangName string `json:"langName,omitempty"`
	Name     string `json:"name,omitempty"`
	Status   string `json:"status,omitempty"`
}

type IssuesTagsObject struct {
	Tags []string `json:"tags,omitempty"`
}

type IssuesAddCommentOption struct {
	Issue string `url:"issue,omitempty"` // Description:"Issue key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Text  string `url:"text,omitempty"`  // Description:"Comment text",ExampleValue:"Won't fix because it doesn't apply to the context"
}

// AddComment Add a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
func (s *IssuesService) AddComment(opt *IssuesAddCommentOption) (v *IssuesAddCommentObject, resp *http.Response, err error) {
	err = s.ValidateAddCommentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/add_comment", opt)
	if err != nil {
		return
	}
	v = new(IssuesAddCommentObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesAssignOption struct {
	Assignee string `url:"assignee,omitempty"` // Description:"Login of the assignee. When not set, it will unassign the issue. Use '_me' to assign to current user",ExampleValue:"admin"
	Issue    string `url:"issue,omitempty"`    // Description:"Issue key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// Assign Assign/Unassign an issue. Requires authentication and Browse permission on project
func (s *IssuesService) Assign(opt *IssuesAssignOption) (v *IssuesAssignObject, resp *http.Response, err error) {
	err = s.ValidateAssignOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/assign", opt)
	if err != nil {
		return
	}
	v = new(IssuesAssignObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesAuthorsOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Ps      string `url:"ps,omitempty"`      // Description:"Page size. Must be greater than 0 and less or equal than 100",ExampleValue:"20"
	Q       string `url:"q,omitempty"`       // Description:"Limit search to authors that contain the supplied string.",ExampleValue:"luke"
}

// Authors Search SCM accounts which match a given query.<br/>Requires authentication.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code.
func (s *IssuesService) Authors(opt *IssuesAuthorsOption) (v *IssuesAuthorsObject, resp *http.Response, err error) {
	err = s.ValidateAuthorsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "issues/authors", opt)
	if err != nil {
		return
	}
	v = new(IssuesAuthorsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesBulkChangeOption struct {
	AddTags           string `url:"add_tags,omitempty"`          // Description:"Add tags",ExampleValue:"security,java8"
	Assign            string `url:"assign,omitempty"`            // Description:"To assign the list of issues to a specific user (login), or un-assign all the issues",ExampleValue:"john.smith"
	Comment           string `url:"comment,omitempty"`           // Description:"Add a comment. The comment will only be added to issues that are affected either by a change of type or a change of severity as a result of the same WS call.",ExampleValue:"Here is my comment"
	DoTransition      string `url:"do_transition,omitempty"`     // Description:"Transition",ExampleValue:"reopen"
	Issues            string `url:"issues,omitempty"`            // Description:"Comma-separated list of issue keys",ExampleValue:"AU-Tpxb--iU5OvuD2FLy,AU-TpxcA-iU5OvuD2FLz"
	RemoveTags        string `url:"remove_tags,omitempty"`       // Description:"Remove tags",ExampleValue:"security,java8"
	SendNotifications string `url:"sendNotifications,omitempty"` // Description:"",ExampleValue:""
	SetSeverity       string `url:"set_severity,omitempty"`      // Description:"To change the severity of the list of issues",ExampleValue:"BLOCKER"
	SetType           string `url:"set_type,omitempty"`          // Description:"To change the type of the list of issues",ExampleValue:"BUG"
}

// BulkChange Bulk change on issues. Up to 500 issues can be updated. <br/>Requires authentication.
func (s *IssuesService) BulkChange(opt *IssuesBulkChangeOption) (v *IssuesBulkChangeObject, resp *http.Response, err error) {
	err = s.ValidateBulkChangeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/bulk_change", opt)
	if err != nil {
		return
	}
	v = new(IssuesBulkChangeObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesChangelogOption struct {
	Issue string `url:"issue,omitempty"` // Description:"Issue key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// Changelog Display changelog of an issue.<br/>Requires the 'Browse' permission on the project of the specified issue.
func (s *IssuesService) Changelog(opt *IssuesChangelogOption) (v *IssuesChangelogObject, resp *http.Response, err error) {
	err = s.ValidateChangelogOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "issues/changelog", opt)
	if err != nil {
		return
	}
	v = new(IssuesChangelogObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesComponentTagsOption struct {
	ComponentUuid string `url:"componentUuid,omitempty"` // Description:"A component UUID",ExampleValue:"7d8749e8-3070-4903-9188-bdd82933bb92"
	CreatedAfter  string `url:"createdAfter,omitempty"`  // Description:"To retrieve tags on issues created after the given date (inclusive). <br>Either a date (server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	Ps            string `url:"ps,omitempty"`            // Description:"The maximum size of the list to return",ExampleValue:"25"
}

// ComponentTags List tags for the issues under a given component (including issues on the descendants of the component)<br/>When issue indexation is in progress returns 503 service unavailable HTTP code.
func (s *IssuesService) ComponentTags(opt *IssuesComponentTagsOption) (v *IssuesComponentTagsObject, resp *http.Response, err error) {
	err = s.ValidateComponentTagsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "issues/component_tags", opt)
	if err != nil {
		return
	}
	v = new(IssuesComponentTagsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesDeleteCommentOption struct {
	Comment string `url:"comment,omitempty"` // Description:"Comment key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
}

// DeleteComment Delete a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
func (s *IssuesService) DeleteComment(opt *IssuesDeleteCommentOption) (v *IssuesDeleteCommentObject, resp *http.Response, err error) {
	err = s.ValidateDeleteCommentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/delete_comment", opt)
	if err != nil {
		return
	}
	v = new(IssuesDeleteCommentObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesDoTransitionOption struct {
	Issue      string `url:"issue,omitempty"`      // Description:"Issue key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Transition string `url:"transition,omitempty"` // Description:"Transition",ExampleValue:""
}

// DoTransition Do workflow transition on an issue. Requires authentication and Browse permission on project.<br/>The transitions 'wontfix' and 'falsepositive' require the permission 'Administer Issues'.<br/>The transitions involving security hotspots require the permission 'Administer Security Hotspot'.
func (s *IssuesService) DoTransition(opt *IssuesDoTransitionOption) (v *IssuesDoTransitionObject, resp *http.Response, err error) {
	err = s.ValidateDoTransitionOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/do_transition", opt)
	if err != nil {
		return
	}
	v = new(IssuesDoTransitionObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesEditCommentOption struct {
	Comment string `url:"comment,omitempty"` // Description:"Comment key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Text    string `url:"text,omitempty"`    // Description:"Comment text",ExampleValue:"Won't fix because it doesn't apply to the context"
}

// EditComment Edit a comment.<br/>Requires authentication and the following permission: 'Browse' on the project of the specified issue.
func (s *IssuesService) EditComment(opt *IssuesEditCommentOption) (v *IssuesEditCommentObject, resp *http.Response, err error) {
	err = s.ValidateEditCommentOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/edit_comment", opt)
	if err != nil {
		return
	}
	v = new(IssuesEditCommentObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesPullOption struct {
	BranchName       string `url:"branchName,omitempty"`       // Description:"Branch name for which issues are fetched.",ExampleValue:"develop"
	ChangedSince     string `url:"changedSince,omitempty"`     // Description:"Timestamp. If present only issues modified after given timestamp are returned (both open and closed). If not present all non-closed issues are returned.",ExampleValue:"1654032306000"
	Languages        string `url:"languages,omitempty"`        // Description:"Comma separated list of languages. If not present all issues regardless of their language are returned.",ExampleValue:"java,cobol"
	ProjectKey       string `url:"projectKey,omitempty"`       // Description:"Project key for which issues are fetched.",ExampleValue:"sonarqube"
	ResolvedOnly     string `url:"resolvedOnly,omitempty"`     // Description:"If true only issues with resolved status are returned",ExampleValue:"true"
	RuleRepositories string `url:"ruleRepositories,omitempty"` // Description:"Comma separated list of rule repositories. If not present all issues regardless of their rule repository are returned.",ExampleValue:"java"
}

// Pull This endpoint fetches and returns all (unless filtered by optional params) the issues for a given branch. The issues returned are not paginated, so the response size can be big. Requires project 'Browse' permission.
func (s *IssuesService) Pull(opt *IssuesPullOption) (v *string, resp *http.Response, err error) {
	err = s.ValidatePullOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "issues/pull", opt)
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

type IssuesPullTaintOption struct {
	BranchName   string `url:"branchName,omitempty"`   // Description:"Branch name for which taint vulnerabilities are fetched.",ExampleValue:"develop"
	ChangedSince string `url:"changedSince,omitempty"` // Description:"Timestamp. If present only taint vulnerabilities modified after given timestamp are returned (both open and closed). If not present all non-closed taint vulnerabilities are returned.",ExampleValue:"1654032306000"
	Languages    string `url:"languages,omitempty"`    // Description:"Comma separated list of languages. If not present all taint vulnerabilities regardless of their language are returned.",ExampleValue:"java,cobol"
	ProjectKey   string `url:"projectKey,omitempty"`   // Description:"Project key for which taint vulnerabilities are fetched.",ExampleValue:"sonarqube"
}

// PullTaint This endpoint fetches and returns all (unless filtered by optional params) the taint vulnerabilities for a given branch. The taint vulnerabilities returned are not paginated, so the response size can be big. Requires project 'Browse' permission.
func (s *IssuesService) PullTaint(opt *IssuesPullTaintOption) (v *string, resp *http.Response, err error) {
	err = s.ValidatePullTaintOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "issues/pull_taint", opt)
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

type IssuesReindexOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// Reindex Reindex issues for a project.<br> Require 'Administer System' permission.
func (s *IssuesService) Reindex(opt *IssuesReindexOption) (resp *http.Response, err error) {
	err = s.ValidateReindexOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/reindex", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}

type IssuesSearchOption struct {
	AdditionalFields    string `url:"additionalFields,omitempty"`    // Description:"Comma-separated list of the optional fields to be returned in response. Action plans are dropped in 5.5, it is not returned in the response.",ExampleValue:""
	Asc                 string `url:"asc,omitempty"`                 // Description:"Ascending sort",ExampleValue:""
	Assigned            string `url:"assigned,omitempty"`            // Description:"To retrieve assigned or unassigned issues",ExampleValue:""
	Assignees           string `url:"assignees,omitempty"`           // Description:"Comma-separated list of assignee logins. The value '__me__' can be used as a placeholder for user who performs the request",ExampleValue:"admin,usera,__me__"
	Author              string `url:"author,omitempty"`              // Description:"SCM accounts. To set several values, the parameter must be called once for each value.",ExampleValue:"author=torvalds@linux-foundation.org&author=linux@fondation.org"
	Branch              string `url:"branch,omitempty"`              // Description:"Branch key. Not available in the community edition.",ExampleValue:"feature/my_branch"
	ComponentKeys       string `url:"componentKeys,omitempty"`       // Description:"Comma-separated list of component keys. Retrieve issues associated to a specific list of components (and all its descendants). A component can be a portfolio, project, module, directory or file.",ExampleValue:"my_project"
	CreatedAfter        string `url:"createdAfter,omitempty"`        // Description:"To retrieve issues created after the given date (inclusive). <br>Either a date (use 'timeZone' attribute or it will default to server timezone) or datetime can be provided. <br>If this parameter is set, createdInLast must not be set",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	CreatedAt           string `url:"createdAt,omitempty"`           // Description:"Datetime to retrieve issues created during a specific analysis",ExampleValue:"2017-10-19T13:00:00+0200"
	CreatedBefore       string `url:"createdBefore,omitempty"`       // Description:"To retrieve issues created before the given date (exclusive). <br>Either a date (use 'timeZone' attribute or it will default to server timezone) or datetime can be provided.",ExampleValue:"2017-10-19 or 2017-10-19T13:00:00+0200"
	CreatedInLast       string `url:"createdInLast,omitempty"`       // Description:"To retrieve issues created during a time span before the current time (exclusive). Accepted units are 'y' for year, 'm' for month, 'w' for week and 'd' for day. If this parameter is set, createdAfter must not be set",ExampleValue:"1m2w (1 month 2 weeks)"
	Cwe                 string `url:"cwe,omitempty"`                 // Description:"Comma-separated list of CWE identifiers. Use 'unknown' to select issues not associated to any CWE.",ExampleValue:"12,125,unknown"
	Directories         string `url:"directories,omitempty"`         // Description:"To retrieve issues associated to a specific list of directories (comma-separated list of directory paths). This parameter is only meaningful when a module is selected. This parameter is mostly used by the Issues page, please prefer usage of the componentKeys parameter. ",ExampleValue:"src/main/java/org/sonar/server/"
	Facets              string `url:"facets,omitempty"`              // Description:"Comma-separated list of the facets to be computed. No facet is computed by default.",ExampleValue:""
	Files               string `url:"files,omitempty"`               // Description:"To retrieve issues associated to a specific list of files (comma-separated list of file paths). This parameter is mostly used by the Issues page, please prefer usage of the componentKeys parameter. ",ExampleValue:"src/main/java/org/sonar/server/Test.java"
	InNewCodePeriod     string `url:"inNewCodePeriod,omitempty"`     // Description:"To retrieve issues created in the new code period.<br>If this parameter is set to a truthy value, createdAfter must not be set and one component uuid or key must be provided.",ExampleValue:""
	Issues              string `url:"issues,omitempty"`              // Description:"Comma-separated list of issue keys",ExampleValue:"5bccd6e8-f525-43a2-8d76-fcb13dde79ef"
	Languages           string `url:"languages,omitempty"`           // Description:"Comma-separated list of languages. Available since 4.4",ExampleValue:"java,js"
	OnComponentOnly     string `url:"onComponentOnly,omitempty"`     // Description:"Return only issues at a component's level, not on its descendants (modules, directories, files, etc). This parameter is only considered when componentKeys is set.",ExampleValue:""
	OwaspAsvs40         string `url:"owaspAsvs-4.0,omitempty"`       // Description:"Comma-separated list of OWASP ASVS v4.0 categories.",ExampleValue:"6,10.1.1"
	OwaspAsvsLevel      string `url:"owaspAsvsLevel,omitempty"`      // Description:"Level of OWASP ASVS categories.",ExampleValue:""
	OwaspTop10          string `url:"owaspTop10,omitempty"`          // Description:"Comma-separated list of OWASP Top 10 2017 lowercase categories.",ExampleValue:""
	OwaspTop102021      string `url:"owaspTop10-2021,omitempty"`     // Description:"Comma-separated list of OWASP Top 10 2021 lowercase categories.",ExampleValue:""
	P                   string `url:"p,omitempty"`                   // Description:"1-based page number",ExampleValue:"42"
	PciDss32            string `url:"pciDss-3.2,omitempty"`          // Description:"Comma-separated list of PCI DSS v3.2 categories.",ExampleValue:"4,6.5.8,10.1"
	PciDss40            string `url:"pciDss-4.0,omitempty"`          // Description:"Comma-separated list of PCI DSS v4.0 categories.",ExampleValue:"4,6.5.8,10.1"
	Projects            string `url:"projects,omitempty"`            // Description:"To retrieve issues associated to a specific list of projects (comma-separated list of project keys). This parameter is mostly used by the Issues page, please prefer usage of the componentKeys parameter. If this parameter is set, projectUuids must not be set.",ExampleValue:"my_project"
	Ps                  string `url:"ps,omitempty"`                  // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	PullRequest         string `url:"pullRequest,omitempty"`         // Description:"Pull request id. Not available in the community edition.",ExampleValue:"5461"
	Resolutions         string `url:"resolutions,omitempty"`         // Description:"Comma-separated list of resolutions",ExampleValue:"FIXED,REMOVED"
	Resolved            string `url:"resolved,omitempty"`            // Description:"To match resolved or unresolved issues",ExampleValue:""
	Rules               string `url:"rules,omitempty"`               // Description:"Comma-separated list of coding rule keys. Format is &lt;repository&gt;:&lt;rule&gt;",ExampleValue:"java:S1144"
	S                   string `url:"s,omitempty"`                   // Description:"Sort field",ExampleValue:""
	SansTop25           string `url:"sansTop25,omitempty"`           // Description:"Comma-separated list of SANS Top 25 categories.",ExampleValue:""
	Scopes              string `url:"scopes,omitempty"`              // Description:"Comma-separated list of scopes. Available since 8.5",ExampleValue:"MAIN,TEST"
	Severities          string `url:"severities,omitempty"`          // Description:"Comma-separated list of severities",ExampleValue:"BLOCKER,CRITICAL"
	SinceLeakPeriod     string `url:"sinceLeakPeriod,omitempty"`     // Description:"To retrieve issues created since the leak period.<br>If this parameter is set to a truthy value, createdAfter must not be set and one component uuid or key must be provided.",ExampleValue:""
	SonarsourceSecurity string `url:"sonarsourceSecurity,omitempty"` // Description:"Comma-separated list of SonarSource security categories. Use 'others' to select issues not associated with any category",ExampleValue:""
	Statuses            string `url:"statuses,omitempty"`            // Description:"Comma-separated list of statuses",ExampleValue:"OPEN,REOPENED"
	Tags                string `url:"tags,omitempty"`                // Description:"Comma-separated list of tags.",ExampleValue:"security,convention"
	TimeZone            string `url:"timeZone,omitempty"`            // Description:"To resolve dates passed to 'createdAfter' or 'createdBefore' (does not apply to datetime) and to compute creation date histogram",ExampleValue:"'Europe/Paris', 'Z' or '+02:00'"
	Types               string `url:"types,omitempty"`               // Description:"Comma-separated list of types.",ExampleValue:"CODE_SMELL,BUG"
}

// Search Search for issues.<br>Requires the 'Browse' permission on the specified project(s). <br>For applications, it also requires 'Browse' permission on its child projects.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code.
func (s *IssuesService) Search(opt *IssuesSearchOption) (v *IssuesSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "issues/search", opt)
	if err != nil {
		return
	}
	v = new(IssuesSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesSetSeverityOption struct {
	Issue    string `url:"issue,omitempty"`    // Description:"Issue key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Severity string `url:"severity,omitempty"` // Description:"New severity",ExampleValue:""
}

// SetSeverity Change severity.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
func (s *IssuesService) SetSeverity(opt *IssuesSetSeverityOption) (v *IssuesSetSeverityObject, resp *http.Response, err error) {
	err = s.ValidateSetSeverityOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/set_severity", opt)
	if err != nil {
		return
	}
	v = new(IssuesSetSeverityObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesSetTagsOption struct {
	Issue string `url:"issue,omitempty"` // Description:"Issue key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Tags  string `url:"tags,omitempty"`  // Description:"Comma-separated list of tags. All tags are removed if parameter is empty or not set.",ExampleValue:"security,cwe,misra-c"
}

// SetTags Set tags on an issue. <br/>Requires authentication and Browse permission on project
func (s *IssuesService) SetTags(opt *IssuesSetTagsOption) (v *IssuesSetTagsObject, resp *http.Response, err error) {
	err = s.ValidateSetTagsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/set_tags", opt)
	if err != nil {
		return
	}
	v = new(IssuesSetTagsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesSetTypeOption struct {
	Issue string `url:"issue,omitempty"` // Description:"Issue key",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	Type  string `url:"type,omitempty"`  // Description:"New type",ExampleValue:""
}

// SetType Change type of issue, for instance from 'code smell' to 'bug'.<br/>Requires the following permissions:<ul>  <li>'Authentication'</li>  <li>'Browse' rights on project of the specified issue</li>  <li>'Administer Issues' rights on project of the specified issue</li></ul>
func (s *IssuesService) SetType(opt *IssuesSetTypeOption) (v *IssuesSetTypeObject, resp *http.Response, err error) {
	err = s.ValidateSetTypeOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "issues/set_type", opt)
	if err != nil {
		return
	}
	v = new(IssuesSetTypeObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}

type IssuesTagsOption struct {
	All     string `url:"all,omitempty"`     // Description:"Indicator to search for all tags or only for tags in the main branch of a project",ExampleValue:""
	Branch  string `url:"branch,omitempty"`  // Description:"Branch key",ExampleValue:"feature/my_branch"
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
	Ps      string `url:"ps,omitempty"`      // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Q       string `url:"q,omitempty"`       // Description:"Limit search to tags that contain the supplied string.",ExampleValue:"misra"
}

// Tags List tags matching a given query
func (s *IssuesService) Tags(opt *IssuesTagsOption) (v *IssuesTagsObject, resp *http.Response, err error) {
	err = s.ValidateTagsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "issues/tags", opt)
	if err != nil {
		return
	}
	v = new(IssuesTagsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
