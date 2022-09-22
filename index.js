const bodyParser = require('body-parser');
const express = require('express');
const http = require('http');
const {v4: uuidv4} = require('uuid');

const globalDataStore = {}

const server = express();

server.use(bodyParser.raw({inflate: true}))

// endpoint routing configuration

server.get('/', (req, res) => {
  res.writeHead(200, { "Content-Type": "application/json"});
  const headersData = JSON.stringify(res.getHeaders());
  res.write(headersData);
  res.end();
});

server.get('/stubbed-process-1', (req, res) => {
  if (Math.random() * 100 > 20) {
    res.status(200);
    res.write('ok');
  } else {
    res.status(500);
    res.write('not ok');
  }
  res.end();
});

server.get('/stubbed-process-2', (req, res) => {
  if (Math.random() * 100 > 40) {
    res.status(200);
    res.write('ok');
  } else {
    res.status(500);
    res.write('not ok');
  }
  res.end();
});

server.get('/stubbed-process-3', (req, res) => {
  if (Math.random() * 100 > 60) {
    res.status(200);
    res.write('ok');
  } else {
    res.status(500);
    res.write('not ok');
  }
  res.end();
});

server.get('/stubbed-process-4', (req, res) => {
  if (Math.random() * 100 > 80) {
    res.status(200);
    res.write('ok');
  } else {
    res.status(500);
    res.write('not ok');
  }
  res.end();
});

server.post('/save', (req, res) => {
  let chunks = '';
  req.on('data', (data) => (data != undefined) ? chunks += data : null);
  req.on('end', () => {
    const requestId = uuidv4();
    globalDataStore[requestId] = String(chunks);
    res.send(requestId);
  });
});

server.get('/load/:uuid', (req, res) => {
  const requestId = req.params.uuid;
  const globalDataInstance = globalDataStore[requestId];
  if (globalDataInstance !== undefined && globalDataInstance !== null) {
    res.status(200);
    res.write(globalDataInstance);
    res.end();
    return;
  }
  res.status(404);
  res.write('not found');
  res.end();
});

// healthcheck function

(function check() {
  const targetServce = 'http://127.0.0.1:3000';
  console.log(`pinging ${targetServce}...`)
  http.get(targetServce, (res) => {
    let chunks = '';
    res.on('data', (data) => (data != undefined) ? chunks += data : null);
    res.on('end', () => {
      const data = String(chunks);
      console.log(`body data: ${data}`);
    })
  })
  setTimeout(check, 1000);
})();

// start the server

server.listen(3000);
