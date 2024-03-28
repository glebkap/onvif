package event

import (
	"context"

	"github.com/glebkap/onvif"
	"github.com/glebkap/onvif/event"
	"github.com/glebkap/onvif/sdk"
	"github.com/juju/errors"
)

// Call_PullMessages forwards the call to dev.CallMethod() then parses the payload of the reply as a PullMessagesResponse.
func Call_PullMessages(ctx context.Context, dev *onvif.Device, request event.PullMessages, headers []onvif.Header) (event.PullMessagesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			PullMessagesResponse event.PullMessagesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethodWithHeaders(request, headers); err != nil {
		return reply.Body.PullMessagesResponse, errors.Annotate(err, "call")
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "PullMessages")
		return reply.Body.PullMessagesResponse, errors.Annotate(err, "reply")
	}
}
