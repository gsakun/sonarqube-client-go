// Get details about Compute Engine tasks.
package sonar

import "net/http"

type AnalysisReportsService struct {
	client *Client
}

// IsQueueEmpty Check if the queue of Compute Engine is empty
func (s *AnalysisReportsService) IsQueueEmpty() (v *string, resp *http.Response, err error) {
	req, err := s.client.NewRequest("GET", "analysis_reports/is_queue_empty", nil)
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
