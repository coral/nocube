<template>
  <div>
    <div id="container" width="800" height="800"></div>
    <p></p>
    <b-input-group prepend="X" class="mt-3">
      <b-form-input v-model="rotX" type="range" min="0" max="10000"/>
    </b-input-group>
    <b-input-group prepend="Y" class="mt-3">
      <b-form-input v-model="rotY" type="range" min="0" max="10000"/>
    </b-input-group>
  </div>
</template>

<script>
import * as Three from "three";
import axios from "axios";

export default {
  name: "Render",
  created: function() {
    this.$options.sockets.onmessage = data => this.update(data);
  },
  data: function() {
    return {
      camera: null,
      scene: null,
      renderer: null,
      mesh: null,
      mapping: null,
      group: null,
      pixelHolder: [],
      gotData: false,
      rotX: 450,
      rotY: 800
    };
  },
  methods: {
    init: function(mapping) {
      let container = document.getElementById("container");

      this.camera = new Three.PerspectiveCamera(70, 800 / 800, 0.01, 200);
      this.camera.position.z = 1.8;
      this.camera.position.x = 0;
      this.camera.position.y = 0;

      this.scene = new Three.Scene();

      var geometry = new Three.SphereGeometry(0.003, 32, 32);

      var group = new Three.Group();
      var pH = this.pixelHolder;
      mapping.forEach(function(pixel) {
        var material = new Three.MeshBasicMaterial({
          color: 0x000000
        });
        var mesh = new Three.Mesh(geometry, material);
        mesh.position.x = -0.5 + pixel.C[0];
        mesh.position.y = -0.5 + pixel.C[1];
        mesh.position.z = -0.5 + pixel.C[2];
        mesh.setColor = function(r, g, b) {
          mesh.material.color = new Three.Color(r, g, b);
        };
        pH[pixel.I] = mesh;
        if (pixel.A) {
          group.add(mesh);
        }
      });

      this.group = group;

      this.scene.add(this.group);

      this.renderer = new Three.WebGLRenderer({
        antialias: true
      });
      this.renderer.setSize(800, 800);
      container.appendChild(this.renderer.domElement);
    },
    update: function(data) {
      var pd = JSON.parse(data.data);
      var pH = this.pixelHolder;
      pd.forEach(function(value, i) {
        pH[i].setColor(value.C[0], value.C[1], value.C[2]);
      });
    },
    animate: function() {
      requestAnimationFrame(this.animate);

      //this.camera.position.x += this.mouseX - this.camera.position.x;
      //this.camera.position.y += this.mouseY - this.camera.position.y;
      this.group.rotation.x = this.rotX / 1000;
      this.group.rotation.y = this.rotY / 1000;
      //this.group.rotation.y += 0.005;
      //this.group.rotation.z += 0.005;
      this.renderer.render(this.scene, this.camera);
    }
  },
  mounted() {
    axios.get("http://127.0.0.1:8000/mapping").then(response => {
      this.mapping = response.data;
      this.init(this.mapping);
      this.animate();
    });

    //window.addEventListener("mousemove", this.mouseIsMoving);
    //window.addEventListener("resize", this.onWindowResize, false);
  }
};
</script>

<style>
</style>
