package obfuscate

import "strings"

// EmailAddressPartially obfuscates the email address by masking everything but the first and last character of each part (except TLD).
// At least some characters of the local part will always be masked.
func EmailAddressPartially(emailAddress string) string {
	parts := strings.Split(emailAddress, "@")

	if len(parts) == 1 {
		return partiallyMaskEmailAddressLocal(parts[0])
	}

	for i, part := range parts {
		if i == len(parts)-1 {
			parts[i] = partiallyMaskEmailAddressDomain(part)
		} else {
			parts[i] = partiallyMaskEmailAddressLocal(part)
		}
	}

	return strings.Join(parts, "@")
}

// Partially mask all items of the local part.
func partiallyMaskEmailAddressLocal(local string) string {
	parts := strings.Split(local, ".")
	for i, part := range parts {
		parts[i] = maskEmailAddressPartItem(part)
	}
	return strings.Join(parts, ".")
}

// Partially mask everything but the last item of the domain part.
func partiallyMaskEmailAddressDomain(domain string) string {
	parts := strings.Split(domain, ".")
	for i := 0; i < len(parts)-1; i++ {
		parts[i] = maskEmailAddressPartItem(parts[i])
	}
	return strings.Join(parts, ".")
}

func maskEmailAddressPartItem(partItem string) string {
	switch len(partItem) {
	case 0:
		return ""
	case 1:
		return "*"
	case 2:
		return partItem[:1] + "*"
	}

	return partItem[:1] + "*" + partItem[len(partItem)-1:]
}
