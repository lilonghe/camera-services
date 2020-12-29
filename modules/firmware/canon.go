package firmware

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"lilonghe.net/camera-services/db"
	"lilonghe.net/camera-services/model"
	"log"
	"strconv"
	"strings"
	"time"
)

type canonSearchItem struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	DetailUrl   string `json:"detail_url"`
	TrackId     string `json:"track_id"`
}

func LoadCanonUpdates() error {
	canonList := make([]model.Camera, 0)
	err := db.Store.Where("company = 'Canon'").Find(&canonList)
	if err != nil {
		return err
	}
	for _, canon := range canonList {
		_, err := db.Store.Where("id = ?", canon.Id).Get(&canon)
		if err != nil {
			return err
		}
		err, updateList := searchMachineByUrl(canon.UpdateUrl)
		if err != nil {
			return err
		}
		for _, info := range updateList {
			err = checkUpdateInfo(info, canon.Id)
			if err != nil {
				return err
			}
		}
		time.Sleep(1 * time.Second)
	}

	return nil
}

func checkUpdateInfo(info canonSearchItem, cameraId int) error {
	var version model.CameraVersion
	has, err := db.Store.Where("track_id = ? and camera_id = ?", info.TrackId, cameraId).Get(&version)
	if err != nil {
		return err
	}
	if !has {
		version.TrackId = info.TrackId
		version.CameraId = cameraId
		version.Description = info.Description
		version.Title = info.Title
		_, err = db.Store.InsertOne(version)
		if err != nil {
			return err
		}
	}
	return nil
}

func searchMachineByUrl(url string) (error, []canonSearchItem) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// run task list
	var nodes []*cdp.Node
	query := `#searchFiles .js_item`
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.Nodes(query, &nodes, chromedp.NodeVisible, chromedp.ByQueryAll, chromedp.AtLeast(0)),
	)
	if err != nil {
		log.Fatal(err)
		return err, nil
	}
	infoList := make([]canonSearchItem, 0)
	for i, node := range nodes {
		infoItem := canonSearchItem{
			DetailUrl: node.AttributeValue("url"),
			TrackId:   getTrackIdByDetailUrl(node.AttributeValue("url")),
		}

		itemQuery := query + `:nth-child(` + strconv.FormatInt(int64(i)+1, 10) + `)`
		if err := chromedp.Run(ctx,
			chromedp.Text(itemQuery+` .js_title`, &infoItem.Title),
			chromedp.InnerHTML(itemQuery+` p`, &infoItem.Description),
		); err != nil {
			log.Fatal(err)
			return err, nil
		} else {
			infoList = append(infoList, infoItem)
		}
	}
	return nil, infoList
}

func getTrackIdByDetailUrl(url string) string {
	// https://www.canon.com.cn/supports/download/simsdetail/0400628605.html
	list := strings.Split(url, "/")
	return strings.TrimRight(list[len(list)-1], ".html")
}
