package tickets

import (
	"errors"
	"fmt"
	"strings"
)

type Ticket struct {
	ID          string
	Title       string
	Description string
	Priority    int
	AssigneeID  string
	Status      string
}

type TicketStore struct {
	items map[string]Ticket
}

func NewTicketStore() *TicketStore {
	ticketStore := TicketStore{
		items: make(map[string]Ticket),
	}
	return &ticketStore
}

func (ticketStore *TicketStore) Create(ticket Ticket) error {
	if strings.TrimSpace(ticket.ID) == "" {
		return errors.New("Invalid Id")
	}
	if _, exists := ticketStore.items[ticket.ID]; exists {
		return errors.New("Ticket with this Id already exists!")
	}

	if strings.TrimSpace(ticket.Title) == "" {
		return errors.New("Empty title")
	}
	if ticket.Priority < 1 || ticket.Priority > 3 {
		return errors.New("Invalid priority")
	}
	if !(strings.HasPrefix(ticket.Status, "OPEN")) {
		return errors.New("Has to start with OPEN")
	}
	ticketStore.items[ticket.ID] = ticket
	return nil
}

func (ticketStore *TicketStore) Assign(ticketID, assigneeID string) error {

	if _, exists := ticketStore.items[ticketID]; !exists {
		return errors.New("Ticket with such Id does not exists!")
	}
	if !strings.HasPrefix(ticketStore.items[ticketID].Status, "OPEN") {
		return errors.New("Ticket is not OPEN")
	}
	if strings.TrimSpace(assigneeID) == "" {
		return errors.New("Assignee is empty")
	}
	currentTicket := ticketStore.items[ticketID]
	currentTicket.AssigneeID = assigneeID
	ticketStore.items[ticketID] = currentTicket
	return nil
}

func (ticketStore *TicketStore) Resolve(ticketID string) error {
	if _, exits := ticketStore.items[ticketID]; !exits {
		return errors.New("Ticket with such Id does not exists!")
	}
	currentTicket := ticketStore.items[ticketID]
	if currentTicket.Status != "OPEN" {
		return errors.New("Ticket is not OPEN")
	}
	currentTicket.Status = "DONE"
	ticketStore.items[ticketID] = currentTicket
	return nil
}

func (ticketStore *TicketStore) ListAll() []Ticket {
	AllTickets := make([]Ticket, 0)
	for _, ticket := range ticketStore.items {
		AllTickets = append(AllTickets, ticket)
	}
	return AllTickets
}

func (ticketStore *TicketStore) ListByStatus(status string) []Ticket {
	SortedTickets := make([]Ticket, 0)
	if !(strings.HasPrefix(status, "OPEN") || status == "DONE") {
		return nil
	}
	for _, ticket := range ticketStore.items {

		if strings.HasPrefix(status, "OPEN") && strings.HasPrefix(ticket.Status, "OPEN") {
			SortedTickets = append(SortedTickets, ticket)
		}

		if ticket.Status == "DONE" && status == "DONE" {
			SortedTickets = append(SortedTickets, ticket)
		}
	}

	return SortedTickets
}

func (ticketStore *TicketStore) ListUnassigned() []Ticket {

	UnassignedTicket := make([]Ticket, 0)

	for _, ticket := range ticketStore.items {

		if ticket.AssigneeID == "Noone" {
			UnassignedTicket = append(UnassignedTicket, ticket)
		}

	}
	return UnassignedTicket

}

type Agent interface {
	GetID() string
	GetName() string
}

type HumanAgent struct {
	ID   string
	Name string
}

func (human HumanAgent) GetID() string {
	return human.ID
}

func (human HumanAgent) GetName() string {
	return human.Name
}

type BotAgent struct {
	ID      string
	Name    string
	Version string
}

func (bot BotAgent) GetID() string {
	return bot.ID
}

func (bot BotAgent) GetName() string {
	return bot.Name
}

func FormatAgent(a Agent) string {
	switch who := a.(type) {
	case HumanAgent:
		return fmt.Sprintf("%s | %s ", who.GetID(), who.GetName())
	case BotAgent:
		return fmt.Sprintf("%s | %s | bot:%s", who.GetID(), who.GetName(), who.Version)
	default:
		return "Wrong input"
	}

}
