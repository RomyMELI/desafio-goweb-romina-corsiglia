package tickets


import(
		"context"
		"desafio-goweb-romina-corsiglia/internal/domain"
)
/*
Desarrollar los métodos correspondientes a la estructura Ticket. 
Uno de ellos debe devolver la cantidad de tickets de un destino. 
El  otro método debe devolver el promedio de personas que viajan 
a un país determinado en un dia:
*/

type Service interface {
	GetAll(ctx context.Context) ([]domain.Ticket, error)
	GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error)
	AverageDestination(ctx context.Context, destination string) (int, error)
}
type service struct {
	repository Repository
}


func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}


func (s *service) GetAll(ctx context.Context) ([]domain.Ticket, error){
	tickets, err := s.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	return tickets, nil
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) ([]domain.Ticket, error){
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil{
		return nil, err
	}

	return tickets, nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (int, error){
	tickets, err := s.repository.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}

	var avg int
	
	for _, ticket := range tickets {
		if ticket.Country == destination {
			avg += 1
		}
	}

	return avg, nil
}