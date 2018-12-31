const util = require("util");

let index = {
  init: function() {
    console.log("Initializing asticode tools...");
    asticode.loader.init();
    asticode.modaler.init();
    asticode.notifier.init();

    asticode.loader.show();
  }
};

const handleError = (notifier, payload) => {
  notifier.error(payload.error);
  console.error(`Raised error: ${payload.error}. Body is following.`);
  console.log(payload);
};

document.addEventListener("astilectron-ready", function() {
  // This will send a message to GO
  astilectron.sendMessage({ name: "ui.ready", payload: {} }, function(message) {
    asticode.loader.hide();
    if (message.payload.error) {
      return handleError(asticode.notifier, message.payload);
    }
    document.querySelector("h1 span").innerHTML = message.payload.FirstName;
    // $("h1").fadeIn(800);
  });

  // This will listen to messages sent by the backend.
  astilectron.onMessage(function(message) {
    // Process message
    if (message.name === "event.name") {
      return { payload: message.message + " world" };
    }
  });
});
