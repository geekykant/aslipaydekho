import amqp, { Channel, Connection } from "amqplib";

import dotenv from "dotenv";
dotenv.config();

let connUrl: string = process.env.RABBIT_MQ_SERVER_URL || "";
let queueName: string = process.env.RABBIT_MQ_CHANNEL_NAME || "";

let conn: Connection| null = null;
let channel: Channel| null = null;

const initializeMQ = (async () => {
  conn = await amqp.connect(connUrl);
  channel = await conn.createChannel();

  await channel.assertQueue(queueName);

  process.once("SIGINT", async () => {
    if (channel != null) await channel.close();
    if (conn != null) await conn.close();
  });
});

const fetchPostFromMQ = async (event: any) => {
  var response = {};

  try {
    await initializeMQ();

    if (conn == null || channel == null) {
      throw Error("Connection couldn't be established with MQ");
    }

    let msgValue = await channel.get(queueName);

    if (msgValue) {
      response = JSON.parse(msgValue.content.toString());
    } else {
      throw Error("No more messages");
    }

    if (channel != null) await channel.close();
    if (conn != null) await conn.close();
  }
  catch (err: any){
      console.warn(err);
  }
  return response;
};

export default defineEventHandler(fetchPostFromMQ);
