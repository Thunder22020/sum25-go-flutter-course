package chatcore

import (
	"context"
	"errors"
	"sync"
	"time"
)

// Message represents a chat message
type Message struct {
	Sender    string
	Recipient string
	Content   string
	Broadcast bool
	Timestamp int64
}

// Broker handles message routing between users
type Broker struct {
	ctx        context.Context
	input      chan Message
	users      map[string]chan Message
	usersMutex sync.RWMutex
	done       chan struct{}
}

// NewBroker creates a new message broker
func NewBroker(ctx context.Context) *Broker {
	return &Broker{
		ctx:   ctx,
		input: make(chan Message, 100),
		users: make(map[string]chan Message),
		done:  make(chan struct{}),
	}
}

// Run starts the broker event loop (fan-in/fan-out pattern)
func (b *Broker) Run() {
	go func() {
		for {
			select {
			case <-b.ctx.Done():
				close(b.done)
				return
			case msg := <-b.input:
				msg.Timestamp = time.Now().Unix()
				if msg.Broadcast {
					b.usersMutex.RLock()
					for _, ch := range b.users {
						select {
						case ch <- msg:
						default:
							// skip if channel full
						}
					}
					b.usersMutex.RUnlock()
				} else {
					b.usersMutex.RLock()
					ch, ok := b.users[msg.Recipient]
					b.usersMutex.RUnlock()
					if ok {
						select {
						case ch <- msg:
						default:
							// skip if full
						}
					}
				}
			}
		}
	}()
}

// SendMessage sends a message to the broker
func (b *Broker) SendMessage(msg Message) error {
	select {
	case <-b.ctx.Done():
		return errors.New("broker stopped")
	case b.input <- msg:
		return nil
	}
}

// RegisterUser adds a user to the broker
func (b *Broker) RegisterUser(userID string, recv chan Message) {
	b.usersMutex.Lock()
	defer b.usersMutex.Unlock()
	b.users[userID] = recv
}

// UnregisterUser removes a user from the broker
func (b *Broker) UnregisterUser(userID string) {
	b.usersMutex.Lock()
	defer b.usersMutex.Unlock()
	delete(b.users, userID)
}
