package campaign

import "strings"

// struct campaign formatter
type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageURL         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_mount"`
	Slug             string `json:"slug"`
}

// single formatter => formatter 1 data
func FormatCampaign(campaign Campaign) CampaignFormatter {
	// object campaign formatter
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.Slug = campaign.Slug
	campaignFormatter.ImageURL = ""

	// check apakah ada image atau tidak
	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}
	return campaignFormatter
}

// multi formatter =? formatter banyak data
func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	campaignsFormatter := []CampaignFormatter{}

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

// Detail campaign formatter
type CampaignDetailFormatter struct {
	ID               int                      `json:"id"`
	Name             string                   `json:"name"`
	ShortDescription string                   `json:"short_description"`
	Description      string                   `json:"description"`
	ImageURL         string                   `json:"image_url"`
	GoalAmount       int                      `json:"goal_amount"`
	CurrentAmount    int                      `json:"current_amount"`
	UserID           int                      `json:"user_id"`
	Slug             string                   `json:"slug"`
	Perks            []string                 `json:"perks"`
	User             CampaignUserFormatter    `json:"user"`
	Images           []CampaignImageFormatter `json:"images"`
}

// campaign user formatter
type CampaignUserFormatter struct {
	Name     string `json:"name"`
	ImageURL string `json:"image_url"`
}

// campaign image formatter
type CampaignImageFormatter struct {
	ImageURL  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	// object campaign detail formatter
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.Slug = campaign.Slug
	campaignDetailFormatter.ImageURL = ""

	// check apakah ada image atau tidak
	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageURL = campaign.CampaignImages[0].FileName
	}

	// mapping perks yang berisi array of string (data string dalam array)
	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	// campaign user formatter
	user := campaign.User
	campaignUserFormatter := CampaignUserFormatter{}
	campaignUserFormatter.Name = user.Name
	campaignUserFormatter.ImageURL = user.AvatarFileName

	campaignDetailFormatter.User = campaignUserFormatter

	// campaign image formatter
	images := []CampaignImageFormatter{}
	for _, image := range campaign.CampaignImages {
		campaignImageFormatter := CampaignImageFormatter{}
		campaignImageFormatter.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}
		campaignImageFormatter.IsPrimary = isPrimary
		images = append(images, campaignImageFormatter)
	}
	campaignDetailFormatter.Images = images

	return campaignDetailFormatter
}
