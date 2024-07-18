package domain

import (
	"fmt"
	"net/url"
)

// DumpDomain will print out all the parts of the parsed domain
func DumpDomain(domain string) {
	u, err := url.Parse(domain)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("dumping domain: %s\n", domain)
	fmt.Printf("Scheme: %s\n", u.Scheme)

	// fmt.Printf("User: %s\n", u.User)
	fmt.Printf("User: %s\n", u.User.Username())
	p, _ := u.User.Password()
	fmt.Printf("Password: %s\n", p)

	fmt.Printf("Host: %s\n", u.Host)
	fmt.Printf("Hostname(): %s\n", u.Hostname())
	fmt.Printf("Path: %s\n", u.Path)
	fmt.Printf("Fragment: %s\n", u.Fragment)
	fmt.Printf("RawQuery: %s\n", u.RawQuery)
	fmt.Printf("Query(): %s\n", u.Query())
	fmt.Printf("Opaque: %s\n", u.Opaque)
}
