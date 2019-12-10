package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/rrylee/go-tinyid/constant"
	"github.com/rrylee/go-tinyid/server/factory"
	"github.com/rrylee/go-tinyid/server/service"
	"net/http"
	"strconv"
)

type Response struct {
	Data    interface{}
	Code    int
	Message string
}

func (r Response) String() string {
	s, _ := json.Marshal(r)
	return string(s)
}

func nextIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bizType := r.URL.Query().Get("bizType")
		if bizType == "" {
			response := Response{
				Data:    nil,
				Code:    constant.BizTypeErr.Code,
				Message: constant.BizTypeErr.Message,
			}
			fmt.Fprint(w, response.String())
			return
		}

		defer recoverError(w)

		idGenerator, err := factory.GetIdGenerator(bizType)
		if err != nil {
			handleError(err, w)
			return
		}
		batchSize := getBatchSize(r)
		if batchSize == 1 {
			id, err := idGenerator.NextId()
			if err != nil {
				handleError(err, w)
				return
			}

			response := Response{
				Data: id,
				Code: 200,
			}
			fmt.Fprint(w, response.String())
			return
		} else {
			ids, err := idGenerator.NextBatchIds(batchSize)
			if err != nil {
				handleError(err, w)
				return
			}
			response := Response{
				Data: ids,
				Code: 200,
			}
			fmt.Fprint(w, response.String())
			return
		}
	})
}

func recoverError(w http.ResponseWriter) {
	if err := recover(); err != nil {
		e := err.(error)
		fmt.Fprint(w, e.Error())
		return
	}
}

func nextIdSimpleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bizType := r.URL.Query().Get("bizType")
		if bizType == "" {
			fmt.Fprint(w, "bizType not empty.")
			return
		}
		batchSize := getBatchSize(r)

		defer recoverError(w)

		idGenerator, err := factory.GetIdGenerator(bizType)
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}

		if batchSize == 1 {
			id, err := idGenerator.NextId()
			if err != nil {
				fmt.Fprint(w, err.Error())
				return
			}
			fmt.Fprint(w, id)
			return
		} else {
			ids, err := idGenerator.NextBatchIds(batchSize)
			if err != nil {
				fmt.Fprint(w, err.Error())
				return
			}
			buf := new(bytes.Buffer)
			n := len(ids)
			for i, id := range ids {
				if i == n-1 {
					buf.WriteString(strconv.FormatInt(id, 10))
				} else {
					buf.WriteString(strconv.FormatInt(id, 10) + ",")
				}
			}
			fmt.Fprint(w, buf.String())
			return
		}
	})
}

func getBatchSize(r *http.Request) int64 {
	batchSizeStr := r.URL.Query().Get("batchSize")
	var batchSize int64 = 1
	if batchSizeStr != "" {
		var err error
		batchSize, err = strconv.ParseInt(batchSizeStr, 10, 64)
		if err != nil {
			return 1
		}
		if batchSize < 0 {
			batchSize = 1
		} else if batchSize > constant.MaxBatchSize {
			batchSize = 100000
		}
	}
	return batchSize
}

func segmentIdHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bizType := r.URL.Query().Get("bizType")
		idGenerator, err := factory.GetIdGenerator(bizType)
		if err != nil {
			handleError(err, w)
			return
		}
		current := idGenerator.GetCurrentSegmentId(bizType)
		next := idGenerator.GetNextSegmentId(bizType)
		var curStr, nextStr string
		if current != nil {
			curStr = current.String()
		}
		if next != nil {
			nextStr = next.String()
		}
		mapping := map[string]interface{}{
			"current": curStr,
			"next":    nextStr,
		}
		s, _ := json.Marshal(mapping)
		w.Header().Set("Content-type", "application/json")
		fmt.Fprintln(w, string(s))
	})
}

func nextSegmentIdSimpleHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bizType := r.URL.Query().Get("bizType")
		segmentIDService := &service.DbSegmentIdService{}
		next, err := segmentIDService.GetNextSegmentId(bizType)
		if err != nil {
			handleError(err, w)
			return
		}
		fmt.Fprintln(w, fmt.Sprintf("%d,%d,%d,%d,%d", next.CurrentId(), next.LoadingId(), next.MaxId(), next.Delta(), next.Remainder()))
	})
}

func handleError(err error, w http.ResponseWriter) {
	response := Response{
		Data:    nil,
		Code:    constant.SysErr.Code,
		Message: err.Error(),
	}
	fmt.Fprint(w, response.String())
}
