package esocial

type Infohomolog struct {
	Sithomolog int64    `xml:"sitHomolog"`
	Dthomolog  []string `xml:"dtHomolog,omitempty"`
}
