<template>
    <v-container fluid>
        <v-row>
            <v-col cols="12">Assignment Bidding</v-col>
        </v-row>

        <v-row>
            <v-col cols="12">
                <v-card
                        class="mx-auto"
                >
                    <v-tabs
                            v-model="tab"
                            background-color="orange darken-3"
                            dark
                            text
                    >

                        <v-tab
                                v-for="territory in territories"
                                :key="territory.name"
                                :href="`#tab-${territory.name}`"
                        >
                            {{ territory.name }}
                        </v-tab>

                        <v-tab-item
                                v-for="territory in territories"
                                :key="territory.name"
                                :value="'tab-' + territory.name"
                        >

                            <v-simple-table>
                                <template v-slot:default>
                                    <thead>
                                    <tr>
                                        <th class="text-left">Impairment Adjusted Price</th>
                                        <th class="text-left">Channels</th>
                                        <th class="text-left">Status</th>
                                        <th class="text-left">Bid Price</th>
                                    </tr>
                                    </thead>
                                    <tbody>
                                    <tr v-for="options in generateOptions(territory.clockPrice, territory.channels, territory.licenses)"
                                        :key="options.id"
                                    >
                                        <td>$ {{ options.price }}</td>
                                        <td>
                                            <v-btn-toggle
                                                    v-model="options.channels"
                                                    multiple
                                            >
                                                <v-btn
                                                        v-for="(item, index) in territory.channels"
                                                        v-bind:key="item.name"
                                                        color="primary"
                                                        disabled
                                                >
                                                    {{ letters[index] }}
                                                </v-btn>
                                            </v-btn-toggle>
                                        </td>
                                        <td> {{ options.status }}</td>
                                        <td>
                                            <v-edit-dialog
                                                    :return-value.sync="options.price"
                                                    large
                                                    @save="save"
                                            >
                                                <div>$ {{ options.bid }}</div>
                                                <template v-slot:input>
                                                    <div class="mt-4 title">Update quantity</div>
                                                </template>
                                                <template v-slot:input>
                                                    <v-text-field
                                                            v-model="options.bid"
                                                            type="number"
                                                            :rules="[q => q >= 0 || 'Invalid price']"
                                                            label="Edit"
                                                            single-line
                                                            autofocus
                                                    ></v-text-field>
                                                </template>
                                            </v-edit-dialog>
                                        </td>
                                    </tr>
                                    </tbody>
                                </template>
                            </v-simple-table>
                        </v-tab-item>
                    </v-tabs>
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
                <v-card-title class="headline">Confirm your assignment bid</v-card-title>
                <v-card-text>
                    <v-simple-table>
                        <template v-slot:default>
                            <thead>
                            <tr>
                                <th class="text-left">Territory</th>
                                <th class="text-left">Channels</th>
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
    import axios from 'axios';

    export default {
        name: 'Dashboard',
        data() {
            return {
                search: '',
                letters: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split(''),
                tab: null,
                confirmDialog: false,
                waitingOverlay: false,
                territories: null,
                totalCommitment: 0,
                eligibility: 0,
                requiredActivity: 0,
                requestedActivity: 0,
            }
        },
        methods: {
            generateOptions: function(channelPrice, channels, wonChannels) {
                const numChannels = channels.length
                let options = []
                for (let i = 0; i <= numChannels-wonChannels; i++) {
                    let chs = []
                    let price = 0
                    for (let j = i; j < i + wonChannels; j++) {
                        price += channelPrice * (100 - channels[j].impairment) / 100
                        chs.push(j)
                    }
                    options.push({
                        id: i,
                        price: price,
                        bid: 0,
                        status: "-",
                        channels: chs,
                    })
                }

                return options
            },
            save () {
                this.totalCommitment = 0
                this.requestedActivity = 0
                for (let i=0; i<this.territories.length; i++) {
                    this.totalCommitment += Number(this.territories[i].quantity) * Number(this.territories[i].minPrice)
                    this.requestedActivity += Number(this.territories[i].quantity)
                }
            },
            prepareBid: function() {
                this.currentBid = {
                    id: "bla",
                    licences: [],
                    status: "submitted"
                }

                for (let i=0; i < this.territories.length; i++) {
                    if (Number(this.territories[i].quantity) > 0) {
                        this.currentBid.licences.push({
                            territory: this.territories[i].name,
                            price:  Number(this.territories[i].openingPrice),
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
        },
        mounted () {
            axios.get('http://localhost:3000/assignment')
                .then(response => {
                    this.territories = response.data.territories
                })
        }
    }
</script>