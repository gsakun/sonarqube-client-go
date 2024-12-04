// Monitoring
package sonar

import "net/http"

type MonitoringService struct {
	client *Client
}

/*
Metrics Return monitoring metrics in Prometheus format.
Support content type 'text/plain' (default) and 'application/openmetrics-text'.
this endpoint can be access using a Bearer token, that needs to be defined in sonar.properties with the 'sonar.web.systemPasscode' key.
*/
func (s *MonitoringService) Metrics() (v *string, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "monitoring/metrics", nil)
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
