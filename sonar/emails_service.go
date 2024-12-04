// Manage emails
package sonar

import "net/http"

type EmailsService struct {
	client *Client
}

type EmailsSendOption struct {
	Message string `url:"message,omitempty"` // Description:"Content of the email",ExampleValue:""
	Subject string `url:"subject,omitempty"` // Description:"Subject of the email",ExampleValue:"Test Message from SonarQube"
	To      string `url:"to,omitempty"`      // Description:"Email address",ExampleValue:"john@doo.com"
}

// Send Test email configuration by sending an email<br>Requires 'Administer System' permission.
func (s *EmailsService) Send(opt *EmailsSendOption) (resp *http.Response, err error) {
	err = s.ValidateSendOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "emails/send", opt)
	if err != nil {
		return
	}
	resp, err = s.client.Do(req, nil)
	if err != nil {
		return
	}
	return
}
