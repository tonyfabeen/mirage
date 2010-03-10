function WebSocketTest()
{
  if ("WebSocket" in window)
  {
      var ws = new WebSocket("ws://localhost:5435/monster");
      ws.onopen = function()
      {
        // Web Socket is connected. You can send data by send() method
        ws.send($("#msg").val()); //Text Input 
      };
      ws.onmessage = function (evt) { 
		var received_msg = evt.data; 
		alert(evt.data);
	  };
      ws.onclose = function() { }; // websocket is closed.
    //alert("WebSocket supported here!   :)\r\n\r\nBrowser: " + navigator.appName + " " + navigator.appVersion + "\r\n\r\ntest by jimbergman.net (based on Google sample code)");
  }
  else
  {
    // the browser doesn't support WebSocket
    alert("WebSocket NOT supported here!   :(\r\n\r\nBrowser: " + navigator.appName + " " + navigator.appVersion + "\r\n\r\ntest by jimbergman.net (based on Google sample code)");
  }
}
