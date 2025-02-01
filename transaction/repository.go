package transaction

import "gorm.io/gorm"

// repository.go -> Untuk akses ke database
type repository struct {
	db *gorm.DB
}

// __construct
type Repository interface {
	GetByCampaignID(campaignID int) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// function GetByCampaignID(campaignID int) ([]Transaction, error)
func (r *repository) GetByCampaignID(campaignID int) ([]Transaction, error) {
	var transactions []Transaction

	err := r.db.Preload("User").Where("campaign_id = ?", campaignID).Order("id desc").Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
