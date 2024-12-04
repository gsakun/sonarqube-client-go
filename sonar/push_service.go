// Endpoints supporting server side events.
package sonar

import "net/http"

type PushService struct {
	client *Client
}

type PushSonarlintEventsOption struct {
	Languages   string `url:"languages,omitempty"`   // Description:"Comma-separated list of languages for which events will be delivered",ExampleValue:"java,cobol"
	ProjectKeys string `url:"projectKeys,omitempty"` // Description:"Comma-separated list of projects keys for which events will be delivered",ExampleValue:"example-project-key,example-project-key2"
}

// SonarlintEvents Endpoint for listening to server side events. Currently it notifies listener about change to activation of a rule
func (s *PushService) SonarlintEvents(opt *PushSonarlintEventsOption) (v *string, resp *http.Response, err error) {
	err = s.ValidateSonarlintEventsOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "push/sonarlint_events", opt)
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
