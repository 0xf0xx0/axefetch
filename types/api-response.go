package types
type ApiInfo struct {
	AsicCount              int     `json:"asicCount"`
	AsicModel              string  `json:"asicModel"`
	BestDiff               string  `json:"bestDiff"`
	BestSessionDiff        string  `json:"bestSessionDiff"`
	BoardVersion           string  `json:"boardVersion"`
	CoreVoltage            int     `json:"coreVoltage"`
	CoreVoltageActual      int     `json:"coreVoltageActual"`
	Current                float64 `json:"current"`
	Display                string  `json:"display"`
	ExpectedHashrate       int     `json:"expectedHashrate"`
	FallbackStratumPort    int     `json:"fallbackStratumPort"`
	FallbackStratumURL     string  `json:"fallbackStratumURL"`
	FallbackStratumUser    string  `json:"fallbackStratumUser"`
	FreeHeap               int     `json:"freeHeap"`
	Frequency              int     `json:"frequency"`
	Hashrate               float64 `json:"hashrate"`
	Hostname               string  `json:"hostname"`
	IdfVersion             string  `json:"idfVersion"`
	IsUsingFallbackStratum int	   `json:"isUsingFallbackStratum"`
	MaxPower               int     `json:"maxPower"`
	NominalVoltage         int     `json:"nominalVoltage"`
	Power                  float64 `json:"power"`
	SharesAccepted         int     `json:"sharesAccepted"`
	SharesRejected         int     `json:"sharesRejected"`
	SmallCoreCount         int     `json:"smallCoreCount"`
	StratumDiff            int     `json:"stratumDiff"`
	StratumPort            int     `json:"stratumPort"`
	StratumURL             string  `json:"stratumURL"`
	StratumUser            string  `json:"stratumUser"`
	Temp                   float64 `json:"temp"`
	UptimeSeconds          int     `json:"uptimeSeconds"`
	Version                string  `json:"version"`
	Voltage                float64 `json:"voltage"`
	VrTemp                 int     `json:"vrTemp"`
	WifiRSSI               int     `json:"wifiRSSI"`
}
