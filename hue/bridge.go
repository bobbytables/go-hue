package hue

type Bridge struct {
	IP       string
	Username string
	Password string
}

type BridgeAuthorization struct {
	DeviceType string `json:"device_type"`
	Username   string `json:"username"`
}

func NewBridge() *Bridge {
	&Bridge{}
}
