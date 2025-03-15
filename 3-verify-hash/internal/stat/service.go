package stat

import (
	"go-ps-adv-homework/pkg/event"
	"log"
)

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatServiceDependencies struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(dependencies StatServiceDependencies) *StatService {
	return &StatService{
		EventBus:       dependencies.EventBus,
		StatRepository: dependencies.StatRepository,
	}
}
func (service *StatService) AddClick() {
	for msg := range service.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if !ok {
				log.Fatalln("Invalid EventLinkVisited message type: ", msg.Data)
				continue
			}
			service.StatRepository.AddClick(id)
		}
	}
}
