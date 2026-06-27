package tls

// CertificateStore manages the loaded certificates
type CertificateStore struct {
	// ... existing fields
}

func (s *CertificateStore) Reload() {
	// Logic to reload certificates from disk into memory
}