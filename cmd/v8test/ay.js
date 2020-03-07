V8Worker2.print("BENIS");

V8Worker2.recv(function(msg) {
  V8Worker2.print("WORKS", msg.byteLength);

  V8Worker2.print(msg);
  V8Worker2.send(new ArrayBuffer([0, 0, 0, 0]));
});
