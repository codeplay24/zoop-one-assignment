# zoop-one-assignment

This is a back-end application that takes a request and serves that request to a registerd endpoint in Round Robin fashion. 

/urls/register - is a endpoint that receives a post call and register's a endpoint. you have to send a json structure in request body. 

json structure
--------------
{
    "url":"end-point"
}

after registeration of a url there is hold time, which is 15 seconds. In these 15 seconds no request will be forwarded to that route. 


 /proxy - this end point does not serve any specific type of http request. It can serve GET, POST, PUT, PATCH, DELETE. This end point will take the request and and send a request to any of the healthy end point in Round Robin fashion. this router forwards everything to the proxy endpoint. and responds with the respond it gets back from the endpoint.
  
