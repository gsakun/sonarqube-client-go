// Provides information about features available in SonarQube
package sonar

import "net/http"

type FeaturesService struct {
	client *Client
}

type FeaturesListObject []string

// List List supported features
func (s *FeaturesService) List() (v *FeaturesListObject, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "features/list", nil)
	if err != nil {
		return
	}
	v = new(FeaturesListObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
