package source

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	. "github.com/logrusorgru/aurora"
)

// Advent adds real Advent of Code input(s) to the queue.
// It checks *.cookie files in inputs/ folder, and uses their content
// as "session" cookie to download your input from www.adventofcode.com.
// Note it can add several inputs, not just one.
func (ins *Inputs) Advent(year, day int) {
	ms, err := filepath.Glob("inputs/*.cookie")
	check(err)
	for _, cookieName := range ms {
		*ins = append(*ins, newAdvent(year, day, cookieName))
	}
}

////////////////////////////////////////////////////////////////////////
// Implementation

func newAdvent(year, day int, cookieName string) input {
	cookie, err := ioutil.ReadFile(cookieName)
	check(err)

	name := filepath.Base(cookieName)
	iext := len(name) - len(filepath.Ext(name))
	name = name[:iext]

	in := &advent{inputBase{name: name}, year, day, string(cookie), nil}
	in.input = in
	return in
}

type advent struct {
	inputBase
	year, day int
	cookie    string
	data      *string
}

func (a *advent) Part(part uint) bool {
	return a.canProcess(part)
}

func (a *advent) String() string {
	if a.data == nil {
		s := getInput(a.year, a.day, a.name, a.cookie)
		a.data = &s
	}
	return *a.data
}

func (a *advent) Submit(part uint, val string) bool {
	a.valid[part-1] = trySubmit(a.name, a.year, a.day, part, val)
	return a.valid[part-1]
}

////////////////////////////////////////////////////////////////////////
// Internals.
//
// Download. Copypasted from 2017.

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// getInput gets input file from cache or from Advent of Code website.
func getInput(year, day int, name, cookie string) string {
	urlGet := fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day)

	c := &http.Client{Timeout: 10 * time.Second}

	inkey := fmt.Sprintf("inputs/%d_%d_%s.txt", year, day, name)

	cache, err := ioutil.ReadFile(inkey)
	if err == nil {
		//fmt.Println("cached", inkey)
		return string(cache)
	}

	fmt.Println("downloading", name, "from", urlGet)
	req, err := http.NewRequest("GET", urlGet, nil)
	req.AddCookie(&http.Cookie{
		Name:   "session",
		Value:  string(cookie),
		Path:   "/",
		Domain: ".adventofcode.com",
	})
	resp, err := c.Do(req)
	check(err)
	defer resp.Body.Close() // needed?

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		panic(resp.Status)
	}

	buf, err := ioutil.ReadAll(resp.Body)
	check(err)

	if len(buf) > 0 && buf[len(buf)-1] == 10 { // remove trailing newline. AoC bug or what?
		buf = buf[:len(buf)-1]
	}

	err = ioutil.WriteFile(inkey, buf, 600)
	check(err)

	return string(buf)
}

////////////////////////////////////////////////////////////////////////
// Submit. Copypasted from 2017.

var lastSubmit time.Time

const submitThrottle time.Duration = 5 * time.Second

func trySubmit(name string, year, day int, part uint, v string) (out bool) {
	out = true
	inkey := fmt.Sprintf("inputs/%d_%d_%s.txt", year, day, name)
	outkey := fmt.Sprintf("results/%d_%d_%d_%s.txt", year, day, part, name)
	result, err := ioutil.ReadFile(outkey)

	fmt.Printf("part%d: %s ", part, Cyan(v))

	if err == nil {
		ex := string(result)
		if v != ex {
			fmt.Printf("%s\n", Red("✗ expected "+ex))
			return false
		}
		fmt.Print(Green("✓"))
		infi, err1 := os.Stat(inkey)
		outfi, err2 := os.Stat(outkey)
		if err1 == nil && err2 == nil {
			fmt.Print(" ", Brown(outfi.ModTime().Sub(infi.ModTime()).Round(time.Second)))
		}
		fmt.Println()
		return
	}

	fmt.Println()

	if dry {
		return
	}
	cookie, err := ioutil.ReadFile("results/" + name + ".cookie")
	if err != nil { // no cookie -> no submit.
		return
	}

	urlPost := fmt.Sprintf("https://adventofcode.com/%d/day/%d/answer", year, day)
	fmt.Println("submitting to:", urlPost, "for", name)

	wait := submitThrottle - time.Since(lastSubmit)
	if wait > 0 {
		fmt.Println("waiting", Cyan(wait))
		time.Sleep(wait)
	}
	data := url.Values{}
	data.Set("level", strconv.Itoa(int(part)))
	data.Set("answer", v)
	encoded := data.Encode()
	req, err := http.NewRequest("POST", urlPost, strings.NewReader(encoded))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	req.AddCookie(&http.Cookie{
		Name:   "session",
		Value:  string(cookie),
		Path:   "/",
		Domain: ".adventofcode.com",
	})

	c := &http.Client{Timeout: 10 * time.Second}
	resp, err := c.Do(req)
	check(err)
	defer resp.Body.Close() // needed?

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		panic(resp.Status)
	}

	lastSubmit = time.Now()

	fmt.Println("status:", resp.Status)
	buf, err := ioutil.ReadAll(resp.Body)
	check(err)
	html := string(buf)
	reg := regexp.MustCompile("(?s)<main>\\s*<article>\\s*<p>(.*)</p>\\s*</article>\\s*</main>")
	m := reg.FindStringSubmatch(html)
	main := html
	if len(m) > 1 {
		main = m[1]
	}

	if strings.Contains(main, "You don't seem to be solving the right level.") {
		fmt.Println("Already submitted.")
		ioutil.WriteFile(outkey, []byte("Unknown value"), 600)
		return
	}

	if strings.Contains(main, "That's the right answer!") {
		fmt.Println(Green("Correct answer."))
		ioutil.WriteFile(outkey, []byte(v), 600)
		fi, err := os.Stat(inkey)
		if err == nil {
			fmt.Println("It took", Brown(time.Since(fi.ModTime()).Round(time.Second)))
		}
		return
	}

	if strings.Contains(main, "That's not the right answer") {
		fmt.Println(Red("Incorrect answer."))
		if strings.Contains(main, "your answer is too low") {
			fmt.Println(Red("- too low."))
		} else if strings.Contains(main, "your answer is too high") {
			fmt.Println(Red("- too high."))
		}
		ioutil.WriteFile(outkey+".err.txt", []byte(main), 600)
		return false
	}

	if strings.Contains(main, "You gave an answer too recently;") {
		fmt.Println(Brown("Submitting too soon. Wait some more."))
		ioutil.WriteFile(outkey+".err.txt", []byte(main), 600)
		return false
	}

	if strings.Contains(main, "Congratulations!  You've finished every puzzle") {
		fmt.Println(Green("Congratulations!  You've finished every puzzle in this year."))
		ioutil.WriteFile(outkey, []byte(v), 600)
		return
	}
	// some unknown response.
	ioutil.WriteFile(outkey+".err.txt", []byte(main), 600)
	fmt.Println("main:", main)
	return
}
