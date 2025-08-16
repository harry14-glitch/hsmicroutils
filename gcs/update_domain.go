package gcs

import "strings"

func GetUpdatedDomain(domain string) string {
	if domain == "" {
		return ""
	}
	return "hs_" + strings.Replace(domain, ".", "_", -1)
}
