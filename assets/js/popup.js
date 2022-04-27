import Vue from 'vue';

import { library } from '@fortawesome/fontawesome-svg-core'
import { faCoffee, faGears, faTimeline, faWallet } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faCoffee)
library.add(faGears)
library.add(faTimeline)
library.add(faWallet)

Vue.component('font-awesome-icon', FontAwesomeIcon)

import Popup from './components/Popup.vue';
const app = new Vue({
    el: '#app',
    render: createElement => createElement(Popup)
});