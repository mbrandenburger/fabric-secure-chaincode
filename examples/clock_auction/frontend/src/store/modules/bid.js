// import axios from 'axios';

const state = {
    submittedBids: []
}

const getters = {
}

const actions = {
    submitBid ({commit}, bid) {
        commit('pushBid', bid)
    }
}

const mutations = {
    pushBid (state, payload) {

        var newBid = payload

        state.submittedBids.push(newBid)
    },
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}