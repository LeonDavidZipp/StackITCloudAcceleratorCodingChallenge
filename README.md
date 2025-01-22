# StackITCloudAcceleratorCodingChallenge

### Requirements
You need a telegram bot token to run the application.
You can create a bot by talking to BotFather on Telegram.
You will also need to own a channel where your bot is added to & can post messages.
A guide on how to get the chat id of the channel can be found
[here](https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id).
The server is accessible both via **http and grpc**.

### Usage
```bash
set -a
source .env
set +a
```
Then either run the application in a container...
```bash
make start
```
...OR locally...
```bash
make start_local
```
...and test it using a tool like Postman. Run the following command if you have the postman cli installed
(it will also start the container):
```bash
make test
```
... OR test the following json bodies manually in the Postman app.
The url is `http://localhost:8080/forward` and the method is `POST`.
This will be forwarded:
```json
{
    "type": "Warning",
    "name": "test",
    "description": "test"
}
```
This won't be forwarded & not return an error:
```json
{
    "type": "info",
    "name": "test",
    "description": "test"
}
```
This will return an error:
```json
{
    "type": "Warning",
    "name": "test"
}
```