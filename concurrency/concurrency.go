package concurrency

import "sync"

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	var wg sync.WaitGroup
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	for r := range resultChannel {
		results[r.string] = r.bool
	}

	return results
}
