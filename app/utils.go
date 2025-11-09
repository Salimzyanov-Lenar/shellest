package main

// package for utils

func findRedirectOutputIndex(commands []string) (int, bool) {
	// - allow to find redirect 'flag' index
	redirectIndex := -1
	for i, p := range commands {
		if p == ">" || p == "1>" {
			redirectIndex = i
			return redirectIndex, true
		}
	}
	return redirectIndex, false
}
