package mxtoolbox

import "fmt"

type HistoryEntry struct {
	Name                   string      `json:"Name"`
	Status                 int         `json:"Status"`
	SubactionStatus        int         `json:"SubactionStatus"`
	Result                 string      `json:"Result"`
	MonitorUID             string      `json:"MonitorUID"`
	TransitionedString     string      `json:"TransitionedString"`
	AlertGroupUID          string      `json:"AlertGroupUID"`
	LookupActionResultUID  string      `json:"LookupActionResultUID"`
	MonitorCommand         string      `json:"MonitorCommand"`
	MonitorCommandArgument string      `json:"MonitorCommandArgument"`
	MonitorLastChecked     string      `json:"MonitorLastChecked"`
	MonitorCreatedOn       string      `json:"MonitorCreatedOn"`
	StatusSummary          interface{} `json:"StatusSummary"`
	DomainLarUID           string      `json:"DomainLarUID"`
	RecordCount            int         `json:"RecordCount"`
	TransitionID           string      `json:"TransitionID"`
	TransitionedOn         string      `json:"TransitionedOn"`
	TransitionedSubaction  struct {
		ID        int         `json:"ID"`
		Name      string      `json:"Name"`
		Info      string      `json:"Info"`
		URL       interface{} `json:"Url"`
		DelistURL interface{} `json:"DelistUrl"`
	} `json:"TransitionedSubaction"`
}

func (s *MonitorsService) History(monitorUUID string) ([]HistoryEntry, error) {
	req, err := s.client.NewRequest("GET", fmt.Sprintf("history/%s", monitorUUID), nil)
	if err != nil {
		// handle err
	}

	var entries []HistoryEntry
	err = s.client.Do(req, &entries)
	if err != nil {
		return nil, err
	}
	return entries, nil
}
