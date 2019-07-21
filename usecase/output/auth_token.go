package output

// AuthTokenOutput represents response body for auth token
type AuthTokenOutput struct {
	AuthToken string `json:"authToken" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjI4NDY5MTUsImlkIjo1fQ.2VZWDiWu3KDULr8p0MlPwxbTKJGnHLhcg3L_Ishx9e4"`
}

// NewAuthTokenOutput returns initialized ConfirmOutput
func NewAuthTokenOutput(authToken string) *AuthTokenOutput {
	return &AuthTokenOutput{AuthToken: authToken}
}
