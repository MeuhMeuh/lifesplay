let index = {
  init: function() {
    // Init
    asticode.loader.init();
    asticode.modaler.init();
    asticode.notifier.init();

    asticode.loader.show();
  }
};

document.addEventListener("astilectron-ready", function() {
  // This will send a message to GO
  astilectron.sendMessage({ name: "event.name", payload: "hello" }, function(
    message
  ) {
    console.log("received " + message.payload);
  });

  // This will listen to messages sent by the backend.
  astilectron.onMessage(function(message) {
    // Process message
    if (message.name === "event.name") {
      return { payload: message.message + " world" };
    }
  });

  asticode.loader.hide();
});
