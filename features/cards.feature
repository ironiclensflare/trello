Feature: Adding cards

    Add a new card to the Trello board

Scenario: Add a card
    Given I am a user
    When I add a card with the text "Test card"
    Then the card is on the board
    And it has the text "Test card"

Scenario: Add a card with a link
    Given I am a user
    When I add a card with the text "Test card with link"
    And I attach the URL "http://www.example.com"
    Then the card is on the board
    And it has the text "Test card with link"
    And it has the attached URL "http://www.example.com"
