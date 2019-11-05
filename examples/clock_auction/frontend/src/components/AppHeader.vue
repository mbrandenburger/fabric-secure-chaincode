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
            <v-btn text>Phase {{auction.currentPhase}}</v-btn>
            <v-btn text>Round {{auction.currentRound}}</v-btn>
            <v-btn text>Round ends: 1:49:33</v-btn>
            <v-btn text>{{ currentTime }}</v-btn>
        </v-toolbar-items>

    </v-app-bar>
</template>

<script>
    import axios from 'axios';
    import moment from 'moment';

    export default {
        name: "AuctionHeader",
        data: function () {
            return {
                auction: null,
                currentTime: null,
            }
        },
        methods: {
            updateCurrentTime() {
                this.currentTime = moment().format('MM/DD/YYYY, h:mm:ss A z')
            },
        },
        beforeDestroy () {
            clearInterval(this.$options.interval);
        },
        mounted () {
            this.updateCurrentTime();
            this.$options.interval = setInterval(this.updateCurrentTime, 1000);
            axios
                .get('http://localhost:3000/auctions/2020')
                .then(response => {
                    this.auction = response.data
                })
                .catch(error => {
                    alert ("error: " + error)
                })
        }
    }
</script>