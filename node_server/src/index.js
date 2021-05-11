var http = require("http");

//create a server object:
http
  .createServer(function (req, res) {
    // res.writeHead(302, {
    //   location: "redir",
    // });
    res.write("Hello Node!"); //write a response to the client
    res.end(); //end the response
  })
  .listen(8080); //the server object listens on port 8080
