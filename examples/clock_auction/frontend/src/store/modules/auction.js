import axios from 'axios';

const state = {
    auction : null,
    currentUser: ""
}

const getters = {
}

const actions = {
    LOAD_AUCTION({commit}) {
        axios
            .get('http://localhost:3000/auctions/2020')
            .then(r => r.data )
            .then(a => {
                console.log(a)
                commit('SET_AUCTION', a)
            })
    },
    UPDATE_CURRENT_PHASE({commit}, phase) {
        commit('SET_CURRENT_PHASE', phase)
    },
    UPDATE_CURRENT_ROUND({commit}, round) {
        commit('SET_CURRENT_ROUND', round)
    },
    UPDATE_CURRENT_USER({commit}, userid) {
        commit('SET_CURRENT_USER', userid)
        console.log("current user updated: " + userid);
    }
}

const mutations = {
    SET_AUCTION : (state, auction) => {
        state.auction = auction
    },
    SET_CURRENT_PHASE : (state, phase) => {
        state.auction.currentPhase = phase
    },
    SET_CURRENT_ROUND : (state, round) => {
        state.auction.currentRound = round
    },
    SET_CURRENT_USER : (state, userid) => {
        state.currentUser = userid
    }
}

export default {
    namespaced: true,
    state,
    getters,
    actions,
    mutations
}