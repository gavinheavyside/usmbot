GOOS=linux GOARCH=amd64 go build && \
zip usmbot.zip usmbot && \
rm usmbot && \
aws --profile $AWS_PROFILE --region eu-west-1 lambda update-function-code --function-name usmbot --zip-file fileb://usmbot.zip && \
rm usmbot.zip
