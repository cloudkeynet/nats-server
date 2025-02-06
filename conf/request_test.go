package httpsuite

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/nats-io/nats-server/v2/server"
	"go.uber.org/automaxprocs/maxprocs"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

type TestRequest struct {
	ID int `json:"id" validate:"required"`
	Name string `json:"name" valdiate:"required"`
}

func (r *TestRequest) SetParam(fieldName, value string) error {
	switch strings.ToLower(fieldName){
	case "id":
		id, err := strconv.Atoi(value)
		if err != nil {
			return errors.New("Invalid ID")
		}
		r.ID = id
	default:
	log.Printf("Parameter %s cannot be set", fieldName)
	}
	return nil
}
func Test_ParseRequest(t *testing.T){
	testSetURLParam := func(r *http.Request. fieldName, value string) *http.Request {
		ctx := chi.NewRouteContext()
		ctx.URLParams.Add(fieldName, value)
		return r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, ctx))
	}

	type args struct {
		w	http.ResponseWriter
		r   *http.Request
		pathParams []string
	}
	test := []testCase[TestRequest]{
		name: "Successfu Request",
		args: args{
			w: httptest.NewRecorder(),
			r: func() *http.Request{
				body, _ := json.Marshal(TestRequest{Name: "Test"})
				req := httptest.NewRequest("POST", "/test/123", bytes.NewBuffer(body))
				req = testSetURLParam(req, "ID", "123")
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			pathParams: []string{"ID"},
		},
		want: &TestRequest{ID: 123, Name: "Test"},
		wantErr: assert.NoError,
	},
	{
		name: "Misssing body",
		args: args{
			w: httptest.NewRecorder(),
			r: func() *http.Request {
				req := httptest.NewRequest("POST", "/test/123", nil)
				req = testSetURLParam(req, "ID", "123")
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			pathParams: []string{"ID"},
		},
		want:    nil,
		wantErr: assert.Error,
	},
	{
		name: "Missing Path Parameter",
		args: args{
			w: httptest.NewRecorder(),
			r: func() *http.Request {
				req := httptest.NewRequest("POST", "/test", nil)
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			pathParams: []string{"ID"},
		},
		want:    nil,
		wantErr: assert.Error,
	},
	{
		name: "Invalid JSON",
		args: args{
			w: httptest.NewRecorder(),
			r: func() *http.Request {
				req := httptest.NewRequest("POST", "/test/123", bytes.NewBufferString("{invalid-json}"))
				req = testSetURLParam(req, "ID", "123")
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			pathParams: []string{"ID"},
		},
		want:    nil,
		wantErr: assert.Error,
	},
	{
		name: "Validation Error for body",
		args: args{
			w: httptest.NewRecorder(),
			r: func() *http.Request {
				body, _ := json.Marshal(TestRequest{})
				req := httptest.NewRequest("POST", "/test/123", bytes.NewBuffer(body))
				req = testSetURLParam(req, "ID", "0")
				req.Header.Set("Content-Type", "application/json")
				return req
			}(),
			want: nil,
			wantErr: assert.Error,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRequest[*TestRequest](tt.args.w, tt.args.r, tt.args.pathParams...)
			if !tt,wantErr(t, err, fmt.Sprintf("parseRequest(%v, %v, %v)", tt.args.w, tt.args.r, tt.args.pathParams)) {
				return
			}
			assert.Equalf(t, tt.want, got, "parseRequest(%v, %v, %v)", tt.args.w, tt.args.r, tt.args.pathParams) 
		})
	}
}