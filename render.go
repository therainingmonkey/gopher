package main

var bodyLines = make([]string, 3)

// TODO Error handling

func loadPage(url string) (err error) {
	bodyLines, err = retrieve(url, "\n")
	return err
}

func loadHomepage() (err error) {
	if homepage == "welcome" {
		bodyLines[0] = "Welcome"
		bodyLines[1] = "to"
		bodyLines[2] = "Gopher!"
		err = nil
		return err
	} else {
		bodyLines, err = retrieve(homepage, "\n")
		return err
	}
}
