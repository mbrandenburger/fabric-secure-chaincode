<template>
    <div>
    <v-toolbar>
        <v-toolbar-title>Place Bid</v-toolbar-title>

        <v-spacer></v-spacer>
        <v-toolbar-items>
            <v-btn text>Required Commitment: $6,000,000</v-btn>
            <v-btn text>Eligibility: </v-btn>
            <v-btn text>Submitted Activity: </v-btn>
            <!--<v-btn depressed small color="primary">Submit</v-btn>-->
            <v-btn v-on:click="submitBid" raised small color="primary">Submit</v-btn>
        </v-toolbar-items>
    </v-toolbar>

    <v-data-table
            :headers="headers"
            :items="desserts"
            :items-per-page="5"
            class="elevation-1"
    ></v-data-table>

    <v-simple-table>
        <template v-slot:default>
            <thead>
            <tr>
                <th class="text-left">Geographic Parcel</th>
                <th class="text-left">Supply</th>
                <th class="text-left">Agg Demand</th>
                <th class="text-left">My Processed Bids</th>
                <th class="text-left">Status</th>
                <th class="text-left">Quantity</th>
                <th class="text-left">Price</th>
                <th class="text-left">Actions</th>
            </tr>
            </thead>
            <tbody>
            <tr v-for="item in territories" :key="item.name">
                <td>{{ item.name }}</td>
                <td>{{ item.supply}}</td>
                <td>{{ item.demand}}</td>
                <td>{{ item.mybids}}</td>
                <td>{{ item.status}}</td>
                <td>{{ item.quantity}}</td>
                <td>${{ item.price}}</td>
                <td>
                    <v-btn depressed small color="primary">Change</v-btn>
                    <v-btn depressed small color="primary" disabled>Add New </v-btn>
                </td>
            </tr>
            </tbody>
        </template>
    </v-simple-table>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'Dashboard',
        props: ['dashboard'],
        data () {
            return {
                territories: [],
            }
        },
        mounted () {
            axios
                .get('http://localhost:3000/getAuctionDetails')
                .then(response => {
                    this.territories = response.data.territories
                })
                .catch(error => {
                    alert ("error "+ error)
                    this.errored = true
                })
                .finally(() => this.loading = false)
        }
    }
</script>
