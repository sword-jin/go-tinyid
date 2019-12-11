package service

import (
	"fmt"
	"github.com/rrylee/go-tinyid/client/config"
	"github.com/rrylee/go-tinyid/constant"
	coreEntity "github.com/rrylee/go-tinyid/core/entity"
	"github.com/rrylee/go-tinyid/core/util"
	"github.com/rrylee/go-tinyid/internal"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type HttpSegmentIdService struct {
	Config *config.Config
}

func (d HttpSegmentIdService) GetNextSegmentId(bizType string) (*coreEntity.SegmentId, error) {
	if d.Config == nil {
		panic("init config first.")
	}
	client := &http.Client{
		Timeout: d.Config.Timeout,
	}

	i := 0
	var err error
	var response *http.Response
	for i = 0; i < constant.MaxRetryFromHTTP; i++ {
		url := d.getNextSegmentIdUrl(bizType)
		response, err = client.Get(url)
		if err != nil {
			internal.Warnf("GetNextSegmentId from http error. err=%v||url=%s||bizType=%s||retry=%d", err, url, bizType, i)
			continue
		}
		break
	}

	if err != nil {
		return nil, fmt.Errorf("GetNextSegmentId from http erro, ret max time. retry=%d", constant.MaxRetryFromHTTP)
	}

	return convertSegmentId(response)
}

func convertSegmentId(response *http.Response) (*coreEntity.SegmentId, error) {
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	arr := strings.Split(string(body), ",")
	currentId, err := parseInt64(arr[0])
	if err != nil {
		return nil, err
	}
	loadingId, err := parseInt64(arr[1])
	if err != nil {
		return nil, err
	}
	maxId, err := parseInt64(arr[2])
	if err != nil {
		return nil, err
	}
	deltaId, err := parseInt64(arr[3])
	if err != nil {
		return nil, err
	}
	remainderId, err := parseInt64(arr[4])
	if err != nil {
		return nil, err
	}
	return coreEntity.NewSegmentId(maxId, currentId, deltaId, remainderId, loadingId), nil
}

func parseInt64(s string) (int64, error) {
	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		return 0, fmt.Errorf("GetNextSegmentId.convertSegmentId.parseInt64 error, err=%v", err)
	}
	return i, nil
}

func (d HttpSegmentIdService) getNextSegmentIdUrl(bizType string) string {
	c := d.Config
	if len(c.TinyIdServer) == 0 {
		return ""
	}
	if len(c.TinyIdServer) == 1 {
		return c.TinyIdServer[0] + "/tinyid/id/nextSegmentIdSimple?token=" + c.TinyIdToken + "&bizType=" + bizType
	}
	serverHost := c.TinyIdServer[util.RandomInt(len(c.TinyIdServer))]
	return serverHost + "/tinyid/id/nextSegmentIdSimple?token=" + c.TinyIdToken + "&bizType=" + bizType
}
