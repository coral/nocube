/**
 * This is the built-in helper functions
 * The workflow is mainly inspired by the javascript-style patterns that Ben Hecke
 * created for Pixelblaze. https://www.bhencke.com/pixelblaze
 * Shoutout to Ben for a great product.
 */

var render = function render() {
  if (typeof beforeRender === "function") {
    beforeRender();
  }

  if (typeof render3D === "function") {
    render3D();
  }
};
