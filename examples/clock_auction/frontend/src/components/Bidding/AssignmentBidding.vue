<template>
    <div>

        <v-card
                class="mx-auto"
        >
            Assignment Bidding
            <v-row v-for="options in generateOptions(7,4)"
                   v-bind:key="options.price"
                   class="px-2 pb-2 ma-0"
                   justify="space-between"
            >
                Impairment Adjusted Price
                <v-btn-toggle
                        v-model="options.option"
                        multiple
                >
                    <v-btn
                            v-for="(item, index) in territories[0].licences"
                            v-bind:key="item.name"
                            color="primary"
                            disabled
                    >
                        {{ letters[index] }}
                    </v-btn>
                </v-btn-toggle>
                {{ options.price }}
                {{ options.option}}
            </v-row>
        </v-card>



    </div>
</template>

<script>
    export default {
        name: 'Dashboard',
        data() {
            return {
                search: '',
                letters: 'ABCDEFGHIJKLMNOPQRSTUVWXYZ'.split(''),
                selections: [0,1],
                territories: [
                    {
                        name: "elbograd",
                        blocks: 1,
                        options: [
                            {price: 0, option: [0]},
                            {price: 0, option: [1]},
                            {price: 0, option: [2]},
                            {price: 0, option: [3]},
                        ],
                        licences: [
                            { impariment: 10 },
                            { impariment: 10 },
                            { impariment: 10 },
                            { impariment: 10 },
                            { impariment: 10 },
                            { impariment: 10 },
                            { impariment: 10 },
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
                var options = []
                for (var i = 0; i <= licenses-blocks; i++) {
                    var opt = []
                    for ( var j = i; j < i + blocks; j++) {
                        opt.push(j)
                    }
                    options.push({
                        price: i,
                        option: opt,
                    })
                }

                return options
            }
        }
    }
</script>