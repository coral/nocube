function render() {
  beforerender();
  return "DONE";
}

function beforerender() {
  var m = 1;
  for (i = 0; i < 1000000; i++) {
    Math.sin(1);
    m++;
  }
}
