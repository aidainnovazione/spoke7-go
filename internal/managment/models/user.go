package models

type UserInfo struct {
	Sub         string   // Subject (unique ID)
	Iss         string   // Issuer
	Aud         string   // Audience
	Name        string   // Preferred username
	DisplayName string   // Full display name
	Email       string   // Email address
	Avatar      string   // Profile picture URL
	Address     []string // Address (optional)
	Phone       string   // Phone number (optional)
	Groups      []string // Group memberships
	Owner       string   //owner
}
