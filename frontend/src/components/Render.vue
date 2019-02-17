<template>
  <div>
    <div id="container" width="500" height="500"></div>
    <p>{{mapping}}</p>
  </div>
</template>

<script>
import * as Three from "three";
import axios from "axios";

export default {
  name: "Render",
  created: function() {
    // this.$options.sockets.onmessage = data => console.log(data);
  },
  data: function() {
    return {
      camera: null,
      scene: null,
      renderer: null,
      mesh: null,
      mapping: null,
      group: null,
      mouseX: 0,
      mouseY: 0,
      windowHalfX: 0,
      windowHalfY: 0
    };
  },
  methods: {
    init: function() {
      let container = document.getElementById("container");

      this.camera = new Three.PerspectiveCamera(70, 500 / 500, 0.01, 100);
      this.camera.position.z = 1;

      this.scene = new Three.Scene();

      var geometry = new Three.SphereGeometry(0.003, 32, 32);
      var material = new Three.MeshBasicMaterial({ color: 0xffff00 });

      var group = new Three.Group();
      this.mapping.forEach(function(pixel) {
        var mesh = new Three.Mesh(geometry, material);
        mesh.position.x = pixel.C[0];
        mesh.position.y = pixel.C[1];
        mesh.position.z = pixel.C[2];
        if (pixel.A) {
          group.add(mesh);
        }
      });

      this.group = group;

      this.scene.add(this.group);

      this.renderer = new Three.WebGLRenderer({ antialias: true });
      this.renderer.setSize(500, 500);
      container.appendChild(this.renderer.domElement);
    },
    animate: function() {
      requestAnimationFrame(this.animate);

      //this.camera.position.x += this.mouseX - this.camera.position.x;
      //this.camera.position.y += this.mouseY - this.camera.position.y;
      this.group.rotation.x += 0.005;
      this.group.rotation.y += 0.005;
      this.group.rotation.z += 0.005;
      this.renderer.render(this.scene, this.camera);
    },
    mouseIsMoving(event) {
      this.mouseX = (event.clientX - this.windowHalfX) * 10;
      this.mouseY = (event.clientY - this.windowHalfY) * 10;
    },
    onWindowResize(event) {
      this.windowHalfX = window.innerWidth / 2;
      this.windowHalfY = window.innerHeight / 2;
    }
  },
  mounted() {
    axios.get("http://127.0.0.1:8000/mapping").then(response => {
      this.mapping = response.data;
      this.init();
      this.animate();
    });

    //window.addEventListener("mousemove", this.mouseIsMoving);
    //window.addEventListener("resize", this.onWindowResize, false);
  }
};
</script>

<style>
</style>
