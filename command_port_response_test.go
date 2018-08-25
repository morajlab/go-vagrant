package vagrant

import (
	"testing"
)

const portOutput = `
1534865103,default,metadata,provider,virtualbox
1534865103,,ui,info,The forwarded ports for the machine are listed below. Please note that\nthese values may differ from values configured in the Vagrantfile if the\nprovider supports automatic port collision detection and resolution.
1534865103,,ui,info,
1534865103,,ui,info,    22 (guest) => 2222 (host)
1534865103,,ui,info,    80 (guest) => 8080 (host)
1534865103,default,forwarded_port,22,2222
1534865103,default,forwarded_port,80,8080
`

func TestPortResponse_handleOutput(t *testing.T) {
	parser := MockOutputParser{}
	data := newPortResponse()
	parser.Run(portOutput, &data)

	if data.Error != nil {
		t.Errorf("Successful vagrant port should not have set an error: %v", data.Error)
	}

	if len(data.ForwardedPorts) != 1 {
		t.Fatalf("Expected forwarded ports for 1 VM; got %v", len(data.ForwardedPorts))
	}

	forwardedPorts, ok := data.ForwardedPorts["default"]
	if !ok {
		t.Fatal("Expected forwarded ports for 'default', but there were none.")
	}

	if len(forwardedPorts) != 2 {
		t.Fatalf("Expected 1 forwarded port; got %v", len(forwardedPorts))
	}
	if forwardedPorts[0].Guest != 22 || forwardedPorts[0].Host != 2222 {
		t.Errorf("Expected guest port 22 -> host 2222; got %v", forwardedPorts[0])
	}
	if forwardedPorts[1].Guest != 80 || forwardedPorts[1].Host != 8080 {
		t.Errorf("Expected guest port 80 -> host 8080; got %v", forwardedPorts[1])
	}
}
