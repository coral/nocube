import Vue from "vue";
import App from "./App.vue";
import router from "./router";

Vue.config.productionTip = false;

import VueNativeSock from "vue-native-websocket";
Vue.use(VueNativeSock, "ws://localhost:9000/stream", {
  format: "json",
  reconnection: true,
  reconnectionAttempts: 50,
  reconnectionDelay: 3000
});

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
