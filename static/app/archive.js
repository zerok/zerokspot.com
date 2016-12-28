import Vue from "vue";
import Archive from "./components/archive.vue";

window.app = new Vue({
    template: "<Archive />",
    components: {
        Archive
    }
}).$mount("#app");
