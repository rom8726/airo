package infra

func init() {
	addInfra("kafka", InfraInfo{
		Code:      "kafka",
		Title:     "Kafka",
		Processor: &KafkaProcessor{},
	})
}

type KafkaProcessor struct {
}

func (k KafkaProcessor) Import() string {
	//TODO implement me
	panic("implement me")
}

func (k KafkaProcessor) Config() string {
	//TODO implement me
	panic("implement me")
}

func (k KafkaProcessor) ConfigField() string {
	//TODO implement me
	panic("implement me")
}

func (k KafkaProcessor) InitInAppConstructor() string {
	//TODO implement me
	panic("implement me")
}

func (k KafkaProcessor) Constructor() string {
	//TODO implement me
	panic("implement me")
}

func (k KafkaProcessor) StructField() string {
	//TODO implement me
	panic("implement me")
}

func (k KafkaProcessor) FillStructField() string {
	//TODO implement me
	panic("implement me")
}

func (k KafkaProcessor) Close() string {
	//TODO implement me
	panic("implement me")
}
