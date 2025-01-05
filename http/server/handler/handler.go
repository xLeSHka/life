package handler

import (
	"context"
	"net/http"

	"github.com/xLeSHka/life/internal/service"
)

// Создадим новый тип для добавления middleware к обработчикам
type Decorator func(http.Handler) http.Handler

// Объект для хранения состояния игры
type LifeStates struct {
	service.LifeService
}

func New(ctx context.Context,
	lifeService service.LifeService,
) (http.Handler, error) {
	serveMux := http.NewServeMux()

	lifeState := LifeStates{
		LifeService: lifeService,
	}

	serveMux.HandleFunc("/nextstate", lifeState.nextState)

	return serveMux, nil
}

// Функция добавления middleware
func Decorate(next http.Handler, ds ...Decorator) http.Handler {
	decorated := next
	for d := len(ds) - 1; d >= 0; d-- {
		decorated = ds[d](decorated)
	}

	return decorated
}

// Получаем очередное состояние игры
func (ls *LifeStates) nextState(w http.ResponseWriter, r *http.Request) {
	worldState := ls.LifeService.NewState()
	w.Write([]byte(worldState.String()))

}
