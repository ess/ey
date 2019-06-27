Feature: Stop A Server
  I have some serves, but I don't always want all of them to be online and
  incurring costs. The thing is, I also don't want to terminate them, because
  booting new instances takes a rather long time.

  To that end, I'd like to be able to stop a server that is running.

  Background:
    Given my Engine Yard API token is configured
    And I have a server with the ID i-00000001

  Scenario: The server is in a "running" state
    Given my server is running
    When I run `ey servers stop i-00000001`
    Then a server stop request is sent upstream
    And it exits successfully
    And the server transitions to a stopped state

  Scenario: The server is not "running"
    Given my server is not running
    When I run `ey servers stop i-00000001`
    Then no server stop request is sent upstream
    And no server state transition occurs
    But it exits successfully

    @failure
  Scenario: The API returns an error
    Given my server is running
    And API is erroring on server stop requests
    When I run `ey servers stop i-00000001`
    Then it exits with an error

    @failure
  Scenario: Invalid server ID
    Given there is no server with the ID i-00000002
    When I run `ey servers stop i-00000002`
    Then it exits with an error

