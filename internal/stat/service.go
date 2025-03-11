package stat

import (
	"log"

	"github.com/arxanev/adv/pkg/event"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository *StatRepository
}

func NewStatService(desp *StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       desp.EventBus,
		StatRepository: desp.StatRepository,
	}
}

func (s *StatService) AddClick() {
	for msg := range s.EventBus.Subscribe() {
		if msg.Type == event.EventLinkVisited {
			id, ok := msg.Data.(uint)
			if ok {
				log.Fatalln("Bad", msg.Data)
				continue
			}
			s.StatRepository.AddClick(id)
		}
	}
}
