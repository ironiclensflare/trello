Feature: Errors

    The application should handle various error scenarios

Scenario: Missing Trello key
    Given I am a user
    And I have not set the environment variable "TRELLO_KEY"
    When I add a card with the text "Test card"
    Then I receive an error saying "Trello API key not set"

Scenario: Missing Trello token
    Given I am a user
    And I have not set the environment variable "TRELLO_TOKEN"
    When I add a card with the text "Test card"
    Then I receive an error saying "Trello token not set"

Scenario: Missing Trello board ID
    Given I am a user
    And I have not set the environment variable "TRELLO_BOARD"
    When I add a card with the text "Test card"
    Then I receive an error saying "Trello board ID not set"

Scenario: Card text blank
    Given I am a user
    When I add a card with the text ""
    Then I receive an error saying "Card has no text"
