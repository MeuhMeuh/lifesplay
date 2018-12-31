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

const displayEvents = events => {
  console.log(events);
  const htmlEvents = [];
  const htmlEventList = document.getElementById("events");
  events.forEach(e => {
    let li = document.createElement("li");
    let eventName = document.createElement("h3");
    eventName.innerHTML = e.summary;
    li.appendChild(eventName);

    // Determining the time (and thus the position) of the event.
    // We actually will ignore the timezone : it is already at the timezone
    // That the user uses, and thus, it means that s.he will want to display the events
    // According to this timezone.
    const startDateTime = new Date(Date.parse(e.start.dateTime));
    const endDateTime = new Date(Date.parse(e.end.dateTime));
    const startDateTimeMinutes =
      startDateTime.getHours() * 60 + startDateTime.getMinutes();

    const eventDuration = Math.abs(endDateTime - startDateTime) / 36e5;

    const topPosition = (startDateTimeMinutes * 180) / 60;
    // We remove 5px from the height to have a delimiter between the date separations
    // And event neighbors.
    const eventHeight = 180 * eventDuration - 5;
    li.style = `top: ${topPosition}px; height: ${eventHeight}px;`;

    htmlEventList.appendChild(li);
  });
  // Flex-start to have the event list start from the top.
  $("main").css("alignItems", "flex-start");
  $(htmlEventList).css("display", "block");
  $(htmlEventList).animate({ opacity: 1 }, 800);

  // Moving to the right event, placing it at the center.
  const now = new Date();
  const nowDateTimeMinutes = now.getHours() * 60 + now.getMinutes();
  const nowPosition = (nowDateTimeMinutes * 180) / 60;
  const scroll = nowPosition - 480 / 2;

  $("body").animate(
    {
      scrollTop: scroll
    },
    1600,
    () => {
      $("#timeDelimiter").animate(
        {
          top: nowPosition,
          opacity: 0.75
        },
        800
      );
    }
  );
};

document.addEventListener("astilectron-ready", function() {
  // This will send a message to GO
  astilectron.sendMessage({ name: "ui.ready", payload: {} }, function(message) {
    asticode.loader.hide();
    if (message.payload.error) {
      return handleError(asticode.notifier, message.payload);
    }
    $("h1 span").html(message.payload.body.firstName);
    $("h1").fadeIn(800, () => {
      $("h2").fadeIn(800, () => {
        $("#splash").fadeOut(800, () =>
          displayEvents(message.payload.body.events)
        );
      });
    });
  });

  // This will listen to messages sent by the backend.
  astilectron.onMessage(function(message) {
    // Process message
    if (message.name === "event.name") {
      return { payload: message.message + " world" };
    }
  });
});
