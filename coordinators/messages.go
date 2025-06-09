package coordinators

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	"take-home-challenge/models"
)

type MessagesCoordinator interface {
	Create(message *models.Message) error
	MarkDeleted(deleteWhen string, messageIDs []int) error
	Read(messageID int) (*models.Message, error)
}

type messagesCoordinator struct {
	db *gorm.DB
}

func NewMessagesCoordinator(db *gorm.DB) MessagesCoordinator {
	return &messagesCoordinator{
		db: db,
	}
}

func (m messagesCoordinator) Create(message *models.Message) error {
	if len(message.UserIDs) == 0 {
		return errors.Wrap(&validationError{message: "UserIDs is invalid."}, "Failed to create message")
	}

	if message.DeletedAt != nil && !message.DeletedAt.IsZero() {
		return errors.Wrap(&validationError{message: "DeletedAt is invalid."}, "Failed to create message")
	}

	if err := m.db.Create(&message).Error; err != nil {
		return errors.Wrap(err, "Failed to create new client message in coordinator")
	}

	fmt.Println("New message!")
	return nil
}

func (m messagesCoordinator) MarkDeleted(deleteWhen string, messageIDs []int) error {
	for i, id := range messageIDs {
		err := m.db.Exec(`
			UPDATE client_messages
			SET deleted_at = ?, deleted_at_order = ?
			WHERE id = ?`,
			deleteWhen, i, id,
		).Error
		if err != nil {
			return errors.Wrapf(err, "failed to mark message ID %d as deleted", id)
		}
		fmt.Println("Marked message as deleted!")
	}
	return nil
}

// Read retrieves a message from the database by its ID.
// Returns the Message if found, otherwise returns an error.
func (m messagesCoordinator) Read(messageID int) (*models.Message, error) {
	var message models.Message

	// Attempt to find the message with the given ID
	err := m.db.Where("id = ?", messageID).First(&message).Error
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read message with ID %d", messageID)
	}
	fmt.Println("Found the message!")
	return &message, nil
}
