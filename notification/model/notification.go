package model

type OneSignalNotificationData struct {
	AppID            string            `json:"app_id"`
	Headings         map[string]string `json:"headings"`
	Contents         map[string]string `json:"contents"`
	IncludePlayerIDs []string          `json:"include_player_ids"`
	Data             map[string]string `json:"data"`
}

type OneSignalNotificationBulkData struct {
	AppID            string              `json:"app_id"`
	Headings         map[string]string   `json:"headings"`
	Contents         map[string]string   `json:"contents"`
	IncludedSegments []string            `json:"included_segments"`
	Data             map[string]string   `json:"data"`
	BigPicture       string              `json:"big_picture"`
	IosAttachments   IosAttachmentsModel `json:"ios_attachments"`
}

type IosAttachmentsModel struct {
	ID1 string `json:"id1"`
}
