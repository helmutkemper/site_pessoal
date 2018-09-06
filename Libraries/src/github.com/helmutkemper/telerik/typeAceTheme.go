package telerik

const (
	ACE_THEME_BRIGHT_CHROME AceTheme = iota + 1
	ACE_THEME_BRIGHT_CLOUDS
	ACE_THEME_BRIGHT_CRIMSON_EDITOR
	ACE_THEME_BRIGHT_DAWN
	ACE_THEME_BRIGHT_DREAMWEAVER
	ACE_THEME_BRIGHT_ECLIPSE
	ACE_THEME_BRIGHT_GITHUB
	ACE_THEME_BRIGHT_IPLASTIC
	ACE_THEME_BRIGHT_SOLARIZED_LIGHT
	ACE_THEME_BRIGHT_TEXTMATE
	ACE_THEME_BRIGHT_TOMORROW
	ACE_THEME_BRIGHT_XCODE
	ACE_THEME_BRIGHT_KUROIR
	ACE_THEME_BRIGHT_KATZENMILCH
	ACE_THEME_BRIGHT_SQL_SERVER
	ACE_THEME_DARK_AMBIANCE
	ACE_THEME_DARK_CHAOS
	ACE_THEME_DARK_CLOUDS_MIDNIGHT
	ACE_THEME_DARK_DRACULA
	ACE_THEME_DARK_COBALT
	ACE_THEME_DARK_GRUVBOX
	ACE_THEME_DARK_GREEN_ON_BLACK
	ACE_THEME_DARK_IDLE_FINGERS
	ACE_THEME_DARK_KRTHEME
	ACE_THEME_DARK_MERBIVORE
	ACE_THEME_DARK_MERBIVORE_SOFT
	ACE_THEME_DARK_MONO_INDUSTRIAL
	ACE_THEME_DARK_MONOKAI
	ACE_THEME_DARK_PASTEL_ON_DARK
	ACE_THEME_DARK_SOLARIZED_DARK
	ACE_THEME_DARK_TERMINAL
	ACE_THEME_DARK_OMORROW_NIGHT
	ACE_THEME_DARK_TOMORROW_NIGHT_BLUE
	ACE_THEME_DARK_TOMORROW_NIGHT_BRIGHT
	ACE_THEME_DARK_TOMORROW_NIGHT_80S
	ACE_THEME_DARK_TWILIGHT
	ACE_THEME_DARK_VIBRANT_INK
)

type AceTheme int

func (el AceTheme) String() string {
	return aceThemes[el]
}

var aceThemes = [...]string{
	"",
	"ace/theme/chrome",
	"ace/theme/clouds",
	"ace/theme/crimson_editor",
	"ace/theme/dawn",
	"ace/theme/dreamweaver",
	"ace/theme/eclipse",
	"ace/theme/github",
	"ace/theme/iplastic",
	"ace/theme/solarized_light",
	"ace/theme/textmate",
	"ace/theme/tomorrow",
	"ace/theme/xcode",
	"ace/theme/kuroir",
	"ace/theme/katzenmilch",
	"ace/theme/sqlserver",
	"ace/theme/ambiance",
	"ace/theme/chaos",
	"ace/theme/clouds_midnight",
	"ace/theme/dracula",
	"ace/theme/cobalt",
	"ace/theme/gruvbox",
	"ace/theme/gob",
	"ace/theme/idle_fingers",
	"ace/theme/kr_theme",
	"ace/theme/merbivore",
	"ace/theme/merbivore_soft",
	"ace/theme/mono_industrial",
	"ace/theme/monokai",
	"ace/theme/pastel_on_dark",
	"ace/theme/solarized_dark",
	"ace/theme/terminal",
	"ace/theme/tomorrow_night",
	"ace/theme/tomorrow_night_blue",
	"ace/theme/tomorrow_night_bright",
	"ace/theme/tomorrow_night_eighties",
	"ace/theme/twilight",
	"ace/theme/vibrant_ink",
}
