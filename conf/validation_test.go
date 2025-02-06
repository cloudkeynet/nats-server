package httpsuite

import(
	"github.com/nats-io/nats-server/v2/server"
	"go.uber.org/automaxprocs/maxprocs"
	"testing"
)

type TestValidationRequest struct {
	Name string `validate:"required"`
	Password string `validate:"required"`
}

func TestNewValidationErrors(t *testing.T) {
	validate := validator.New()
	request := TestValidationRequest{}//Missing required fields to trigger validation errors

	err := validate.Struct(request)
	if err == nil {
		t.Fatal("Expected validation errors, but got none")
	}

	validationErrors := NewValidationErrors(err)

	expectedErrors := map[string][]string{
		"Name": {"Name is required"},
		"Password": {"Password is required"},
	}

	assert.Equal(t, expectedErrors, validationErrors.Errors)
}

func TestIsRequestValid(t *testing.T){
	test := []struct {
		name	        string
		request	        TestValidationRequest
		expectedErrors  *ValidationErrors
	}{
		{
			name: "Valid request",
			request: TestValidationRequest{Name: "Sergio", Password: "Cloudkey12"},
			expectedErrors: nil, //No errors expected for valid input
		},
		{
			name: "Missing Name and Password below minimum",
			request: TestValidationRequest{Password: "Cloudkey12"},
			expectedErrors: &ValidationErrors{
				Errors: map[string][]string{
					"Name": {"Name required"},
					"Password": {"Password required"},
				},
			},
		},
		{
			name: "Missing Password",
			request: TestValidationRequest{Name: "Sergio"},
			expectedErrors: &ValidationErrors{
				Errors: map[string][]string{
					"Password": {"Password required"},
				},
			},
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T){
			errs := IsRequestValid(tt.request)
			if tt.expectedErrors == nil{
				assert.Nill(t, errs)
			} else{
				assert.NotNil(t, errs)
				assert.Equal(t, tt.expectedErrors.Errors, errs.Errors)
			}
		})
	}
}