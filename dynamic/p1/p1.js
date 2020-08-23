var _ = require("lodash");

module.exports.beforeRender = function () {
  console.log("LOL");
};

function render3D(denis) {
  var x = new Float64Array(mapping.Length);

  mapping.Pixels.forEach(pixel);
  x[0] = 0.9;
  x[4] = 0.4;

  x[128] = 1.8;
  return x.buffer;
}

function pixel(d) {
  console.log(d.Y);
}
