package esocial

type Basescp struct {
	Vrbccp00       float64  `xml:"vrBcCp00"`
	Vrbccp15       float64  `xml:"vrBcCp15"`
	Vrbccp20       float64  `xml:"vrBcCp20"`
	Vrbccp25       float64  `xml:"vrBcCp25"`
	Vrsuspbccp00   float64  `xml:"vrSuspBcCp00"`
	Vrsuspbccp15   float64  `xml:"vrSuspBcCp15"`
	Vrsuspbccp20   float64  `xml:"vrSuspBcCp20"`
	Vrsuspbccp25   float64  `xml:"vrSuspBcCp25"`
	Vrbccp00va     []string `xml:"vrBcCp00VA,omitempty"`
	Vrbccp15va     []string `xml:"vrBcCp15VA,omitempty"`
	Vrbccp20va     []string `xml:"vrBcCp20VA,omitempty"`
	Vrbccp25va     []string `xml:"vrBcCp25VA,omitempty"`
	Vrsuspbccp00va []string `xml:"vrSuspBcCp00VA,omitempty"`
	Vrsuspbccp15va []string `xml:"vrSuspBcCp15VA,omitempty"`
	Vrsuspbccp20va []string `xml:"vrSuspBcCp20VA,omitempty"`
	Vrsuspbccp25va []string `xml:"vrSuspBcCp25VA,omitempty"`
	Vrdescsest     float64  `xml:"vrDescSest"`
	Vrcalcsest     float64  `xml:"vrCalcSest"`
	Vrdescsenat    float64  `xml:"vrDescSenat"`
	Vrcalcsenat    float64  `xml:"vrCalcSenat"`
	Vrsalfam       float64  `xml:"vrSalFam"`
	Vrsalmat       float64  `xml:"vrSalMat"`
}
