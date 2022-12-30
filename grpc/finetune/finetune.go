package finetune

type FinetuneIndex struct {
	Id   string
	User string
}

type FinetuneInfo struct {
	Duration string
	Status   string
}

type FinetuneService interface {
	SetFinetuneInfo(*FinetuneIndex, *FinetuneInfo) error
}
