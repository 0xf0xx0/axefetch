package types

type ApiInfo struct {
	AsicCount              uint8   `json:"asicCount"`
	AsicModel              string  `json:"asicModel"`
	BestDiff               uint64  `json:"bestDiff"`
	BestSessionDiff        uint64  `json:"bestSessionDiff"`
	BoardFamily            string  `json:"boardFamily"`
	BoardVersion           string  `json:"boardVersion"`
	BoardVendor            string  `json:"boardVendor"`
	CoreVoltage            uint16  `json:"coreVoltage"`
	CoreVoltageActual      uint16  `json:"coreVoltageActual"`
	Current                float64 `json:"current"`
	Display                string  `json:"display"`
	ExpectedHashrate       float64  `json:"expectedHashrate"`
	FallbackStratumPort    uint16  `json:"fallbackStratumPort"`
	FallbackStratumURL     string  `json:"fallbackStratumURL"`
	FallbackStratumUser    string  `json:"fallbackStratumUser"`
	FreeHeap               uint32  `json:"freeHeap"`
	Frequency              float64  `json:"frequency"`
	Hashrate               float64 `json:"hashrate"`
	Hostname               string  `json:"hostname"`
	IdfVersion             string  `json:"idfVersion"`
	IsUsingFallbackStratum bool    `json:"isUsingFallbackStratum"`
	MaxPower               uint8   `json:"maxPower"`
	NominalVoltage         uint    `json:"nominalVoltage"`
	Power                  float64 `json:"power"`
	SharesAccepted         uint    `json:"sharesAccepted"`
	SharesRejected         uint    `json:"sharesRejected"`
	SmallCoreCount         uint16  `json:"smallCoreCount"`
	StratumDiff            uint    `json:"stratumDiff"`
	StratumPort            uint16  `json:"stratumPort"`
	StratumURL             string  `json:"stratumURL"`
	StratumUser            string  `json:"stratumUser"`
	Temp                   float64 `json:"temp"`
	UptimeSeconds          uint32    `json:"uptimeSeconds"`
	Version                string  `json:"version"`
	Voltage                float64 `json:"voltage"`
	VrTemp                 float64 `json:"vrTemp"`
	WifiRSSI               int8    `json:"wifiRSSI"`
}
