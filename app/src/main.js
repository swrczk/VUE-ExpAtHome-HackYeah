import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';

Vue.config.productionTip = false

new Vue({
  vuetify,
  render: h => h(App),
        data () {
        return {
            items: [
                { title: 'Dashboard', icon: 'dashboard' },
                { title: 'Account', icon: 'account_box' },
                { title: 'Admin', icon: 'gavel' },
            ],
        }
    },
}
).$mount('#app')
