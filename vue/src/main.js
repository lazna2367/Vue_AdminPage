// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import 'core-js/es6/promise'
import 'core-js/es6/string'
import 'core-js/es7/array'
// import cssVars from 'css-vars-ponyfill'
import Vue from 'vue'
import BootstrapVue from 'bootstrap-vue'
import App from './App'
import router from './router'
import axios from 'axios'
import store from './store'
import VueApexCharts from 'vue-apexcharts'
import VueGoogleCharts from 'vue-google-charts'

import moment from 'moment'
import Web3 from 'web3'

// todo
// cssVars()
Vue.use(axios)
Vue.prototype.$axios = axios

Vue.use(BootstrapVue)

Vue.use(VueApexCharts)
Vue.component('apexchart', VueApexCharts)
Vue.use(VueGoogleCharts)

Vue.use(moment)
Vue.prototype.$moment = moment

Vue.use(Web3)
Vue.prototype.$web3 = Web3


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  template: '<App/>',
  components: {
    App
  }
})
