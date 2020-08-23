var _ = require("lodash");

function beforeRender(numPixels) {
  //console.log("before render, number of pixels is: " + numPixels);
  //console.log("PEPEPGA");
  //console.log(_.isNumber(3));
}

function render3D(index, x, y, z) {
  rgb(index, 1 * x, 1 * z, 1 * y);
}
