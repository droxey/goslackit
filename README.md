<p align="center">
  <img src="gopher-typing.gif" height="250">
</p>

# goslackit

[![Go Report Card](https://goreportcard.com/badge/github.com/droxey/goslackit)](https://goreportcard.com/report/github.com/droxey/goslackit) [![Codacy Badge](https://api.codacy.com/project/badge/Grade/7ed40f9f3ecf46709879d5fbac28fd9b)](https://www.codacy.com/app/droxey/goslackit?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=droxey/goslackit&amp;utm_campaign=Badge_Grade)

[BEW2.5] Fork this repo to begin the Slackbot goroutines challenge presented in class on [Day 7](https://github.com/Make-School-Courses/BEW-2.5-Strongly-Typed-Ecosystems/blob/master/Lessons/Lesson07.md).

## Instructions

1. Create a new application here. https://api.slack.com/apps?new_classic_app=1 - Set the name to be whatever you want.
1. Go to OAuth and Permissions on the left side under "Features", scroll down to "Scopes".
1. Add the three following scopes. `channels:history` , `channels:read` , `chat:write:bot`. 
1. Go to "App Home" on the left side, and then add a "legacy bot user.
1. Go to Install App on the left side under "Settings". Go through the steps and add the bot into the "Make School Students" slack channel.
1. This will give you an OAuth token. Save the one that starts with `xoxb-` as you will be needing it in your code.
1. Go to the #golang-slackbots channel, and @ the bot that you just invited into the workspace. Invite the bot into the channel and now we can move onto the code!
