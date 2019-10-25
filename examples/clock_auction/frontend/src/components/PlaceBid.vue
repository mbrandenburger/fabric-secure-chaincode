<template>
    <div>
    <v-toolbar>
        <v-toolbar-title>Place Bid</v-toolbar-title>

        <v-spacer></v-spacer>

        <v-toolbar-items>
            <v-btn text>Stage 1</v-btn>
            <v-btn text>Round 4</v-btn>
            <v-btn text>Round ends: 1:49:33</v-btn>
            <v-btn text>2:39 PM ET 10/25/2019</v-btn>
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
            <tr v-for="item in desserts" :key="item.name">
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
    export default {
        name: 'Dashboard',
        props: ['dashboard'],
        data () {
            return {
                desserts: [
                    {
                        name: 'Elbograd',
                        supply: 10,
                        demand: 4,
                        mybids: 3,
                        status: "check",
                        quantity: 3,
                        price: 10000,
                    },
                    {
                        name: 'Mudberg',
                        supply: 10,
                        demand: 4,
                        mybids: 3,
                        status: "check",
                        quantity: 3,
                        price: 6000,
                    },
                    {
                        name: 'Deserton',
                        supply: 9,
                        demand: 4,
                        mybids: 3,
                        status: "check",
                        quantity: 3,
                        price: 6000,
                    },
                    {
                        name: 'Phlimsk',
                        supply: 6,
                        demand: 4,
                        mybids: 3,
                        status: "check",
                        quantity: 3,
                        price: 4000,
                    },

                ],
            }
        },
    }
</script>
