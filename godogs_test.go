package main

import "github.com/cucumber/godog"

func iAddACardWithTheText(text string) error {
	return godog.ErrPending
}

func iAmAUser() error {
	return godog.ErrPending
}

func itHasTheText(text string) error {
	return godog.ErrPending
}

func theCardIsOnTheBoard() error {
	return godog.ErrPending
}

func iAttachTheURL(url string) error {
	return godog.ErrPending
}

func itHasTheAttachedURL(arg1 string) error {
	return godog.ErrPending
}

func iHaveNotSetTheEnvironmentVariable(varName string) error {
	return godog.ErrPending
}

func iReceiveAnErrorSaying(errorText string) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I add a card with the text "([^"]*)"$`, iAddACardWithTheText)
	s.Step(`^I am a user$`, iAmAUser)
	s.Step(`^it has the text "([^"]*)"$`, itHasTheText)
	s.Step(`^the card is on the board$`, theCardIsOnTheBoard)
	s.Step(`^I attach the URL "([^"]*)"$`, iAttachTheURL)
	s.Step(`^it has the attached URL "([^"]*)"$`, itHasTheAttachedURL)
	s.Step(`^I have not set the environment variable "([^"]*)"$`, iHaveNotSetTheEnvironmentVariable)
	s.Step(`^I receive an error saying "([^"]*)"$`, iReceiveAnErrorSaying)
}
