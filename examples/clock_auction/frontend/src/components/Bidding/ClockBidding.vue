<template>
    <v-container fluid>
        <v-row>
            <v-col cols="12">Clock Bidding</v-col>
        </v-row>
        <v-row>
            <v-col cols="3">
                <v-card dark>
                    <div class="d-flex flex-no-wrap justify-space-between">
                        <div>
                            <v-card-title>Total commitment</v-card-title>
                            <v-card-subtitle class="headline">$ {{ totalCommitment }}</v-card-subtitle>
                        </div>
                    </div>
                </v-card>
            </v-col>
            <v-col cols="3">
                <v-card dark >
                    <div class="d-flex flex-no-wrap justify-space-between">
                        <div>
                            <v-card-title>Eligibility</v-card-title>
                            <v-card-subtitle class="headline">{{ eligibility }}</v-card-subtitle>
                        </div>
                    </div>
                </v-card>
            </v-col>
            <v-col cols="3">
                <v-card dark >
                    <div class="d-flex flex-no-wrap justify-space-between">
                        <div>
                            <v-card-title>Required activity</v-card-title>
                            <v-card-subtitle class="headline">{{ requiredActivity }}</v-card-subtitle>
                        </div>
                    </div>
                </v-card>
            </v-col>
            <v-col cols="3">
                <v-card dark >
                    <div class="d-flex flex-no-wrap justify-space-between">
                        <div>
                            <v-card-title>Requested activity</v-card-title>
                            <v-card-subtitle class="headline">{{ requestedActivity }}</v-card-subtitle>
                        </div>
                    </div>
                </v-card>
            </v-col>
        </v-row>

        <v-row>
            <v-col cols="12">
                <v-card>
                    <v-card-title>
                        <v-spacer></v-spacer>
                        <v-text-field
                                v-model="search"
                                append-icon="fa-search"
                                label="Search"
                                single-line
                                hide-details
                        ></v-text-field>
                    </v-card-title>
                    <v-card-text class="pa-0">
                        <v-data-table v-if="auction.currentRound === 1"
                                      :headers="tableHeader"
                                      :items="territories"
                                      :search="search"
                                      hide-default-footer
                        >
                            <template v-slot:item.minPrice="props">
                                $ {{ props.item.minPrice}}
                            </template>

                            <template v-slot:item.quantity="props">
                                <v-edit-dialog
                                        :return-value.sync="props.item.quantity"
                                        large
                                        persistent
                                        @save="save"
                                >
                                    <div>{{ props.item.quantity}}</div>
                                    <template v-slot:input>
                                        <div class="mt-4 title">Update quantity</div>
                                    </template>
                                    <template v-slot:input>
                                        <v-text-field
                                                v-model="props.item.quantity"
                                                type="number"
                                                :rules="[q => q >= 0 && q <= props.item.supply || 'Invalid quantity']"
                                                label="Edit"
                                                single-line
                                                autofocus
                                        ></v-text-field>
                                    </template>
                                </v-edit-dialog>
                            </template>
                        </v-data-table>
                        <v-data-table v-else
                                      :headers="tableHeader"
                                      :items="territories"
                                      :search="search"
                                      hide-default-footer
                        >
                            <template v-slot:item.minPrice="props">
                                $ {{ props.item.minPrice }}
                            </template>

                            <template v-slot:item.clockPrice="props">
                                $ {{ props.item.clockPrice }}
                            </template>

                            <template v-slot:item.price="props">
                                <v-edit-dialog
                                        :return-value.sync="props.item.price"
                                        large
                                        persistent
                                        @save="save"
                                >
                                    <div>{{ props.item.price}}</div>
                                    <template v-slot:input>
                                        <div class="mt-4 title">Update bid price</div>
                                    </template>
                                    <template v-slot:input>
                                        <v-text-field
                                                v-model="props.item.price"
                                                type="number"
                                                :rules="[q => q >= props.item.minPrice && q <= props.item.clockPrice || 'Invalid price']"
                                                label="Edit"
                                                single-line
                                                autofocus
                                        ></v-text-field>
                                    </template>
                                </v-edit-dialog>
                            </template>

                            <template v-slot:item.quantity="props">
                                <v-edit-dialog
                                        :return-value.sync="props.item.quantity"
                                        large
                                        persistent
                                        @save="save"
                                >
                                    <div>{{ props.item.quantity}}</div>
                                    <template v-slot:input>
                                        <div class="mt-4 title">Update quantity</div>
                                    </template>
                                    <template v-slot:input>
                                        <v-text-field
                                                v-model="props.item.quantity"
                                                type="number"
                                                :rules="[q => q >= 0 && q <= props.item.supply || 'Invalid quantity']"
                                                label="Edit"
                                                single-line
                                                autofocus
                                        ></v-text-field>
                                    </template>
                                </v-edit-dialog>
                            </template>

                        </v-data-table>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>

        <v-row>
            <v-col cols="12" right>
                <div class="d-flex justify-end">
                    <v-btn color="primary" @click="prepareBid">Submit your bid</v-btn>
                </div>
            </v-col>
        </v-row>


        <v-overlay :value="waitingOverlay">
            <v-progress-circular indeterminate size="64"></v-progress-circular>
        </v-overlay>

        <v-dialog
                v-model="confirmDialog"
                max-width="400"
        >
            <v-card>
                <v-card-title class="headline">Confirm your clock bid</v-card-title>
                <v-card-text>
                    <v-simple-table>
                        <template v-slot:default>
                            <thead>
                            <tr>
                                <th class="text-left">Territory</th>
                                <th class="text-left">Quantity</th>
                                <th class="text-left">Price</th>
                            </tr>
                            </thead>
                            <tbody>
                            <tr v-for="license in currentBid.licences" :key="license.territory">
                                <td>{{ license.territory}}</td>
                                <td>$ {{ license.price }}</td>
                                <td>{{ license.quantity }}</td>
                            </tr>
                            </tbody>
                        </template>
                    </v-simple-table>
                </v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn
                            color="red darken-1"
                            text
                            @click="confirmDialog = false"
                    >
                        Cancel
                    </v-btn>
                    <v-btn
                            color="green darken-1"
                            text
                            @click="submitBid"
                    >
                        Submit
                    </v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>

    </v-container>
</template>

<script>
    import { mapState } from 'vuex'
    import axios from 'axios';

    export default {
        name: 'Dashboard',
        data () {
            return {
                search: '',
                confirmDialog: false,
                waitingOverlay: false,
                territories: [],
                totalCommitment: 0,
                eligibility: 0,
                requiredActivity: 0,
                requestedActivity: 0,
                currentBid: null,
            }
        },
        watch: {
            // this is a helper function to emulate some waiting :D
            waitingOverlay (val) {
                val && setTimeout(() => {
                    this.waitingOverlay = false
                }, 1500)
            },
        },
        computed: {
            tableHeader: function () {
                if (this.auction.currentRound === 1) {
                    return [
                        { text: 'Territory', value: 'name', },
                        { text: 'Supply', value: 'supply' },
                        { text: 'Opening Price', value: 'minPrice' },
                        { text: 'Quantity', value: 'quantity' },
                    ]
                } else {
                    return [
                        { text: 'Territory', value: 'name', },
                        { text: 'Supply', value: 'supply' },
                        { text: 'Agg. Demand', value: 'demand' },
                        { text: 'Min Price', value: 'minPrice' },
                        { text: 'Clock Price', value: 'clockPrice' },
                        { text: 'Bid Price', value: 'price' },
                        { text: 'Quantity', value: 'quantity' },
                    ]
                }
            },
            ...mapState({
                auction: state  => state.auction.auction
            })
        },
        mounted () {
            this.waitingOverlay = true
            axios
                .get('http://localhost:3000/clock_rounds/' + this.auction.currentRound)
                .then(response => {
                    this.territories = response.data.territories
                    this.eligibility = response.data.participants[0].eligibility
                    this.requiredActivity = response.data.participants[0].eligibility * 0.8
                })
                .catch(error => {
                    alert ("error "+ error)
                }).finally(
                    this.waitingOverlay = false
                )
        },
        methods: {
            save () {
                this.totalCommitment = 0
                this.requestedActivity = 0
                for (let i=0; i<this.territories.length; i++) {
                    this.totalCommitment += Number(this.territories[i].quantity) * Number(this.territories[i].minPrice)
                    this.requestedActivity += Number(this.territories[i].quantity) * Number(this.territories[i].biddingUnits)
                }
            },

            prepareBid: function() {
                this.currentBid = {
                    id: "bla",
                    round: this.auction.currentRound,
                    licences: [],
                    status: "submitted"
                }

                for (let i=0; i < this.territories.length; i++) {
                    if (Number(this.territories[i].quantity) > 0) {
                        this.currentBid.licences.push({
                            territory: this.territories[i].name,
                            price:  Number(this.territories[i].price),
                            quantity: Number(this.territories[i].quantity)
                        })
                    }
                }
                this.confirmDialog = true
            },

            submitBid: function () {
                this.confirmDialog = false
                this.waitingOverlay = true
                this.$store.dispatch('bid/submitBid', this.bid)
                    .then(() => {
                        this.currentBid = null
                    })
            }
        }
    }
</script>
