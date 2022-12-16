package utils

func ErrorMsg(msg string) string {
	return ColorRed + "✗ " + msg + ColorReset + "\n" // Run for some time to simulate work
}

func SuccessMsg(msg string) string {
	return ColorGreen + "✗ " + msg + ColorReset + "\n" // Run for some time to simulate work
}

func WarningMsg(msg string) string {
	return ColorYellow + "✗ " + msg + ColorReset + "\n" // Run for some time to simulate work
}
