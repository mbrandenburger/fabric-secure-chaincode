<template>
    <v-app-bar
            :clipped-left=true
            app
    >
        <v-toolbar-title class="headline text-uppercase">
            <span>Elbonia Communication Commmission</span>
            <span class="font-weight-light"> - Spectrum Auction</span>
        </v-toolbar-title>
        <v-spacer></v-spacer>

        <v-toolbar-items v-if="auction">
            <v-menu offset-y>
                <template v-slot:activator="{ on }">
                    <v-btn
                            color="primary"
                            dark
                            v-on="on"
                    >
                        Phase {{ auction.currentPhase }}
                    </v-btn>
                </template>
                <v-list>
                    <v-list-item
                            v-for="index in 2"
                            :key="index"
                            @click="UPDATE_CURRENT_PHASE(index)"
                    >
                        <v-list-item-title>Phase {{ index }}</v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>

            <v-menu offset-y v-if="auction.currentPhase === 1">
                <template v-slot:activator="{ on }">
                    <v-btn
                            color="secondary"
                            dark
                            v-on="on"
                    >
                        Round {{ auction.currentRound }}
                    </v-btn>
                </template>
                <v-list>
                    <v-list-item
                            v-for="index in 3"
                            :key="index"
                            @click="UPDATE_CURRENT_ROUND(index)"
                    >
                        <v-list-item-title>Round {{ index }}</v-list-item-title>
                    </v-list-item>
                </v-list>
            </v-menu>

            <v-btn text>{{ currentTime }}</v-btn>
        </v-toolbar-items>

    </v-app-bar>
</template>

<script>
    import { mapState, mapActions } from 'vuex'
    import moment from 'moment';

    export default {
        name: "AuctionHeader",
        data: function () {
            return {
                currentTime: null,
            }
        },
        computed: mapState({
            auction: state => state.auction.auction
        }),
        methods: {
            updateCurrentTime() {
                this.currentTime = moment().format('MM/DD/YYYY, h:mm:ss A z')
            },
            ...mapActions('auction', [
                'UPDATE_CURRENT_PHASE',
                'UPDATE_CURRENT_ROUND'
            ])
        },
        beforeDestroy () {
            clearInterval(this.$options.interval);
        },
        mounted () {
            this.updateCurrentTime();
            this.$options.interval = setInterval(this.updateCurrentTime, 1000);
        }
    }
</script>