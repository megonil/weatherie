package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"weatherie/internal/utils"
)

type FormatType uint

const (
	FormatDefault FormatType = iota
	FormatTable              // table
	FormatPlain              // just like default but without "code" block
	FormatSmall              // breaf 1-line info
	FormatCustom             // TODO:
)

const _FormatTypeName = "FormatDefaultFormatTableFormatPlainFormatSmallFormatCustom"

var _FormatTypeIndex = [...]uint8{0, 13, 24, 35, 46, 58}

const _FormatTypeLowerName = "formatdefaultformattableformatplainformatsmallformatcustom"

func (i FormatType) String() string {
	if i >= FormatType(len(_FormatTypeIndex)-1) {
		return fmt.Sprintf("FormatType(%d)", i)
	}
	return _FormatTypeName[_FormatTypeIndex[i]:_FormatTypeIndex[i+1]]
}

type InfoNodeType uint

const (
	InfoTemp InfoNodeType = iota
	InfoLocation
	InfoPressure
	InfoUV
	InfoWind
	InfoCloud
)

const _InfoNodeTypeName = "InfoTempInfoLocationInfoPressureInfoUVInfoWindInfoCloud"

var _InfoNodeTypeIndex = [...]uint8{0, 8, 20, 32, 38, 46, 55}

const _InfoNodeTypeLowerName = "infotempinfolocationinfopressureinfouvinfowindinfocloud"

func (i InfoNodeType) String() string {
	if i >= InfoNodeType(len(_InfoNodeTypeIndex)-1) {
		return fmt.Sprintf("InfoNodeType(%d)", i)
	}
	return _InfoNodeTypeName[_InfoNodeTypeIndex[i]:_InfoNodeTypeIndex[i+1]]
}

type UserConfig struct {
	FormatType              FormatType
	InfoNodes               []InfoNodeType
	SavedLocation, ShowIcon bool
	Location                string
}

var defaultInfoNodes = []InfoNodeType{InfoLocation, InfoTemp, InfoWind, InfoCloud}

func (u *UserConfig) JSON() []byte {
	b, err := json.Marshal(u)
	if err != nil {
		log.Fatalf("Error marshalling UserConfig struct: %e\n", err)
	}

	return b
}

func userHash(userID int64) string {
	return fmt.Sprintf("user:%d", userID)
}

func SaveConfig(config *UserConfig, userID int64) error {
	b := config.JSON()

	return utils.RedisSet(userHash(userID), b).Err()
}

func GetConfig(userID int64) *UserConfig {
	r := utils.RedisGet(userHash(userID))
	if r.Err() != nil {
		log.Printf("Redis get error: %e\n", r.Err())
		return nil
	}
	b, err := r.Bytes()
	if err != nil {
		log.Printf("Unable to get bytes from redis: %e\n", err)
		return nil
	}

	var config UserConfig
	err = json.Unmarshal(b, &config)
	if err != nil {
		log.Printf("Unable to unmarshal config: %e\n", err)
		return nil
	}

	return &config
}

func GetConfigSafe(userID int64) *UserConfig {
	config := GetConfig(userID)

	if config == nil {
		InitConfig(userID)
		config = NewUserConfig()
	}

	return config
}

func NewUserConfig() *UserConfig {
	return &UserConfig{
		FormatType:    FormatDefault,
		InfoNodes:     defaultInfoNodes,
		SavedLocation: false,
		ShowIcon:      true,
	}
}

func InitConfig(userID int64) error {
	return SaveConfig(NewUserConfig(), userID)
}

func (u *UserConfig) String() string {
	var b strings.Builder
	fmt.Fprintf(&b, "Format Type: %s\n", u.FormatType.String())
	fmt.Fprintf(&b, "Show Icon: %v\n", u.ShowIcon)
	if u.SavedLocation {
		fmt.Fprintf(&b, "Location: %s\n", u.Location)
	}

	b.WriteString("Info Layout:\n")
	for _, e := range u.InfoNodes {
		fmt.Fprintf(&b, "\t- %s\n", e.String())
	}
	b.WriteByte('\n')

	return b.String()
}
