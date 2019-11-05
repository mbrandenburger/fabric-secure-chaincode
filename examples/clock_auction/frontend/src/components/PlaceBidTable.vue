<template>
        <v-container fluid>
            <v-row>
                <v-col cols="3">
                    <v-card dark >
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
                                <v-card-subtitle class="headline">{{ eligibility }} BUs</v-card-subtitle>
                            </div>
                        </div>
                    </v-card>
                </v-col>
                <v-col cols="3">
                    <v-card dark >
                        <div class="d-flex flex-no-wrap justify-space-between">
                            <div>
                                <v-card-title>Required activity</v-card-title>
                                <v-card-subtitle class="headline">{{ requiredActivity }} BUs</v-card-subtitle>
                            </div>
                        </div>
                    </v-card>
                </v-col>
                <v-col cols="3">
                    <v-card dark >
                        <div class="d-flex flex-no-wrap justify-space-between">
                            <div>
                                <v-card-title>Requested activity</v-card-title>
                                <v-card-subtitle class="headline">{{ requestedActivity }} BUs</v-card-subtitle>
                            </div>
                        </div>
                    </v-card>
                </v-col>
            </v-row>
            <v-row>
                <v-col cols="12">
                    <v-card>
                        <v-card-title>Build your bid</v-card-title>
                        <v-card-text class="pa-0">
                            <v-data-table
                                    :headers="headers"
                                    :items="territories"
                                    hide-default-footer
                            >
                                <template v-slot:item.openingPrice="props">
                                    $ {{ props.item.openingPrice }}
                                </template>

                                <template v-slot:item.quantity="props">
                                    <v-edit-dialog
                                            :return-value.sync="props.item.quantity"
                                            large
                                            persistent
                                            @save="save"
                                            @cancel="cancel"
                                            @open="open"
                                            @close="close"
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
                        <v-card-actions>
                            <v-spacer></v-spacer>
                            <v-btn color="primary">Submit your bid</v-btn>
                        </v-card-actions>
                    </v-card>
                </v-col>
            </v-row>
        </v-container>




</template>

<script>
    import axios from 'axios';

    export default {
        name: 'Dashboard',
        props: ['dashboard'],
        data () {
            return {
                pagination: {},
                headers: [
                    { text: 'Territory', value: 'name', },
                    { text: 'Supply', value: 'supply' },
                    { text: 'BUs', value: 'biddingUnits' },
                    { text: 'Price', value: 'openingPrice' },
                    { text: 'Quantity', value: 'quantity' },
                ],
                territories: [],
                totalCommitment: 0,
                eligibility: 0,
                requiredActivity: 0,
                requestedActivity: 0,
            }
        },
        mounted () {
            axios
                .get('http://localhost:3000/auctions/2020')
                .then(response => {
                    this.territories = response.data.territories
                    this.eligibility = response.data.participants[0].eligibility
                    this.requiredActivity = response.data.participants[0].eligibility * 0.8
                })
                .catch(error => {
                    alert ("error "+ error)
                })
        },
        methods: {
            save () {
                this.totalCommitment = 0
                this.requestedActivity = 0
                for (var i=0; i<this.territories.length;i++) {
                    this.totalCommitment += Number(this.territories[i].quantity) * Number(this.territories[i].openingPrice)
                    this.requestedActivity += Number(this.territories[i].quantity) * Number(this.territories[i].biddingUnits)
                }
            },
            submitBid: function () {
                //alert (this.territories[0].name)
                for (var i=0; i<this.territories.length;i++) {
                    this.bidTotal += Number(this.territories[i].quantity)
                    this.bidPriceTotal += Number(this.territories[i].price)
                }
                alert ("bidTotal "+this.bidTotal+", bidTotalPrice "+this.bidPriceTotal)
            }
        }
    }
</script>
