from flask import Flask, request
import requests

app = Flask(__name__)

@app.route('/api/command', methods=['POST'])
def handle_command():
    # Get the command from the request
    command = request.json['command']

    # Pass the command to the other API
    response = requests.post('https://other-api.com/api/function', json={'command': command})

    # Return the result of the other API's function to the original requestor
    return response.json()

if __name__ == '__main__':
    app.run()
#To make a curl request against this API, you can use the following command:
#curl -X POST https://python-api.com/api/command -d '{"command": "YOUR_COMMAND"}'

#This will send a POST request to the /api/command endpoint with the specified command in the request body. The API will pass this command to the other API,
# and return the result to the original requestor. You can then use the -o or --output flag to save the result to a file, like this:
#curl -X POST https://python-api.com/api/command -d '{"command": "YOUR_COMMAND"}' -o output.txt

#Nginx config This configuration will route all requests to example.com to the API, which is running on localhost on port 5000.
#server {
#    listen 80;
#    server_name example.com;
#
#    location / {
#       proxy_pass http://localhost:5000;
#    }
#}
#to Curl curl -X POST http://example.com/api/command -d '{"command": "YOUR_COMMAND"}'
