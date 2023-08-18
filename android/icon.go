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
	DPI    uint   `json:"dpi"`
}

type AndroidIconSpecDictionary map[AndroidIconType]AndroidIconSpec

func NewAndroidIconSpecDictionary() AndroidIconSpecDictionary {
	return AndroidIconSpecDictionary{
		AndroidIconTypeLauncher:     AndroidIconSpec{Prefix: "ic_launcher", DPI: 48},
		AndroidIconTypeAction:       AndroidIconSpec{Prefix: "ic_action", DPI: 32},
		AndroidIconTypeTab:          AndroidIconSpec{Prefix: "ic_tab", DPI: 32},
		AndroidIconTypeMenu:         AndroidIconSpec{Prefix: "ic_menu", DPI: 48},
		AndroidIconTypeDialog:       AndroidIconSpec{Prefix: "ic_dialog", DPI: 32},
		AndroidIconTypeSmall:        AndroidIconSpec{Prefix: "ic_small", DPI: 16},
		AndroidIconTypeStatusBar:    AndroidIconSpec{Prefix: "ic_status_bar", DPI: 24},
		AndroidIconTypeToolbar:      AndroidIconSpec{Prefix: "ic_toolbar", DPI: 32},
		AndroidIconTypeNotification: AndroidIconSpec{Prefix: "ic_notification", DPI: 24},
	}
}
