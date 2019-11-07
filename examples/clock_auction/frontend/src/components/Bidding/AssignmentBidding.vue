<template>
    <div>

        <v-card
                class="mx-auto"
        >
            <v-card-title>
                Assignment Bidding
            </v-card-title>

            <v-list-item two-line
                         v-for="territory in territories"
                         v-bind:key="territory.name"
                         class="px-2 pb-2 ma-2"
                         justify="space-between"
            >
                <v-list-item-content>
                    <v-list-item-title> {{ territory.name }}</v-list-item-title>
                    <v-list-item-subtitle>
                        <v-simple-table>
                            <template v-slot:default>
                                <thead>
                                <tr>
                                    <th class="text-left">Impairment Adjusted Price</th>
                                    <th class="text-left">Channels</th>
                                    <th class="text-left">Status</th>
                                    <th class="text-left">Price</th>
                                </tr>
                                </thead>
                                <tbody>
                                <tr v-for="options in generateOptions(territory.licences, territory.blocks)" :key="options.price">
                                    <td>$ 0</td>
                                    <td>
                                        <v-btn-toggle
                                                v-model="options.channels"
                                                multiple
                                        >
                                            <v-btn
                                                    v-for="(item, index) in territory.licences"
                                                    v-bind:key="item.name"
                                                    color="primary"
                                                    disabled
                                            >
                                                {{ letters[index] }}
                                            </v-btn>
                                        </v-btn-toggle>
                                    </td>
                                    <td> status </td>
                                    <td>$ {{ options.price }}</td>
                                </tr>
                                </tbody>
                            </template>
                        </v-simple-table>
                    </v-list-item-subtitle>
                </v-list-item-content>
            </v-list-item>
        </v-card>

    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'Dashboard',
        data() {
            return {
                search: '',
                letters: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split(''),
                territories: [
                    {
                        name: "Elbograd",
                        blocks: 4,
                        clockPrce: 10000,
                        licences: [
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 10 },
                            { impariment: 10 },
                            { impariment: 0 },
                        ]
                    },
                    {
                        name: "Mudberg",
                        blocks: 2,
                        clockPrce: 10000,
                        licences: [
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                        ]
                    },
                    {
                        name: "Deserton",
                        blocks: 3,
                        clockPrce: 10000,
                        licences: [
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 10 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                        ]
                    },
                    {
                        name: "Phlimsk",
                        blocks: 2,
                        clockPrce: 10000,
                        licences: [
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 0 },
                            { impariment: 10 },
                            { impariment: 0 },
                        ]
                    },
                ],
                totalCommitment: 0,
                eligibility: 0,
                requiredActivity: 0,
                requestedActivity: 0,
            }
        },
        methods: {
            generateOptions: function(licenses, blocks) {
                const numLicenses = licenses.length
                let options = []
                for (let i = 0; i <= numLicenses-blocks; i++) {
                    let chs = []
                    for (let j = i; j < i + blocks; j++) {
                        chs.push(j)
                    }
                    options.push({
                        price: 0,
                        channels: chs,
                    })
                }

                return options
            }
        },
        mounted () {
            axios.get('http://localhost:3000/assignment')
                .then(response => {
                    this.territories = response.data.territories
                })
            // .catch(error => {
            //     alert ("error "+ error)
            // })
        }
    }
</script>