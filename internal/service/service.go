package service

import (
	"math/rand"
	"time"

	"github.com/xLeSHka/life/pkg/life"
)

// Для хранения состояния
type LifeService struct {
	currentWorld *life.World
	nextWorld    *life.World
}

func New(height, width int) (*LifeService, error) {
	rand.NewSource(time.Now().UTC().UnixNano())

	currentWorld, err := life.NewWorld(height, width)
	if err != nil {
		return nil, err
	}
	// Заполним случайными показателями, чтобы упростить пример
	currentWorld.RandInit(10)

	newWorld, err := life.NewWorld(height, width)
	if err != nil {
		return nil, err
	}

	ls := LifeService{
		currentWorld: currentWorld,
		nextWorld:    newWorld,
	}

	return &ls, nil
}

// Получаем очередное состояние игры
func (ls *LifeService) NewState() *life.World {
	life.NextState(ls.currentWorld, ls.nextWorld)

	ls.currentWorld = ls.nextWorld

	return ls.currentWorld
}
