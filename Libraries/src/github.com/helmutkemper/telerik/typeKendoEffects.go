package telerik

type KendoEffect int

var kendoEffects = [...]string{
	"",
	"expand:out",
	"expand:in",
	"fade:out",
	"fade:in",
	"flip:out",
	"flip:in",
	"pageturn:out",
	"pageturn:in",
	"replace:out",
	"replace:in",
	"slade:out",
	"slade:in",
	"tile:out",
	"tile:in",
	"transfer:out",
	"transfer:in",
	"zoom:out",
	"zoom:in",
}

func (el KendoEffect) String() string {
	return kendoEffects[el]
}

const (
	EFFECT_EXPAND_OUT KendoEffect = iota + 1
	EFFECT_EXPAND_IN
	EFFECT_FADE_OUT
	EFFECT_FADE_IN
	EFFECT_FLIP_OUT
	EFFECT_FLIP_IN
	EFFECT_PAGE_TURN_OUT
	EFFECT_PAGE_TURN_IN
	EFFECT_REPLACE_OUT
	EFFECT_REPLACE_IN
	EFFECT_SLADE_OUT
	EFFECT_SLADE_IN
	EFFECT_TILE_OUT
	EFFECT_TILE_IN
	EFFECT_TRANSFER_OUT
	EFFECT_TRANSFER_IN
	EFFECT_ZOOM_OUT
	EFFECT_ZOOM_IN
)
