package gopro

import (
	"regexp"
	"time"
)

type GoProVersion struct {
	InfoVersion        string `json:"info version"`
	FirmwareVersion    string `json:"firmware version"`
	WifiMac            string `json:"wifi mac"`
	CameraType         string `json:"camera type"`
	CameraSerialNumber string `json:"camera serial number"`
}

type Directory string

const (
	DCIM Directory = "DCIM"
	MISC Directory = "MISC"
)

type GoProType string

const (
	V2  GoProType = "v2"
	MAX GoProType = "max"
	V1  GoProType = "v1"

	UNKNOWN GoProType = "unknown"
)

type FileOfInterest string

const (
	GetStarted FileOfInterest = "Get_started_with_GoPro.url"
	Version    FileOfInterest = "version.txt"
)

type SortOptions struct {
	ByDays             bool
	ByLocation         bool
	SkipAuxiliaryFiles bool
	AddHiLightTags     bool
	ByCamera           bool
	DateFormat         string
	BufferSize         int
	Prefix             string
	DateRange          []time.Time
}

type FileType string

const (
	Video              FileType = "video"
	Photo              FileType = "photo"
	PowerPano          FileType = "powerpano"
	ChapteredVideo     FileType = "chapteredvideo"
	Multishot          FileType = "multishot"
	LowResolutionVideo FileType = "lrv"
	Thumbnail          FileType = "thm"
	RawPhoto           FileType = "gpr"
)

type FileTypeMatch struct {
	Regex    *regexp.Regexp
	Type     FileType
	HeroMode bool
}

type Camera struct {
	Model                  int      `json:"model"`
	ModelString            string   `json:"model_string"`
	Name                   string   `json:"name"`
	Version                string   `json:"version"`
	URL                    string   `json:"url"`
	ReleaseDate            string   `json:"release_date"`
	Sha1                   string   `json:"sha1"`
	RequiredFreeLocalSpace int      `json:"required_free_local_space"`
	CriticalUpdate         bool     `json:"criticalUpdate"`
	Languages              []string `json:"languages"`
	ReleaseHTML            string   `json:"release_html"`
	LicenseHTML            string   `json:"license_html"`
}

type FirmwareCatalog struct {
	Version string   `json:"version"`
	Magic   string   `json:"magic"`
	Group   string   `json:"group"`
	Cameras []Camera `json:"cameras"`
}

type cameraInfo struct {
	Info struct {
		ModelNumber             int      `json:"model_number"`
		ModelName               string   `json:"model_name"`
		FirmwareVersion         string   `json:"firmware_version"`
		SerialNumber            string   `json:"serial_number"`
		BoardType               string   `json:"board_type"`
		ApMac                   string   `json:"ap_mac"`
		ApSsid                  string   `json:"ap_ssid"`
		ApHasDefaultCredentials string   `json:"ap_has_default_credentials"`
		Capabilities            string   `json:"capabilities"`
		LensCount               string   `json:"lens_count"`
		UpdateRequired          string   `json:"update_required"`
		RequiredNetworks        []string `json:"required_networks"`
		HTTPKeepalive           int      `json:"http_keepalive"`
	} `json:"info"`
}

type goProMediaList struct {
	ID    string `json:"id"`
	Media []struct {
		D  string `json:"d"`
		Fs []struct {
			N    string        `json:"n"`
			Cre  string        `json:"cre"`
			Mod  string        `json:"mod"`
			Glrv string        `json:"glrv,omitempty"`
			Ls   string        `json:"ls,omitempty"`
			S    string        `json:"s"`
			G    string        `json:"g,omitempty"`
			B    string        `json:"b,omitempty"`
			L    string        `json:"l,omitempty"`
			T    string        `json:"t,omitempty"`
			M    []interface{} `json:"m,omitempty"`
			Raw  string        `json:"raw,omitempty"`
		} `json:"fs"`
	} `json:"media"`
}

type GoProConnectDevice struct {
	IP   string
	Info cameraInfo
}

type goProMediaMetadata struct {
	Cre       string        `json:"cre"`
	S         string        `json:"s"`
	Us        string        `json:"us"`
	Mos       []interface{} `json:"mos"`
	Eis       string        `json:"eis"`
	Pta       string        `json:"pta"`
	Ao        string        `json:"ao"`
	Tr        string        `json:"tr"`
	Mp        string        `json:"mp"`
	Gumi      string        `json:"gumi"`
	Ls        string        `json:"ls"`
	Cl        string        `json:"cl"`
	Hc        string        `json:"hc"`
	Hi        []interface{} `json:"hi"`
	Dur       string        `json:"dur"`
	W         string        `json:"w"`
	H         string        `json:"h"`
	Fps       int           `json:"fps,string"`
	FpsDenom  int           `json:"fps_denom,string"`
	Prog      string        `json:"prog"`
	Subsample string        `json:"subsample"`
}
