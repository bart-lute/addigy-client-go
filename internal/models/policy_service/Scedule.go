package policy_service

type Schedule struct {
	Fr   bool         `json:"fr"`
	From ScheduleTime `json:"from"`
	Mo   bool         `json:"mo"`
	Sa   bool         `json:"sa"`
	Su   bool         `json:"su"`
	Th   bool         `json:"th"`
	To   ScheduleTime `json:"to"`
	Tu   bool         `json:"tu"`
	We   bool         `json:"we"`
}
