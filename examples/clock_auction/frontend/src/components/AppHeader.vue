<template>
    <v-app-bar
            :clipped-left="true"
            app
    >
        <v-toolbar-title class="headline text-uppercase">
            <span>Elbonia Communication Commmission</span>
            <span class="font-weight-light"> - Spectrum Auction</span>
        </v-toolbar-title>
        <v-spacer></v-spacer>

        <v-toolbar-items>
            <v-btn text>Phase {{auctionDetails.currentPhaseStatus}}</v-btn>
            <v-btn text>Round {{auctionDetails.currentRoundNumber}}</v-btn>
            <v-btn text>Round ends: 1:49:33</v-btn>
            <v-btn text>{{ new Date() }}</v-btn>
        </v-toolbar-items>

    </v-app-bar>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "AuctionHeader",
        data: function () {
            return {
                auctionDetails: null,
            }
        },
        mounted () {
            axios
                .get('http://localhost:3000/getAuctionDetails')
                .then(response => {
                    this.auctionDetails = response.data
                })
                .catch(error => {
                    alert ("error "+ error)
                    this.errored = true
                })
                .finally(() => this.loading = false)
        }
    }
</script>