<template>
    <div>
        <v-toolbar>
            <v-toolbar-title>Auction Info</v-toolbar-title>

            <v-spacer></v-spacer>

        </v-toolbar>

        <v-container grid-list-xl fluid>
            <v-layout row wrap>
                <v-flex lg4 sm6 xs12>
                    <v-card>
                        <v-card-text class="pa-0">
                            <v-container class="pa-0">
                                <div class="layout row ma-0">
                                    <div class="sm6 xs6 flex py-3" :class="color">
                                        <div class="headline">{{ auctionDetails.auctionName }}</div>
                                        <span class="caption">conducted by {{ auctionDetails.owner }}</span>
                                    </div>
                                </div>
                            </v-container>
                        </v-card-text>
                    </v-card>
                </v-flex>
                <v-flex lg2 sm6 xs12>
                    <v-card>
                        <v-card-text class="pa-0">
                            <v-container class="pa-0">
                                <div class="layout row ma-0">
                                    <div class="sm6 xs6 flex">
                                        <div class="layout column ma-0 justify-center align-center">
                                            <v-icon size="56px" :color="color">fa-hashtag</v-icon>
                                        </div>
                                    </div>
                                    <div class="sm6 xs6 flex text-sm-center py-3" :class="color">
                                        <div class="headline">Round</div>
                                        <span class="caption">{{ auctionDetails.currentRoundNumber }}</span>
                                    </div>
                                </div>
                            </v-container>
                        </v-card-text>
                    </v-card>
                </v-flex>
                <v-flex lg2 sm6 xs12>
                    <v-card>
                        <v-card-text class="pa-0">
                            <v-container class="pa-0">
                                <div class="layout row ma-0">
                                    <div class="sm6 xs6 flex">
                                        <div class="layout column ma-0 justify-center align-center">
                                            <v-icon size="56px" :color="color">fa-hashtag</v-icon>
                                        </div>
                                    </div>
                                    <div class="sm6 xs6 flex text-sm-center py-3" :class="color">
                                        <div class="headline">Phase</div>
                                        <span class="caption">{{ auctionDetails.currentPhaseStatus }}</span>
                                    </div>
                                </div>
                            </v-container>
                        </v-card-text>
                    </v-card>
                </v-flex>


                <v-flex lg4 sm6 xs12>
                    <v-card>
                        <v-card-text class="pa-0">
                            <v-container class="pa-0">
                                <v-simple-table>
                                    <template v-slot:default>
                                        <thead>
                                        <tr>
                                            <th class="text-left">Name</th>
                                            <th class="text-left">Details</th>
                                        </tr>
                                        </thead>
                                        <tbody>
                                        <tr v-for="item in auctionDetails.bidders" :key="item.name">
                                            <td>{{ item.name }}</td>
                                            <td> - </td>
                                        </tr>
                                        </tbody>
                                    </template>
                                </v-simple-table>
                            </v-container>
                        </v-card-text>
                    </v-card>
                </v-flex>


                <v-flex lg4 sm6 xs12>
                    <v-card>
                        <v-card-text class="pa-0">
                            <v-container class="pa-0">
                                <v-simple-table>
                                    <template v-slot:default>
                                        <thead>
                                        <tr>
                                            <th class="text-left">Territory</th>
                                            <th class="text-left">Supply</th>
                                            <th class="text-left">Opening bidding price</th>
                                        </tr>
                                        </thead>
                                        <tbody>
                                        <tr v-for="item in auctionDetails.territories" :key="item.name">
                                            <td>{{ item.name }}</td>
                                            <td>{{ item.numLicenses }}</td>
                                            <td>{{ item.minimumPrice }}</td>
                                        </tr>
                                        </tbody>
                                    </template>
                                </v-simple-table>
                            </v-container>
                        </v-card-text>
                    </v-card>
                </v-flex>

                <v-flex lg12 sm6 xs12>
                    <v-card>
                        <v-card-text class="pa-0">
                            <v-container class="pa-0">
                                {{ auctionDetails }}
                            </v-container>
                        </v-card-text>
                    </v-card>
                </v-flex>
            </v-layout>
        </v-container>

    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'Dashboard',
        props: ['dashboard'],
        data () {
            return {
                auctionDetails: [],
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
