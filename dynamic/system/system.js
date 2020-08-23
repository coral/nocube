var x = new Float64Array(mapping.Length * 3);

var numPixels = mapping.Length;

module.exports.render = function () {
  beforeRender(mapping.Length);

  mapping.Pixels.forEach((pixel) =>
    render3D(pixel.Index, pixel.X, pixel.Y, pixel.Z)
  );

  return x;
};

function rgb(i, r, g, b) {
  x[i * 3] = r;
  x[i * 3 + 1] = g;
  x[i * 3 + 2] = b;
}

function hsv(i, r, g, b) {}
