package services

import (
	"net/http"
	"time"

	"github.com/AvaterClasher/zexd/internal/daos"
	"github.com/AvaterClasher/zexd/internal/model"
	"github.com/medama-io/go-useragent"
)

func RecordClick(uid int, r *http.Request) error {
	agent := useragent.NewParser()
	ua := agent.Parse(r.UserAgent())
	deviceType := "Desktop"
	if ua.IsMobile() {
		deviceType = "Mobile"
	} else if ua.IsTablet() {
		deviceType = "Tablet"
	}

	clickData := model.ClickMetadata{
		UrlUid:          uid,
		ClickedAt:       time.Now(),
		IpAddress:       r.RemoteAddr,
		DeviceType:      deviceType,
		OperatingSystem: ua.GetOS(),
		Referrer:        r.Referer(),
		Browser:         ua.GetBrowser(),
	}

	return daos.InsertClickMetadata(clickData)
}
