// Repository => untuk mengambil data atau memanipulasi data dari database
package campaign

import (
	"gorm.io/gorm"
)

type Repository interface {
	// Fungsi dengan ([]Campaign), karena akan mengembalikan lebih dari 1 data campaign yang ada dari database jadinya menngunakan array slice
	FindAll() ([]Campaign, error)
	FindByUserID(userID int) ([]Campaign, error)
	FindByID(ID int) (Campaign, error)
	Save(campaign Campaign) (Campaign, error)                       // save/create campaign baru
	Update(campaign Campaign) (Campaign, error)                     // update/edit campaign
	CreateImage(campaignImage CampaignImage) (CampaignImage, error) // create campaign upload image
	MarkAllImagesAsNonPrimary(campaignID int) (bool, error)         // ubah is_primary true ke false (ini kasus kalau kita melakukan upload data image jika is_primary yang sebelumnya true maka kita ubah menjadi false)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Repository: Find All
func (r *repository) FindAll() ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

// Repository: Find By User ID
func (r *repository) FindByUserID(userID int) ([]Campaign, error) {
	var campaigns []Campaign

	err := r.db.Where("user_id = ?", userID).Preload("CampaignImages", "campaign_images.is_primary = 1").Find(&campaigns).Error
	if err != nil {
		return campaigns, err
	}
	return campaigns, nil
}

// Repository: Find By ID
func (r *repository) FindByID(ID int) (Campaign, error) {
	var campaign Campaign

	err := r.db.Preload("User").Preload("CampaignImages").Where("id = ?", ID).Find(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

// Repository: save/create campaign baru
func (r *repository) Save(campaign Campaign) (Campaign, error) {
	err := r.db.Create(&campaign).Error
	if err != nil {
		return campaign, err
	}
	return campaign, nil
}

// Repository: update/edit campaign
func (r *repository) Update(campaign Campaign) (Campaign, error) {
	err := r.db.Save(&campaign).Error

	// check jika error
	if err != nil {
		return campaign, err
	}

	// jika sukses
	return campaign, nil
}

// Repository: create upload image campaign
func (r *repository) CreateImage(campaignImage CampaignImage) (CampaignImage, error) {
	err := r.db.Create(&campaignImage).Error
	if err != nil {
		return campaignImage, err
	}
	return campaignImage, nil
}

func (r *repository) MarkAllImagesAsNonPrimary(campaignID int) (bool, error) {
	// repository:
	// 1. create image/save data image ke dalam table campaign_images
	// 2. ubah is_primary true ke false (ini kasus kalau kita melakukan upload data image jika is_primary yang sebelumnya true maka kita ubah menjadi false)

	// UPDATE campaign_images SET is_primary = false WHERE campaign_id = 1
	err := r.db.Model(&CampaignImage{}).Where("campaign_id = ?", campaignID).Update("is_primary", false).Error
	if err != nil {
		return false, nil
	}
	return true, nil
}
