GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap && \
zip usmbot.zip bootstrap && \
rm bootstrap && \
aws --profile $AWS_PROFILE --region eu-west-1 lambda update-function-code --function-name usmbot-mastodon --zip-file fileb://usmbot.zip && \
rm usmbot.zip
