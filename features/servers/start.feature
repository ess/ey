Feature: Start A Server
  I have some serves, but I don't always want all of them to be online and
  incurring costs. The thing is, I also don't want to terminate them, because
  booting new instances takes a rather long time.

  To that end, I'd like to be able to start a server that is stopped.

  Background:
    Given my Engine Yard API token is configured
    And I have a server with the ID i-00000001

  Scenario: The server is in a "stopped" state
    Given my server is stopped
    When I run `ey servers start i-00000001`
    Then a server start request is sent upstream
    And it exits successfully
    And the server transitions to a running state

  Scenario: The server is not "stopped"
    Given my server is not stopped
    When I run `ey servers start i-00000001`
    Then no server start request is sent upstream
    And no server state transition occurs
    But it exits successfully

    @failure
  Scenario: The API returns an error
    Given my server is stopped
    And API is erroring on server start requests
    When I run `ey servers start i-00000001`
    Then it exits with an error

    @failure
  Scenario: Invalid server ID
    Given there is no server with the ID i-00000002
    When I run `ey servers start i-00000002`
    Then it exits with an error
