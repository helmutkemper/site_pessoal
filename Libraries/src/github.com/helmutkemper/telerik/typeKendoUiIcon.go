package telerik

type CssClassIcon int

func (el CssClassIcon) String() string {
	return CssClassIcons[el]
}

var CssClassIcons = [...]string{
	"",
	"k-icon k-i-undo",
	"k-icon k-i-redo",
	"k-icon k-i-reset",
	"k-icon k-i-reload",
	"k-icon k-i-refresh",
	"k-icon k-i-recurrence",
	"k-icon k-i-non-recurrence",
	"k-icon k-i-reset-sm",
	"k-icon k-i-reload-sm",
	"k-icon k-i-refresh-sm",
	"k-icon k-i-recurrence-sm",
	"k-icon k-i-clock",
	"k-icon k-i-calendar",
	"k-icon k-i-save",
	"k-icon k-i-floppy",
	"k-icon k-i-print",
	"k-icon k-i-printer",
	"k-icon k-i-edit",
	"k-icon k-i-pencil",
	"k-icon k-i-delete",
	"k-icon k-i-trash",
	"k-icon k-i-attachment",
	"k-icon k-i-clip",
	"k-icon k-i-attachment-45",
	"k-icon k-i-clip-45",
	"k-icon k-i-link-horizontal",
	"k-icon k-i-hyperlink",
	"k-icon k-i-unlink-horizontal",
	"k-icon k-i-hyperlink-remove",
	"k-icon k-i-link-vertical",
	"k-icon k-i-unlink-vertical",
	"k-icon k-i-lock",
	"k-icon k-i-unlock",
	"k-icon k-i-cancel",
	"k-icon k-i-cancel-outline",
	"k-icon k-i-cancel-circle",
	"k-icon k-i-check",
	"k-icon k-i-checkmark",
	"k-icon k-i-check-outline",
	"k-icon k-i-checkmark-outline",
	"k-icon k-i-success",
	"k-icon k-i-check-circle",
	"k-icon k-i-checkmark-circle",
	"k-icon k-i-close",
	"k-icon k-i-x",
	"k-icon k-i-close-outline",
	"k-icon k-i-x-outline",
	"k-icon k-i-error",
	"k-icon k-i-close-circle",
	"k-icon k-i-x-circle",
	"k-icon k-i-plus",
	"k-icon k-i-plus-outline",
	"k-icon k-i-plus-circle",
	"k-icon k-i-minus",
	"k-icon k-i-kpi-trend-equal",
	"k-icon k-i-minus-outline",
	"k-icon k-i-minus-circle",
	"k-icon k-i-sort-asc",
	"k-icon k-i-sort-desc",
	"k-icon k-i-unsort",
	"k-icon k-i-sort-clear",
	"k-icon k-i-sort-asc-sm",
	"k-icon k-i-sort-desc-sm",
	"k-icon k-i-filter",
	"k-icon k-i-filter-clear",
	"k-icon k-i-filter-sm",
	"k-icon k-i-filter-sort-asc-sm",
	"k-icon k-i-filter-sort-desc-sm",
	"k-icon k-i-filter-add-expression",
	"k-icon k-i-filter-add-group",
	"k-icon k-i-login",
	"k-icon k-i-logout",
	"k-icon k-i-download",
	"k-icon k-i-upload",
	"k-icon k-i-hyperlink-open",
	"k-icon k-i-hyperlink-open-sm",
	"k-icon k-i-launch",
	"k-icon k-i-window",
	"k-icon k-i-window-maximize",
	"k-icon k-i-windows",
	"k-icon k-i-window-restore",
	"k-icon k-i-tiles",
	"k-icon k-i-window-minimize",
	"k-icon k-i-gear",
	"k-icon k-i-cog",
	"k-icon k-i-custom",
	"k-icon k-i-gears",
	"k-icon k-i-cogs",
	"k-icon k-i-wrench",
	"k-icon k-i-settings",
	"k-icon k-i-preview",
	"k-icon k-i-eye",
	"k-icon k-i-zoom",
	"k-icon k-i-search",
	"k-icon k-i-zoom-in",
	"k-icon k-i-zoom-out",
	"k-icon k-i-pan",
	"k-icon k-i-move",
	"k-icon k-i-calculator",
	"k-icon k-i-cart",
	"k-icon k-i-shopping-cart",
	"k-icon k-i-connector",
	"k-icon k-i-plus-sm",
	"k-icon k-i-minus-sm",
	"k-icon k-i-kpi-status-deny",
	"k-icon k-i-kpi-status-hold",
	"k-icon k-i-kpi-status-open",
	"k-icon k-i-notification",
	"k-icon k-i-bell",
	"k-icon k-i-information",
	"k-icon k-i-info",
	"k-icon k-i-question",
	"k-icon k-i-help",
	"k-icon k-i-warning",
	"k-icon k-i-exception",
	"k-icon k-i-page-properties",
	"k-icon k-i-bold",
	"k-icon k-i-italic",
	"k-icon k-i-underline",
	"k-icon k-i-font-family",
	"k-icon k-i-foreground-color",
	"k-icon k-i-convert-lowercase",
	"k-icon k-i-convert-uppercase",
	"k-icon k-i-strikethrough",
	"k-icon k-i-sub-script",
	"k-icon k-i-sup-script",
	"k-icon k-i-div",
	"k-icon k-i-all",
	"k-icon k-i-h1",
	"k-icon k-i-h2",
	"k-icon k-i-h3",
	"k-icon k-i-h4",
	"k-icon k-i-h5",
	"k-icon k-i-h6",
	"k-icon k-i-list-ordered",
	"k-icon k-i-list-numbered",
	"k-icon k-i-list-unordered",
	"k-icon k-i-list-bulleted",
	"k-icon k-i-indent-increase",
	"k-icon k-i-indent",
	"k-icon k-i-indent-decrease",
	"k-icon k-i-outdent",
	"k-icon k-i-insert-up",
	"k-icon k-i-insert-top",
	"k-icon k-i-insert-middle",
	"k-icon k-i-insert-down",
	"k-icon k-i-insert-bottom",
	"k-icon k-i-align-top",
	"k-icon k-i-align-middle",
	"k-icon k-i-align-bottom",
	"k-icon k-i-align-left",
	"k-icon k-i-align-center",
	"k-icon k-i-align-right",
	"k-icon k-i-align-justify",
	"k-icon k-i-align-remove",
	"k-icon k-i-text-wrap",
	"k-icon k-i-rule-horizontal",
	"k-icon k-i-table-align-top-left",
	"k-icon k-i-table-align-top-center",
	"k-icon k-i-table-align-top-right",
	"k-icon k-i-table-align-middle-left",
	"k-icon k-i-table-align-middle-center",
	"k-icon k-i-table-align-middle-right",
	"k-icon k-i-table-align-bottom-left",
	"k-icon k-i-table-align-bottom-center",
	"k-icon k-i-table-align-bottom-right",
	"k-icon k-i-table-align-remove",
	"k-icon k-i-borders-all",
	"k-icon k-i-borders-outside",
	"k-icon k-i-borders-inside",
	"k-icon k-i-borders-inside-horizontal",
	"k-icon k-i-borders-inside-vertical",
	"k-icon k-i-border-top",
	"k-icon k-i-border-bottom",
	"k-icon k-i-border-left",
	"k-icon k-i-border-right",
	"k-icon k-i-border-no",
	"k-icon k-i-borders-show-hide",
	"k-icon k-i-form",
	"k-icon k-i-border",
	"k-icon k-i-form-element",
	"k-icon k-i-code-snippet",
	"k-icon k-i-select-all",
	"k-icon k-i-button",
	"k-icon k-i-select-box",
	"k-icon k-i-calendar-date",
	"k-icon k-i-group-box",
	"k-icon k-i-textarea",
	"k-icon k-i-textbox",
	"k-icon k-i-textbox-hidden",
	"k-icon k-i-password",
	"k-icon k-i-paragraph-add",
	"k-icon k-i-edit-tools",
	"k-icon k-i-template-manager",
	"k-icon k-i-change-manually",
	"k-icon k-i-track-changes",
	"k-icon k-i-track-changes-enable",
	"k-icon k-i-track-changes-accept",
	"k-icon k-i-track-changes-accept-all",
	"k-icon k-i-track-changes-reject",
	"k-icon k-i-track-changes-reject-all",
	"k-icon k-i-document-manager",
	"k-icon k-i-custom-icon",
	"k-icon k-i-dictionary-add",
	"k-icon k-i-image-light-dialog",
	"k-icon k-i-image-insert",
	"k-icon k-i-image-edit",
	"k-icon k-i-image-map-editor",
	"k-icon k-i-comment",
	"k-icon k-i-comment-remove",
	"k-icon k-i-comments-remove-all",
	"k-icon k-i-silverlight",
	"k-icon k-i-media-manager",
	"k-icon k-i-video-external",
	"k-icon k-i-flash-manager",
	"k-icon k-i-find-and-replace",
	"k-icon k-i-find",
	"k-icon k-i-copy",
	"k-icon k-i-files",
	"k-icon k-i-cut",
	"k-icon k-i-paste",
	"k-icon k-i-paste-as-html",
	"k-icon k-i-paste-from-word",
	"k-icon k-i-paste-from-word-strip-file",
	"k-icon k-i-paste-html",
	"k-icon k-i-paste-markdown",
	"k-icon k-i-paste-plain-text",
	"k-icon k-i-apply-format",
	"k-icon k-i-clear-css",
	"k-icon k-i-copy-format",
	"k-icon k-i-strip-all-formating",
	"k-icon k-i-strip-css-format",
	"k-icon k-i-strip-font-elements",
	"k-icon k-i-strip-span-elements",
	"k-icon k-i-strip-word-formatting",
	"k-icon k-i-format-code-block",
	"k-icon k-i-style-builder",
	"k-icon k-i-module-manager",
	"k-icon k-i-hyperlink-light-dialog",
	"k-icon k-i-hyperlink-insert",
	"k-icon k-i-hyperlink-globe",
	"k-icon k-i-hyperlink-globe-remove",
	"k-icon k-i-hyperlink-email",
	"k-icon k-i-anchor",
	"k-icon k-i-table-light-dialog",
	"k-icon k-i-table-insert",
	"k-icon k-i-table",
	"k-icon k-i-table-unmerge",
	"k-icon k-i-table-properties",
	"k-icon k-i-table-cell",
	"k-icon k-i-table-cell-properties",
	"k-icon k-i-table-column-insert-left",
	"k-icon k-i-table-column-insert-right",
	"k-icon k-i-table-row-insert-above",
	"k-icon k-i-table-row-insert-below",
	"k-icon k-i-table-column-delete",
	"k-icon k-i-table-row-delete",
	"k-icon k-i-table-cell-delete",
	"k-icon k-i-table-delete",
	"k-icon k-i-cells-merge",
	"k-icon k-i-cells-merge-horizontally",
	"k-icon k-i-cells-merge-vertically",
	"k-icon k-i-cell-split-horizontally",
	"k-icon k-i-cell-split-vertically",
	"k-icon k-i-pane-freeze",
	"k-icon k-i-row-freeze",
	"k-icon k-i-column-freeze",
	"k-icon k-i-toolbar-float",
	"k-icon k-i-spell-checker",
	"k-icon k-i-validation-xhtml",
	"k-icon k-i-validation-data",
	"k-icon k-i-toggle-full-screen-mode",
	"k-icon k-i-formula-fx",
	"k-icon k-i-sum",
	"k-icon k-i-symbol",
	"k-icon k-i-dollar",
	"k-icon k-i-currency",
	"k-icon k-i-percent",
	"k-icon k-i-custom-format",
	"k-icon k-i-decimal-increase",
	"k-icon k-i-decimal-decrease",
	"k-icon k-i-font-size",
	"k-icon k-i-image-absolute-position",
	"k-icon k-i-folder",
	"k-icon k-i-folder-open",
	"k-icon k-i-folder-add",
	"k-icon k-i-folder-up",
	"k-icon k-i-folder-more",
	"k-icon k-i-fields-more",
	"k-icon k-i-aggregate-fields",
	"k-icon k-i-file",
	"k-icon k-i-file-vertical",
	"k-icon k-i-file-add",
	"k-icon k-i-file-txt",
	"k-icon k-i-txt",
	"k-icon k-i-file-csv",
	"k-icon k-i-csv",
	"k-icon k-i-file-excel",
	"k-icon k-i-file-xls",
	"k-icon k-i-excel",
	"k-icon k-i-xls",
	"k-icon k-i-file-word",
	"k-icon k-i-file-doc",
	"k-icon k-i-word",
	"k-icon k-i-doc",
	"k-icon k-i-file-mdb",
	"k-icon k-i-mdb",
	"k-icon k-i-file-ppt",
	"k-icon k-i-ppt",
	"k-icon k-i-file-pdf",
	"k-icon k-i-pdf",
	"k-icon k-i-file-psd",
	"k-icon k-i-psd",
	"k-icon k-i-file-flash",
	"k-icon k-i-flash",
	"k-icon k-i-file-config",
	"k-icon k-i-config",
	"k-icon k-i-file-ascx",
	"k-icon k-i-ascx",
	"k-icon k-i-file-bac",
	"k-icon k-i-bac",
	"k-icon k-i-file-zip",
	"k-icon k-i-zip",
	"k-icon k-i-film",
	"k-icon k-i-css3",
	"k-icon k-i-html5",
	"k-icon k-i-html",
	"k-icon k-i-source-code",
	"k-icon k-i-view-source",
	"k-icon k-i-css",
	"k-icon k-i-js",
	"k-icon k-i-exe",
	"k-icon k-i-csproj",
	"k-icon k-i-vbproj",
	"k-icon k-i-cs",
	"k-icon k-i-vb",
	"k-icon k-i-sln",
	"k-icon k-i-cloud",
	"k-icon k-i-file-horizontal",
	"k-icon k-i-photo-camera",
	"k-icon k-i-image",
	"k-icon k-i-photo",
	"k-icon k-i-image-export",
	"k-icon k-i-photo-export",
	"k-icon k-i-zoom-actual-size",
	"k-icon k-i-zoom-best-fit",
	"k-icon k-i-image-resize",
	"k-icon k-i-crop",
	"k-icon k-i-mirror",
	"k-icon k-i-flip-horizontal",
	"k-icon k-i-flip-vertical",
	"k-icon k-i-rotate",
	"k-icon k-i-rotate-right",
	"k-icon k-i-rotate-left",
	"k-icon k-i-brush",
	"k-icon k-i-palette",
	"k-icon k-i-paint",
	"k-icon k-i-droplet",
	"k-icon k-i-background",
	"k-icon k-i-line",
	"k-icon k-i-shape-line",
	"k-icon k-i-brightness-contrast",
	"k-icon k-i-saturation",
	"k-icon k-i-invert-colors",
	"k-icon k-i-transperancy",
	"k-icon k-i-opacity",
	"k-icon k-i-greyscale",
	"k-icon k-i-blur",
	"k-icon k-i-sharpen",
	"k-icon k-i-shape",
	"k-icon k-i-round-corners",
	"k-icon k-i-front-element",
	"k-icon k-i-back-element",
	"k-icon k-i-forward-element",
	"k-icon k-i-backward-element",
	"k-icon k-i-align-left-element",
	"k-icon k-i-align-center-element",
	"k-icon k-i-align-right-element",
	"k-icon k-i-align-top-element",
	"k-icon k-i-align-middle-element",
	"k-icon k-i-align-bottom-element",
	"k-icon k-i-thumbnails-up",
	"k-icon k-i-thumbnails-right",
	"k-icon k-i-thumbnails-down",
	"k-icon k-i-thumbnails-left",
	"k-icon k-i-full-screen",
	"k-icon k-i-fullscreen",
	"k-icon k-i-full-screen-exit",
	"k-icon k-i-fullscreen-exit",
	"k-icon k-i-reset-color",
	"k-icon k-i-paint-remove",
	"k-icon k-i-background-remove",
	"k-icon k-i-arrow-45-up-right",
	"k-icon k-i-collapse-ne",
	"k-icon k-i-resize-ne",
	"k-icon k-i-arrow-45-down-right",
	"k-icon k-i-collapse-se",
	"k-icon k-i-resize-se",
	"k-icon k-i-arrow-45-down-left",
	"k-icon k-i-collapse-sw",
	"k-icon k-i-resize-sw",
	"k-icon k-i-arrow-45-up-left",
	"k-icon k-i-collapse-nw",
	"k-icon k-i-resize-new",
	"k-icon k-i-arrow-60-up",
	"k-icon k-i-kpi-trend-increase",
	"k-icon k-i-arrow-60-right",
	"k-icon k-i-arrow-60-down",
	"k-icon k-i-kpi-trend-decrease",
	"k-icon k-i-arrow-60-left",
	"k-icon k-i-arrow-end-up",
	"k-icon k-i-arrow-end-right",
	"k-icon k-i-arrow-end-down",
	"k-icon k-i-arrow-end-left",
	"k-icon k-i-arrow-double-60-up",
	"k-icon k-i-arrow-seek-up",
	"k-icon k-i-arrow-double-60-right",
	"k-icon k-i-forward-sm",
	"k-icon k-i-arrow-seek-right",
	"k-icon k-i-arrow-double-60-down",
	"k-icon k-i-arrow-seek-down",
	"k-icon k-i-arrow-double-60-left",
	"k-icon k-i-rewind-sm",
	"k-icon k-i-arrows-kpi",
	"k-icon k-i-kpi",
	"k-icon k-i-arrows-no-change",
	"k-icon k-i-arrow-overflow-down",
	"k-icon k-i-arrow-chevron-up",
	"k-icon k-i-arrow-chevron-right",
	"k-icon k-i-arrow-chevron-down",
	"k-icon k-i-arrow-chevron-left",
	"k-icon k-i-arrow-up",
	"k-icon k-i-arrow-right",
	"k-icon k-i-arrow-down",
	"k-icon k-i-arrow-left",
	"k-icon k-i-arrow-drill",
	"k-icon k-i-arrow-parent",
	"k-icon k-i-arrow-root",
	"k-icon k-i-arrows-resizing",
	"k-icon k-i-arrows-dimensions",
	"k-icon k-i-arrows-swap",
	"k-icon k-i-drag-and-drop",
	"k-icon k-i-categorize",
	"k-icon k-i-grid",
	"k-icon k-i-grid-layout",
	"k-icon k-i-group",
	"k-icon k-i-ungroup",
	"k-icon k-i-handler-drag",
	"k-icon k-i-layout",
	"k-icon k-i-layout-1-by-4",
	"k-icon k-i-layout-2-by-2",
	"k-icon k-i-layout-side-by-side",
	"k-icon k-i-layout-stacked",
	"k-icon k-i-columns",
	"k-icon k-i-rows",
	"k-icon k-i-reorder",
	"k-icon k-i-menu",
	"k-icon k-i-more-vertical",
	"k-icon k-i-more-horizontal",
	"k-icon k-i-globe-outline",
	"k-icon k-i-globe",
	"k-icon k-i-marker-pin",
	"k-icon k-i-marker-pin-target",
	"k-icon k-i-pin",
	"k-icon k-i-unpin",
	"k-icon k-i-play",
	"k-icon k-i-pause",
	"k-icon k-i-stop",
	"k-icon k-i-rewind",
	"k-icon k-i-forward",
	"k-icon k-i-volume-down",
	"k-icon k-i-volume-up",
	"k-icon k-i-volume-off",
	"k-icon k-i-hd",
	"k-icon k-i-subtitles",
	"k-icon k-i-playlist",
	"k-icon k-i-audio",
	"k-icon k-i-share",
	"k-icon k-i-user",
	"k-icon k-i-inbox",
	"k-icon k-i-blogger",
	"k-icon k-i-blogger-box",
	"k-icon k-i-delicious",
	"k-icon k-i-delicious-box",
	"k-icon k-i-digg",
	"k-icon k-i-digg-box",
	"k-icon k-i-email",
	"k-icon k-i-envelop",
	"k-icon k-i-letter",
	"k-icon k-i-email-box",
	"k-icon k-i-envelop-box",
	"k-icon k-i-letter-box",
	"k-icon k-i-facebook",
	"k-icon k-i-facebook-box",
	"k-icon k-i-google",
	"k-icon k-i-google-box",
	"k-icon k-i-google-plus",
	"k-icon k-i-google-plus-box",
	"k-icon k-i-linkedin",
	"k-icon k-i-linkedin-box",
	"k-icon k-i-myspace",
	"k-icon k-i-myspace-box",
	"k-icon k-i-pinterest",
	"k-icon k-i-pinterest-box",
	"k-icon k-i-reddit",
	"k-icon k-i-reddit-box",
	"k-icon k-i-stumble-upon",
	"k-icon k-i-stumble-upon-box",
	"k-icon k-i-tell-a-friend",
	"k-icon k-i-tell-a-friend-box",
	"k-icon k-i-tumblr",
	"k-icon k-i-tumblr-box",
	"k-icon k-i-twitter",
	"k-icon k-i-twitter-box",
	"k-icon k-i-yammer",
	"k-icon k-i-yammer-box",
	"k-icon k-i-behance",
	"k-icon k-i-behance-box",
	"k-icon k-i-dribbble",
	"k-icon k-i-dribbble-box",
	"k-icon k-i-rss",
	"k-icon k-i-rss-box",
	"k-icon k-i-vimeo",
	"k-icon k-i-vimeo-box",
	"k-icon k-i-youtube",
	"k-icon k-i-youtube-box",
	"k-icon k-i-heart-outline",
	"k-icon k-i-fav-outline",
	"k-icon k-i-favorite-outline",
	"k-icon k-i-heart",
	"k-icon k-i-fav",
	"k-icon k-i-favorite",
	"k-icon k-i-star-outline",
	"k-icon k-i-bookmark-outline",
	"k-icon k-i-star",
	"k-icon k-i-bookmark",
	"k-icon k-i-checkbox",
	"k-icon k-i-shape-rect",
	"k-icon k-i-checkbox-checked",
	"k-icon k-i-tri-state-indeterminate",
	"k-icon k-i-tri-state-null",
	"k-icon k-i-circle",
	"k-icon k-i-radiobutton",
	"k-icon k-i-shape-circle",
	"k-icon k-i-radiobutton-checked",
}

const (
	ICON_UNDO CssClassIcon = iota + 1
	ICON_REDO
	ICON_RESET
	ICON_RELOAD
	ICON_REFRESH
	ICON_RECURRENCE
	ICON_NON_RECURRENCE
	ICON_RESET_SM
	ICON_RELOAD_SM
	ICON_REFRESH_SM
	ICON_RECURRENCE_SM
	ICON_CLOCK
	ICON_CALENDAR
	ICON_SAVE
	ICON_FLOPPY
	ICON_PRINT
	ICON_PRINTER
	ICON_EDIT
	ICON_PENCIL
	ICON_DELETE
	ICON_TRASH
	ICON_ATTACHMENT
	ICON_CLIP
	ICON_ATTACHMENT_45
	ICON_CLIP_45
	ICON_LINK_HORIZONTAL
	ICON_HYPERLINK
	ICON_UNLINK_HORIZONTAL
	ICON_HYPERLINK_REMOVE
	ICON_LINK_VERTICAL
	ICON_UNLINK_VERTICAL
	ICON_LOCK
	ICON_UNLOCK
	ICON_CANCEL
	ICON_CANCEL_OUTLINE
	ICON_CANCEL_CIRCLE
	ICON_CHECK
	ICON_CHECKMARK
	ICON_CHECK_OUTLINE
	ICON_CHECKMARK_OUTLINE
	ICON_SUCCESS
	ICON_CHECK_CIRCLE
	ICON_CHECKMARK_CIRCLE
	ICON_CLOSE
	ICON_X
	ICON_CLOSE_OUTLINE
	ICON_X_OUTLINE
	ICON_ERROR
	ICON_CLOSE_CIRCLE
	ICON_X_CIRCLE
	ICON_PLUS
	ICON_PLUS_OUTLINE
	ICON_PLUS_CIRCLE
	ICON_MINUS
	ICON_KPI_TREND_EQUAL
	ICON_MINUS_OUTLINE
	ICON_MINUS_CIRCLE
	ICON_SORT_ASC
	ICON_SORT_DESC
	ICON_UNSORT
	ICON_SORT_CLEAR
	ICON_SORT_ASC_SM
	ICON_SORT_DESC_SM
	ICON_FILTER
	ICON_FILTER_CLEAR
	ICON_FILTER_SM
	ICON_FILTER_SORT_ASC_SM
	ICON_FILTER_SORT_DESC_SM
	ICON_FILTER_ADD_EXPRESSION
	ICON_FILTER_ADD_GROUP
	ICON_LOGIN
	ICON_LOGOUT
	ICON_DOWNLOAD
	ICON_UPLOAD
	ICON_HYPERLINK_OPEN
	ICON_HYPERLINK_OPEN_SM
	ICON_LAUNCH
	ICON_WINDOW
	ICON_WINDOW_MAXIMIZE
	ICON_WINDOWS
	ICON_WINDOW_RESTORE
	ICON_TILES
	ICON_WINDOW_MINIMIZE
	ICON_GEAR
	ICON_COG
	ICON_CUSTOM
	ICON_GEARS
	ICON_COGS
	ICON_WRENCH
	ICON_SETTINGS
	ICON_PREVIEW
	ICON_EYE
	ICON_ZOOM
	ICON_SEARCH
	ICON_ZOOM_IN
	ICON_ZOOM_OUT
	ICON_PAN
	ICON_MOVE
	ICON_CALCULATOR
	ICON_CART
	ICON_SHOPPING_CART
	ICON_CONNECTOR
	ICON_PLUS_SM
	ICON_MINUS_SM
	ICON_KPI_STATUS_DENY
	ICON_KPI_STATUS_HOLD
	ICON_KPI_STATUS_OPEN
	ICON_NOTIFICATION
	ICON_BELL
	ICON_INFORMATION
	ICON_INFO
	ICON_QUESTION
	ICON_HELP
	ICON_WARNING
	ICON_EXCEPTION
	ICON_PAGE_PROPERTIES
	ICON_BOLD
	ICON_ITALIC
	ICON_UNDERLINE
	ICON_FONT_FAMILY
	ICON_FOREGROUND_COLOR
	ICON_CONVERT_LOWERCASE
	ICON_CONVERT_UPPERCASE
	ICON_STRIKETHROUGH
	ICON_SUB_SCRIPT
	ICON_SUP_SCRIPT
	ICON_DIV
	ICON_ALL
	ICON_H1
	ICON_H2
	ICON_H3
	ICON_H4
	ICON_H5
	ICON_H6
	ICON_LIST_ORDERED
	ICON_LIST_NUMBERED
	ICON_LIST_UNORDERED
	ICON_LIST_BULLETED
	ICON_INDENT_INCREASE
	ICON_INDENT
	ICON_INDENT_DECREASE
	ICON_OUTDENT
	ICON_INSERT_UP
	ICON_INSERT_TOP
	ICON_INSERT_MIDDLE
	ICON_INSERT_DOWN
	ICON_INSERT_BOTTOM
	ICON_ALIGN_TOP
	ICON_ALIGN_MIDDLE
	ICON_ALIGN_BOTTOM
	ICON_ALIGN_LEFT
	ICON_ALIGN_CENTER
	ICON_ALIGN_RIGHT
	ICON_ALIGN_JUSTIFY
	ICON_ALIGN_REMOVE
	ICON_TEXT_WRAP
	ICON_RULE_HORIZONTAL
	ICON_TABLE_ALIGN_TOP_LEFT
	ICON_TABLE_ALIGN_TOP_CENTER
	ICON_TABLE_ALIGN_TOP_RIGHT
	ICON_TABLE_ALIGN_MIDDLE_LEFT
	ICON_TABLE_ALIGN_MIDDLE_CENTER
	ICON_TABLE_ALIGN_MIDDLE_RIGHT
	ICON_TABLE_ALIGN_BOTTOM_LEFT
	ICON_TABLE_ALIGN_BOTTOM_CENTER
	ICON_TABLE_ALIGN_BOTTOM_RIGHT
	ICON_TABLE_ALIGN_REMOVE
	ICON_BORDERS_ALL
	ICON_BORDERS_OUTSIDE
	ICON_BORDERS_INSIDE
	ICON_BORDERS_INSIDE_HORIZONTAL
	ICON_BORDERS_INSIDE_VERTICAL
	ICON_BORDER_TOP
	ICON_BORDER_BOTTOM
	ICON_BORDER_LEFT
	ICON_BORDER_RIGHT
	ICON_BORDER_NO
	ICON_BORDERS_SHOW_HIDE
	ICON_FORM
	ICON_BORDER
	ICON_FORM_ELEMENT
	ICON_CODE_SNIPPET
	ICON_SELECT_ALL
	ICON_BUTTON
	ICON_SELECT_BOX
	ICON_CALENDAR_DATE
	ICON_GROUP_BOX
	ICON_TEXTAREA
	ICON_TEXTBOX
	ICON_TEXTBOX_HIDDEN
	ICON_PASSWORD
	ICON_PARAGRAPH_ADD
	ICON_EDIT_TOOLS
	ICON_TEMPLATE_MANAGER
	ICON_CHANGE_MANUALLY
	ICON_TRACK_CHANGES
	ICON_TRACK_CHANGES_ENABLE
	ICON_TRACK_CHANGES_ACCEPT
	ICON_TRACK_CHANGES_ACCEPT_ALL
	ICON_TRACK_CHANGES_REJECT
	ICON_TRACK_CHANGES_REJECT_ALL
	ICON_DOCUMENT_MANAGER
	ICON_CUSTOM_ICON
	ICON_DICTIONARY_ADD
	ICON_IMAGE_LIGHT_DIALOG
	ICON_IMAGE_INSERT
	ICON_IMAGE_EDIT
	ICON_IMAGE_MAP_EDITOR
	ICON_COMMENT
	ICON_COMMENT_REMOVE
	ICON_COMMENTS_REMOVE_ALL
	ICON_SILVERLIGHT
	ICON_MEDIA_MANAGER
	ICON_VIDEO_EXTERNAL
	ICON_FLASH_MANAGER
	ICON_FIND_AND_REPLACE
	ICON_FIND
	ICON_COPY
	ICON_FILES
	ICON_CUT
	ICON_PASTE
	ICON_PASTE_AS_HTML
	ICON_PASTE_FROM_WORD
	ICON_PASTE_FROM_WORD_STRIP_FILE
	ICON_PASTE_HTML
	ICON_PASTE_MARKDOWN
	ICON_PASTE_PLAIN_TEXT
	ICON_APPLY_FORMAT
	ICON_CLEAR_CSS
	ICON_COPY_FORMAT
	ICON_STRIP_ALL_FORMATING
	ICON_STRIP_CSS_FORMAT
	ICON_STRIP_FONT_ELEMENTS
	ICON_STRIP_SPAN_ELEMENTS
	ICON_STRIP_WORD_FORMATTING
	ICON_FORMAT_CODE_BLOCK
	ICON_STYLE_BUILDER
	ICON_MODULE_MANAGER
	ICON_HYPERLINK_LIGHT_DIALOG
	ICON_HYPERLINK_INSERT
	ICON_HYPERLINK_GLOBE
	ICON_HYPERLINK_GLOBE_REMOVE
	ICON_HYPERLINK_EMAIL
	ICON_ANCHOR
	ICON_TABLE_LIGHT_DIALOG
	ICON_TABLE_INSERT
	ICON_TABLE
	ICON_TABLE_UNMERGE
	ICON_TABLE_PROPERTIES
	ICON_TABLE_CELL
	ICON_TABLE_CELL_PROPERTIES
	ICON_TABLE_COLUMN_INSERT_LEFT
	ICON_TABLE_COLUMN_INSERT_RIGHT
	ICON_TABLE_ROW_INSERT_ABOVE
	ICON_TABLE_ROW_INSERT_BELOW
	ICON_TABLE_COLUMN_DELETE
	ICON_TABLE_ROW_DELETE
	ICON_TABLE_CELL_DELETE
	ICON_TABLE_DELETE
	ICON_CELLS_MERGE
	ICON_CELLS_MERGE_HORIZONTALLY
	ICON_CELLS_MERGE_VERTICALLY
	ICON_CELL_SPLIT_HORIZONTALLY
	ICON_CELL_SPLIT_VERTICALLY
	ICON_PANE_FREEZE
	ICON_ROW_FREEZE
	ICON_COLUMN_FREEZE
	ICON_TOOLBAR_FLOAT
	ICON_SPELL_CHECKER
	ICON_VALIDATION_XHTML
	ICON_VALIDATION_DATA
	ICON_TOGGLE_FULL_SCREEN_MODE
	ICON_FORMULA_FX
	ICON_SUM
	ICON_SYMBOL
	ICON_DOLLAR
	ICON_CURRENCY
	ICON_PERCENT
	ICON_CUSTOM_FORMAT
	ICON_DECIMAL_INCREASE
	ICON_DECIMAL_DECREASE
	ICON_FONT_SIZE
	ICON_IMAGE_ABSOLUTE_POSITION
	ICON_FOLDER
	ICON_FOLDER_OPEN
	ICON_FOLDER_ADD
	ICON_FOLDER_UP
	ICON_FOLDER_MORE
	ICON_FIELDS_MORE
	ICON_AGGREGATE_FIELDS
	ICON_FILE
	ICON_FILE_VERTICAL
	ICON_FILE_ADD
	ICON_FILE_TXT
	ICON_TXT
	ICON_FILE_CSV
	ICON_CSV
	ICON_FILE_EXCEL
	ICON_FILE_XLS
	ICON_EXCEL
	ICON_XLS
	ICON_FILE_WORD
	ICON_FILE_DOC
	ICON_WORD
	ICON_DOC
	ICON_FILE_MDB
	ICON_MDB
	ICON_FILE_PPT
	ICON_PPT
	ICON_FILE_PDF
	ICON_PDF
	ICON_FILE_PSD
	ICON_PSD
	ICON_FILE_FLASH
	ICON_FLASH
	ICON_FILE_CONFIG
	ICON_CONFIG
	ICON_FILE_ASCX
	ICON_ASCX
	ICON_FILE_BAC
	ICON_BAC
	ICON_FILE_ZIP
	ICON_ZIP
	ICON_FILM
	ICON_CSS3
	ICON_HTML5
	ICON_HTML
	ICON_SOURCE_CODE
	ICON_VIEW_SOURCE
	ICON_CSS
	ICON_JS
	ICON_EXE
	ICON_CSPROJ
	ICON_VBPROJ
	ICON_CS
	ICON_VB
	ICON_SLN
	ICON_CLOUD
	ICON_FILE_HORIZONTAL
	ICON_PHOTO_CAMERA
	ICON_IMAGE
	ICON_PHOTO
	ICON_IMAGE_EXPORT
	ICON_PHOTO_EXPORT
	ICON_ZOOM_ACTUAL_SIZE
	ICON_ZOOM_BEST_FIT
	ICON_IMAGE_RESIZE
	ICON_CROP
	ICON_MIRROR
	ICON_FLIP_HORIZONTAL
	ICON_FLIP_VERTICAL
	ICON_ROTATE
	ICON_ROTATE_RIGHT
	ICON_ROTATE_LEFT
	ICON_BRUSH
	ICON_PALETTE
	ICON_PAINT
	ICON_DROPLET
	ICON_BACKGROUND
	ICON_LINE
	ICON_SHAPE_LINE
	ICON_BRIGHTNESS_CONTRAST
	ICON_SATURATION
	ICON_INVERT_COLORS
	ICON_TRANSPERANCY
	ICON_OPACITY
	ICON_GREYSCALE
	ICON_BLUR
	ICON_SHARPEN
	ICON_SHAPE
	ICON_ROUND_CORNERS
	ICON_FRONT_ELEMENT
	ICON_BACK_ELEMENT
	ICON_FORWARD_ELEMENT
	ICON_BACKWARD_ELEMENT
	ICON_ALIGN_LEFT_ELEMENT
	ICON_ALIGN_CENTER_ELEMENT
	ICON_ALIGN_RIGHT_ELEMENT
	ICON_ALIGN_TOP_ELEMENT
	ICON_ALIGN_MIDDLE_ELEMENT
	ICON_ALIGN_BOTTOM_ELEMENT
	ICON_THUMBNAILS_UP
	ICON_THUMBNAILS_RIGHT
	ICON_THUMBNAILS_DOWN
	ICON_THUMBNAILS_LEFT
	ICON_FULL_SCREEN
	ICON_FULLSCREEN
	ICON_FULL_SCREEN_EXIT
	ICON_FULLSCREEN_EXIT
	ICON_RESET_COLOR
	ICON_PAINT_REMOVE
	ICON_BACKGROUND_REMOVE
	ICON_ARROW_45_UP_RIGHT
	ICON_COLLAPSE_NE
	ICON_RESIZE_NE
	ICON_ARROW_45_DOWN_RIGHT
	ICON_COLLAPSE_SE
	ICON_RESIZE_SE
	ICON_ARROW_45_DOWN_LEFT
	ICON_COLLAPSE_SW
	ICON_RESIZE_SW
	ICON_ARROW_45_UP_LEFT
	ICON_COLLAPSE_NW
	ICON_RESIZE_NEW
	ICON_ARROW_60_UP
	ICON_KPI_TREND_INCREASE
	ICON_ARROW_60_RIGHT
	ICON_ARROW_60_DOWN
	ICON_KPI_TREND_DECREASE
	ICON_ARROW_60_LEFT
	ICON_ARROW_END_UP
	ICON_ARROW_END_RIGHT
	ICON_ARROW_END_DOWN
	ICON_ARROW_END_LEFT
	ICON_ARROW_DOUBLE_60_UP
	ICON_ARROW_SEEK_UP
	ICON_ARROW_DOUBLE_60_RIGHT
	ICON_FORWARD_SM
	ICON_ARROW_SEEK_RIGHT
	ICON_ARROW_DOUBLE_60_DOWN
	ICON_ARROW_SEEK_DOWN
	ICON_ARROW_DOUBLE_60_LEFT
	ICON_REWIND_SM
	ICON_ARROWS_KPI
	ICON_KPI
	ICON_ARROWS_NO_CHANGE
	ICON_ARROW_OVERFLOW_DOWN
	ICON_ARROW_CHEVRON_UP
	ICON_ARROW_CHEVRON_RIGHT
	ICON_ARROW_CHEVRON_DOWN
	ICON_ARROW_CHEVRON_LEFT
	ICON_ARROW_UP
	ICON_ARROW_RIGHT
	ICON_ARROW_DOWN
	ICON_ARROW_LEFT
	ICON_ARROW_DRILL
	ICON_ARROW_PARENT
	ICON_ARROW_ROOT
	ICON_ARROWS_RESIZING
	ICON_ARROWS_DIMENSIONS
	ICON_ARROWS_SWAP
	ICON_DRAG_AND_DROP
	ICON_CATEGORIZE
	ICON_GRID
	ICON_GRID_LAYOUT
	ICON_GROUP
	ICON_UNGROUP
	ICON_HANDLER_DRAG
	ICON_LAYOUT
	ICON_LAYOUT_1_BY_4
	ICON_LAYOUT_2_BY_2
	ICON_LAYOUT_SIDE_BY_SIDE
	ICON_LAYOUT_STACKED
	ICON_COLUMNS
	ICON_ROWS
	ICON_REORDER
	ICON_MENU
	ICON_MORE_VERTICAL
	ICON_MORE_HORIZONTAL
	ICON_GLOBE_OUTLINE
	ICON_GLOBE
	ICON_MARKER_PIN
	ICON_MARKER_PIN_TARGET
	ICON_PIN
	ICON_UNPIN
	ICON_PLAY
	ICON_PAUSE
	ICON_STOP
	ICON_REWIND
	ICON_FORWARD
	ICON_VOLUME_DOWN
	ICON_VOLUME_UP
	ICON_VOLUME_OFF
	ICON_HD
	ICON_SUBTITLES
	ICON_PLAYLIST
	ICON_AUDIO
	ICON_SHARE
	ICON_USER
	ICON_INBOX
	ICON_BLOGGER
	ICON_BLOGGER_BOX
	ICON_DELICIOUS
	ICON_DELICIOUS_BOX
	ICON_DIGG
	ICON_DIGG_BOX
	ICON_EMAIL
	ICON_ENVELOP
	ICON_LETTER
	ICON_EMAIL_BOX
	ICON_ENVELOP_BOX
	ICON_LETTER_BOX
	ICON_FACEBOOK
	ICON_FACEBOOK_BOX
	ICON_GOOGLE
	ICON_GOOGLE_BOX
	ICON_GOOGLE_PLUS
	ICON_GOOGLE_PLUS_BOX
	ICON_LINKEDIN
	ICON_LINKEDIN_BOX
	ICON_MYSPACE
	ICON_MYSPACE_BOX
	ICON_PINTEREST
	ICON_PINTEREST_BOX
	ICON_REDDIT
	ICON_REDDIT_BOX
	ICON_STUMBLE_UPON
	ICON_STUMBLE_UPON_BOX
	ICON_TELL_A_FRIEND
	ICON_TELL_A_FRIEND_BOX
	ICON_TUMBLR
	ICON_TUMBLR_BOX
	ICON_TWITTER
	ICON_TWITTER_BOX
	ICON_YAMMER
	ICON_YAMMER_BOX
	ICON_BEHANCE
	ICON_BEHANCE_BOX
	ICON_DRIBBBLE
	ICON_DRIBBBLE_BOX
	ICON_RSS
	ICON_RSS_BOX
	ICON_VIMEO
	ICON_VIMEO_BOX
	ICON_YOUTUBE
	ICON_YOUTUBE_BOX
	ICON_HEART_OUTLINE
	ICON_FAV_OUTLINE
	ICON_FAVORITE_OUTLINE
	ICON_HEART
	ICON_FAV
	ICON_FAVORITE
	ICON_STAR_OUTLINE
	ICON_BOOKMARK_OUTLINE
	ICON_STAR
	ICON_BOOKMARK
	ICON_CHECKBOX
	ICON_SHAPE_RECT
	ICON_CHECKBOX_CHECKED
	ICON_TRI_STATE_INDETERMINATE
	ICON_TRI_STATE_NULL
	ICON_CIRCLE
	ICON_RADIOBUTTON
	ICON_SHAPE_CIRCLE
	ICON_RADIOBUTTON_CHECKED
)
