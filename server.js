const net = require('net');

net.createServer((sock) => {
  console.log('connected', sock);
  sock.on('data', (data) => {
    console.log('data:', data, sock);
    sock.write('you said:', data);
  });
  sock.on('close', (data) => {
    console.log('close', data);
  });
}).listen(6969, '127.0.0.1');

console.log('Server start');

