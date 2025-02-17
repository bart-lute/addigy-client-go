package fact_entities

type FactResponse struct {
	Fact        Fact        `json:"fact"`
	Instruction Instruction `json:"instruction"`
}
