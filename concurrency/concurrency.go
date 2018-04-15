package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

// calling this function in a goroutin executes it in a lightweight thread
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func myGoroutine() {
	go say("world") // executes "say" in a separate goroutine and continues immediately
	say("hello")
}

// ========

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func myChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

// ========

// Closing a channel
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x // Send operation
		x, y = y, x+y
	}
	close(c)
}

// Iterate over messages in a channel
func myChannelRange() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c { // Receive operation
		fmt.Println(i)
	}
}

// =======

// The select statement lets a goroutine wait on multiple communication operations.
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.

// Calculates fibonacci until a value is received in the "quit" channel
func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // Send operation
			x, y = y, x+y
		case <-quit: // Receive operation
			fmt.Println("quit")
			return
		}
	}
}

// Starts a goroutine that prints 10 values received in the "c" channel and then sends a value to the "quit" channel
func myChannelSelect() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c) // Receive operation
		}
		quit <- 0 // Send operation
	}()
	fibonacci2(c, quit)
}

// ========

// Goroutine exercise

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, cap(ch1))
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < cap(ch1); i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}
	return true
}

// ========

// Mutex

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mux.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mux.Unlock()
	return c.v[key]
}

func myMutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// ========

// Mutex exercise

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	time.Sleep(time.Millisecond * 500) // Added by myself to see if the parallelization worked
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url) // fmt.Errorf creates an error object
}

type cache struct {
	result map[string]fakeResult
	mux    sync.Mutex
}

var myCache = cache{
	result: make(map[string]fakeResult),
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	if depth <= 0 || myCache.result[url].body != "" {
		return
	}
	// Fill before fetching, because fetching takes some time and multiple goroutines could start fetching the same URL at the same time
	myCache.fill(url, "/", nil)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	myCache.fill(url, body, urls)
	for _, u := range urls {
		go Crawl(u, depth-1, fetcher)
	}
	return
}

func (c *cache) fill(url string, body string, urls []string) {
	// Lock, fill and unlock cache
	c.mux.Lock()
	c.result[url] = fakeResult{
		body: body,
		urls: urls,
	}
	// Unlock directly instead of via "defer", because this method calls itself recursively
	c.mux.Unlock()
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// ========

func main() {
	myGoroutine()

	myChannel()

	// Buffered channels
	ch := make(chan int, 2)
	ch <- 1 // Doesn't block
	ch <- 2 // Doesn't block
	// Sending another message through the channel would overfill the buffer and block the current thread
	// (only until another goroutine reads from the channel, which is not the case in this example code)
	// ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	myChannelRange()

	myChannelSelect()

	// The "default" case in a select statement is run if no other case is ready
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	var br bool
	for br == false {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			br = true
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}

	ch = make(chan int, 10) // Exercise says tree.New() creates a tree with 10 values
	go Walk(tree.New(1), ch)
	for i := 0; i < cap(ch); i++ {
		fmt.Println(<-ch)
	}
	same := Same(tree.New(1), tree.New(1))
	if same == false {
		panic("same is false but should be true")
	}
	same = Same(tree.New(1), tree.New(2))
	if same == true {
		panic("same is true but should be false")
	}

	myMutex()

	Crawl("https://golang.org/", 4, fetcher)
	time.Sleep(time.Second * 2)
}
