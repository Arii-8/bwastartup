package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"time"
)

// entity.go -> Untuk membuat struct yang mewakili atau yang merupakan representasi dari table transaction yang ada di dalam database
type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	PaymentURL string
	User       user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
