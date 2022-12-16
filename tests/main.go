// Example application that uses all of the available API options.
package main

import (
	"time"
	"vmctl/src/utils"

	"github.com/briandowns/spinner"
)

func main() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Color("white")
	s.Suffix = " Checking KVM Installation..."
	s.Start()
	// Do something here
	time.Sleep(1 * time.Second)
	// Do something here
	s.Color("green")
	s.FinalMSG = utils.ColorGreen + "✔ KVM Installation check." + utils.ColorReset + "\n"
	s.Stop()

	s2 := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s2.Color("white")
	s2.Suffix = " Checking Some Other thing..." // Set the spinner color to red
	s2.Start()                                  // Start the spinner
	// Do something here
	time.Sleep(1 * time.Second)
	// Do something here
	s2.Color("green")
	s2.FinalMSG = utils.ColorYellow + "⚠ Warning check." + utils.ColorReset + "\n" // Run for some time to simulate work
	s2.Stop()

	s3 := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s3.Color("white")
	s3.Suffix = " Third step thing..." // Set the spinner color to red
	s3.Start()                         // Start the spinner
	// Do something here
	time.Sleep(1 * time.Second)
	// Do something here
	s3.Color("green")
	s3.FinalMSG = utils.ColorRed + "✗ Something went wrong." + utils.ColorReset + "\n" // Run for some time to simulate work
	s3.Stop()

	// s.Restart()
	// s.Color("white")
	// s.Suffix = " Checking KVM Installation..." // Set the spinner color to red
	// s.Start()                                  // Start the spinner
	// time.Sleep(1 * time.Second)
	// s.Color("green")
	// s.FinalMSG = utils.ColorGreen + "✔ KVM Installation check." + utils.ColorReset // Run for some time to simulate work
	// s.Stop()
	// s2 := spinner.New(spinner.CharSets[9], 100*time.Millisecond) // Build our new spinner
	// s2.Color("red")                                              // Set the spinner color to red
	// s2.Start()                                                   // Start the spinner
	// time.Sleep(4 * time.Second)                                  // Run for some time to simulate work
	// s2.UpdateCharSet(spinner.CharSets[9])                        // Update spinner to use a different character set
	// s2.UpdateSpeed(100 * time.Millisecond)                       // Update the speed the spinner spins at
	// s2.Prefix = "prefixed text: "                                // Prefix text before the spinner
	// s2.Prefix = ""
	// s2.Suffix = " :appended text" // Append text after the spinner
	// time.Sleep(4 * time.Second)

	// s.Prefix = "Colors: "

	// if err := s.Color("yellow"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.Start()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("red"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[20])
	// s.Reverse()
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("blue"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[3])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("cyan"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[28])
	// s.Reverse()
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("green"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[25])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("magenta"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.UpdateCharSet(spinner.CharSets[32])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// if err := s.Color("white"); err != nil {
	// 	log.Fatalln(err)
	// }

	// s.FinalMSG = "Complete!\nNew line!\nAnother one!\n"

	// s.UpdateCharSet(spinner.CharSets[31])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// s.Stop() // Stop the spinner

	// s.Prefix = "Earth! "
	// s.UpdateCharSet(spinner.CharSets[39])
	// s.Restart()

	// time.Sleep(4 * time.Second) // Run for some time to simulate work

	// s.Stop() // Stop the spinner

	// println("")
}
