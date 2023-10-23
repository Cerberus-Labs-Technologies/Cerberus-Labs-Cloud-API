package reports

// All report related packets

type PlayerReportedPacket struct {
	Name   string `json:"name"`
	Report Report `json:"report"`
}

type ReportInvestigatedPacket struct {
	Name   string `json:"name"`
	Report Report `json:"report"`
}

type ReportUpdatedPacket struct {
	Name      string `json:"name"`
	OldReport Report `json:"oldReport"`
	NewReport Report `json:"newReport"`
}

type ReportDeletedPacket struct {
	Name   string `json:"name"`
	Report Report `json:"report"`
}

type ReportCreatedPacket struct {
	Name   string `json:"name"`
	Report Report `json:"report"`
}
