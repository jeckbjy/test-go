package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/c-bata/go-prompt"
	"github.com/gosuri/uiprogress"
	"github.com/jedib0t/go-pretty/progress"
	"github.com/manifoldco/promptui"
)

func main() {
	// example1()
	// example2()
	// example3()
	// example4()
	example5()
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "users", Description: "Store the username and age"},
		{Text: "articles", Description: "Store the article text posted by user"},
		{Text: "comments", Description: "Store the text commented to articles"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func example1() {
	fmt.Println("Please select table.")
	t := prompt.Input("> ", completer)
	fmt.Println("You selected " + t)
}

func example2() {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Number",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func example3() {
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday",
			"Saturday", "Sunday"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

var (
	autoStop    = flag.Bool("auto-stop", false, "Auto-stop rendering?")
	numTrackers = flag.Int("num-trackers", 13, "Number of Trackers")
)

func trackSomething(pw progress.Writer, idx int64) {
	total := idx * idx * idx * 250
	incrementPerCycle := idx * int64(*numTrackers) * 250

	var units *progress.Units
	switch {
	case idx%5 == 0:
		units = &progress.UnitsCurrencyPound
	case idx%4 == 0:
		units = &progress.UnitsCurrencyDollar
	case idx%3 == 0:
		units = &progress.UnitsBytes
	default:
		units = &progress.UnitsDefault
	}

	var message string
	switch units {
	case &progress.UnitsBytes:
		message = fmt.Sprintf("Downloading File    #%3d", idx)
	case &progress.UnitsCurrencyDollar, &progress.UnitsCurrencyEuro, &progress.UnitsCurrencyPound:
		message = fmt.Sprintf("Transferring Amount #%3d", idx)
	default:
		message = fmt.Sprintf("Calculating Total   #%3d", idx)
	}
	tracker := progress.Tracker{Message: message, Total: total, Units: *units}

	pw.AppendTracker(&tracker)

	c := time.Tick(time.Millisecond * 100)
	for !tracker.IsDone() {
		select {
		case <-c:
			tracker.Increment(incrementPerCycle)
		}
	}
}

func example4() {
	flag.Parse()
	fmt.Printf("Tracking Progress of %d trackers ...\n\n", *numTrackers)

	// instantiate a Progress Writer and set up the options
	pw := progress.NewWriter()
	pw.SetAutoStop(*autoStop)
	pw.SetTrackerLength(25)
	pw.ShowOverallTracker(true)
	pw.ShowTime(true)
	pw.ShowTracker(true)
	pw.ShowValue(true)
	pw.SetMessageWidth(24)
	pw.SetNumTrackersExpected(*numTrackers)
	pw.SetSortBy(progress.SortByPercentDsc)
	pw.SetStyle(progress.StyleDefault)
	pw.SetTrackerPosition(progress.PositionRight)
	pw.SetUpdateFrequency(time.Millisecond * 100)
	pw.Style().Colors = progress.StyleColorsExample
	pw.Style().Options.PercentFormat = "%4.1f%%"

	// call Render() in async mode; yes we don't have any trackers at the moment
	go pw.Render()

	// add a bunch of trackers with random parameters to demo most of the
	// features available; do this in async too like a client might do (for ex.
	// when downloading a bunch of files in parallel)
	for idx := int64(1); idx <= int64(*numTrackers); idx++ {
		go trackSomething(pw, idx)

		// in auto-stop mode, the Render logic terminates the moment it detects
		// zero active trackers; but in a manual-stop mode, it keeps waiting and
		// is a good chance to demo trackers being added dynamically while other
		// trackers are active or done
		if !*autoStop {
			time.Sleep(time.Millisecond * 100)
		}
	}

	// wait for one or more trackers to become active (just blind-wait for a
	// second) and then keep watching until Rendering is in progress
	time.Sleep(time.Second)
	for pw.IsRenderInProgress() {
		// for manual-stop mode, stop when there are no more active trackers
		if !*autoStop && pw.LengthActive() == 0 {
			pw.Stop()
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Println("\nAll done!")
}

func example5() {
	waitTime := time.Millisecond * 100
	uiprogress.Start()

	// start the progress bars in go routines
	var wg sync.WaitGroup

	bar1 := uiprogress.AddBar(20).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for bar1.Incr() {
			time.Sleep(waitTime)
		}
	}()

	bar2 := uiprogress.AddBar(40).AppendCompleted().PrependElapsed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for bar2.Incr() {
			time.Sleep(waitTime)
		}
	}()

	time.Sleep(time.Second)
	bar3 := uiprogress.AddBar(20).PrependElapsed().AppendCompleted()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= bar3.Total; i++ {
			bar3.Set(i)
			time.Sleep(waitTime)
		}
	}()

	// wait for all the go routines to finish
	wg.Wait()
}
