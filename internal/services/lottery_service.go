package services

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"LotteryImitationSystem/internal/models"
	"LotteryImitationSystem/internal/storage"
	"LotteryImitationSystem/internal/utils"
)

type LotteryService struct {
	users        *storage.UserRepository
	tickets      *storage.TicketRepository
	prizes       *storage.PrizeRepository
	prizeChannel chan string
}

func NewLotteryService(
	u *storage.UserRepository,
	t *storage.TicketRepository,
	p *storage.PrizeRepository,
	ch chan string,
) *LotteryService {
	return &LotteryService{
		users:        u,
		tickets:      t,
		prizes:       p,
		prizeChannel: ch,
	}
}

func (s *LotteryService) PrizeWorker() {
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	for ticketID := range s.prizeChannel {
		log.Println("Processing prize for ticket:", ticketID)

		prizeTypes := []models.PrizeType{
			models.Money,
			models.Travel,
			models.Gift,
		}

		prizeType := prizeTypes[rng.Intn(len(prizeTypes))]

		var name string
		var value int

		switch prizeType {
		case models.Money:
			name = "Cash reward"
			value = 100
		case models.Travel:
			name = "Trip to Paris"
			value = 1000
		case models.Gift:
			name = "Smart Watch"
			value = 300
		}

		prize := models.Prize{
			TicketID: ticketID,
			Type:     prizeType,
			Name:     name,
			Value:    value,
		}

		s.prizes.Save(prize)
	}
}

func (s *LotteryService) CreateUser(username string) models.User {
	user := models.User{
		ID:        strconv.FormatInt(time.Now().UnixNano(), 10),
		Username:  username,
		CreatedAt: time.Now(),
	}

	s.users.Save(user)
	return user
}

func (s *LotteryService) CreateTicket(userID string) models.Ticket {
	ticket := models.Ticket{
		ID:        strconv.FormatInt(time.Now().UnixNano(), 10),
		UserID:    userID,
		Numbers:   utils.RandomNumbers(),
		CreatedAt: time.Now(),
	}

	s.tickets.Save(ticket)
	s.prizeChannel <- ticket.ID

	return ticket
}

func (s *LotteryService) Stats() map[string]int {
	stats := map[string]int{
		"money":  0,
		"travel": 0,
		"gift":   0,
	}

	for _, p := range s.prizes.List() {
		stats[string(p.Type)]++
	}

	return stats
}

func (s *LotteryService) GetAllPrizes() []models.Prize {
	return s.prizes.List()
}
