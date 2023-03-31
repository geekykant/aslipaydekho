import express from "express";

const app = express();
app.use(express.json());

function getHello(event) {
  var value = new Date();
  
  return {
    date: value
  };
}

export default defineEventHandler(getHello);


