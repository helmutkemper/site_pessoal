package telerik

type KendoOfflineStorage struct {
	GetItem string `jsObject:"getItem"`
	SetItem string `jsObject:"setItem"`
}

func (el KendoOfflineStorage) String() string {
	return `{ getItem: function() { return JSON.parse(sessionStorage.getItem("` + el.GetItem + `")); }, setItem: function(item) { sessionStorage.setItem("` + el.SetItem + `", JSON.stringify(item)); } }`
}
