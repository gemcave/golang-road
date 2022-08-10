
package solution

import (
	"fmt"
)

// BEGIN (write your solution here)
func DomainForLocale(domain, locale string) string {
	if(locale == "") {
		return fmt.Sprintf("%s.%s", "en", domain)
	}
	return fmt.Sprintf("%s.%s", locale, domain)

}

// END
