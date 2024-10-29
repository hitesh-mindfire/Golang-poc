const nats = require("nats");

async function main() {
  // Connect to NATS;
  const nc = await nats.connect({ servers: "nats-server:4222" });

  console.log("Notification Service listening for order events...");

  // Subscribe to "order.created" subject
  const sub = nc.subscribe("order.created");
  for await (const m of sub) {
    console.log(`Received order notification: ${m.data}`);
  }

  nc.closed().then((err) => {
    if (err) {
      console.log(`Notification Service exited with an error: ${err.message}`);
    } else {
      console.log("Notification Service closed gracefully.");
    }
  });
}

main().catch((err) => {
  console.error(`Error starting Notification Service: ${err.message}`);
});
