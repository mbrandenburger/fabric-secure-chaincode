// import axios from 'axios';

const state = {
    submittedBids: [],
}

const getters = {
}

const actions = {
    submitBid ({commit}, bid) {
        commit('pushBid', bid)
    }
}

const mutations = {
    pushBid (state, bid ) {
        state.submittedBids.push(bid)
    },
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}