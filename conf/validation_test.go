package httpsuite

import (
	"github.com/nats-io/nats-server/v2/server"
	"go.uber.org/automaxprocs/maxprocs"
	"testing"
)
// TestValidationRequest define una estructura con reglas de validación.
type TestValidationRequest struct {
	Name     string `validate:"required"`  // El campo Name es obligatorio
	Password string `validate:"required"`  // El campo Password es obligatorio
}
// TestNewValidationErrors prueba la validación de la estructura TestValidationRequest.
// Se espera que falle si los campos obligatorios no están presentes.
func TestNewValidationErrors(t *testing.T) {
	validate := validator.New()
	request := TestValidationRequest{} // Falta llenar los campos obligatorios para generar errores de validación
	err := validate.Struct(request)
	if err == nil {
		t.Fatal("Expected validation errors, but got none")
	}
	// Se obtiene la lista de errores de validación
	validationErrors := NewValidationErrors(err)
	expectedErrors := map[string][]string{
		"Name":     {"Name is required"},
		"Password": {"Password is required"},
	}
	// Se comparan los errores esperados con los obtenidos
	assert.Equal(t, expectedErrors, validationErrors.Errors)
}

// TestIsRequestValid prueba diferentes casos de validación de una solicitud.
func TestIsRequestValid(t *testing.T) {
	test := []struct {
		name           string
		request        TestValidationRequest
		expectedErrors *ValidationErrors
	}{
		{
			name:          "Valid request",
			request:       TestValidationRequest{Name: "Sergio", Password: "Cloudkey12"},
			expectedErrors: nil, // No se esperan errores para una entrada válida
		},
		{
			name: "Missing Name and Password below minimum",
			request: TestValidationRequest{Password: "Cloudkey12"},
			expectedErrors: &ValidationErrors{
				Errors: map[string][]string{
					"Name":     {"Name required"},
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

	// Se ejecutan los casos de prueba
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			errs := IsRequestValid(tt.request)
			if tt.expectedErrors == nil {
				assert.Nil(t, errs)
			} else {
				assert.NotNil(t, errs)
				assert.Equal(t, tt.expectedErrors.Errors, errs.Errors)
			}
		})
	}
}