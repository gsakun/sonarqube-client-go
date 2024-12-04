// Return data needed by SonarLint.
package sonar

import "net/http"

type DevelopersService struct {
	client *Client
}

type DevelopersSearchEventsObject struct {
	Events []DevelopersSearchEventsObject_sub1 `json:"events,omitempty"`
}

type DevelopersSearchEventsObject_sub1 struct {
	Category string `json:"category,omitempty"`
	Link     string `json:"link,omitempty"`
	Message  string `json:"message,omitempty"`
	Project  string `json:"project,omitempty"`
}

type DevelopersSearchEventsOption struct {
	From     string `url:"from,omitempty"`     // Description:"Comma-separated list of datetimes. Filter events created after the given date (exclusive).",ExampleValue:"2017-10-19T13:00:00+0200"
	Projects string `url:"projects,omitempty"` // Description:"Comma-separated list of project keys to search notifications for",ExampleValue:"my_project,another_project"
}

// SearchEvents Search for events.<br/>Requires authentication.<br/>When issue indexation is in progress returns 503 service unavailable HTTP code.
func (s *DevelopersService) SearchEvents(opt *DevelopersSearchEventsOption) (v *DevelopersSearchEventsObject, resp *http.Response, err error) {
	err = s.ValidateSearchEventsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "developers/search_events", opt)
	if err != nil {
		return
	}
	v = new(DevelopersSearchEventsObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
