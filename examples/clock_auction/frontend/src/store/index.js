import Vue from 'vue'
import Vuex from 'vuex'
import auction from './modules/auction'
import bid from './modules/bid'

Vue.use(Vuex)

export default new Vuex.Store({
    modules: {
        auction,
        bid
    },
    strict: process.env.NODE_ENV !== 'production',
})