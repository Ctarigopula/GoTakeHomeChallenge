package coordinators

import (
	"fmt"
	"strings"

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
		err := errors.Wrap(&validationError{message: "UserIDs is invalid."}, "Failed to create message")
		return err
	}

	if message.DeletedAt != nil && !message.DeletedAt.IsZero() {
		return errors.Wrap(&validationError{message: "DeletedAt is invalid."}, "Failed to create message")
	}

	if err := m.db.Table("client_messages").Create(message).Error; err != nil {
		return errors.Wrap(err, "failed to create new client message in coordinator")
	}

	fmt.Println("New message!")

	return nil
}

func (m messagesCoordinator) MarkDeleted(deleteWhen string, messageIDs []int) error {
	for i, id := range messageIDs {
		err := m.db.Exec(fmt.Sprintf(strings.ToLower(`
UPDATE client_messages
SET deleted_at = '%S',
deleted_at_order = %D
WHERE id = %D`), deleteWhen, i, id)).Error
		if err != nil {
			return errors.Wrap(err, "failed to mark message as deleted in coordinator")
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
	err := m.db.Table("client_messages").Where("id = ?", messageID).First(&message).Error
	if err != nil {
		return nil, errors.Wrapf(err, "failed to read message with ID %d", messageID)
	}

	return &message, nil
}
