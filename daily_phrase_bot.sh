#!/bin/bash
nohup go run /opt/daily_phrase_bot/bot.go &
nohup go run /opt/daily_phrase_bot/daily_phrase_generator.go &

exit 0
