// Webhooks allow to notify external services when a project analysis is done
package sonarqube // import "github.com/magicsong/generate-go-for-sonarqube/pkg/sonargo"

type WebhooksService struct {
	client *Client
}

// [TODO] you should call the <Create> func manually and complete the fields of this struct
type WebhooksCreateResp struct{}
type WebhooksCreateOption struct {
	Name    string `url:"name,omitempty"`    // Description:"Name displayed in the administration console of webhooks",ExampleValue:"My Webhook"
	Project string `url:"project,omitempty"` // Description:"The key of the project that will own the webhook",ExampleValue:"my_project"
	Url     string `url:"url,omitempty"`     // Description:"Server endpoint that will receive the webhook payload, for example 'http://my_server/foo'. If HTTP Basic authentication is used, HTTPS is recommended to avoid man in the middle attacks. Example: 'https://myLogin:myPassword@my_server/foo'",ExampleValue:"https://www.my-webhook-listener.com/sonar"
}

// Create Create a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) Create(opt *WebhooksCreateOption) (Resp *WebhooksCreateResp, err error) {
	err := s.ValidateCreateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "webhooks/create", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Delete> func manually and complete the fields of this struct
type WebhooksDeleteResp struct{}
type WebhooksDeleteOption struct {
	Webhook string `url:"webhook,omitempty"` // Description:"The key of the webhook to be deleted,auto-generated value can be obtained through api/webhooks/create or api/webhooks/list",ExampleValue:"my_project"
}

// Delete Delete a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) Delete(opt *WebhooksDeleteOption) (Resp *WebhooksDeleteResp, err error) {
	err := s.ValidateDeleteOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "webhooks/delete", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Deliveries> func manually and complete the fields of this struct
type WebhooksDeliveriesResp struct{}
type WebhooksDeliveriesOption struct {
	CeTaskId     string `url:"ceTaskId,omitempty"`     // Description:"Id of the Compute Engine task",ExampleValue:"AU-Tpxb--iU5OvuD2FLy"
	ComponentKey string `url:"componentKey,omitempty"` // Description:"Key of the project",ExampleValue:"my-project"
	P            string `url:"p,omitempty"`            // Description:"1-based page number",ExampleValue:"42"
	Ps           string `url:"ps,omitempty"`           // Description:"Page size. Must be greater than 0 and less than 500",ExampleValue:"20"
	Webhook      string `url:"webhook,omitempty"`      // Description:"Key of the webhook that triggered those deliveries,auto-generated value that can be obtained through api/webhooks/create or api/webhooks/list",ExampleValue:"AU-TpxcA-iU5OvuD2FLz"
}

// Deliveries Get the recent deliveries for a specified project or Compute Engine task.<br/>Require 'Administer' permission on the related project.<br/>Note that additional information are returned by api/webhooks/delivery.
func (s *WebhooksService) Deliveries(opt *WebhooksDeliveriesOption) (Resp *WebhooksDeliveriesResp, err error) {
	err := s.ValidateDeliveriesOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webhooks/deliveries", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Delivery> func manually and complete the fields of this struct
type WebhooksDeliveryResp struct{}
type WebhooksDeliveryOption struct {
	DeliveryId string `url:"deliveryId,omitempty"` // Description:"Id of delivery",ExampleValue:"AU-TpxcA-iU5OvuD2FL3"
}

// Delivery Get a webhook delivery by its id.<br/>Require 'Administer System' permission.<br/>Note that additional information are returned by api/webhooks/delivery.
func (s *WebhooksService) Delivery(opt *WebhooksDeliveryOption) (Resp *WebhooksDeliveryResp, err error) {
	err := s.ValidateDeliveryOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webhooks/delivery", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <List> func manually and complete the fields of this struct
type WebhooksListResp struct{}
type WebhooksListOption struct {
	Project string `url:"project,omitempty"` // Description:"Project key",ExampleValue:"my_project"
}

// List Search for global webhooks or project webhooks. Webhooks are ordered by name.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) List(opt *WebhooksListOption) (Resp *WebhooksListResp, err error) {
	err := s.ValidateListOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "webhooks/list", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}

// [TODO] you should call the <Update> func manually and complete the fields of this struct
type WebhooksUpdateResp struct{}
type WebhooksUpdateOption struct {
	Name    string `url:"name,omitempty"`    // Description:"new name of the webhook",ExampleValue:"My Webhook"
	Url     string `url:"url,omitempty"`     // Description:"new url to be called by the webhook",ExampleValue:"https://www.my-webhook-listener.com/sonar"
	Webhook string `url:"webhook,omitempty"` // Description:"The key of the webhook to be updated,auto-generated value can be obtained through api/webhooks/create or api/webhooks/list",ExampleValue:"my_project"
}

// Update Update a Webhook.<br>Requires 'Administer' permission on the specified project, or global 'Administer' permission.
func (s *WebhooksService) Update(opt *WebhooksUpdateOption) (Resp *WebhooksUpdateResp, err error) {
	err := s.ValidateUpdateOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("POST", "webhooks/update", Opt)
	if err != nil {
		return
	}
	err = s.client.Do(req, Resp)
	if err != nil {
		return
	}
	return
}
