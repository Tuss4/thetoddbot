# The Todd Bot
A Slack Outgoing Webhook integration in Go, that listens and serves scrubs clips of the todd based on the query. This can be easily reconfigured into any type of bot you want.

[![Build Status](https://drone.io/github.com/Tuss4/thetoddbot/status.png)](https://drone.io/github.com/Tuss4/thetoddbot/latest)

# Install
+ `go get github.com/tuss4/thetoddbot`
+ Add `export SLACK_TOKEN=[your_slack_token]` to your .bashrc or your server's user-data.
+ Be sure to update the `trigger` var to what you actually want your hook's trigger to be.

# Run tests
`go test` from within the directory.
