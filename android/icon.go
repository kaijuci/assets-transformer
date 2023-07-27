package android

type AndroidIconType string

const (
	AndroidIconTypeLauncher     AndroidIconType = "launcher"
	AndroidIconTypeAction       AndroidIconType = "action"
	AndroidIconTypeTab          AndroidIconType = "tab"
	AndroidIconTypeMenu         AndroidIconType = "menu"
	AndroidIconTypeDialog       AndroidIconType = "dialog"
	AndroidIconTypeSmall        AndroidIconType = "small"
	AndroidIconTypeStatusBar    AndroidIconType = "statusbar"
	AndroidIconTypeToolbar      AndroidIconType = "toolbar"
	AndroidIconTypeNotification AndroidIconType = "notification"
	AndroidIconTypeSplash       AndroidIconType = "splash"
	AndroidIconTypeWeb          AndroidIconType = "web"
	AndroidIconTypeStore        AndroidIconType = "store"
	AndroidIconTypeFeature      AndroidIconType = "feature"
	AndroidIconTypeAny          AndroidIconType = "any"
)

type AssetDPIList []AssetDPI

type AndroidIconSpec struct {
	Prefix string `json:"prefix"`
	Width  uint   `json:"width"`
	Height uint   `json:"height"`
}

type AndroidIconSpecDictionary map[AndroidIconType]AndroidIconSpec

func newAndroidIconSpecDictionary() AndroidIconSpecDictionary {
	return AndroidIconSpecDictionary{
		AndroidIconTypeLauncher:     AndroidIconSpec{Prefix: "ic_launcher", Width: 512, Height: 512},
		AndroidIconTypeAction:       AndroidIconSpec{Prefix: "ic_action", Width: 192, Height: 192},
		AndroidIconTypeTab:          AndroidIconSpec{Prefix: "ic_tab", Width: 128, Height: 128},
		AndroidIconTypeMenu:         AndroidIconSpec{Prefix: "ic_menu", Width: 96, Height: 96},
		AndroidIconTypeDialog:       AndroidIconSpec{Prefix: "ic_dialog", Width: 96, Height: 96},
		AndroidIconTypeSmall:        AndroidIconSpec{Prefix: "ic_small", Width: 48, Height: 48},
		AndroidIconTypeStatusBar:    AndroidIconSpec{Prefix: "ic_status_bar", Width: 24, Height: 24},
		AndroidIconTypeToolbar:      AndroidIconSpec{Prefix: "ic_toolbar", Width: 48, Height: 48},
		AndroidIconTypeNotification: AndroidIconSpec{Prefix: "ic_notification", Width: 24, Height: 24},
		AndroidIconTypeSplash:       AndroidIconSpec{Prefix: "ic_splash", Width: 512, Height: 512},
		AndroidIconTypeWeb:          AndroidIconSpec{Prefix: "ic_web", Width: 512, Height: 512},
		AndroidIconTypeStore:        AndroidIconSpec{Prefix: "ic_store", Width: 512, Height: 512},
		AndroidIconTypeFeature:      AndroidIconSpec{Prefix: "ic_feature", Width: 1024, Height: 500},
		AndroidIconTypeAny:          AndroidIconSpec{Prefix: "ic_any", Width: 512, Height: 512},
	}
}
