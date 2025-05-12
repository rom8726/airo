package infra

func init() {
	addInfra("etcd", InfraInfo{
		Code:      "etcd",
		Title:     "Etcd",
		Processor: &EtcdProcessor{},
	})
}

type EtcdProcessor struct {
}

func (e EtcdProcessor) Import() string {
	//TODO implement me
	panic("implement me")
}

func (e EtcdProcessor) Config() string {
	//TODO implement me
	panic("implement me")
}

func (e EtcdProcessor) ConfigField() string {
	//TODO implement me
	panic("implement me")
}

func (e EtcdProcessor) InitInAppConstructor() string {
	//TODO implement me
	panic("implement me")
}

func (e EtcdProcessor) Constructor() string {
	//TODO implement me
	panic("implement me")
}

func (e EtcdProcessor) StructField() string {
	//TODO implement me
	panic("implement me")
}

func (e EtcdProcessor) FillStructField() string {
	//TODO implement me
	panic("implement me")
}

func (e EtcdProcessor) Close() string {
	//TODO implement me
	panic("implement me")
}
