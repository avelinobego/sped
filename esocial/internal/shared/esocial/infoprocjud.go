package esocial

type Infoprocjud struct {
	Tpproc      int64    `xml:"tpProc"`
	Nrproc      string   `xml:"nrProc"`
	Codsusp     int64    `xml:"codSusp"`
	Vrcpsusp    []string `xml:"vrCPSusp,omitempty"`
	Vrratsusp   []string `xml:"vrRatSusp,omitempty"`
	Vrsenarsusp []string `xml:"vrSenarSusp,omitempty"`
	Dtsent      string   `xml:"dtSent"`
	Ufvara      string   `xml:"ufVara"`
	Codmunic    int64    `xml:"codMunic"`
	Idvara      int64    `xml:"idVara"`
}
